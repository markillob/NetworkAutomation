#/home/mbarrera/github.com/NetworkAutomation/bin/python3
#coding=utf-8
import sys
word1 = "oso"
word2 = "hola"
print("First String:", word1)
print ("Second String:", word2)
#access a String word[0]
#check if strings are equal
def equal_strings(word1,word2):
    if len(word1) != len(word2):
        print ("Strins are not equal")
    string_size = len(word1)
    for i in range(string_size):
        if word1[i] != word2[i]:
           print ("Words are not equal")
           sys.exit(1)
    print ("Words are equal")
    return

def same_letters(word1,word2):
    string_size = len(word1)
    letter={}
    count=1
    for i in range(string_size):
        if word1[i] in letter:
            count+=1
            letter[word1[i]]=count
        else:
           letter[word1[i]]= 1
    print (letter)
    letter2={}
    count2=1
    string_size2 = len(word2)
    for i in range(string_size2):
        if word2[i] in letter2:
            count2+=1
            letter2[word2[i]]=count2
        else:
            letter2[word2[i]]=1
    print (letter2)
    if letter == letter2:
        return "True"
    else:
        return "False"
print (same_letters(word1, word2))

#equal_strings (word1,word2)
