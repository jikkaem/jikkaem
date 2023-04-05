import random


f = open("titles.txt", 'r')
lines = f.readlines()
f.close()

random.shuffle(lines)

f2 = open("randomisedTitles.txt", 'a')
for line in lines:
    f2.write(line)
f2.close()
