#!/bin/bash
year=$(date +"%Y")
day=$(date +"%-d")
SESSION=$(cat session.txt)
useragent="Advent of Code by https://github.com/t0nyandre/advent-of-code"

while getopts y:d:o: flag
do
    case "${flag}" in
        y) year=${OPTARG};;
        d) day=${OPTARG};;
        o) output=${OPTARG};;
    esac
done

if [ -z "$output" ]
then
    echo "-o (output) is required"; exit $ERRCODE
else
    curl https://adventofcode.com/${year}/day/${day}/input --cookie "session=$SESSION" -A "$useragent"  > "$output/Day$(printf "%02d" $day).txt"
fi
