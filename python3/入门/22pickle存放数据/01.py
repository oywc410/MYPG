import pickle

a_dice = {'a': 'bbb', '21':[1,2,3]}

file = open('aaa.pickle', 'wb')
pickle.dump(a_dice, file)
file.close()

file = open('aaa.pickle', 'rb')
a_dice1 = pickle.load(file)
file.close()

print(a_dice1)