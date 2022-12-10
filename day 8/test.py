f = open("input8.txt", "r")
ls = []
for i in f.read().split("\n"):
    if i == "":
        break
    ls.append([*i])

visible = 0
total = []
for y in range(0, len(ls)):
    for x in range(0, len(ls[0])):
        if x == 0 or x == len(ls[0])-1 or y == 0 or y == len(ls)-1:
            continue
        curr = ls[y][x]
    
        x_total = 0
        x2_total = 0
        for xr in range(0, x):
            if ls[y][x - 1 - xr] < curr:
                x_total += 1
            else:
                x_total+=1
                break
        for xr in range(x+1, len(ls[0])):
            if ls[y][xr] < curr:
                x2_total += 1
            else:
                x2_total+=1
                break

        y_total = 0
        y2_total = 0
        for yr in range(0, y):
            if ls[y - 1 - yr][x] < curr:
                y_total += 1
            else:
                y_total += 1
                break
        
        for yr in range(y+1, len(ls)):
            if ls[yr][x] < curr:
                y2_total += 1
            else:
                y2_total += 1
                break

        total.append(x_total*y_total*x2_total*y2_total)

print(max(total))
# print(ls)
    

