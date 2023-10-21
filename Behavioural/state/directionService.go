package state

import "fmt"

/*
This is the class that powers our mapping
app. It provides two methods for calculating the estimated time of
arrival (ETA) and the direction between two points.
Identify the problems in this implementation. Then, refactor the code
to use the state
*/

type DirectionService struct {
	travelMode string
}

func (d DirectionService) getEta() int {
	if d.travelMode == "DRIVING" {
		fmt.Println("Calculating ETA (driving)")
		return 1
	} else if d.travelMode == "BICYCLING" {
		fmt.Println("Calculating ETA (bicycling)")
		return 2
	} else if d.travelMode == "TRANSIT" {
		fmt.Println("Calculating ETA (transit)")
		return 3
	} else {
		fmt.Println("Calculating ETA (walking)")
		return 4
	}
}

func (d DirectionService) getDirection() int {
	if d.travelMode == "DRIVING" {
		fmt.Println("Calculating Direction (driving)")
		return 1
	} else if d.travelMode == "BICYCLING" {
		fmt.Println("Calculating Direction (bicycling)")
		return 2
	} else if d.travelMode == "TRANSIT" {
		fmt.Println("Calculating Direction (transit)")
		return 3
	} else {
		fmt.Println("Calculating Direction (walking)")
		return 4
	}
}
