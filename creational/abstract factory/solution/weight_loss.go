package solution

type Plan interface {
	CreateWorkoutPlan() WorkoutPlan
	CreateMealPlan() MealPlan
}

type MealPlan interface{}
type WorkoutPlan interface{}

type BuildMuscleMealPlan struct{}
type BuildMuscleWorkout struct{}

type WeighLossMealPlan struct{}
type WeightLossWorkout struct{}

type BuildMuscle struct{}

func (BuildMuscle) CreateWorkoutPlan() WorkoutPlan {
	return BuildMuscleWorkout{}
}

func (BuildMuscle) CreateMealPlan() MealPlan {
	return BuildMuscleMealPlan{}
}

type WeightLoss struct{}

func (WeightLoss) CreateWorkoutPlan() WorkoutPlan {
	return WeightLossWorkout{}
}

func (WeightLoss) CreateMealPlan() MealPlan {
	return WeighLossMealPlan{}
}

type HomePage struct {
	workoutPlan WorkoutPlan
	mealPlan    MealPlan
}

func (h *HomePage) setGoal(plan Plan) {
	h.workoutPlan = plan.CreateWorkoutPlan()
	h.mealPlan = plan.CreateMealPlan()
}
