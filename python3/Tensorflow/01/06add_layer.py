import tensorflow as tf
import numpy as np

#添加神经层的函数
def add_layer(inputs, in_size, out_size, activation_function=None):
    Weights = tf.Variable(tf.random_normal([in_size, out_size])) #变量矩阵
    biases = tf.Variable(tf.zeros([1, out_size]) + 0.1) #列表 初始为0 + 0.1
    Wx_plus_b = tf.matmul(inputs, Weights) + biases # 乘法
    if activation_function is None :
        outputs = Wx_plus_b
    else:
        outputs = activation_function(Wx_plus_b)
    return outputs

# 300 行
x_data = np.linspace(-1, 1, 300)[:, np.newaxis] # 插入维度
#print(x_data)
#x_data.shape 维度 (300, 1)
noise = np.random.normal(0, 0.05, x_data.shape) #噪点 模拟真实数据
# np.square 计算各元素的平方 等于array**2
y_data = np.square(x_data) - 0.5 + noise #二次方


xs = tf.placeholder(tf.float32, [None, 1]) #输入 [None, 1] 对于noise 结构
ys = tf.placeholder(tf.float32, [None, 1]) #输出

# 隐藏层
# in_size 输入 out_size 神经元个数 #tf.nn.relu 非线性方程
l1 = add_layer(xs, 1, 10, activation_function=tf.nn.relu)

l2 = add_layer(l1, 10, 10, activation_function=tf.nn.relu)

#输出层 10 隐藏层个数 1 输出个数
predition = add_layer(l2, 10, 1, activation_function=None)

# 检查值 求平均值误差
loss = tf.reduce_mean(tf.reduce_sum(tf.square(ys - predition), reduction_indices=[1]))

# GradientDescentOptimizer 学习进度 minimize 减少误差
train_step = tf.train.GradientDescentOptimizer(0.01).minimize(loss)

# 初始化变量
init = tf.global_variables_initializer()
sess = tf.Session()
sess.run(init)

for i in range(1000):
    sess.run(train_step, feed_dict={xs:x_data, ys:y_data})
    if i % 50 == 0:
        print(sess.run(loss, feed_dict={xs:x_data, ys:y_data}))
sess.close()
