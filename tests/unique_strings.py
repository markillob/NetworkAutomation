#coding=utf-8
string_list=["white","black", "white","marco","orange","red","white"]
def equal_strings(strings_list):
    string_dictionary={}
    string_list2=string_list
    string_list3=[]
    for i in range(len(string_list)):
       string_dictionary[string_list[i]]="1"
    for i in range(len(string_dictionary)):
        string_list3.append(string_list2[i+1])
    print(string_list3)
    return string_dictionary

print (equal_strings(string_list))

