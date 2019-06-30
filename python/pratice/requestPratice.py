# coding=utf-8
import requests
import re
import MySQLdb
from bs4 import BeautifulSoup


comma = ","
jump = "\n"
single = "'"

url = "https://tw.voicetube.com/videos/print_words/75338/first"
req = requests.get(url)
req.encoding = 'utf8'

soup = BeautifulSoup(req.text, "html.parser")
words = soup.select("ol.wordlist.controls")
w = []
for word in words:
    w = word.select("li")

db = MySQLdb.connect(host="localhost",
                     user="root", passwd="", db="english_words", charset="utf8")
cursor = db.cursor()

for value in w:
    x = []
    tmp = 0
    x = value.text.split(" ")
    if(len(x) == 1):
        continue
    else:
        sql = "INSERT INTO `words` (word,translation) VALUES (" + \
            single+x[0]+single+comma+single+x[1]+single+");"
        cursor.execute(sql)
        db.commit()


db.close()
