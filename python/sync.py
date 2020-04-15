import time
import asyncio
import requests


async def getPage(name, url):
	print("正在爬取%s......" % name)
	response = requests.get(url=url).text
	with open("%s.html" % name, "w", encoding="utf-8") as fp:
		fp.write(response)
	print("%s爬取完毕......" % name)


if __name__ == '__main__':
	startTime = time.time()
	urlDict = {
		"百度搜索": "https://www.baidu.com/",
		"百度翻译": "https://fanyi.baidu.com/",
		"CSDN": "https://www.csdn.net/",
		"博客园": "https://www.cnblogs.com/",
		"哔哩哔哩": "https://www.bilibili.com/",
		"码云": "https://gitee.com/",
		"拉勾网": "https://www.lagou.com/",
	}
	taskList = []
	for key, value in urlDict.items():
		request = getPage(key, value)
		task = asyncio.ensure_future(request)
		taskList.append(task)

	loop = asyncio.get_event_loop()
	loop.run_until_complete(asyncio.wait(taskList))
	print("Time consuming:", time.time() - startTime)
