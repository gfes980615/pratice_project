import os

print('\n\n\t\t新增conf檔參數工具\n\n')


title = input('Enter titile=')
parameter = []
value = []
count = 0
try:
    while True:
        parameter.append(input('Enter parameter:'))
        value.append(input('Enter value:'))
except EOFError:
    pass

env = ['dev']
brand = ['3h','lv','ls','cdd','hy','bh','c7','c8','co','dm','qm','sc','tz','xpj']
apath=os.path.abspath('.')

for e in env:
    for b in brand:
        path = '\\'+e+'\\'+b
        if not os.path.isdir(apath+path):
            os.makedirs(apath+path)
        with open(apath+path+'\\'+'app.conf','a') as f :
            f.write('\n\n')
            f.write('[' + title + ']' + '\n')
            for index in range(len(parameter)):
                f.write(parameter[index] + '=' + value[index] + '\n')

input('新增完畢 按任意鍵結束程式...')