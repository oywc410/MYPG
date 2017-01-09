from numpy import genfromtxt
import numpy as np
from sklearn import datasets, linear_model

dataPath = r"./data.csv"
deliverData = genfromtxt(dataPath, delimiter=",")

print("data")
print(deliverData)

X = deliverData[:, :-1]
Y = deliverData[:, -1]

print("X:")
print(X)
print("Y:")
print(Y)

regr = linear_model.LinearRegression()

regr.fit(X, Y)

print(X, Y)

print("coefficients")
print(regr.coef_)
print("intercept: ")
print(regr.intercept_)

xPred = [102, 6]
yPred = regr.predict(xPred)
print("predicted y: ")
print(yPred)

# 当线性回归中的属性涉及到有分类属性时 详见图02.png (将类型分为多属性 并默认值为0 只有该类型为1)