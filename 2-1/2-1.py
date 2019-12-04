import math
import os


def calculate(operation, input1, input2):
    if (operation == 1):
        return input1 + input2
    elif (operation == 2):
        return input1 * input2


path = os.path.dirname(os.path.realpath(__file__))
input = open(path + '\\input.txt', 'r')

with input:
    inputList = list(map(int, input.readline().split(',')))
    i = 0

    while (i < len(inputList)):
        operation = inputList[i]
        if (operation == 99):
            break
        input1Pos = inputList[i+1]
        input1Value = inputList[input1Pos]
        input2Pos = inputList[i+2]
        input2Value = inputList[input2Pos]
        result = calculate(operation, input1Value, input2Value)
        outputPos = inputList[i+3]
        inputList[outputPos] = result
        i += 4

    print(inputList[0])
