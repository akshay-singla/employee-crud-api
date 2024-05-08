package module

// Employee struct represents an employee record
type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Position string  `json:"position" binding:"required"`
	Salary   float64 `json:"salary" binding:"required"`
}
