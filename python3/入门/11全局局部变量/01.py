APPLE = 100 #全局变量
b = 1

def func():
	global b #调用全局变量
	b = 10
	a = APPLE
	return a + 100

print(func())
print(b)