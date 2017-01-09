import tensorflow as tf

# 定义变量值为0 名称为counter
state = tf.Variable(0, name='counter') #print(state.name)
# 定义常量
one = tf.constant(1)

#进行加法运算
new_value = tf.add(state, one)
#重新赋值
update = tf.assign(state, new_value)

# 在定义Variable情况下一定要使用 initialize_all_variables 来进行初始化
init = tf.global_variables_initializer()

with tf.Session() as sess:
    sess.run(init)
    for _ in range(3):
        sess.run(update)
        print(sess.run(state))
