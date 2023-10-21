package solution

/*
This is the class that powers our mapping
app. It provides two methods for calculating the estimated time of
arrival (ETA) and the direction between two points.
Identify the problems in this implementation. Then, refactor the code
to use the state
*/

type DirectionService struct {
}

func (d DirectionService) GetEta(e etaService) int {
	return e.calculateEta()
}

func (d DirectionService) GetDirection(dir directionService) int {
	return dir.calculateDirection()
}
