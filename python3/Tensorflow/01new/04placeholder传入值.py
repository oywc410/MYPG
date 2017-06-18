import tensorflow as tf

input1 = tf.placeholder(tf.float32) #声明输入类型
input2 = tf.placeholder(tf.float32)

output = tf.mul(input1, input2) #乘法运算

with tf.Session() as sess:
    print(sess.run(output, feed_dict={input1:[7.], input2:[2.]})) #运行时输入值