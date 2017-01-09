test = "test\ntest"

my_file = open('test.txt', 'w')
my_file.write(test)
my_file.close()

# 追加写入
my_file2 = open('test.txt', 'a')
my_file2.write(test)
my_file2.close()