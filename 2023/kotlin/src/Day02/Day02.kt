package Day02

import println
import readInput

const val day: String = "Day02"

const val MAX_RED: Int = 12
const val MAX_BLUE: Int = 14
const val MAX_GREEN: Int = 13

fun main() {
    fun getMatches(input: String): Pair<Int, Sequence<MatchResult>> {
        var gameNum = input.split(": ")[0].split(" ")[1].toInt()
        val pattern = "\\b(\\d+) \\b(red|blue|green)"
        val matches = Regex(pattern).findAll(input)
        return Pair(gameNum, matches)
    }

    fun checkPossible(matches: Sequence<MatchResult>): Boolean {
        for (match in matches) {
            val num = match.value.split(" ")[0].toInt()
            val cube = match.value.split(" ")[1]
            if ((cube.equals("green")) && (num.compareTo(MAX_GREEN) > 0))
                return false
            if ((cube.equals("red")) && (num.compareTo(MAX_RED) > 0))
                return false
            if ((cube.equals("blue")) && (num.compareTo(MAX_BLUE) > 0))
                return false
        }
        return true
    }

    fun checkLowestPossible(matches: Sequence<MatchResult>): Int {
        var min_green: Int = 0
        var min_red: Int = 0
        var min_blue: Int = 0

        for (match in matches) {
            val num = match.value.split(" ")[0].toInt()
            val cube = match.value.split(" ")[1]
            if ((cube.equals("green")) && (num > min_green))
                min_green = num
            if ((cube.equals("red")) && (num > min_red))
                min_red = num
            if ((cube.equals("blue")) && (num > min_blue))
                min_blue = num
        }
        return min_green * min_red * min_blue
    }

    fun part1(input: List<String>): Int {
        var score: Int = 0
        input.forEach({
            val (gameNum, matches) = getMatches(it)
            if (checkPossible(matches)) {
                score += gameNum
            }
        })
        return score
    }

    fun part2(input: List<String>): Int {
        var power: Int = 0
        input.forEach({
            val (_, matches) = getMatches(it)
            power += checkLowestPossible(matches)
        })
        return power
    }

    val input = readInput("Day02", day)
    part1(input).println()
    part2(input).println()
}
