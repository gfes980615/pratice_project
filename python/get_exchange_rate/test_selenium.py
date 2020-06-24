from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import time
from bs4 import BeautifulSoup

options = Options()
webdriver_path = '.\chromedriver.exe'
options.binary_location = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
driver = webdriver.Chrome(executable_path=webdriver_path, options=options)
driver.get("https://tw.beanfun.com/maplestory/main") #前往這個網址


soup = BeautifulSoup(driver.page_source, "html.parser")
driver.close()
annSet = soup.select("a.mBulletin-items-link")
for se in annSet:
    date = se.select("div.mBulletin-items-date")
    print(date[0].text)
    category = se.select("div.mBulletin-items-cate")
    print(category[0].text)
    content = se.select("div.mBulletin-items-title")
    print(content[0].text)