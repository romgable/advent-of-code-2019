import os
import math


def calculate(value):
    return math.floor(value / 3) - 2


path = os.path.dirname(os.path.realpath(__file__))
input = open(path + '\\input.txt', 'r')
sum = 0

for line in input:
    sum = sum + calculate(int(line))

print(sum)
