import requests
import re
import MySQLdb
from bs4 import BeautifulSoup
import rate_tool.rate_tool as tool

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
countRate = 0

realtimeRateBuy, realtimeRateSell = tool.GetRateResult(RealtimeRates)

blank = "\t | \t"
print("幣別\t"+blank+"現金匯率"+blank+"\t即期匯率")
print("\t"+blank+"買入"+blank+"賣出"+blank+"買入"+blank+"賣出")
print("--------------------------------------------------------------------------------")
for i in range(len(countrys)):
    print(countrys[i] + blank + cashRateBuy[i] + blank +
          cashRateSell[i] + blank + realtimeRateBuy[i] + blank +
          realtimeRateSell[i])
