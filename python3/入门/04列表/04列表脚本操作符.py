print (len([1, 2, 3]))

lists = [1, 2, 3] + [4, 5, 6]
print (lists)

print (['Hi!'] * 4) #重复

print (3 in [1, 2, 3]) #元素是否存在于列表中

for x in [1, 2, 3]: \
	print (x)
#	迭代

"""
操作方法
序号	函数
1	len(list)
列表元素个数
2	max(list)
返回列表元素最大值
3	min(list)
返回列表元素最小值
4	list(seq)
将元组转换为列表
Python包含以下方法:
序号	方法
1	list.append(obj)
在列表末尾添加新的对象
2	list.count(obj)
统计某个元素在列表中出现的次数
3	list.extend(seq)
在列表末尾一次性追加另一个序列中的多个值（用新列表扩展原来的列表）
4	list.index(obj)
从列表中找出某个值第一个匹配项的索引位置
5	list.insert(index, obj)
将对象插入列表
6	list.pop(obj=list[-1])
移除列表中的一个元素（默认最后一个元素），并且返回该元素的值
7	list.remove(obj)
移除列表中某个值的第一个匹配项
8	list.reverse()
反向列表中元素
9	list.sort([func])
对原列表进行排序
10	list.clear()
清空列表
11	list.copy()
复制列表
"""