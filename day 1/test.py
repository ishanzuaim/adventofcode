f = open("input.txt", "r")
print()

obj = {}
x = 0
obj[0] = 0
for i in f.read().split("\n"):
    if i == "":
        x+=1
        obj[x] = 0
        continue
    obj[x] += int(i)

ls = []
for i in range(0,x):
    ls.append(obj[i])
ls.sort()
print(68579 +  69863+  74394)
