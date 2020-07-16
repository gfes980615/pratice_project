package athena

import (
	"errors"

	"gitlab.paradise-soft.com.tw/backend/shark/platform/common/glob/bisystemsettings"
	"gitlab.paradise-soft.com.tw/backend/yaitoo/mvc"
	utilsMVC "gitlab.paradise-soft.com.tw/glob/utils/mvc"
	"gitlab.paradise-soft.com.tw/glob/utils/network"
)

var (
	biApiWriter mvc.ViewWriter
)

// convert to athena api
func ApisToAthenaV2(ctx *mvc.Context, api string) {
	params := utilsMVC.GetMapFromContext(ctx)
	client := network.NewClient(bisystemsettings.BIServerHost + api)

	var buf []byte
	var err error
	switch ctx.Request.Route.Method {
	case "GET":
		buf, err = client.Get(params)
	case "PUT":
		buf, err = client.Put(params)
	case "POST":
		buf, err = client.Post(params)
	case "DELETE":
		buf, err = client.Del(params)
	default:
		buf = nil
		err = errors.New("illegal method")
	}

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Response.Body = string(buf)
}
