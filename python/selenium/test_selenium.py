from selenium import webdriver 
from selenium.webdriver.chrome.options import Options
import time

options = Options()
webdriver_path = '.\chromedriver.exe'
options.binary_location = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"

chrome_options = webdriver.ChromeOptions()
chrome_options.add_argument('--headless')
chrome_options.add_argument('--disable-gpu')

driver = webdriver.Chrome(executable_path=webdriver_path, options=options,chrome_options=chrome_options)

def main():
    getLatLong("桃園")

def getLatLong(name):
    driver.get('https://www.google.com.tw/maps') 
    driver.find_element_by_id('searchboxinput').send_keys(name)
    driver.find_element_by_id('searchbox-searchbutton').click()
    time.sleep(3)
    print(driver.current_url)
    driver.close()

if __name__ == "__main__":
    main()