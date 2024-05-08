package repo

import (
	"database/sql"
	"log"

	"github.com/akshay-singla/employee-crud-api/employee/module"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	r := &Repository{
		db: db,
	}
	return r
}

func (r *Repository) Create(emp module.Employee) error {
	_, err := r.db.Exec("INSERT INTO employees (name, position, salary) VALUES ($1, $2, $3)",
		emp.Name, emp.Position, emp.Salary)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Update(emp module.Employee) error {
	_, err := r.db.Exec("UPDATE employees SET name = $1, position = $2, salary = $3 WHERE id = $4",
		emp.Name, emp.Position, emp.Salary, emp.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) GetEmployeeByID(id string) (*module.Employee, error) {
	emp := module.Employee{}

	err := r.db.QueryRow("SELECT id, name, position, salary FROM employees WHERE id = $1", id).
		Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &emp, nil
}

func (r *Repository) ListEmployee(pageSize, offset int) ([]module.Employee, error) {
	employees := make([]module.Employee, 0)
	rows, err := r.db.Query("SELECT id, name, position, salary FROM employees ORDER BY id LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var emp module.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary); err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	return employees, nil
}

// CountEmployees counts total number of employee records
func (r *Repository) CountEmployees() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM employees").Scan(&count)
	if err != nil {
		log.Println(err)
		return count, err
	}

	return count, nil
}
