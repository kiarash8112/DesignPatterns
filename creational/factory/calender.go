package factory

import (
	"fmt"
	"time"
)

/*
You’re building a framework for scheduling tasks.The Scheduler struct is tightly
coupled to the Calendar struct which represents the Gregorian calendar. With
this design, this framework cannot be effectively used in Arabian countries.
Use the factory method to add exibility to this design.

*/

type Event struct{}
type Calender struct{}

func (Calender) addEvent(event Event, date time.Time) {
	fmt.Println("Adding an event on the given date.")
}

type Scheduler struct {
	calender Calender
}

func (s *Scheduler) schdule(event Event) {
	today := time.Now()
	s.calender.addEvent(event, today)
}
