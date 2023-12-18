package entities

type TaskWithAssignedUser struct {
	Task Task `json:"Task"`
	User User `json:"User"`
}
