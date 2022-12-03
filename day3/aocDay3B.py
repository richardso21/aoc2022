f = open("./input.txt", "r")

line = f.readline()
res = 0
while line:
    letters = set()
    for _ in range(3):
        line = line.strip()
        currLetters = set(line)
        letters = letters.intersection(currLetters) if letters else currLetters
        line = f.readline()

    letter = letters.pop()
    res += (ord(letter) - 96) if letter.islower() else ((ord(letter) - 64) + 26)

f.close()
print(res)
