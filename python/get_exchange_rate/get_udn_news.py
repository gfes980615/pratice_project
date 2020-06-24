import requests
import re
from bs4 import BeautifulSoup

url = "https://udn.com/news/cate/2/6638"
req = requests.get(url)
req.encoding = 'utf8'

soup = BeautifulSoup(req.text, "html.parser")
titles = soup.select("div.story-list__news")
for title in titles:
    tmpTitle = title.select("div.story-list__text h3 a")
    if len(tmpTitle) == 0:
        tmpTitle = title.select("div.story-list__text h2 a")

    tmp = tmpTitle[0].text
    print(tmp)
