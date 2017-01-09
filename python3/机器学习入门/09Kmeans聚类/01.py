import numpy as np

"""
x 数控
k 分类数目
maxIt 最大迭代数目
"""
def kmeans(X, k, maxIt):
    
    numPoints, numDim = X.shape
    
    #创建矩阵
    dataSet = np.zeros((numPoints, numDim + 1))
    dataSet[:, :-1] = X
    
    #随机选择矩阵中的k行
    centroids = dataSet[np.random.randint(numPoints, size = k)]
#    centroids = dataSet[0:2, :]
    
    #初始化类型 1, 2, 3 .....
    centroids[:, -1] = range(1, k + 1)
    
    #循环次数
    iterations = 0
    #旧的中心点
    oldCentroids = None
    
    while not shouldStop(oldCentroids, centroids, iterations, maxIt):
        print("iteration: \n", iterations)
        print("dataSet: \n", dataSet)
        print("centroids: \n", centroids)

        #拷贝数据
        oldCentroids = np.copy(centroids)
        iterations += 1
        
        #分类
        updateLabels(dataSet, centroids)
        
        #计算新的中心点
        centroids = getCentroids(dataSet, k)
        
    return dataSet

#判断迭代是否停止
def shouldStop(oldCentroids, centroids, iterations, maxIt):
    if iterations > maxIt:
        return True
    return np.array_equal(oldCentroids, centroids)  

def updateLabels(dataSet, centroids):
    # For each element in the dataset, chose the closest centroid. 
    # Make that centroid the element's label.
    #centroids中心点
    #dataSet判断的数据
    numPoints, numDim = dataSet.shape
    for i in range(0, numPoints):
        dataSet[i, -1] = getLabelFromClosestCentroid(dataSet[i, :-1], centroids)
#计算类型
def getLabelFromClosestCentroid(dataSetRow, centroids):
    label = centroids[0, -1];
    minDist = np.linalg.norm(dataSetRow - centroids[0, :-1])
    for i in range(1 , centroids.shape[0]):
        #计算距离
        dist = np.linalg.norm(dataSetRow - centroids[i, :-1])
        if dist < minDist:
            minDist = dist
            label = centroids[i, -1]
    print("minDist:", minDist)
    return label

def getCentroids(dataSet, k):
    # Each centroid is the geometric mean of the points that
    # have that centroid's label. Important: If a centroid is empty (no points have
    # that centroid's label) you should randomly re-initialize it.
    #初始化矩阵
    result = np.zeros((k, dataSet.shape[1]))
    for i in range(1, k + 1):
        oneCluster = dataSet[dataSet[:, -1] == i, :-1]
        #求每个类的中心值 axis 对行求均值
        result[i - 1, :-1] = np.mean(oneCluster, axis = 0)
        result[i - 1, -1] = i
    
    return result

    
x1 = np.array([1, 1])
x2 = np.array([2, 1])
x3 = np.array([4, 3])
x4 = np.array([5, 4])
testX = np.vstack((x1, x2, x3, x4))

result = kmeans(testX, 2, 10)
print("final result:")
print(result)