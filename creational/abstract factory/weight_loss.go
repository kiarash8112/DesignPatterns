package absfactory

/*
You’re building a weight loss app. Your app currently supports two kinds of
goals: Weight Loss and Build Muscle. There’s a plan to support other kinds of
goals like Strength Training, Getting Active and so on in the near future.
Depending on the goal the user selects,
they’ll see a different mean plan and workout routine.
What are the problems with this design? Refactor this design using the
abstract factory pattern.

*/

type MealPlan interface{}
type WorkoutPlan interface{}

type BuildMuscleMealPlan struct{}
type BuildMuscleWorkout struct{}

type WeighLossMealPlan struct{}
type WeightLossWorkout struct{}

type HomePage struct {
	workoutPlan WorkoutPlan
	mealPlan    MealPlan
}

func (h *HomePage) setGoal(goal string) {
	if goal == "BUILD_MUSCLE" {
		h.workoutPlan = BuildMuscleWorkout{}
		h.mealPlan = BuildMuscleMealPlan{}
	} else if goal == "WEIGHT_LOSS" {
		h.workoutPlan = WeightLossWorkout{}
		h.mealPlan = WeighLossMealPlan{}
	}
}
