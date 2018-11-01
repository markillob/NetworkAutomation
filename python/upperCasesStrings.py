#coding: utf-8

def upperPhrase(str):
    countString = 0
    index=len(str)
    tempString = ''
    for i in range(index) :
        if i == countString :
            if str[countString] == ' ' :
                tempString = tempString + str[countString]
                countString += 1
                tempString = tempString + str[countString].upper()
                countString += 1
                continue
            else :
                if countString == 0 :
                    tempString = tempString + str[countString].upper()
                    countString += 1
                    continue
                tempString = tempString + str[countString]
                countString += 1
        else:
            continue
    return tempString
print (upperPhrase('arriba las chivas'))