import numpy as np

array = np.array([[1,2,3], [2,3,4]])

print(array)
print("number of dim:", array.ndim) #数组维度
print("shape:", array.shape) #数组形状
print("size:", array.size)

print("============创建============")
a = np.array([2, 23, 4], dtype=np.int64)
print(a)
print(a.dtype)

a = np.array([[2, 23, 4], [4, 12, 43]])
print(a)
print(a.dtype)
print(a.shape)

#生成为0的矩阵
a = np.zeros((3, 4)) #noes 所有为1 emprty 空
print(a)

#有序
a = np.arange(10, 22, 2).reshape((2, 3)) #reshape 重新定义结构
print(a)

a = np.linspace(1, 10, 20)
print(a)

print("===========基础计算============")
a=np.array([10,20,30,40])   # array([10, 20, 30, 40])
b=np.arange(4)              # array([0, 1, 2, 3])

c=a-b  # array([10, 19, 28, 37])
c=a+b   # array([10, 21, 32, 43])
c=a*b   # array([  0,  20,  60, 120])

c=b**2  # array([0, 1, 4, 9]) 每个元素进行求平方
c=10*np.sin(a)
print(b<3) # 对每个元素进行逻辑判断

#多维
a=np.array([[1,1],[0,1]])
b=np.arange(4).reshape((2,2))

print(a)
# array([[1, 1],
#       [0, 1]])

print(b)
# array([[0, 1],
#       [2, 3]])

dot = np.dot(a,b) #标准的矩阵乘法运算，即对应行乘对应列得到相应元素
# array([[2, 4],
#       [2, 3]])

a=np.random.random((2,4))
print(a)
# array([[ 0.94692159,  0.20821798,  0.35339414,  0.2805278 ],
#       [ 0.04836775,  0.04023552,  0.44091941,  0.21665268]])


np.sum(a)   # 4.4043622002745959
np.min(a)   # 0.23651223533671784
np.max(a)   # 0.90438450240606416

# 当axis的值为0的时候，将会以列作为查找单元， 当axis的值为1的时候，将会以行作为查找单元。
print("a =",a)
# a = [[ 0.23651224  0.41900661  0.84869417  0.46456022]
# [ 0.60771087  0.9043845   0.36603285  0.55746074]]

print("sum =",np.sum(a,axis=1))
# sum = [ 1.96877324  2.43558896]

print("min =",np.min(a,axis=0))
# min = [ 0.23651224  0.41900661  0.36603285  0.46456022]

print("max =",np.max(a,axis=1))
# max = [ 0.84869417  0.9043845 ]

# argmin() 和 argmax() 两个函数分别对应着求矩阵中最小元素和最大元素的索引
A = np.arange(2, 14).reshape((3, 4))

# array([[ 2, 3, 4, 5]
#        [ 600, 7, 8, 9]
#        [10,11,12,13]])

print(np.argmin(A))  # 0
print(np.argmax(A))  # 11

#计算统计中的均值
print(np.mean(A))        # 7.5
print(np.average(A))     # 7.5

#求解中位数的函数
print(np.median(A))       # 7.5

#累加函数 
print(np.cumsum(A))

#累差运算
print(np.diff(A))

#找出非0数
print(np.nonzero(A))

#排序
print(np.sort(A))

#矩阵反置
print(np.transpose(A))

#所有大于10的数变为10 所有小于5的数变为5
print(np.clip(A, 5, 10))

#所有操作可以添加参数axis=0是对矩阵列进行计算 =1 时对列进行计算

print("============索引===================")

A = np.arange(3,15)

# array([3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14])
         
print(A[3])    # 6

# 让我们将矩阵转换为二维的，此时进行同样的操作
A = np.arange(3,15).reshape((3,4))
print(A[2])         
# [11 12 13 14]
print(A[1][1])      # 8
print(A[1, 1])      # 8

print(A[1, 1:3])    # [8 9]

for row in A:
    print(row)
    """    
    [ 3,  4,  5, 6]
    [ 7,  8,  9, 10]
    [11, 12, 13, 14]
    """

#A.T 矩阵的反置
for column in A.T:
    print(column)
    """  
    [ 3,  7,  11]
    [ 4,  8,  12]
    [ 5,  9,  13]
    [ 6, 10,  14]
    """

#flatten是一个展开性质的函数，将多维的矩阵进行展开成1行的数列
print(A.flatten())   
# array([3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14])

for item in A.flat:
    print(item)
    # 3
    # 4
    # 14

print("===========array合并================")
A = np.array([1,1,1])
B = np.array([2,2,2])
          
print(np.vstack((A,B)))    # vertical stack
"""
[[1,1,1]
[2,2,2]]
"""

D = np.hstack((A,B))       # horizontal stack

print(D)
# [1,1,1,2,2,2]

#添加维度
print(A[np.newaxis,:])
# [[1 1 1]]

print(A[np.newaxis,:].shape)
# (1,3)

print(A[:,np.newaxis])
"""
[[1]
[1]
[1]]
"""

print(A[:,np.newaxis].shape)
# (3,1)

#当你的合并操作需要针对多个矩阵或序列时，借助concatenate函数可能会让你使用起来比前述的函数更加方便：
C = np.concatenate((A,B,B,A),axis=0)

print(C)
"""
array([[1],
[1],
[1],
[2],
[2],
[2],
[2],
[2],
[2],
[1],
[1],
[1]])
"""
D = np.concatenate((A,B,B,A),axis=1)
print(D)
"""
array([[1, 2, 2, 1],
    [1, 2, 2, 1],
    [1, 2, 2, 1]])
"""

#axis参数很好的控制了矩阵的纵向或是横向打印，相比较vstack和hstack函数显得更加方便。

print("=============array分割=====================")

A = np.arange(12).reshape((3, 4))
print(A)
"""
array([[ 0,  1,  2,  3],
    [ 4,  5,  6,  7],
    [ 8,  9, 10, 11]])
"""

print(np.split(A, 2, axis=1))
"""
[array([[0, 1],
        [4, 5],
        [8, 9]]), array([[ 2,  3],
        [ 6,  7],
        [10, 11]])]
"""

print(np.split(A, 3, axis=0))

# [array([[0, 1, 2, 3]]), array([[4, 5, 6, 7]]), array([[ 8,  9, 10, 11]])]

#错误的分割 范例的Array只有4列，只能等量对分，因此输入以上程序代码后Python就会报错。
print(np.split(A, 3, axis=1))

#不等量的分割
print(np.array_split(A, 3, axis=1))
"""
[array([[0, 1],
        [4, 5],
        [8, 9]]), array([[ 2],
        [ 6],
        [10]]), array([[ 3],
        [ 7],
        [11]])]
"""

#其他的分割方式
print(np.vsplit(A, 3)) #等于 print(np.split(A, 3, axis=0))

# [array([[0, 1, 2, 3]]), array([[4, 5, 6, 7]]), array([[ 8,  9, 10, 11]])]


print(np.hsplit(A, 2)) #等于 print(np.split(A, 2, axis=1))
"""
[array([[0, 1],
       [4, 5],
       [8, 9]]), array([[ 2,  3],
        [ 6,  7],
        [10, 11]])]
"""

print("=======Numpy copy & deep copy==========")

a = np.arange(4)
# array([0, 1, 2, 3])

b = a
c = a
d = b

#改变a的第一个值，b、c、d的第一个值也会同时改变。
a[0] = 11
print(a)
# array([11,  1,  2,  3])

#确认b、c、d是否与a相同。
b is a  # True
c is a  # True
d is a  # True

#同样更改d的值，a、b、c也会改变。
d[1:3] = [22, 33]   # array([11, 22, 33,  3])
print(a)            # array([11, 22, 33,  3])
print(b)            # array([11, 22, 33,  3])
print(c)            # array([11, 22, 33,  3])

#copy()的赋值方式没有关联性
b = a.copy()    # deep copy
print(b)        # array([11, 22, 33,  3])
a[3] = 44
print(a)        # array([11, 22, 33, 44])
print(b)        # array([11, 22, 33,  3])



