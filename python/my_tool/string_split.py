import re
import sys

def sub_split(string, word):
    return string.split(word)


def remove_empty(res):
    for index in res:
        if index == '':
            res.remove(index)
    return res


def split_chinese(resource):
    try:
        string = resource.replace(" ", "")
        string = string.replace("'", "")
        print(string)
        chineseRes = re.split('[a-z(.)]+', string)
        chineseRes = remove_empty(chineseRes)
        wordRes = []
        tmpArr = []
        for index in range(len(chineseRes)):
            if index == 0:
                tmpArr = sub_split(string, chineseRes[index])
                wordRes.append(tmpArr[0])
                continue
            if index == len(chineseRes)-1:
                string = tmpArr[1]
                tmpArr = sub_split(string, chineseRes[index])
                wordRes.append(tmpArr[0])
                if len(tmpArr) > 1:
                    wordRes.append(tmpArr[1])
                break
            string = tmpArr[1]
            tmpArr = sub_split(string, chineseRes[index])
            wordRes.append(tmpArr[0])
        wordRes = remove_empty(wordRes)

        # 額外處理
        excepthandle = wordRes[0].split('(', 1)
        wordRes[0] = excepthandle[0]
        if len(excepthandle) > 1:
            wordRes.append('('+excepthandle[1])
    except:
        type, message, traceback = sys.exc_info()
        while traceback:
            print('..........')
            print(type)
            print(message)
            print('function or module？', traceback.tb_frame.f_code.co_name)
            print('file？', traceback.tb_frame.f_code.co_filename)
            traceback = traceback.tb_next

    return wordRes, chineseRes


def merge_list(res_list):
    res = ""
    for value in res_list:
        res = res + value + ' '
    return res
