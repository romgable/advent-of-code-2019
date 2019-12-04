import math
import os


def calculate(operation, input1, input2):
    if (operation == 1):
        return input1 + input2
    elif (operation == 2):
        return input1 * input2


path = os.path.dirname(os.path.realpath(__file__))
input = open(path + '\\input.txt', 'r')
resultToFind = 19690720
result = 0
noun = 0
verb = 0

with input:
    inputListOriginal = list(map(int, input.readline().split(',')))

    while (noun <= 99 and verb <= 99):
        result = 0
        inputList = inputListOriginal.copy()

        if (noun == 99):
            verb = verb + 1
            noun = 0
        else:
            noun = noun + 1

        inputList[1] = noun
        inputList[2] = verb
        i = 0

        while (i < len(inputList)):
            operation = inputList[i]
            if (operation == 99):
                break
            input1Pos = inputList[i+1]
            input1Value = inputList[input1Pos]
            input2Pos = inputList[i+2]
            input2Value = inputList[input2Pos]
            outputPos = inputList[i+3]
            inputList[outputPos] = calculate(operation, input1Value, input2Value)
            i += 4
        
        if (inputList[0] == resultToFind):
            result = 100 * noun + verb
            break

    print(result)
