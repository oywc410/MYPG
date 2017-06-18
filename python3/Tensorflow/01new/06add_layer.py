import tensorflow as tf
import matplotlib.pyplot as plt
import numpy as np

tf.set_random_seed(1)
np.random.seed(1)

x = np.linspace(-1, 1, 100)[:, np.newaxis] #生成随机数 并转化为数组
noise = np.random.normal(0, 0.1, size=x.shape)  #噪点 模拟真实数据
y = np.power(x, 2) + noise  #二次方

# plot data 打印数据
plt.scatter(x, y)
plt.show()

tf_x = tf.placeholder(tf.float32, x.shape)     # input x
tf_y = tf.placeholder(tf.float32, y.shape)     # input y

# neural network layers
l1 = tf.layers.dense(tf_x, 10, tf.nn.relu)          # 添加隐藏层
output = tf.layers.dense(l1, 1)                     # 输出层

loss = tf.losses.mean_squared_error(tf_y, output)   # 计算差
optimizer = tf.train.GradientDescentOptimizer(learning_rate=0.5) #梯度下降算法
train_op = optimizer.minimize(loss) # GradientDescentOptimizer 学习进度 minimize 减少误差

sess = tf.Session()
sess.run(tf.global_variables_initializer())

plt.ion()

for step in range(100):
    # train and net output
    _, l, pred = sess.run([train_op, loss, output], {tf_x: x, tf_y: y})
    if step % 5 == 0:
        # plot and show learning process
        plt.cla()
        plt.scatter(x, y)
        plt.plot(x, pred, 'r-', lw=5)
        plt.text(0.5, 0, 'Loss=%.4f' % l, fontdict={'size': 20, 'color': 'red'})
        plt.pause(0.1)

plt.ioff()
plt.show()