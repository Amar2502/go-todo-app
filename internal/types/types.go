package types

type Todo struct {
	ID    int
	Task  string `validate:"required"` 
	StartTime string `validate:"required"`
}