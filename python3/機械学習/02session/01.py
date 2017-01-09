import tensorflow as tf

# 常量
matrix1 = tf.constant([[3,3]])
matrix2 = tf.constant([[2],[2]])

# 矩阵乘法
product = tf.matmul(matrix1, matrix2) # matrix multply np.dot(m1, m2)

#方法1
sess = tf.Session()
result = sess.run(product) #执行结构
print(result)
sess.close()

# 方法2
with tf.Session() as sess:
	resutlt = sess.run(product)
	print(resutlt)