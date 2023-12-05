package aoc2023

import println
import readInput

fun main() {
    fun getNumbers(input: String): Pair<ArrayList<String>, ArrayList<String>> {
        val card = input.split(":")[1].split("|")
        val winningNumbers = ArrayList(card[0].split(" "))
        val numbersOnCard = ArrayList(card[1].split(" "))
        winningNumbers.removeIf({ it == "" || it == " " })
        numbersOnCard.removeIf({ it == "" || it == " " })
        return Pair(winningNumbers, numbersOnCard)
    }

    fun calculateScore(winning: ArrayList<String>, numbers: ArrayList<String>): Pair<Int, Int> {
        var score: Int = 0
        var matches: Int = 0
        numbers.forEach({
            if (winning.contains(it)) {
                if (score == 0) {
                    score = 1
                } else {
                    score *= 2
                }
                matches++
            }
        })

        return Pair(score, matches)
    }

    fun part1(input: List<String>): Int {
        var score: Int = 0
        input.forEach({
            val (winning, numbers) = getNumbers(it)
            val (calScore, _) = calculateScore(winning, numbers)
            score += calScore
        })
        return score
    }

    fun part2(input: List<String>): Int {
        val copies = IntArray(input.size) { 1 }
        for ((index, line) in input.withIndex()) {
            val (winning, numbers) = getNumbers(line)
            val (_, matches) = calculateScore(winning, numbers)
            if (matches > 0) {
                for (i in 1..matches) {
                    if (index + i >= copies.size) break
                    copies[index + i] += copies[index]
                }
            }
        }
        return copies.sum()
    }

    val input = readInput("Day04")
    part1(input).println()
    part2(input).println()
}
