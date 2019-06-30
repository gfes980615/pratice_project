import my_tool.string_split as tool

import MySQLdb

chinese = []
word = []

chinese,word = tool.split_chinese('tense (a.)緊張的,拉緊的(vt.)(vi.)(使)緊張,(使)拉緊時態')
print(word)
print(chinese)



# db = MySQLdb.connect(host="localhost",
#                      user="root", passwd="", db="english_words", charset="utf8")
# cursor = db.cursor()
# level = "first"
# table_name = level + '_words'
# cursor.execute("select COUNT(*) as count from INFORMATION_SCHEMA.TABLES where TABLE_NAME='"+table_name+"';")
# results = cursor.fetchall()
# print(results[0][0])

# db.close()
