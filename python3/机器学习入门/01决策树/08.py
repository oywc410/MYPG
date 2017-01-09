from sklearn.feature_extraction import DictVectorizer
import csv
from sklearn import preprocessing
from sklearn import tree
from sklearn.externals.six import StringIO

with open('test_data.csv') as f:
    f_csv = csv.reader(f)
    headers = next(f_csv)
    
    # 数据装换
    labelList = [] # 结果
    featureList = [] # 判断数据
    
    for row in f_csv:
        labelList.append(row[len(row) - 1])
        rowDict = {}
        for i in range(1, len(row) - 1):
            rowDict[headers[i]] = row[i]
        featureList.append(rowDict)
        
    print(labelList)
    print(featureList)
    
    # 转化为分析数据 
    vec = DictVectorizer()
    dummyX = vec.fit_transform(featureList).toarray()
    
    print("dummyX:" + str(dummyX))
    print(vec.get_feature_names())
    
    print("labellist: " + str(labelList))
    
    # 转化目标值
    lb = preprocessing.LabelBinarizer()
    dummyY = lb.fit_transform(labelList)
    print("dummyY: " + str(dummyY))
    
    # 分类器 创建决策树  criterion 选取root点方法/度量标准 
    clf = tree.DecisionTreeClassifier(criterion='entropy')
    clf = clf.fit(dummyX, dummyY)
    print("clf:" + str(clf))
    
    with open("test.dot", 'w') as f2:
        # 可视化数据
        f2 = tree.export_graphviz(clf, feature_names=vec.get_feature_names(), out_file = f2)
        # dot 文件 转化为 pdf 文件     命令: dot -T pdf test.dot -o output.pdf
        
    #预测值 (取第一行)
    oneRowX = dummyX[0, :]
    print("oneRowX:" + str(oneRowX))
    
    newRowX = oneRowX
    #创建新数据
    newRowX[0] = 1
    newRowX[2] = 0 
    print("newRowX:" + str(newRowX))
    
    #预测数据
    predictedY = clf.predict(newRowX)
    print("predictedY:" + str(predictedY))
    