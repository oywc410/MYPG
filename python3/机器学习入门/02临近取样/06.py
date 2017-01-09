from sklearn import neighbors
from sklearn import datasets

#分类器
knn = neighbors.KNeighborsClassifier()

# 内置数据
iris = datasets.load_iris()

#内置数据的结构
#print(iris)

# 载入数据  iris.data 多维特征值 iris.target一维 每个实例对于的象s
knn.fit(iris.data, iris.target)
# 预测类型
item = [[0.1, 0.2, 0.3, 0.4]]
predictedLabel = knn.predict(item)

print("判定花の属性: \nsepal length (cm): " + str(item[0][0]))
i = 0
for key in iris.feature_names:
	print(key + ":" + str(item[0][i]))
	i = i + 1

print("予測の花タイプは:" + iris.target_names[predictedLabel][0])