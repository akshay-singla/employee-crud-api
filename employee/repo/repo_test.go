package repo_test

import (
	"database/sql"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/akshay-singla/employee-crud-api/employee/module"
	"github.com/akshay-singla/employee-crud-api/employee/repo"

	"errors"

	"github.com/stretchr/testify/assert"
)

// Custom error type for user not found
var ErrUserNotFound = errors.New("sql: no rows in result set")

func TestGetEmployeeByID(t *testing.T) {
	testCases := []struct {
		name          string
		id            string
		expectedEmp   *module.Employee // Define your Employee struct type
		expectedError error
	}{
		{
			name:          "No rows found",
			id:            "1",
			expectedEmp:   nil,
			expectedError: ErrUserNotFound,
		},
		{
			name: "Data found",
			id:   "1",
			expectedEmp: &module.Employee{
				ID:       1,
				Name:     "John Doe",
				Position: "Manager",
				Salary:   50000,
			},
			expectedError: nil,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Initialize mock database
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error creating mock database: %s", err)
			}
			defer db.Close()

			// Create repository with mock database
			repo := repo.NewRepository(db)

			// Define mock query and return appropriate rows or error
			if tc.expectedEmp == nil {
				mock.ExpectQuery("SELECT id, name, position, salary FROM employees WHERE id = ?").
					WithArgs(tc.id).
					WillReturnError(sql.ErrNoRows)
			} else {
				expectedRows := sqlmock.NewRows([]string{"id", "name", "position", "salary"}).
					AddRow("1", "John Doe", "Manager", 50000)

				mock.ExpectQuery("SELECT id, name, position, salary FROM employees WHERE id = ?").
					WithArgs(tc.id).
					WillReturnRows(expectedRows)
			}

			// Call the business logic
			emp, err := repo.GetEmployeeByID(tc.id)

			// Assertions
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedEmp, emp)

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
