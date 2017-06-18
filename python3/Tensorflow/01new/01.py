import tensorflow as tf
import numpy as np

#creat data
x_data = np.random.rand(100).astype(np.float32)
y_data = x_data * 0.1 + 0.3 #学习目标

#学习过程#
#设置变量 建立结构
Weights = tf.Variable(tf.random_uniform([1], -1.0, 1.0)) #初始化
biases = tf.Variable(tf.zeros([1])) # 初始化

y = Weights * x_data + biases #预测目标结果

loss = tf.reduce_mean(tf.square(y-y_data)) #预测结果与真实值对比结果

optimizer = tf.train.GradientDescentOptimizer(0.1) #设置梯度下降算法与学习率
train = optimizer.minimize(loss) #

#初始化变量
init = tf.global_variables_initializer()

#创建session
sess = tf.Session()
sess.run(init) #执行

for step in range(2000):
    sess.run(train)
    if step % 20 == 0:
        print(step, sess.run(Weights), sess.run(biases))