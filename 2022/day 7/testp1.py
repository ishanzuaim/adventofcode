f = open("input7.txt", "r")

obj = {}
curr_dir = ""
mode = ""
for i in f.read().split("\n"):
    if i == "":
        break

    cmd = i.split(" ")
    if cmd[0] == "$":
        if cmd[1] == "cd":
            if cmd[2] == "..":
                curr_dir = curr_dir.rsplit("/", 1)[0]
            elif cmd[2] == "/":
                curr_dir = "/"
                obj["/"] = 0
            else:
                curr_dir += ("/" + cmd[2])

        continue

    if cmd[0] == "dir":
        temp_value = curr_dir
        temp_value += ("/" + cmd[1])
        if temp_value not in obj.keys():
            obj[temp_value] = 0
    else:
        obj[curr_dir] += float(cmd[0])
        value = float(cmd[0])
        # print("temp", curr_dir)
        temp = curr_dir.rsplit("/", 1)[0]
        while temp != "":
            # print("temp", temp)
            obj[temp] += value
            temp = temp.rsplit("/", 1)[0]
        # print(obj)

print(obj)

total = 0
for i in obj.values():
    if i <= 100000:
        total += i

free_space = obj["/"] - 40000000

print(obj["/"], free_space)
print(list(filter(lambda v: (v >= free_space), obj.values())))



