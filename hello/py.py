import time 

sum = 0
st = time.time()
for i in range(1000000):
	sum += i
print(sum)
ed = time.time()
#print(st*10)
print(ed-st)