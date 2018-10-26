#coding: utf-8

def upperPhrase(str):
    countString = 0
    index=len(str)
    tempString = ''
    for i in range(index):
        if i == 0:
            tempString = tempString + str[i].upper()
        elif str[i] == ' ':
            tempString = tempString + str[i].upper()
            continue
        else:
            tempString = tempString + str[i]
        countString += 1
    return tempString
print (upperPhrase('arriba las chivas'))