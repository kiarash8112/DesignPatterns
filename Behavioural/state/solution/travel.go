package solution

import "fmt"

type etaService interface {
	calculateEta() int
}

type directionService interface {
	calculateDirection() int
}

var _ etaService = Driving{}
var _ etaService = BICYCLING{}
var _ etaService = TRANSIT{}
var _ etaService = WALKING{}

var _ directionService = Driving{}
var _ directionService = BICYCLING{}
var _ directionService = TRANSIT{}
var _ directionService = WALKING{}

type Driving struct {
}

func (Driving) calculateEta() int {
	fmt.Println("Calculating ETA (driving)")
	return 1
}

func (Driving) calculateDirection() int {
	fmt.Println("Calculating Direction (driving)")
	return 1
}

type BICYCLING struct {
}

func (BICYCLING) calculateEta() int {
	fmt.Println("Calculating ETA (bicycling)")
	return 2
}

func (BICYCLING) calculateDirection() int {
	fmt.Println("Calculating Direction (bicycling)")
	return 2
}

type TRANSIT struct {
}

func (TRANSIT) calculateEta() int {
	fmt.Println("Calculating ETA (transit)")
	return 3
}

func (TRANSIT) calculateDirection() int {
	fmt.Println("Calculating Direction (transit)")
	return 3
}

type WALKING struct {
}

func (WALKING) calculateEta() int {
	fmt.Println("Calculating ETA (walking)")
	return 4
}

func (WALKING) calculateDirection() int {
	fmt.Println("Calculating Direction (walking)")
	return 4
}
