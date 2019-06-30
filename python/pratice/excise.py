import csv

from my_tool.string_split import split_chinese

# with open('csvFile.csv') as f:
#     myCsv = csv.reader(f)
#     headers = next(myCsv)
#     for row in myCsv:
#         print(row[0])

# with open('./csvFile.csv') as f:
#     myCsvDic = csv.DictReader(f)
#     for row in myCsvDic:
#         print(row['Year'])

chinese, word = split_chinese('threat威脅,恐嚇,造成威脅的事物')
print(chinese)
print(word)

# with open('./searchContent.html', 'w+') as f:
#   f.write(r.text.encode("utf8").decode("cp950", "ignore"))
#   print('saved')
