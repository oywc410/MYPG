# 第三方
import tensorflow as tf
from sklearn.metrics import confusion_matrix
import numpy as np

# 我们自己
import load

train_samples, train_labels = load._train_samples, load._train_labels
test_samples, test_labels = load._test_samples,  load._test_labels

print('Training set', train_samples.shape, train_labels.shape)
print('    Test set', test_samples.shape, test_labels.shape)

image_size = load.image_size
num_labels = load.num_labels
num_channels = load.num_channels

def get_chunk(samples, labels, chunkSize):
	'''
	Iterator/Generator: get a batch of data
	这个函数是一个迭代器/生成器，用于每一次只得到 chunkSize 这么多的数据
	用于 for loop， just like range() function
	'''
	if len(samples) != len(labels):
		raise Exception('Length of samples and labels must equal')
	stepStart = 0	# initial step
	i = 0
	while stepStart < len(samples):
		stepEnd = stepStart + chunkSize
		if stepEnd < len(samples):
			yield i, samples[stepStart:stepEnd], labels[stepStart:stepEnd]
			i += 1
		stepStart = stepEnd


class Network():
	def __init__(self, num_hidden, batch_size, patch_size, conv_depth, pooling_scale, optimizeMethod='adam'):
		'''
		@num_hidden: 隐藏层的节点数量
		@batch_size：因为我们要节省内存，所以分批处理数据。每一批的数据量。
		'''
		self.batch_size = batch_size
		self.test_batch_size = 500

		self.optimizeMethod = optimizeMethod;

		# Hyper Parameters
		self.num_hidden = num_hidden
		self.patch_size = patch_size  # 滑窗大小
		self.conv1_depth = conv_depth  # 输出结果层数
		self.conv2_depth = conv_depth  # 卷积层数
		self.conv3_depth = conv_depth
		self.conv4_depth = conv_depth
		self.last_conv_depth = self.conv3_depth
		self.pooling_scale = pooling_scale  # poling图片缩小
		self.pooling_stride = self.pooling_scale

		# Graph Related
		self.graph = tf.Graph()
		self.tf_train_samples = None
		self.tf_train_labels = None
		self.tf_test_samples = None
		self.tf_test_labels = None
		self.tf_test_prediction = None

		self.fc_weights = []
		self.fc_biases = []

		# 统计
		self.merged = None

		# 初始化
		self.define_graph()
		self.session = tf.Session(graph=self.graph)
		self.writer = tf.summary.FileWriter("logs/", self.graph)

	def apply_regularization(self, _lambda):
		# L2 regularization for the fully connected parameters
		regularization = 0.0
		for weights, biases in zip(self.fc_weights, self.fc_biases):
			regularization += tf.nn.l2_loss(weights) + tf.nn.l2_loss(biases)
		# 1e5
		return _lambda * regularization

	def define_graph(self):
		'''
		定义我的的计算图谱
		'''
		with self.graph.as_default():
			# 这里只是定义图谱中的各种变量
			with tf.name_scope('inputs'):
				self.tf_train_samples = tf.placeholder(
					tf.float32, shape=(self.batch_size, image_size, image_size, num_channels), name='tf_train_samples'
				)
				self.tf_train_labels  = tf.placeholder(
					tf.float32, shape=(self.batch_size, num_labels), name='tf_train_labels'
				)
				self.tf_test_samples  = tf.placeholder(
					tf.float32, shape=(self.test_batch_size, image_size, image_size, num_channels), name='tf_test_samples'
				)

			with tf.name_scope('conv1'):
				# conv1_depth 为卷积层数
				conv1_weights = tf.Variable(
				    tf.truncated_normal([self.patch_size, self.patch_size, num_channels, self.conv1_depth], stddev=0.1))
				conv1_biases = tf.Variable(tf.zeros([self.conv1_depth]))

			with tf.name_scope('conv2'):
				conv2_weights = tf.Variable(
				    tf.truncated_normal([self.patch_size, self.patch_size, self.conv1_depth, self.conv2_depth], stddev=0.1))
				conv2_biases = tf.Variable(tf.constant(0.1, shape=[self.conv2_depth]))

			with tf.name_scope('conv3'):
				conv3_weights = tf.Variable(
				    tf.truncated_normal([self.patch_size, self.patch_size, self.conv2_depth, self.conv3_depth], stddev=0.1))
				conv3_biases = tf.Variable(tf.constant(0.1, shape=[self.conv3_depth]))

			with tf.name_scope('conv4'):
				conv4_weights = tf.Variable(
				    tf.truncated_normal([self.patch_size, self.patch_size, self.conv3_depth, self.conv4_depth], stddev=0.1))
				conv4_biases = tf.Variable(tf.constant(0.1, shape=[self.conv4_depth]))

			# fully connected layer 1, fully connected
			with tf.name_scope('fc1'):
				fc1_weights = tf.Variable(
					tf.truncated_normal([image_size * image_size, self.num_hidden], stddev=0.1), name='fc1_weights'
				)
				fc1_biases = tf.Variable(tf.constant(0.1, shape=[self.num_hidden]), name='fc1_biases')
				tf.summary.histogram('fc1_weights', fc1_weights)
				tf.summary.histogram('fc1_biases', fc1_biases)

			# fully connected layer 2 --> output layer
			with tf.name_scope('fc2'):
				fc2_weights = tf.Variable(
					tf.truncated_normal([self.num_hidden, num_labels], stddev=0.1), name='fc2_weights'
				)
				fc2_biases = tf.Variable(tf.constant(0.1, shape=[num_labels]), name='fc2_biases')
				tf.summary.histogram('fc2_weights', fc2_weights)
				tf.summary.histogram('fc2_biases', fc2_biases)


			# 想在来定义图谱的运算
			def model(data, train=True):
				# fully connected layer 1
				shape = data.get_shape().as_list()
				reshape = tf.reshape(data, [shape[0], shape[1] * shape[2] * shape[3]])

				with tf.name_scope('fc1_model'):
					self.fc_weights.append(fc1_weights)
					self.fc_biases.append(fc1_biases)
					fc1_model = tf.matmul(reshape, fc1_weights) + fc1_biases
					hidden = tf.nn.relu(fc1_model)
					if train:
						hidden = tf.nn.dropout(hidden, 0.5, seed=4926)  # 随机丢弃神经节点来减少拟合

				# fully connected layer 2
				with tf.name_scope('fc2_model'):
					self.fc_weights.append(fc1_weights)
					self.fc_biases.append(fc1_biases)
					return tf.matmul(hidden, fc2_weights) + fc2_biases

			# Training computation.
			logits = model(self.tf_train_samples)
			with tf.name_scope('loss'):
				self.loss = tf.reduce_mean(
					tf.nn.softmax_cross_entropy_with_logits(logits=logits, labels=self.tf_train_labels)
				)
				self.loss += self.apply_regularization(_lambda=5e-4) # 0.0005  l2 regularization    使用l1 或者 l2来减少拟合
				tf.summary.scalar('Loss', self.loss)

			# learning rate decay 学习率衰退 (学习率)　　0.9   1 x 0.9 =0.9 => 0.9 x 0.9 = 0.81 => 0.81 x 0.81 (衰退方法)
			global_step = tf.Variable(0)
			lr = 0.001
			dr = 0.99
			learning_rate = tf.train.exponential_decay(
				learning_rate=lr, #开始
				global_step=global_step * self.batch_size, #总共走了多少步 (一次批处理算做一步)
				decay_steps=100, #多少步衰减一次
				decay_rate=dr, #衰减量(因子)
				staircase=True
			)

			# Optimizer.
			with tf.name_scope('optimizer'):
				if (self.optimizeMethod == 'gradient'):
					self.optimizer = tf.train \
						.GradientDescentOptimizer(learning_rate) \
						.minimize(self.loss)
				elif (self.optimizeMethod == 'momentum'):
					self.optimizer = tf.train \
						.MomentumOptimizer(learning_rate, 0.5) \
						.minimize(self.loss)
				elif (self.optimizeMethod == 'adam'):
					self.optimizer = tf.train \
						.AdamOptimizer(learning_rate) \
						.minimize(self.loss)

			# Predictions for the training, validation, and test data.
			with tf.name_scope('predictions'):
				self.train_prediction = tf.nn.softmax(logits, name='train_prediction')
				self.test_prediction = tf.nn.softmax(model(self.tf_test_samples, True), name='test_prediction')

			self.merged = tf.summary.merge_all()

	def run(self):
		'''
		用到Session
		'''
		# private function
		def print_confusion_matrix(confusionMatrix):
			print('Confusion    Matrix:')
			for i, line in enumerate(confusionMatrix):
				print(line, line[i]/np.sum(line))
			a = 0
			for i, column in enumerate(np.transpose(confusionMatrix, (1, 0))):
				a += (column[i]/np.sum(column))*(np.sum(column)/26000)
				print(column[i]/np.sum(column),)
			print('\n',np.sum(confusionMatrix), a)


		with self.session as session:
			tf.global_variables_initializer().run()

			### 训练
			print('Start Training')
			# batch 1000
			for i, samples, labels in get_chunk(train_samples, train_labels, chunkSize=self.batch_size):
				_, l, predictions, summary = session.run(
					[self.optimizer, self.loss, self.train_prediction, self.merged],
					feed_dict={self.tf_train_samples: samples, self.tf_train_labels: labels}
				)
				self.writer.add_summary(summary, i)
				# labels is True Labels
				accuracy, _ = self.accuracy(predictions, labels)
				if i % 50 == 0:
					print('Minibatch loss at step %d: %f' % (i, l))
					print('Minibatch accuracy: %.1f%%' % accuracy)
			###

			### 测试
			accuracies = []
			confusionMatrices = []
			for i, samples, labels in get_chunk(test_samples, test_labels, chunkSize=self.test_batch_size):
				result = self.test_prediction.eval(feed_dict={self.tf_test_samples: samples})
				accuracy, cm = self.accuracy(result, labels, need_confusion_matrix=True)
				accuracies.append(accuracy)
				confusionMatrices.append(cm)
				print('Test Accuracy: %.1f%%' % accuracy)
			print(' Average  Accuracy:', np.average(accuracies))
			print('Standard Deviation:', np.std(accuracies))
			print_confusion_matrix(np.add.reduce(confusionMatrices))
			###

	def accuracy(self, predictions, labels, need_confusion_matrix=False):
		'''
		计算预测的正确率与召回率
		@return: accuracy and confusionMatrix as a tuple
		'''
		_predictions = np.argmax(predictions, 1)
		_labels = np.argmax(labels, 1)
		cm = confusion_matrix(_labels, _predictions) if need_confusion_matrix else None
		# == is overloaded for numpy array
		accuracy = (100.0 * np.sum(_predictions == _labels) / predictions.shape[0])
		return accuracy, cm


if __name__ == '__main__':
	net = Network(num_hidden=128, batch_size=100, patch_size=3, conv_depth=16, pooling_scale=2)
	net.run()