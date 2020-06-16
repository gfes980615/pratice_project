import requests
import re
import MySQLdb
from bs4 import BeautifulSoup

url = "https://udn.com/news/cate/2/6638"
req = requests.get(url)
req.encoding = 'utf8'

soup = BeautifulSoup(req.text, "html.parser")
titles = soup.select("div.story-list__news")
for title in titles:
    tmp = title.select("div.story-list__text h3 a")[0].text
    print(tmp)
