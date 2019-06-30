# coding=utf-8
import requests
import re
import MySQLdb
from bs4 import BeautifulSoup
import my_tool.string_split as tool
import configparser
 
config = configparser.ConfigParser()

config.read('Config.ini')
db_name = config.get('mysql', 'db_name') 
db_host = config.get('mysql', 'db_host') 
db_pass = config.get('mysql', 'db_pass') 
db_user = config.get('mysql', 'db_user') 

comma = ","
jump = "\n"
single = "'"

video_id = input("\n輸入影片id:")
print("\n多益單字: vol\n英檢初級單字: first\n英檢中級單字: mid\n英檢高級: high")
level = input("\n輸入想要取得的單字難度:")

url = "https://tw.voicetube.com/videos/print_words/"+video_id+"/"+level
req = requests.get(url)
req.encoding = 'utf8'

soup = BeautifulSoup(req.text, "html.parser")
words = soup.select("ol.wordlist.controls")
wordlist = []
for word in words:
    wordlist = word.select("li")

db = MySQLdb.connect(host=db_host,
                     user=db_user, passwd=db_pass, db=db_name, charset="utf8")
cursor = db.cursor()
table_name = level + '_words'
cursor.execute("select COUNT(*) as count from INFORMATION_SCHEMA.TABLES where TABLE_NAME='"+table_name+"';")
results = cursor.fetchall()
if results[0][0] == 0:
    cursor.execute("CREATE TABLE " + table_name + " (\
                    `id` int(11) NOT NULL AUTO_INCREMENT,\
                    `word` char(100) DEFAULT NULL,\
                    `translation` char(100) DEFAULT NULL,\
                    `parts` char(100) DEFAULT NULL,\
                    PRIMARY KEY (`id`)\
                    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8\
                ")

for value in wordlist:
    words, chinese = tool.split_chinese(value.text)
    word = words[0]
    translation = tool.merge_list(chinese)
    words.remove(word)
    parts = tool.merge_list(words)
    sql = "INSERT INTO "+ table_name +" (word,translation,parts) VALUES (" + \
        single+word+single+comma+single+translation + \
        single+comma+single+parts+single+");"
    cursor.execute(sql)
    db.commit()

input('新增完畢 按任意鍵結束...')
db.close()
