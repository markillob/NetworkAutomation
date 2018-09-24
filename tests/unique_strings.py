#coding=utf-8
string_list=["brown","white","white","black", "brown","white","marco","orange","red","white"]
def equal_strings(strings_list):
    string_dictionary={}
    string_list2=string_list
    string_list3=[]
    for i in range(len(string_list)):
       string_dictionary[string_list[i]]="1"
    for i in range(len(string_dictionary)):
        string_list3.append(string_list2[i+1])
    string_list4=[]
    for i,a in string_dictionary.items():
        string_list4.append(i)    
    return string_list4

print (equal_strings(string_list))

