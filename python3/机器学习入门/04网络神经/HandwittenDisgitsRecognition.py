import numpy as np 
from sklearn.datasets import load_digits 
from sklearn.metrics import confusion_matrix, classification_report 
from sklearn.preprocessing import LabelBinarizer 
from NeuralNetwork import NeuralNetwork
from sklearn.cross_validation import train_test_split

#下载数据
digits = load_digits()  
X = digits.data#数据集
y = digits.target#结果集
X -= X.min() #将所有数据转化为 0 - 1之间的值 (神经算法要求) 
X /= X.max()

nn = NeuralNetwork([64,100,10],'logistic')#由于是8 x 8 的图片及输入为64 输出的数字为0 - 9及输出为10
X_train, X_test, y_train, y_test = train_test_split(X, y)#将数据集分组为训练集 和 测试集
labels_train = LabelBinarizer().fit_transform(y_train)  # 转化为运算的结果集  例: 2  => [0,0,1,0,0,0,0,0,0,0]
labels_test = LabelBinarizer().fit_transform(y_test)
print("start fitting")
nn.fit(X_train,labels_train,epochs=3000)#训练
#预测结果
predictions = []  
for i in range(X_test.shape[0]):  
    o = nn.predict(X_test[i] )  
    predictions.append(np.argmax(o))#获取最大值的位置 例:[  2.15258127e-04   4.55731727e-02   2.01712666e-03   9.78470277e-03 8.33249280e-03   1.09374505e-02   9.29067994e-05   6.33198745e-01   3.34788810e-02   5.13996088e-02]

print(confusion_matrix(y_test,predictions))  
"""
预测结果 
   0  1  2  3  4  5  6  7  8  9  
[[45  0  0  0  0  0  1  0  0  0]0
 [ 0 41  0  0  0  0  1  0  0  4]1
 [ 0  1 38  1  0  0  0  0  0  0]2
 [ 0  1  0 48  0  1  0  2  0  0]3
 [ 0  2  0  0 49  0  0  0  0  0]4
 [ 0  0  0  0  0 43  0  0  0  0]5
 [ 0  2  0  0  0  0 44  0  0  0]6
 [ 0  1  0  0  3  2  0 38  0  0]6
 [ 0  5  1  1  0  2  1  0 28  2]7
 [ 0  2  0  2  0  0  0  0  0 38]]8
"""
print(classification_report(y_test,predictions))
"""
            结果    所有预测为0图片的结果   所以真实是0的图片预测为0了
            precision    recall  f1-score   support

          0       1.00      0.98      0.99        46
          1       0.75      0.89      0.81        46
          2       0.97      0.95      0.96        40
          3       0.92      0.92      0.92        52
          4       0.94      0.96      0.95        51
          5       0.90      1.00      0.95        43
          6       0.94      0.96      0.95        46
          7       0.95      0.86      0.90        44
          8       1.00      0.70      0.82        40
          9       0.86      0.90      0.88        42

avg / total       0.92      0.92      0.92       450

"""