import os
import re
import requests

if __name__ == '__main__':
    requestHeaders = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Safari/537.36"
    }
    if not os.path.exists("./image/test"):
        os.mkdir("./image/test")
    # 首先使用for 循環遍歷第1 頁到第22 頁，將頁碼與URL 進行拼接；
    for pageIndex in range(1, 2):
        pageURL = "http://wuming3175.lofter.com/?page=%d" % pageIndex
        # 然後爬取每個頁面的所有元素，並使用正則表達式過濾img 的src 屬性同時簡化其URL；
        requestPageResponse = requests.get(
            url=pageURL, headers=requestHeaders).text
        regex = '<div class="block photo">.*?<img src="(.*?)?imageView.*?</div>'
        imagesURLList = re.findall(regex, requestPageResponse, re.S)
        for imageURL in imagesURLList:
            imageName = imageURL.split("/")[- 1].strip("?")
            # 最後通過URL 獲取圖片二進制數據並保存到本地。
            imageData = requests . get(
                url=imageURL, headers=requestHeaders).content
            with open("./image/test/%s" % imageName, "wb") as fp:
                fp.write(imageData)
        print("第%d頁爬取完畢！" % pageIndex)
