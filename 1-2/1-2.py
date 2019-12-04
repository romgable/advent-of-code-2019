import math
import os


def calculate(value):
    result = math.floor(value / 3) - 2
    if result <= 0:
        return 0
    else:
        return result


path = os.path.dirname(os.path.realpath(__file__))
input = open(path + '\\input.txt', 'r')
sum = 0

for line in input:
    fuelInput = int(line)
    fuelSum = 0
    while fuelInput > 0:
        fuelResult = calculate(fuelInput)
        fuelSum = fuelSum + fuelResult
        fuelInput = fuelResult
    sum = sum + fuelSum

print(sum)
