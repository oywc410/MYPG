
char_list = ['a', 'b', 'c', 'c', 'd', 'd', 'd']

print(set(char_list))

sentence = "Hello Hello aaaaa"

print(set(sentence))

unique = set(sentence)
unique.add('b')
unique.remove('a')
#unique.clear()

print(unique)

set1 = unique
set2 = {'a', 'e', 'i'}
print(set1.difference(set2))