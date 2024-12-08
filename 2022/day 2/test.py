f = open("input2.txt", "r")

score = 0
obj = {
    "A": 1,
    "B": 2,
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3
}
score = 0
for i in f.read().split("\n"):
    if i == "":
        break
    [p1, p2] = i.split(" ")
    
    if obj[p2] == 1:
        #lose
        x = obj[p1]-1
        if x == 0:
            score += 3
        else:
            score+=x

    if obj[p2] == 2:
        #draw
        score += 3
        score += obj[p1]

    if obj[p2] == 3:
        #win
        x = obj[p1]+1
        if x == 4:
            score += 1
        else:
            score+=x
        score += 6
        

    

print(score)
