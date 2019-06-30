import requests
from bs4 import BeautifulSoup

url = "https://www.8591.com.tw/mallList-list-859.html?searchGame=859&searchServer=944?&group=1&priceSort=0&ratios=0&searchGame=859&searchServer=944&firstRow="
req = requests.get(url)
soup = BeautifulSoup(req.text, "html.parser")
u = 0
raw = 21
dir_name = "8591/page"
count = soup.find("span", class_="R")
count = int(count.text)
page = count / raw
file_page = page / 10
m = 1
for j in range(int(file_page)):
    file_name = "file" + str(j + 1) + ".txt"
    data = open(dir_name + "/" + file_name, "w", encoding='UTF-8')
    for i in range(10):
        req = requests.get(url + str(u))
        u = u + raw
        soup = BeautifulSoup(req.text, "html.parser")
        soup_span = soup.find_all("span", class_="ml-item-title")
        p = "page" + str(m)
        data.write(p + "\n")
        item = 1
        for n in soup_span:
            data.write(str(item) + ": " + n.text + "\n")
            item = item + 1
        data.write(
            "-------------------------------------------------------------\n")
        count = count - raw
        m = m + 1
