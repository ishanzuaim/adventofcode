f = open("input4.txt", "r")

count = 0
for i in f.read().split("\n"):
    if i == "":
        break

    [pair1, pair2] = i.split(",")
    [start1, end1] = pair1.split("-")
    [start2, end2] = pair2.split('-')

    ls1 = list(range(int(start1), int(end1)+1))
    ls2 = list(range(int(start2), int(end2)+1))

    if len(set(ls1+ls2)) != len(ls1+ls2):
        count+=1
print(count)
