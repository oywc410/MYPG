import copy

a = [1,2,3]
b = a

# 指针也被复制(a改变同时b也被改变)
print(id(a))
print(id(b))

b[1] = 7

print(a)
print(b)

c = copy.copy(a) #有二维时 二维还是属于指针

print(id(c))
c[1] = 10

print(a)
print(c)

e = copy.deepcopy(a) #完全copy (无指针拷贝)

