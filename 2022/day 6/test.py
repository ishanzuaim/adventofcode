f = open("input6.txt", "r")
count = 0
ls = [*f.read().split("\n")[0]]



for i in range(0, len(ls)):
    test_ls = set()
    for j in range(i, len(ls)):
        old_length = len(test_ls)
        test_ls.add(ls[j])
        if old_length == len(test_ls):
            break
        if(len(test_ls) == 14):
            print(j+1)
            exit(1);

