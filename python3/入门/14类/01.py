class className:
	name = "aaa"
	def __init__(self, b = 1):
		print("new class")
	def funcname(self, a = 1):
		print(self.name)
		print(a)

classN = className()
classN.name
classN.funcname()