f = open("./input.txt", "r")

line = f.readline()
res = 0
while line:
    line = line.strip()
    mid = int(len(line) / 2)
    left = set(line[:mid])
    right = set(line[mid:])

    letter = left.intersection(right).pop()

    res += (ord(letter) - 96) if letter.islower() else ((ord(letter) - 64) + 26)

    line = f.readline()

f.close()
print(res)
