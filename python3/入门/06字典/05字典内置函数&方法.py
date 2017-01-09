"""
len(dict)
计算字典元素个数，即键的总数。	

str(dict)
输出字典以可打印的字符串表示。

type(variable)
返回输入的变量类型，如果变量是字典就返回字典类型。

1	radiansdict.clear()
删除字典内所有元素
2	radiansdict.copy()
返回一个字典的浅复制
3	radiansdict.fromkeys()
创建一个新字典，以序列seq中元素做字典的键，val为字典所有键对应的初始值
4	radiansdict.get(key, default=None)
返回指定键的值，如果值不在字典中返回default值
5	key in dict
如果键在字典dict里返回true，否则返回false
6	radiansdict.items()
以列表返回可遍历的(键, 值) 元组数组
7	radiansdict.keys()
以列表返回一个字典所有的键
8	radiansdict.setdefault(key, default=None)
和get()类似, 但如果键不存在于字典中，将会添加键并将值设为default
9	radiansdict.update(dict2)
把字典dict2的键/值对更新到dict里
10	radiansdict.values()
以列表返回字典中的所有值
"""