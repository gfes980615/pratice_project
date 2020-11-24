package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"gitlab.paradise-soft.com.tw/backend/yaitoo/db"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/tracer"

	"gitlab.paradise-soft.com.tw/backend/athena/restful/glob"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/core"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/core/helper"

	"gitlab.paradise-soft.com.tw/backend/athena/restful/services/mgmt"
	"gitlab.paradise-soft.com.tw/backend/athena/restful/srv"
	"gitlab.paradise-soft.com.tw/glob/utils/timetool"
)

const (
	platformMessage          = "『%s』， %s ~ %s， 投注金额： %.2f，营利金额： %.2f ， 请留意追踪。"
	productKeepLossMessage   = "『%s』 频道： %s ， 产品： %s ， %s ~ %s， 持续亏损 %d 小时， 请留意追踪。"
	productProportionMessage = "『%s』 频道： %s ， 产品： %s ， %s ~ %s， 投注金额： ¥%.2f，亏损金额： ¥%.2f， 请留意追踪。"
	productMessage           = "『%s』 频道： %s ， 产品： %s ， %s ~ %s， 亏损金额： ¥%.2f， 请留意追踪。"
	sampleMessage            = "『%s』， %s ~ %s 的%s： %.2f， %s前七日单日平均值： %.2f， 请留意追踪。"
	winlost                  = "输赢监控"
	withdraw                 = "出款监控"
	deposit                  = "入款监控"
	member                   = "会员监控"

	sevenDay int = 7
	onDay    int = 1
	now      int = 0
)

var (
	rwMutex     *sync.RWMutex
	taskChan    map[string](chan *mgmt.RiskNotificationEvent)
	listenerMap map[string]*ListenerMeta

	eventHandleMap = map[string]func(task *mgmt.RiskNotificationEvent){
		glob.PlatformEarning:     platformEarningNotification,
		glob.ProductKeepLoss:     productKeepLossNotification,
		glob.ProductLossPercent:  productEarningPercentNotification,
		glob.ProductLossAmount:   productEarningNotification,
		glob.WithdrawAmount:      withdrawHigherEvent,
		glob.DepositAmountHigher: depositEvent,
		glob.DepositAmountLower:  depositEvent,
		glob.MemberRegister:      registerMember,
	}

	thresholdName = map[string]string{
		glob.WithdrawAmount:      "出款金额",
		glob.DepositAmountHigher: "入款金额",
		glob.DepositAmountLower:  "入款金额",
		glob.MemberRegister:      "注册会员数",
	}

	timeMap = map[string]time.Duration{
		"now":         0,
		"oneDayAgo":   -1 * 24,
		"eightDayAgo": -8 * 24,
	}
)

type ListenerMeta struct {
	m               *sync.RWMutex
	LastTime        core.DateTime
	TriggerInterval time.Duration
}

func init() {
	rwMutex = &sync.RWMutex{}
	taskChan = make(map[string](chan *mgmt.RiskNotificationEvent))
	taskChan[glob.PlatformEarning] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.ProductKeepLoss] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.ProductLossPercent] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.ProductLossAmount] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.WithdrawAmount] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.DepositAmountHigher] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.DepositAmountLower] = make(chan *mgmt.RiskNotificationEvent)
	taskChan[glob.MemberRegister] = make(chan *mgmt.RiskNotificationEvent)
	listenerMap = make(map[string]*ListenerMeta, 10)
	CheckEventTask()
}

func GetListenerEvent() {
	if items, err := srv.RiskNotificationMgmt.QueryEventItems(cc, "get_all_event", ""); err != nil {
		tracer.Error("notification", "GetListenerEvent", err.Error())
	} else {
		for _, item := range items {
			if item.MonitorStatus == 1 {
				taskChan[item.EventCode] <- item
			}
		}
	}
}

// 測試用func
func ExecuteNotificationAllEvents() {
	if items, err := srv.RiskNotificationMgmt.QueryEventItems(cc, "get_all_event", ""); err != nil {
		tracer.Error("notification", "GetListenerEvent", err.Error())
	} else {
		for _, item := range items {
			eventHandleMap[item.EventCode](item)
		}
	}
}

func CheckEventTask() {
	for i := 0; i < 5; i++ {
		go func() {
			for {
				select {
				case task := <-taskChan[glob.PlatformEarning]:
					check(task, eventHandleMap[glob.PlatformEarning])
				case task := <-taskChan[glob.ProductKeepLoss]:
					check(task, eventHandleMap[glob.ProductKeepLoss])
				case task := <-taskChan[glob.ProductLossPercent]:
					check(task, eventHandleMap[glob.ProductLossPercent])
				case task := <-taskChan[glob.ProductLossAmount]:
					check(task, eventHandleMap[glob.ProductLossAmount])
				case task := <-taskChan[glob.WithdrawAmount]:
					check(task, eventHandleMap[glob.WithdrawAmount])
				case task := <-taskChan[glob.DepositAmountHigher]:
					check(task, eventHandleMap[glob.DepositAmountHigher])
				case task := <-taskChan[glob.DepositAmountLower]:
					check(task, eventHandleMap[glob.DepositAmountLower])
				case task := <-taskChan[glob.MemberRegister]:
					check(task, eventHandleMap[glob.MemberRegister])
				}
			}
		}()
	}
}

func (this *ListenerMeta) checkTrigger() bool {
	this.m.RLock()
	defer this.m.RUnlock()
	nowUTC := time.Now().UTC()
	res := nowUTC.Sub(this.LastTime.UTC()).Seconds() >= this.TriggerInterval.Seconds()
	if res {
		this.LastTime.Parse(nowUTC)
	}
	return res
}

func (this *ListenerMeta) setTriggerInterval(interval int) {
	this.m.Lock()
	defer this.m.Unlock()
	this.TriggerInterval = time.Duration(interval) * time.Second
}

func getTriggerTime(frequency int, frequencyType string) int {
	switch frequencyType {
	case "s":
		return frequency
	case "m":
		return 60 * frequency
	case "h":
		return 60 * 60 * frequency
	case "d":
		return 24 * 60 * 60 * frequency
	}
	return 0
}

func check(task *mgmt.RiskNotificationEvent, handle func(task *mgmt.RiskNotificationEvent)) {
	rwMutex.RLock()
	listener, found := listenerMap[task.EventCode]
	rwMutex.RUnlock()

	if found {
		if !listener.checkTrigger() {
			return
		}
	}

	listener = &ListenerMeta{m: &sync.RWMutex{}}
	listener.LastTime.Parse(time.Now().UTC())
	frequency := getTriggerTime(task.Frequency, task.FrequencyType)
	listener.setTriggerInterval(frequency)

	rwMutex.Lock()
	listenerMap[task.EventCode] = listener
	rwMutex.Unlock()

	handle(task)
}

func platformEarningNotification(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	sTime := getTime(timeMap["oneDayAgo"], helper.Format_DateTime)
	eTime := getTime(timeMap["now"], helper.Format_DateTime)
	for subTarget, logic := range eventLogicMap {
		products, err := srv.RiskNotificationMgmt.GetProductEarning(sTime, eTime, subTarget, "notification/get_all_product_earning")
		if err != nil {
			tracer.Error("notification", "productKeepLossNotification", err.Error())
			return
		}

		if !checkLogic(glob.FloatRound(products[0].Earning/products[0].Total), 1, logic) {
			message := getMessage(products[0], task, sTime, eTime, task.EventCode)
			glob.MessageToSlack(winlost, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
			insertMessageToDB(task.ID, message, "")
		}
	}
}

func productKeepLossNotification(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	for subTarget, logic := range eventLogicMap {
		sTime := getTime(-1*time.Duration(logic[0].Threshold), helper.Format_DateTime)
		eTime := getTime(timeMap["now"], helper.Format_DateTime)
		products, err := srv.RiskNotificationMgmt.GetProductEarning(sTime, eTime, subTarget, "notification/get_product_by_time")
		if err != nil {
			tracer.Error("notification", "productKeepLossNotification", err.Error())
			return
		}
		for _, product := range products {
			if !checkLogic(float64(product.ProductCount), 0, logic) {
				task.Threshold = logic[0].Threshold
				message := getMessage(product, task, sTime, eTime, task.EventCode)
				glob.MessageToSlack(winlost, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
				insertMessageToDB(task.ID, message, product.ChannelCode+"/"+product.ProductCode)
			}
		}
	}
}

func productEarningNotification(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	sTime := getTime(timeMap["oneDayAgo"], helper.Format_DateTime)
	eTime := getTime(timeMap["now"], helper.Format_DateTime)
	for subTarget, logic := range eventLogicMap {
		products, err := srv.RiskNotificationMgmt.GetProductEarning(sTime, eTime, subTarget, "notification/get_product")
		if err != nil {
			tracer.Error("notification", "productEarningNotification", err.Error())
			return
		}

		for _, product := range products {
			if !checkLogic(product.Earning, 0, logic) {
				message := getMessage(product, task, sTime, eTime, task.EventCode)
				glob.MessageToSlack(winlost, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
				insertMessageToDB(task.ID, message, product.ChannelCode+"/"+product.ProductCode)
			}
		}
	}
}

func productEarningPercentNotification(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	sTime := getTime(timeMap["oneDayAgo"], helper.Format_DateTime)
	eTime := getTime(timeMap["now"], helper.Format_DateTime)

	for subTarget, logic := range eventLogicMap {
		products, err := srv.RiskNotificationMgmt.GetProductEarning(sTime, eTime, subTarget, "notification/get_product")
		if err != nil {
			tracer.Error("notification", "productEarningPercentNotification", err.Error())
			return
		}

		for _, product := range products {
			if !checkLogic(product.Earning, product.Total, logic) {
				message := getMessage(product, task, sTime, eTime, task.EventCode)
				glob.MessageToSlack(winlost, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
				insertMessageToDB(task.ID, message, product.ChannelCode+"/"+product.ProductCode)
			}
		}
	}
}

func withdrawHigherEvent(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	sTime := getTime(timeMap["oneDayAgo"], helper.Format_DateTime)
	eTime := getTime(timeMap["now"], helper.Format_DateTime)

	compareSTime, compareETime := getPastRangeTime(timeMap["eightDayAgo"])
	oneDayAmount := getCompareValue(sTime, eTime, "get_withdraw")
	compareAmount := getCompareValue(compareSTime, compareETime, "get_withdraw")
	compareAverage := getAverage(compareAmount, sevenDay)
	for _, logic := range eventLogicMap {
		if !checkLogic(oneDayAmount, compareAverage, logic) {
			value := setValueItems(oneDayAmount, compareAverage)
			message := getMessage(value, task, sTime, eTime, task.EventCode)
			glob.MessageToSlack(withdraw, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
			insertMessageToDB(task.ID, message, "")
		}
	}
}

func depositEvent(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	sTime := getTime(timeMap["oneDayAgo"], helper.Format_DateTime)
	eTime := getTime(timeMap["now"], helper.Format_DateTime)

	compareSTime, compareETime := getPastRangeTime(timeMap["eightDayAgo"])
	oneDayAmount := getCompareValue(sTime, eTime, "get_deposit")
	compareAmount := getCompareValue(compareSTime, compareETime, "get_deposit")
	compareAverage := getAverage(compareAmount, sevenDay)
	for _, logic := range eventLogicMap {
		if !checkLogic(oneDayAmount, compareAverage, logic) {
			value := setValueItems(oneDayAmount, compareAverage)
			message := getMessage(value, task, sTime, eTime, task.EventCode)
			glob.MessageToSlack(deposit, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
			insertMessageToDB(task.ID, message, "")
		}
	}
}

func registerMember(task *mgmt.RiskNotificationEvent) {
	eventLogicMap := getLogicMap(strconv.Itoa(task.ID))

	sTime := getTime(timeMap["oneDayAgo"], helper.Format_DateTime)
	eTime := getTime(timeMap["now"], helper.Format_DateTime)

	compareSTime, compareETime := getPastRangeTime(timeMap["eightDayAgo"])
	oneDayAmount := getCompareValue(sTime, eTime, "get_register_member")
	compareAmount := getCompareValue(compareSTime, compareETime, "get_register_member")
	compareAverage := getAverage(compareAmount, sevenDay)
	for _, logic := range eventLogicMap {
		if !checkLogic(oneDayAmount, compareAverage, logic) {
			value := setValueItems(oneDayAmount, compareAverage)
			message := getMessage(value, task, sTime, eTime, task.EventCode)
			glob.MessageToSlack(member, message, task.WorkSpace, task.NotifyChannel, task.PushStatus)
			insertMessageToDB(task.ID, message, "")
		}
	}
}

type valueItem struct {
	average float64
	value   float64
}

func convertTimeZone(sTime, eTime string) (string, string) {
	start, _ := time.Parse(helper.Format_DateTime, sTime)
	end, _ := time.Parse(helper.Format_DateTime, eTime)
	return start.Format(helper.Format_DateTime), end.Format(helper.Format_DateTime)
}

func setValueItems(value, average float64) valueItem {
	return valueItem{
		average: average,
		value:   value,
	}
}

func getPastRangeTime(t time.Duration) (string, string) {
	start := getTime(t, helper.Format_Date)
	end := getTime(timeMap["oneDayAgo"], helper.Format_Date)
	return start, end
}

func getTime(t time.Duration, layout string) string {
	return time.Now().UTC().Add(-4 * time.Hour).Add(t * time.Hour).Format(layout)
}

func getAverage(value float64, day int) float64 {
	return glob.FloatRound(value / float64(day))
}

func getCompareValue(sTime, eTime, sql string) float64 {
	conds := &mgmt.RiskNotificationConds{}
	startTime, endTime := timetool.ConvertDateTimeRngByTimezone(sTime, eTime, timetool.TIME_ZONE_THE_EAST)
	conds.StartTime.ParsePtr(startTime)
	conds.EndTime.ParsePtr(endTime)

	amountItems, err := srv.RiskNotificationMgmt.QueryAmountItems(nil, conds, sql)
	if err != nil {
	}

	var totalAmount float64
	for _, item := range amountItems {
		totalAmount = glob.FloatRound(item.Value + totalAmount)
	}

	return totalAmount

}

func checkLogic(leftValue, rightValue float64, eventLogics []*mgmt.RiskNotificationLogic) bool {
	checkTime := 0
	for _, event := range eventLogics {
		switch event.ThresholdType {
		case "percent":
			if checkThershold(leftValue, float64(rightValue*(float64(event.Threshold)/100)), event.Operator) {
				checkTime++
			}
		case "amount", "hour":
			if checkThershold(leftValue, float64(event.Threshold), event.Operator) {
				checkTime++
			}
		default:
			tracer.Error("Notification", "checkLogic", errors.New("no condition_code"))
		}
	}

	if checkTime == len(eventLogics) {
		return false
	}

	return true
}

func checkThershold(value, threshold float64, operator string) bool {
	switch operator {
	case "gt":
		if value > threshold {
			return true
		}
	case "lt":
		if value < threshold {
			return true
		}
	case "gte":
		if value >= threshold {
			return true
		}
	default:
		tracer.Error("Notification", "checkThershold", errors.New("no operator"))
	}
	return false
}

func getMessage(item interface{}, event *mgmt.RiskNotificationEvent, sTime, eTime, eventCode string) string {
	start, end := convertTimeZone(sTime, eTime)
	switch eventCode {
	case glob.PlatformEarning:
		platform := item.(*mgmt.ProductAnalysisItem)
		message := fmt.Sprintf(platformMessage, event.Name, start, end, platform.Total, platform.Earning)
		return message
	case glob.ProductKeepLoss:
		product := item.(*mgmt.ProductAnalysisItem)
		message := fmt.Sprintf(productKeepLossMessage, event.Name, product.ChannelName, product.ProductName, start, end, event.Threshold)
		return message
	case glob.ProductLossPercent:
		product := item.(*mgmt.ProductAnalysisItem)
		message := fmt.Sprintf(productProportionMessage, event.Name, product.ChannelName, product.ProductName, start, end, product.Total, product.Earning)
		return message
	case glob.ProductLossAmount:
		product := item.(*mgmt.ProductAnalysisItem)
		message := fmt.Sprintf(productMessage, event.Name, product.ChannelName, product.ProductName, start, end, product.Earning)
		return message
	case glob.DepositAmountLower, glob.DepositAmountHigher, glob.WithdrawAmount, glob.MemberRegister:
		value := item.(valueItem)
		message := fmt.Sprintf(sampleMessage, event.Name, start, end, thresholdName[eventCode], value.value,
			setSymbolMessage(event.Operator), value.average)
		return message
	}
	return ""
}

func setSymbolMessage(symbol string) string {
	switch symbol {
	case "lt":
		return "小于"
	case "gt":
		return "大于"
	}
	return ""
}

func getLogicMap(event string) map[string][]*mgmt.RiskNotificationLogic {
	eventLogicMap := make(map[string][]*mgmt.RiskNotificationLogic)

	eventLogics, err := srv.RiskNotificationMgmt.QueryEventItems(nil, "get_event", event)
	if err != nil {
		tracer.Error("notification", "getLogicMap", err.Error())
		return eventLogicMap
	}

	for _, logic := range eventLogics {
		tmpLogic := &mgmt.RiskNotificationLogic{
			EventID:    logic.EventID,
			PushStatus: logic.PushStatus,
		}
		tmpLogic.Operator = logic.Operator
		tmpLogic.ThresholdType = logic.ThresholdType
		tmpLogic.ThresholdName = logic.ThresholdName
		tmpLogic.Threshold = logic.Threshold
		eventLogicMap[logic.SubTarget] = append(eventLogicMap[logic.SubTarget], tmpLogic)
	}
	return eventLogicMap
}

func insertMessageToDB(id int, message, subTarget string) {
	dc := glob.DBContext.Clone(cc)
	cmd := dc.NewCommand()

	cmd.Value("event_id", id, true)
	cmd.Value("sub_target", subTarget, true)
	cmd.Value("message", message, true)

	cmd.RawQuery(db.LoadSQL("riskcron/notification/insert_notification_message"))

	if _, err := cmd.Exec(); err != nil {
		tracer.Errorf("insertProductMessageToDB", "DB Error : %v", err.Error())
		return
	}
}
