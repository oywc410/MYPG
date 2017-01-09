from sklearn import svm

# 数据 (特征向量)
x = [[2, 0], [1, 1], [2, 3]]
# 各个数据的类型 (标记)
y = [0, 0, 1]
clf = svm.SVC(kernel = 'linear')
clf.fit(x, y) # 建立模型 计算超平面

print(clf)

#查看支持向量 [[1, 1][2, 3]]
print(clf.support_vectors_)

#支持向量的位置[1, 3]
print(clf.support_)

#两个类的 支持向量的个数 
print(clf.n_support_)

#判断类型
print(clf.predict([[2, 0], [3, 0]]))