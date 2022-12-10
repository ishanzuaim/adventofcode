f = open("input5.txt", "r")
ls = []
flag = 0
for i in f.read().split("\n"):
    if i == "":
        continue
    

    if(i[1] == "1"):
        flag = 1
        continue
    if not flag:
        r = 0
        count = 0
        for j in i.split(" "):
            if j != "" and j[0] == "[":
                for i in range(len(ls), r+1):
                    ls.append([])
                
                ls[r].append(j.strip("[]"))
            else:
                count+=1
            
            if count %4 == 0:
                r +=1
    else:
        amount = int(i.split(" ")[1])
        start = int(i.split(" ")[3]) - 1
        end = int(i.split(" ")[5]) - 1

        temp_ls = []
        for lsx in range(0, amount):
            temp_ls.append(ls[start].pop(0))

        temp_ls.extend(ls[end])
        ls[end] = temp_ls
        
        
for n in ls:
    print(n[0], end='')
