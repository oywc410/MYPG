import numpy as np
import random


# 梯度下降算法
"""
x 实例
y 结果值
theta 要求的值
alpha 学习概率
m 总共的实例
numIterations 梯度算法计算多少次
"""
# m denotes the number of examples here, not the number of features
def gradientDescent(x, y, theta, alpha, m, numIterations):
    xTrans = x.transpose()
    for i in range(0, numIterations):
        hypothesis = np.dot(x, theta)
        loss = hypothesis - y
        # avg cost per example (the 2 in 2*m doesn't really matter here.
        # But to be consistent with the gradient, I include it)
        cost = np.sum(loss ** 2) / (2 * m)
        print("Iteration %d | Cost: %f" % (i, cost))
        # avg gradient per example
        gradient = np.dot(xTrans, loss) / m
        # update
        theta = theta - alpha * gradient
    return theta

"""
产生数据
numPoints 实例
bias偏好
variance 方差
"""
def genData(numPoints, bias, variance):
    x = np.zeros(shape=(numPoints, 2)) #默认值为0 的 numPoints x 2 的矩阵
    y = np.zeros(shape=numPoints)
    # basically a straight line
    for i in range(0, numPoints):
        # bias feature
        x[i][0] = 1
        x[i][1] = i
        # our target variable 产生结果值
        y[i] = (i + bias) + random.uniform(0, 1) * variance
    return x, y

# gen 100 points with a bias of 25 and 10 variance as a bit of noise
x, y = genData(100, 25, 10)
m, n = np.shape(x)# 查询生成矩阵的大小 100 x 2
numIterations= 100000
alpha = 0.0005
theta = np.ones(n)
theta = gradientDescent(x, y, theta, alpha, m, numIterations)
print(theta)

#预测
print(2 * theta[0] + 3 * theta[1])
