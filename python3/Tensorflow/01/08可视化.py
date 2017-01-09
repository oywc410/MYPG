import tensorflow as tf
import numpy as np

#添加神经层的函数
def add_layer(inputs, in_size, out_size, activation_function=None):
    with tf.name_scope('layer'):
        with tf.name_scope('weights'):
            Weights = tf.Variable(tf.random_normal([in_size, out_size]), name='W') #变量矩阵
        with tf.name_scope('biases'):
            biases = tf.Variable(tf.zeros([1, out_size]) + 0.1, name='b') #列表 初始为0 + 0.1
        with tf.name_scope('Wx_plus_b'):
            Wx_plus_b = tf.add(tf.matmul(inputs, Weights), biases) # 乘法
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

with tf.name_scope('inputs'):
    xs = tf.placeholder(tf.float32, [None, 1], name='x_input') #输入 [None, 1] 对于noise 结构
    ys = tf.placeholder(tf.float32, [None, 1], name='y_input') #输出

# 隐藏层
# in_size 输入 out_size 神经元个数 #tf.nn.relu 非线性方程
l1 = add_layer(xs, 1, 10, activation_function=tf.nn.relu)

l2 = add_layer(l1, 10, 10, activation_function=tf.nn.relu)

#输出层 10 隐藏层个数 1 输出个数
predition = add_layer(l2, 10, 1, activation_function=None)

# 检查值 求平均值误差
with tf.name_scope('loss'):
    loss = tf.reduce_mean(tf.reduce_sum(tf.square(ys - predition), reduction_indices=[1]))

# GradientDescentOptimizer 学习进度 minimize 减少误差
with tf.name_scope('train'):
    train_step = tf.train.GradientDescentOptimizer(0.01).minimize(loss)

# 初始化变量
init = tf.global_variables_initializer()
sess = tf.Session()

writer = tf.summary.FileWriter("logs/", sess.graph)

sess.run(init)

#tensorboard --logdir='logs/'
#游览器中 graphs 查看

###
#for i in range(1000):
#    sess.run(train_step, feed_dict={xs:x_data, ys:y_data})
#    if i % 50 == 0:
#        print(sess.run(loss, feed_dict={xs:x_data, ys:y_data}))
#sess.close()
###