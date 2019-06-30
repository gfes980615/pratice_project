import requests
from bs4 import BeautifulSoup
url="https://www.8591.com.tw/mallList-list-859.html?searchGame=859&searchServer=944?&group=1&priceSort=0&ratios=0&searchGame=859&searchServer=944&firstRow="
page=0

req = requests.get(url)

soup = BeautifulSoup(req.text,"html.parser")

soup_page = soup.find("span",class_="R")
f = open("getAllHtml.txt","w",encoding='UTF-8')
f.write(soup_page.text)
f.close

data = open("page1.txt","w",encoding='UTF-8')
soup_span = soup.find_all("span",class_="ml-item-title")
for n in soup_span:
    data.write(n.text + "\n")
data.close