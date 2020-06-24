import requests
import re
from bs4 import BeautifulSoup

url = "https://tw.beanfun.com/maplestory/script/m-mainmenu.js"
# url = "https://tw.beanfun.com/MapleStory/main?section=mBulletin&cate=all"
session = requests.Session()
headers = {
"Accept-Ranges": "bytes",
"Content-Length": "5183",
"Content-Security-Policy": "frame-ancestors 'self' http://beanfun.com https://beanfun.com http://*.beanfun.com https://*.beanfun.com http://*.gungho-gamania.com https://*.gungho-gamania.com",
"Content-Type": "application/javascript",
"Date": "Tue, 23 Jun 2020 03:25:40 GMT",
"ETag": "1d648651d086ebf",
"Last-Modified": "Mon, 22 Jun 2020 07:16:57 GMT",
"Server": "Microsoft-IIS/7.5",
"X-Frame-Options": "SAMEORIGIN",
"X-Powered-By": "ASP.NET",
"X-UA-Compatible": "IE=Edge",
}
req = session.get(url,headers=headers)
req.encoding = 'utf8'

fp = open("page.html","w",encoding="utf-8")
fp.write(req.text)
fp.close()


