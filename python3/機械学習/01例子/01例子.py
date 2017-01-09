import tensorflow as tf
import numpy as np

# 创建测试数据  tensorflow 输入数据多数为float32
x_data = np.random.rand(100).astype(np.float32)
# 预测目
y_data = x_data * 0.1 + 0.3

# 创建结构 -------------------------------------------------- ###
# 生成随机预测值 (1维, 取值范围-1 ~ 1) Weights预测值最终接近0.1
Weights = tf.Variable(tf.random_uniform([1], -1.0, 1.0))
# 初始值0 biases 预测值最终接近0.3
biases = tf.Variable(tf.zeros([1]))
# y为预测
y = x_data * Weights + biases

# 取得差别
loss = tf.reduce_mean(tf.square(y-y_data))
#优化器 0.5 学习效率(小于1)
optimizer = tf.train.GradientDescentOptimizer(0.5)
train = optimizer.minimize(loss)

#初始化结构
init = tf.initialize_all_variables()

# 创建结构结束 ----------------------------------------------- ###

sess = tf.Session()
sess.run(init)

for step in range(201):
	sess.run(train)
	if step % 20 == 0:
		print(step, sess.run(Weights), sess.run(biases))