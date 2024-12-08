
f = open("input7.txt", "r")

obj = {}
curr_dir = ""
mode = ""
deep = 0
for i in f.read().split("\n"):
    if i == "":
        break

    par = i.split(" ")
    print(par, deep, curr_dir)
    if par[0] == "$":
        if par[1] == "cd":
            val = par[2]
            if val == "..":
                # print(obj[curr_dir], curr_dir)
                curr_dir = obj[curr_dir]["parent"]
                deep -= 1
            else:
                deep +=1
                curr_dir = val
        elif par[1] == "ls":
            mode = "ls"
        continue

    if mode == "ls":
        if par[0] == "dir":
            if par[1] not in obj.keys():
                obj[par[1]] = {}
                obj[par[1]]["parent"] = curr_dir
                obj[par[1]]["deep"] = deep

        else:
            if curr_dir not in obj.keys():
                obj[curr_dir] = {}
                obj[curr_dir]["deep"] = deep-1

            if "value" not in obj[curr_dir].keys():
                obj[curr_dir]["value"] = float(par[0])
            else:
                obj[curr_dir]["value"] += float(par[0])

new_obj = {}
print(obj)

curr = deep

while curr != -1:
    for i in obj.keys():
        if obj[i]["deep"] == curr:
            if i not in new_obj.keys():
                new_obj[i] = obj[i]["value"]
            else:
                new_obj[i] += obj[i]["value"]
                
            if "parent" not in obj[i].keys():
                break
            if obj[i]["parent"] not in new_obj.keys():
                new_obj[obj[i]["parent"]] = new_obj[i]
            else:
                new_obj[obj[i]["parent"]] += new_obj[i]
    
    curr -= 1

print(new_obj)

total = 0
for i in new_obj.keys():
    if new_obj[i] < 100000:
        total += new_obj[i]



print(total)
