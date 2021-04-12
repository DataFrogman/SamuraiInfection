def int2base(n, b):
    if n == 0:
        return [0]
    digits = []
    while n:
        digits.append(int(n % b))
        n //= b
    return digits[::-1]

def listToString(s):
    str1 = ""
    for e in s:
        if e == 0:
            str1 += "A"
        elif e == 1:
            str1 += "C"
        elif e == 2:
            str1 += "G"
        else:
            str1 += "T"
    return str1

f = open("binary", "r")
f2 = open("base4", "w")
for x in f.readlines():
    if len(x) == 49:
        one = x[0:8]
        two = x[8:16]
        three = x[16:24]
        four = x[24:32]
        five = x[32:40]
        six = x[40:48]

        one = int(one, 2)
        two = int(two, 2)
        three = int(three, 2)
        four = int(four, 2)
        five = int(five, 2)
        six = int(six, 2)

        one = int2base(one, 4)
        two = int2base(two, 4)
        three = int2base(three, 4)
        four = int2base(four, 4)
        five = int2base(five, 4)
        six = int2base(six, 4)

        f2.write(listToString(one) + " ")
        f2.write(listToString(two) + " ")
        f2.write(listToString(three) + " ")
        f2.write(listToString(four) + " ")
        f2.write(listToString(five) + " ")
        f2.write(listToString(six) + " ")
    else:
        one = x[0:8]
        two = x[8:16]

        one = int(one, 2)
        two = int(two, 2)

        one = int2base(one, 4)
        two = int2base(two, 4)

        f2.write(listToString(one) + " ")
        f2.write(listToString(two) + " ")

f.close()
f2.close()
    