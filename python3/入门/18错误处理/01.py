try:
	file = open('aaa', 'r+')
except Exception as e:
	print(e)
else:
	file.write('sss') #未发生错误则尝试写入
file.close()