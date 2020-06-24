from selenium import webdriver

driver = webdriver.PhantomJS(executable_path='./phantomjs/bin/phantomjs')
# 以PChome購物搜尋 macbook 為例
driver.get('https://tw.beanfun.com/MapleStory/main')
pageSource = driver.page_source
print(pageSource)
driver.close()