import mnist_loader
import network

traing_data, validation_data, test_data = mnist_loader.load_data_wrapper()
print("training data")
print(type(traing_data))
print(len(traing_data))
print(traing_data[0][0].shape)
print(traing_data[0][1].shape)
#print(traing_data[0])

print("validation data")
print(len(validation_data))
print("test data")
print(len(test_data))

net = network.Network([784, 30, 100, 60, 30, 10])
net.SGD(traing_data, 30, 10, 2.0, test_data=test_data)

#net = network.Network([784, 30, 10])
#net.SGD(traing_data, 30, 10, 3.0, test_data=test_data)

#net = network.Network([784, 100, 10])
#net.SGD(traing_data, 30, 10, 3.0, test_data=test_data)

#net = network.Network([784, 100, 10])
#net.SGD(traing_data, 30, 10, 0.001, test_data=test_data)

#net = network.Network([784, 10])
#net.SGD(traing_data, 30, 10, 3.0, test_data=test_data)