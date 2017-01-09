#!/usr/bin/python3

var1 = 100
if var1:
   print ("1 - if 表达式条件为 true")
   print (var1)

var2 = 0
if var2:
   print ("2 - if 表达式条件为 true")
   print (var2)

age = int(input("请输入你家狗狗的年龄: "))
print("")
if age < 0:
	print("你是在逗我吧!")
elif age == 1:
	print("相当于 14 岁的人。")
elif age == 2:
	print("相当于 22 岁的人。")
elif age > 2:
	human = 22 + (age -2)*5
	print("对应人类年龄: ", human)

### 退出提示
input('点击 enter 键退出')