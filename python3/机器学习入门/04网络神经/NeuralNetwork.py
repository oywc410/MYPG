import numpy as np

# sigmoid函数
def tanh(x):#双曲函数
    return np.tanh(x)

def tanh_deriv(x):#tanh 函数的导数
    return 1.0 - np.tanh(x) * np.tanh(x)

def logistic(x): #逻辑函数
    return 1/(1 + np.exp(-x))

def logistic_derivative(x):  
    return logistic(x)*(1-logistic(x))

class NeuralNetwork:
    def __init__(self, layers, activation='tanh'):
        """
        layers 网络每层神经层神经个数  10 2 2
        activation sigmoid函数
        """
        if activation == 'logistic':
            self.activation = logistic
            self.activation_deriv = logistic_derivative
        elif activation == 'tanh':
            self.activation = tanh
            self.activation_deriv = tanh_deriv
        
        #weights 俩层神经元间的权重
        self.weights = []  
        #初始化神经元之间的权重 0到1之间
        for i in range(1, len(layers) - 1):  
            self.weights.append((2*np.random.random((layers[i - 1] + 1, layers[i] + 1))-1)*0.25)  
            self.weights.append((2*np.random.random((layers[i] + 1, layers[i + 1]))-1)*0.25)
    
    def fit(self, X, y, learning_rate=0.2, epochs=10000):
        """
        X 训练集二维数组 每一行为一个实例, 每一列为特征值
        y 预测值分类 标记  一位数组
        learning_rate 学习率 (比例太大可能会越过极值 太小效率有问题)
        epochs 随机抽样数 默认最多10000次
        """
        X = np.atleast_2d(X)#转化为np二维类型
        temp = np.ones([X.shape[0], X.shape[1]+1])#创建默认值为1的矩阵 X的行 X的列+1 
        temp[:, 0:-1] = X
        X = temp #偏量的赋值
        y = np.array(y)#装换数据类型
        
        for k in range(epochs):  
            #随机抽取一行
            i = np.random.randint(X.shape[0])  
            a = [X[i]]
        
            #神经元的正向更新
            for l in range(len(self.weights)):  #对每层的神经网络进行计算
                #每个特征值与权重进行相乘并求和 再进行线性函数计算 得出下一个神经的输入值 (见图06 05)
                a.append(self.activation(np.dot(a[l], self.weights[l])))
                
            #神经元的反向更新 见图07
            error = y[i] - a[-1]  #取得差值  a[-1]:最后的输出值
            deltas = [error * self.activation_deriv(a[-1])] #进行 函数的导数 的运算
            
            #进行反向更新
            for l in range(len(a) - 2, 0, -1): # we need to begin at the second to last layer     
                #Compute the updated error (i,e, deltas) for each node going from top layer to input layer 
                deltas.append(deltas[-1].dot(self.weights[l].T)*self.activation_deriv(a[l]))  
            deltas.reverse()  #颠倒顺序
            #进行权重的更新
            for i in range(len(self.weights)):  
                layer = np.atleast_2d(a[i])  
                delta = np.atleast_2d(deltas[i])  
                self.weights[i] += learning_rate * layer.T.dot(delta)
    #接受新的实例 进行预测                
    def predict(self, x):       
        #实际就是神经元的正向计算的结果  
        x = np.array(x)         
        temp = np.ones(x.shape[0]+1)         
        temp[0:-1] = x         
        a = temp         
        for l in range(0, len(self.weights)):             
            a = self.activation(np.dot(a, self.weights[l]))         
        return a
    