package Day01

import readInput
import println

fun main() {
    val day = "Day01"

    val words: Map<String, Int> = mapOf(
        "one" to 1,
        "two" to 2,
        "three" to 3,
        "four" to 4,
        "five" to 5,
        "six" to 6,
        "seven" to 7,
        "eight" to 8,
        "nine" to 9
    )

    fun getNumber(input: String): Int =
        "${input.first { it.isDigit() }}${input.last { it.isDigit() }}".toInt()

    fun String.isWordNumberAt(startingAt: Int): List<String> =
        (3..5).map { len ->
            substring(startingAt, (startingAt + len).coerceAtMost(length))
        }

    fun part1(input: List<String>): Int =
        input.sumOf { getNumber(it) }
    

    fun part2(input: List<String>): Int =
        input.sumOf { row ->
            getNumber(
                row.mapIndexedNotNull { index, n ->
                    if (n.isDigit()) n
                    else
                        row.isWordNumberAt(index).firstNotNullOfOrNull { num: String ->
                            words[num]
                        }
                }.joinToString()
            )
        }

    val input = readInput("Day01", day)
    part1(input).println()
    part2(input).println()
}
