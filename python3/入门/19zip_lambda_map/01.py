#-----------zip

a = [1,2,3]
b = [4,5,6]

for i,j in zip(a,b): #同时按顺序处理俩个元素
	print(i/2, j*2)

print(list(zip(a, b))) #返回重新组合成的list

#-----------lambda

def func1(x, y):
	return(x+y)

func2 = lambda x,y : x+y

print(func1(2,3))
print(func2(2,3))

#-----------map
#回调
print(list(map(func1, [1,3], [2,5])))