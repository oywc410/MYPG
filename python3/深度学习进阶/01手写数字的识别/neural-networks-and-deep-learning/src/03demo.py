import mnist_loader 
training_data, validation_data, test_data = mnist_loader.load_data_wrapper() 
import network2 
# Regularized cross-entropy  Regularized L2 减少overfitting
# Softmax 算法
net = network2.Network([784, 30, 10], cost=network2.CrossEntropyCost)
net.large_weight_initializer() #用原始形式初始化w b N(1, 0)     新方法初始化w b N(0, 1/sqrt(n_in))默认 把此行删掉即可
net.SGD(training_data[:1000], 400, 10, 0.5,
        evaluation_data=test_data, lmbda = 0.1, #lmbda: overfitting中的参数  随着训练集的增大 需要调节lmbda变大
        monitor_evaluation_cost=True, monitor_evaluation_accuracy=True,
        monitor_training_cost=True, monitor_training_accuracy=True)

net.SGD(training_data, 400, 10, 0.5,
        evaluation_data=test_data, lmbda = 0.5, #lmbda: overfitting中的参数  随着训练集的增大 需要调节lmbda变大
        monitor_evaluation_cost=True, monitor_evaluation_accuracy=True,
        monitor_training_cost=True, monitor_training_accuracy=True)



training_data, validation_data, test_data = mnist_loader.load_data_wrapper()
net = network2.Network([784, 30, 10], cost = network2.CrossEntropyCost)

net.SGD(training_data, 30, 10, 0.5, 5.0, 
        evaluation_data = validation_data, monitor_evaluation_accuracy=True, 
        monitor_evaluation_cost=True, monitor_training_accuracy=True, 
        monitor_training_cost=True)
