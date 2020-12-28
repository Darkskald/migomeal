package domain

import "time"

type MealPlan struct {
	Plan map[time.Time][]Recipe
}
