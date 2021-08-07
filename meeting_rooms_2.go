package main

import "sort"

func minMeetingRooms(intervals [][]int) int {
	startTimes := make([]int, len(intervals))
	endTimes := make([]int, len(intervals))

	for currIntervalIdx, currInterval := range intervals {
		startTimes[currIntervalIdx] = currInterval[0]
		endTimes[currIntervalIdx] = currInterval[1]
	}

	sort.Ints(startTimes)
	sort.Ints(endTimes)

	startIdx := 0
	endIdx := 0
	roomsRequired := 0

	for startIdx < len(intervals) {
		if startTimes[startIdx] >= endTimes[endIdx] {
			endIdx++
			roomsRequired--
		}


		startIdx++
		roomsRequired++
	}

	return roomsRequired
}

