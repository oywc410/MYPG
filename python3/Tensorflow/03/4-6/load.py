from scipy.io import loadmat as load
import matplotlib.pyplot as plt
import numpy as np

traindata = load('../data/train_32x32.mat') #训练数据
#testdata = load('../data/test_32x32.mat') #测试数据
#extradata = load('../data/extra_32x32.mat') #更多的训练数据

print('Train Data Samples Shape:', traindata['X'].shape)
print('Train Data Labels Shap:', traindata['y'].shape)

#print('Test Data Samples Shape:', testdata['X'].shape)
#print('Test Data Labels Shap:', testdata['y'].shape)

#print('Extra Data Samples Shape:', extradata['X'].shape)
#print('Extra Data Labels Shap:', extradata['y'].shape)

train_samples = traindata['X']
train_labels = traindata['y']
#test_samples = testdata['X']
#test_labels = testdata['y']
#extra_samples = extradata['X']
#extra_labels = extradata['y']

def reformat(samples, labels):
	# 改变原始数据的形状
	#  0       1       2      3          3       0       1      2
	# (图片高，图片宽，通道数，图片数) -> (图片数，图片高，图片宽，通道数)
	new = np.transpose(samples, (3, 0, 1, 2)).astype(np.float32)

	# labels 变成 one-hot encoding, [2] -> [0, 0, 1, 0, 0, 0, 0, 0, 0, 0]
	# 原始数字0 被标记为 10
	# labels 变成 one-hot encoding, [10] -> [1, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	labels = np.array([x[0] for x in labels])	# slow code, whatever
	one_hot_labels = []
	for num in labels:
		one_hot = [0.0] * 10
		if num == 10:
			one_hot[0] = 1.0
		else:
			one_hot[num] = 1.0
		one_hot_labels.append(one_hot)
	labels = np.array(one_hot_labels).astype(np.float32)
	return new, labels

def normalize(samples):
	'''
	并且灰度化: 从三色通道 -> 单色通道     省内存 + 加快训练速度
	(R + G + B) / 3
	将图片从 0 ~ 255 线性映射到 -1.0 ~ +1.0
	@samples: numpy array
	'''
	a = np.add.reduce(samples, keepdims=True, axis=3)  # shape (图片数，图片高，图片宽，通道数)
	a = a/3.0 #灰度
	return a/128.0 - 1.0 #映射到 -1.0 ~ +1.0

def inspect(dataset, labels, i):
	# 显示图片看看
	if dataset.shape[3] == 1:
		shape = dataset.shape
		dataset = dataset.reshape(shape[0], shape[1], shape[2])
	print(labels[i])
	plt.imshow(dataset[i])
	plt.show()


_train_samples, _train_labels = reformat(train_samples, train_labels)
#_test_samples, _test_labels = reformat(test_samples, test_labels)
#_extra_samples, _extra_labels = reformat(extra_samples, extra_labels)

if __name__ == '__main__':
        #inspect(_train_samples, _train_labels, 0)
        inspect(_train_samples, _train_labels, 1234)
        _train_samples = normalize(_train_smples)
        inspect(_train_samples, _train_labels, 1234)

