package Day06

import java.math.BigInteger
import println
import readInput

const val day = "Day06"

fun main() {
    fun ways(time: BigInteger, distance: BigInteger): Int {
        var count: Int = 0
        for (i in 1..time.toInt()) {
            if ((time - i.toBigInteger()) * i.toBigInteger() > distance) {
                count += 1
            }
        }
        return count
    }

    fun convertInputToList(input: String): List<BigInteger> {
        val inputString = input.split(":")[1]
        val numbers = Regex("[0-9]+").findAll(inputString).map({ it.value.toBigInteger() }).toList()

        return numbers
    }

    fun convertInputToBigNumber(input: String): BigInteger {
        val inputString = input.split(":")[1]
        val number = inputString.filter { it.isDigit() }.toBigInteger()

        return number
    }

    fun part1(input: List<String>): Int {
        val time = convertInputToList(input[0])
        val distance = convertInputToList(input[1])
        val result = ArrayList<Int>()

        for ((t, d) in time.zip(distance)) {
            result.add(ways(t, d))
        }

        var answer = 1
        for (ways in result) {
            answer *= ways
        }

        return answer
    }

    fun part2(input: List<String>): Int {
        val time = convertInputToBigNumber(input[0])
        val distance = convertInputToBigNumber(input[1])

        return ways(time, distance)
    }

    val input = readInput("Day06", day)
    part1(input).println()
    part2(input).println()
}
