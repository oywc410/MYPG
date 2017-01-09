file = open('test.txt', 'r')
content = file.read()
print(content)

print("------------------------------")

my_list = []

file = open('test.txt', 'r')
# 只读一行
content = file.readline()
#一次多去多行
contents = file.readlines()
print(content)
print(contents)
