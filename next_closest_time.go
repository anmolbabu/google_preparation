package main

import (
	"fmt"
	"math"
	"strings"
)

func getTimeStr(epoch int) string {
	hours := epoch/60
	mins := epoch % 60
	return normaliseTime(fmt.Sprintf("%d:%d", hours, mins))
}

func normaliseTime(timeStr string) string {
	timeParts := strings.Split(timeStr, ":")

	hours := timeParts[0]
	if len(hours) < 2 {
		hours = fmt.Sprintf("0%s", hours)
	}

	mins := timeParts[1]
	if len(mins) < 2 {
		mins = fmt.Sprintf("0%s", mins)
	}

	return fmt.Sprintf("%s:%s", hours, mins)
}

func getDigits(normalisedTime string) []int {
	digits := make([]int, 4)

	digitIterator := 0
	for idx := 0; idx < len(normalisedTime); idx++ {
		if normalisedTime[idx] == ':' {
			continue
		}
		digits[digitIterator] = int(normalisedTime[idx] - '0')
		digitIterator++
	}

	return digits
}

func getEpochTimeFromTime(str string) int {
	timeDigits := getDigits(str)
	return getEpochTimeFromDigits(timeDigits[0], timeDigits[1], timeDigits[2], timeDigits[3])
}

func getEpochTimeFromDigits(h1 int, h2 int, m1 int, m2 int) int {
	hours := (h1 * 10) + h2
	mins := (m1 * 10) + m2

	return (hours * 60) + mins
}

func isValidTime(h1 int, h2 int, m1 int, m2 int) bool {
	hours := (h1 * 10) + h2
	mins := (m1 * 10) + m2

	return (hours < 24) && (mins < 60)
}

func nextClosestTime(time string) string {
	closestTime := normaliseTime(time)

	timeDigits := getDigits(time)
	elapsed := float64(24 * 60)

	for _, h1 := range timeDigits{
		for _, h2 := range timeDigits {
			for _, m1 := range timeDigits {
				for _, m2 := range timeDigits {
					if isValidTime(h1, h2, m1, m2) {
						currTime := getEpochTimeFromTime(time)
						candidateTime := getEpochTimeFromDigits(h1, h2, m1, m2)

						currTimeElapsed := getTimeDiff(candidateTime, currTime)

						if (currTimeElapsed > 0) && (currTimeElapsed < elapsed) {
							closestTime = getTimeStr(candidateTime)
							elapsed = currTimeElapsed
						}
					}
				}
			}
		}
	}

	return closestTime
}

func getTimeDiff(candidateTime int, currTime int) float64 {
	currTimeElapsed := math.Mod(float64(candidateTime-currTime), 24*40)
	if currTimeElapsed < 0 {
		currTimeElapsed = (24 * 60) + currTimeElapsed
	}
	return currTimeElapsed
}

func main() {
	fmt.Println(nextClosestTime("19:34"))
	fmt.Println(nextClosestTime("23:59"))
	fmt.Println(nextClosestTime("00:00"))
	fmt.Println(nextClosestTime("01:00"))
}