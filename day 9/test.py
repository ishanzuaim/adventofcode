f = open("input9.txt", "r")

pos = []
h_pos = [1000, 1000]
t_pos = [[1000, 1000] for x in range(0, 9)]
pos.append(''.join(str(x) for x in t_pos[0]))
for x in f.read().split("\n"):
    if x == "":
        break
    [dire, amt] = x.split(" ")
    # print(dire, amt)
    for n in range(0, int(amt)):
        if dire == "R":
            h_pos[0] += 1
        elif dire == "U":
            h_pos[1] += 1
        elif dire == "L":
            h_pos[0] -= 1
        elif dire == "D":
            h_pos[1] -= 1
        
        x_diff = h_pos[0] - t_pos[8][0]
        y_diff = h_pos[1] - t_pos[8][1]
        if abs(x_diff) > 1 and y_diff == 0:
            if x_diff > 0:
                t_pos[8][0] += (abs(x_diff) - 1)
            else:
                t_pos[8][0] -= (abs(x_diff) - 1)
        elif abs(y_diff) > 1 and x_diff == 0:
            if y_diff > 0:
                t_pos[8][1] += (abs(y_diff) - 1)
            else:
                t_pos[8][1] -= (abs(y_diff) - 1)
        elif (abs(x_diff) == 1 and abs(y_diff) > 1) or (abs(x_diff) > 1 and abs(y_diff) == 1):
            if y_diff > 0 and x_diff > 0:
                t_pos[8][0] += 1
                t_pos[8][1] += 1
            elif y_diff > 0 and x_diff < 0:
                t_pos[8][0] -= 1
                t_pos[8][1] += 1
            elif y_diff < 0 and x_diff > 0:
                t_pos[8][0] += 1
                t_pos[8][1] -= 1
            elif y_diff < 0 and x_diff < 0:
                t_pos[8][0] -= 1
                t_pos[8][1] -= 1

        for i in range(7, -1, -1):
            x_diff = t_pos[i+1][0] - t_pos[i][0]
            y_diff = t_pos[i+1][1] - t_pos[i][1]

            if abs(x_diff) <= 1 and abs(y_diff) <= 1:
                continue
            # if i == 0: print(t_pos, x_diff, y_diff, h_pos)
            if abs(x_diff) > 1 and y_diff == 0:
                if x_diff > 0:
                    t_pos[i][0] += (abs(x_diff) - 1)
                else:
                    t_pos[i][0] -= (abs(x_diff) - 1)
            elif abs(y_diff) > 1 and x_diff == 0:
                if y_diff > 0:
                    t_pos[i][1] += (abs(y_diff) - 1)
                else:
                    t_pos[i][1] -= (abs(y_diff) - 1)
            else:
                if y_diff > 0 and x_diff > 0:
                    t_pos[i][0] += 1
                    t_pos[i][1] += 1
                elif y_diff > 0 and x_diff < 0:
                    t_pos[i][0] -= 1
                    t_pos[i][1] += 1
                elif y_diff < 0 and x_diff > 0:
                    t_pos[i][0] += 1
                    t_pos[i][1] -= 1
                elif y_diff < 0 and x_diff < 0:
                    t_pos[i][0] -= 1
                    t_pos[i][1] -= 1
        pos.append(''.join(str(x) for x in t_pos[0]))
    
print(len(set(pos)))

