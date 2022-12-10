f = open("input3.txt", "r")
total = 0
count = 0
ls = []
for i in f.read().split("\n"):
    if i == "":
        break
    
    count+=1
    ls.append(i)
    if count % 3 != 0:
        continue
    

    commons = set()

    # print(ls)
    for x in ls[0]:
        if x in [*ls[1]] and x in [*ls[2]]:
            # print(x)
            commons.add(x)

    # print(commons)
    for l in commons:
        if ord(l) >= 97:
            total += (ord(l) - 96)
        else:
            total+=(ord(l) - 38)

    ls = []

print(total) 

