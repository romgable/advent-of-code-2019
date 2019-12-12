import java.io.File
import java.io.InputStream

fun main(args: Array<String>) {
    val fileName = System.getProperty("user.dir") + "\\input.txt"
    val inputStream: InputStream = File(fileName).inputStream()

    val directOrbits = mutableMapOf<String, MutableList<String>>()
    
    inputStream.bufferedReader().useLines { lines ->
        lines.forEach {
            var line = it
            var orbits = line.split(')').toList()
            
            var list = directOrbits.getOrDefault(orbits[0], mutableListOf<String>())
            if (!list.contains(orbits[1])) {
                list.add(orbits[1])
            }
            directOrbits.put(orbits[0], list)
        }
    }

    var sum: Int = 0;
    directOrbits.forEach {k, _ -> println(getIndirectOrbit(k, directOrbits, 0))}

    // getIndirectOrbits("C", orbits, 0)

    println(sum)
}

fun getIndirectOrbit(planet: String, orbits: Map<String, MutableList<String>>, indirectSum: Int): Int {
    for ((k, v) in orbits) {
        if (v.contains(planet)) {
            println("$planet orbits in $k")
            return getIndirectOrbit(k, orbits, indirectSum + 1)
        }
    }
    return indirectSum
}
