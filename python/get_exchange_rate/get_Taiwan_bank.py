import requests
import re
import MySQLdb
from bs4 import BeautifulSoup
import rate_tool.rate_tool as tool

class ExchangeRate:
    country = ""           #國家
    cashRateBuy = 0         #現金買入
    cashRateSell = 0        #現金賣出
    realtimeRateBuy = 0     #即期買入
    realtimeRateSell = 0    #即期賣出
    def __init__(self, country="", cashRateBuy=0, cashRateSell=0, realtimeRateBuy=0, realtimeRateSell=0):
        self.country = country
        self.cashRateBuy = cashRateBuy
        self.cashRateSell = cashRateSell
        self.realtimeRateBuy = realtimeRateBuy
        self.realtimeRateSell = realtimeRateSell

comma = ","
jump = "\n"
single = "'"

url = "https://rate.bot.com.tw/xrt?Lang=zh-TW"
req = requests.get(url)
req.encoding = 'utf8'

soup = BeautifulSoup(req.text, "html.parser")
countryCoins = soup.select("td.currency.phone-small-font")  #取得幣別
countrys = []
countCountry = 0
for countryCoin in countryCoins:
    country = countryCoin.select("div.hidden-phone.print_show")
    countrys.append(country[0].text.strip())
    countCountry = countCountry + 1

# for country in countrys:
#     print(country)

CashRates = soup.select("td.rate-content-cash.text-right.print_hide")  #取得現金匯率
RealtimeRates = soup.select(
    "td.rate-content-sight.text-right.print_hide")  #取得即期匯率

cashRateBuy, cashRateSell = tool.GetRateResult(CashRates)

realtimeRateBuy, realtimeRateSell = tool.GetRateResult(RealtimeRates)

blank = "\t | \t"
print("幣別\t"+blank+"現金匯率"+blank+"\t即期匯率")
print("\t"+blank+"買入"+blank+"賣出"+blank+"買入"+blank+"賣出")
print("--------------------------------------------------------------------------------")
for i in range(len(countrys)):
    print(countrys[i] + blank + cashRateBuy[i] + blank +
          cashRateSell[i] + blank + realtimeRateBuy[i] + blank +
          realtimeRateSell[i])

rate_list = []
for i in range(len(countrys)):
    rate_obj = ExchangeRate(countrys[i],cashRateBuy[i],cashRateSell[i],realtimeRateBuy[i],realtimeRateSell[i])
    rate_list.append(rate_obj)

for i in range(len(rate_list)):
    print(rate_list[i].country)

