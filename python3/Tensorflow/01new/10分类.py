import tensorflow as tf
from tensorflow.examples.tutorials.mnist import input_data
from tensorflow.contrib.labeled_tensor import batch

#学习数据
mnist = input_data.read_data_sets('MNIST_data', one_hot=True)

def add_layer(inputs, in_size, out_size, activation_function=None):
    Weights = tf.Variable(tf.random_normal([in_size, out_size])) #变量矩阵
    biases = tf.Variable(tf.zeros([1, out_size]) + 0.1) #列表 初始为0 + 0.1
    Wx_plus_b = tf.matmul(inputs, Weights) + biases # 乘法
    if activation_function is None :
        outputs = Wx_plus_b
    else:
        outputs = activation_function(Wx_plus_b)
    return outputs

#输出准确率
def compute_accuracy(v_xs, v_ys):
    global prediction
    y_pre = sess.run(prediction, feed_dict={xs: v_xs})#预测
    correct_prediction = tf.equal(tf.argmax(y_pre,1), tf.argmax(v_ys,1))#对比预测值与真实值
    accuracy = tf.reduce_mean(tf.cast(correct_prediction, tf.float32))
    result = sess.run(accuracy, feed_dict={xs: v_xs, ys: v_ys})
    return result

xs = tf.placeholder(tf.float32, [None, 784]) # 输入数据 28 x 28 的像素点
ys = tf.placeholder(tf.float32, [None, 10]) #输出结果 [0,0,0,1,0,0,0,0,0,0,0]

l1 = add_layer(xs, 784, 30, activation_function=tf.nn.sigmoid) #第一层使用 sigmoid

prediction = add_layer(l1, 30, 10, activation_function=tf.nn.softmax) #第二层使用softmax 

cross_entropy = tf.reduce_mean(-tf.reduce_sum(ys * tf.log(prediction), reduction_indices=[1])) #loss

train_step = tf.train.GradientDescentOptimizer(0.5).minimize(cross_entropy)

sess = tf.Session()
sess.run(tf.global_variables_initializer())

for i in range(5000):
    batch_xs, batch_ys = mnist.train.next_batch(100) #随机取出 100 个一组数据 学习 1000组
    sess.run(train_step, feed_dict={xs: batch_xs, ys: batch_ys})
    if i % 50 == 0:
        print(compute_accuracy(
            mnist.test.images, mnist.test.labels))
        
        