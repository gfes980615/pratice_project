import os

print('\n\n\t\t新增conf檔參數工具\n\n')


title = input('Enter titile=')
parameter = []
value = []

try:
    while True:
        parameter.append(input('Enter parameter:'))
        value.append(input('Enter value:'))
except EOFError:
    pass

env = ['dev','prod','uat']
brand = ['lv','ls','cdd','hy']
apath=os.path.abspath('.')

for e in env:
    for b in brand:
        path = '\\'+e+'\\'+b
        if not os.path.isdir(apath+path):
            os.makedirs(apath+path)
        with open(apath+path+'\\'+'app.conf','w+') as f :
            f.write('\n')
            f.write('[' + title + ']' + '\n')
            for index in range(len(parameter)):
                f.write(parameter[index] + '=' + value[index] + '\n')

input('新增完畢 按任意鍵結束程式...')

