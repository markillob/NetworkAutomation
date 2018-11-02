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

def FirstFactorial( int ) :
    contTemp = 1
    output = 1
    if int  > 18 or int == 0 :
        print ("Greater than 18 or value = 0"  , int )
        return int
    else:
        for i in range( int ):
            contTemp = contTemp * output
            output+=1
        return contTemp
print (FirstFactorial(0))

print (upperPhrase('arriba las chivas'))