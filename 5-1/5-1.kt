import java.io.File
import java.io.InputStream

enum class Operation(val value: Int, val parameters: Int) {
    ABORT(99, 0),
    ADDITION(1, 3),
    MULTIPLICATION(2, 3),
    INPUT(3, 1),
    OUTPUT(3, 1)
}

enum class Mode(val value: Int) {
    POSITION(0),
    IMMEDIATE(1)
}

fun main(args: Array<String>) {
    val fileName = System.getProperty("user.dir") + "\\input.txt"
    val inputStream: InputStream = File(fileName).inputStream()

    inputStream.bufferedReader().useLines { lines ->
        lines.forEach {
            var inputs: MutableList<Int> = it.split(',').map { x -> Integer.valueOf(x) }.toMutableList()
            var i = 0
            while (i < inputs.size) {
                var value = Integer.valueOf(inputs.get(i))
                if (value == Operation.ABORT.value) break

                var modes = listOf(value - 1000, value / 100)

                when (value) {
                    Operation.ADDITION.value -> {
                        calculate(inputs, i, Operation.ADDITION, modes)
                        i += Operation.ADDITION.parameters + 1
                    }
                    Operation.MULTIPLICATION.value -> {
                        calculate(inputs, i, Operation.MULTIPLICATION, modes)
                        i += Operation.MULTIPLICATION.parameters + 1
                    }
                }
            }

            println(inputs.get(0))
        }
    }
}

fun calculate(inputs: MutableList<Int>, i: Int, op: Operation, modes: List<Int>) {
    var parameters: List<Int> = getParameters(inputs, i, op)
    when (op) {
        Operation.ADDITION -> inputs[parameters.get(2)] = getValue(inputs, parameters, 0, modes.get(0)) + getValue(inputs, parameters, 1, modes.get(1))
        Operation.MULTIPLICATION -> inputs[parameters.get(2)] = getValue(inputs, parameters, 0, modes.get(0)) * getValue(inputs, parameters, 1, modes.get(1))
        Operation.ABORT -> return
    }
}

fun getValue(inputs: List<Int>, parameters: List<Int>, number: Int, mode: Int): Int {
    when(mode) {
        Mode.POSITION.value -> return inputs.get(parameters.get(number))
        Mode.IMMEDIATE.value -> return parameters.get(number)
    }
    return 0
}

fun getParameters(inputs: List<Int>, i: Int, op: Operation) = inputs.filterIndexed { index, _ -> index in i + 1..i + op.parameters }
