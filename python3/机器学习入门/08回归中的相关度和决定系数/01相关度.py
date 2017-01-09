import numpy as np
from astropy.units import Ybarn
import math

#计算相关度 见图02.png
def computeCorrelation(X, Y):
    xBar = np.mean(X)#计算均值
    yBar = np.mean(Y)#计算均值
    SSR = 0
    varX = 0
    varY = 0
    for i in range(0 , len(X)):
        diffXXBar = X[i] - xBar
        diffYYBar = Y[i] - yBar
        SSR += (diffXXBar * diffYYBar)
        varX +=  diffXXBar**2
        varY += diffYYBar**2
    
    SST = math.sqrt(varX * varY) #开方
    return SSR / SST

testX = [1, 3, 8, 7, 9]
testY = [10, 12, 24, 21, 34]

print("r:", computeCorrelation(testX, testY))
print("r^2:", computeCorrelation(testX, testY)**2)

#计算决定系数
"""
degree 回归次数 x平方 等
"""
def polyfit(x, y, degree):
    results = {}
    
    #numpy 线性回归计算(返回 b0 ,b1, b2....)
    coeffs = np.polyfit(x, y, degree) #简单与多元线性回归都可用
    
    results['polynomial'] = coeffs.tolist()
    
    #p 为线性回归的模型方程
    p = np.poly1d(coeffs)
    
    yhat = p(x)
    ybar = np.sum(y)/len(y)
    ssreg = np.sum((yhat-ybar)**2)
    sstot = np.sum((y - ybar)**2)
    results['determination'] = ssreg / sstot
    
    return results

print(polyfit(testX, testY, 1))

    