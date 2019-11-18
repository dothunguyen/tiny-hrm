package repository

import (
	"container/list"
	"database/sql"
	"fmt"
	"io/ioutil"
)

// Employee contains employee details and list of directly managed employee
type Employee struct {
	ID        int        `form:"id" json:"id"`
	Name      string     `form:"name" json:"name"`
	ManagerID *int       `form:"managerId" json:"managerId"`
	ManagerOf []Employee `json:"managerOf"`
}

// Organisation hold the structure & all freelancers & managed-by-unknown
type Organisation struct {
	Topmanagers     []Employee   `json:"topmanagers"`
	Freelancers     []Employee   `json:"freelancers"`
	ManagedByUnkown []Employee   `json:"managedByUnknown"`
	Circles         [][]Employee `json:"circles"`
}

// Repository provides app with an interface to database for persistent
type Repository struct {
	db *sql.DB
}

// NewRepository create a new instance of Repository
func NewRepository() (*Repository, error) {
	sqlitedb, err := sql.Open("sqlite3", ".run/tiny-hrm.db")

	if err == nil {
		return &Repository{db: sqlitedb}, nil
	}
	return nil, err
}

// InitDB create the schema
func (con *Repository) InitDB() {
	// create data tables
	statement, _ := con.db.Prepare("CREATE TABLE IF NOT EXISTS employee (id INTEGER UNIQUE, name TEXT, managerId INTEGER)")
	statement.Exec()
	sql, err := ioutil.ReadFile("./src/resources/initdb.sql")
	if err == nil {
		statement, _ = con.db.Prepare(string(sql))
		statement.Exec()
	}
}

// GetAllEmployees returns a list of all employees.
func (con *Repository) GetAllEmployees() (*list.List, error) {
	return con.query("SELECT * FROM employee")
}

// GetTopManagers return the top managers who are not managed by anyone but directly manage some employees
func (con *Repository) GetTopManagers() (*list.List, error) {
	sql := "SELECT id, name, managerId FROM employee e1 WHERE e1.managerId IS NULL AND e1.id IN (SELECT managerId FROM employee e2 WHERE e2.managerId = e1.id)"
	return con.query(sql)
}

// GetEmployeeManagedByID return the list of employees directly managed by the employee with the provided id
func (con *Repository) GetEmployeeManagedByID(id string) (*list.List, error) {
	sql := fmt.Sprintf("SELECT id, name, managerId FROM employee e1 WHERE e1.managerId = %s AND e1.id <> e1.managerId", id)
	return con.query(sql)
}

// GetFreelancers return the list of employees that are not managed by anyone and don't manage anyone as well
func (con *Repository) GetFreelancers() (*list.List, error) {
	sql := "SELECT id, name, managerId FROM employee e1 WHERE e1.managerId IS NULL AND e1.id NOT IN (SELECT managerId FROM employee e2 WHERE e2.managerId = e1.id)"
	return con.query(sql)
}

// GetManagedByUnknown return the list of employees that are not managed by anyone (or by themselves) and don't manage anyone as well.
func (con *Repository) GetManagedByUnknown() (*list.List, error) {
	sql := "SELECT id, name, managerId FROM employee e1 WHERE e1.managerId NOT IN (SELECT id FROM employee) OR e1.id = e1.managerId"
	return con.query(sql)
}

// AddEmployee insert an employee record into sqlite database
func (con *Repository) AddEmployee(e Employee) (*Employee, error) {
	var sql = "INSERT INTO employee (id, name, managerId) VALUES (?, ?, ?)"
	statement, _ := con.db.Prepare(string(sql))

	_, err := statement.Exec(e.ID, e.Name, e.ManagerID)
	if err == nil {
		return &e, nil
	} else {
		return nil, err
	}
}

// DeleteEmployee remove an employee record from sqlite database
func (con *Repository) DeleteEmployee(id string) error {
	var sql = "DELETE FROM employee WHERE id = ?"
	statement, _ := con.db.Prepare(string(sql))

	_, err := statement.Exec(id)
	return err
}

func (con *Repository) query(sql string) (*list.List, error) {
	rows, err := con.db.Query(sql)
	if err == nil {
		l := list.New()
		for rows.Next() {
			var id int
			var name string
			var managerID *int
			rows.Scan(&id, &name, &managerID)
			l.PushBack(Employee{ID: id, Name: name, ManagerID: managerID})
		}
		return l, nil
	} else {
		return nil, err
	}
}

// Close db connection, free resource
func (con *Repository) Close() {
	con.db.Close()
}
