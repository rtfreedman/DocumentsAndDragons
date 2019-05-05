package store

// Ability represents an ability belonging to a character
type Ability struct {
	ID                int
	Name              string
	Description       string
	Charges           int
	MaxCharges        int
	RechargeCondition int
	Use               Filter
}
