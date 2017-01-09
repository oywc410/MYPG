import numpy as np
import pylab as pl
from sklearn import svm

#创建40个点
np.random.seed(0) 
#randn产生随机矩阵 20 X 2
X = np.r_[np.random.randn(20, 2) - [2, 2], np.random.randn(20, 2) + [2, 2]]
# 各个数据的类型 (标记)
Y = [0] * 20 + [1] * 20

print(X)
print(Y)

#建立模型
clf = svm.SVC(kernel='linear')
clf.fit(X, Y)

#集合
w = clf.coef_[0]
#平面斜率
a = -w[0] / w[1]

#产生-5 到 5的连续值
xx = np.linspace(-5, 5)
print(xx)

#超平面
yy = a * xx - (clf.intercept_[0]) / w[1]

#第一个支持向量
b = clf.support_vectors_[0]
yy_down = a * xx + (b[1] - a * b[0])
#取到另一个支持向量
b = clf.support_vectors_[-1]
yy_up = a * xx + (b[1] - a * b[0])

print("w:", w)
print("a:", a)

print("support_vectors:", clf.support_vectors_)
print("clf.coef_:", clf.coef_)

pl.plot(xx, yy, 'k-')
pl.plot(xx, yy_down, 'k--')
pl.plot(xx, yy_up, 'k--')

pl.scatter(clf.support_vectors_[:, 0], clf.support_vectors_[:, -1], s=80, facecolors='none')
pl.scatter(X[:, 0], X[:, 1], c=Y, cmap=pl.cm.Paired)
pl.axis('tight')
p1.show();