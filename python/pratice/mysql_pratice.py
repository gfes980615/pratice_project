import MySQLdb

db = MySQLdb.connect(host="localhost",
    user="root", passwd="", db="test2")
cursor = db.cursor()

cursor.execute("SELECT * FROM plan")

# 取回所有查詢結果
results = cursor.fetchall()

# 輸出結果
for record in results:
  col1 = record[0]
  col2 = record[1]
  print(col1,col2)

# 關閉連線
db.close()