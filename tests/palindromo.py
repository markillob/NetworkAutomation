#/home/mbarrera/github.com/NetworkAutomation/bin/python3
#coding=utf-8
import sys
word1 = "oso"
word2 = "osi"
print("First String:", word1)
print ("Second String:", word2)
#access a String word[0]
if len(word1) != len(word2):
    print ("Strins are not equal")
string_size = len(word1)
for i in range(string_size):
    if word1[i] != word2[i]:
       print ("Words are not equal")
       sys.exit(1)
print ("Words are equal")
