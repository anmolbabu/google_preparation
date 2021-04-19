package main

import (
	"fmt"
	"math"
)

func getOccupiedSeats(seats []int) []int {
	occupiedSeats := []int{}

	for seatIdx, seat := range seats {
		seatOccupied := (seat == 1)
		if seatOccupied {
			occupiedSeats = append(occupiedSeats, seatIdx)
		}
	}

	return occupiedSeats
}

func maxDistToClosest(seats []int) int {
	maxDistance := 0
	occupiedSeats := getOccupiedSeats(seats)

	prevOccupiedSeat := math.MinInt32

	nxtOccupiedSeat := math.MaxInt32
	nxtOccupiedSeatIdx := 0
	if len(occupiedSeats) > 0 {
		nxtOccupiedSeat = occupiedSeats[nxtOccupiedSeatIdx]
	}

	for seatIdx, seat := range seats {
		closestSeatDistance := -1
		seatOccupied := (seat == 1)
		if !seatOccupied {
			fmt.Print(prevOccupiedSeat, seatIdx, nxtOccupiedSeat)
			closestSeatDistance = min(seatIdx - prevOccupiedSeat, nxtOccupiedSeat - seatIdx)
			fmt.Println("\t\t\t", closestSeatDistance)
		} else {
			prevOccupiedSeat = nxtOccupiedSeat
			nxtOccupiedSeatIdx++
			if len(occupiedSeats) > nxtOccupiedSeatIdx {
				nxtOccupiedSeat = occupiedSeats[nxtOccupiedSeatIdx]
			} else {
				nxtOccupiedSeat = math.MaxInt32
			}
		}

		if closestSeatDistance != -1 {
			maxDistance = max(maxDistance, closestSeatDistance)
		}
	}

	return maxDistance
}

