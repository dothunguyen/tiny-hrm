package services

import (
	"container/list"
	"strconv"
	"tiny-hrm/repository"
)

// EmployeeService provides applicatin with business logic implementation
type EmployeeService struct {
	repo *repository.Repository
}

// NewEmployeeService constructor of type EmployeeService
func NewEmployeeService(repo *repository.Repository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

// GetOrganisation returns the hierarchy structure of company together
// with list of 'freelancers' employees, employees has invalid managerId
// and employess whose form a circle of management (a -> b, b->c, c -> a)
func (es *EmployeeService) GetOrganisation() (*repository.Organisation, error) {
	rs, err := es.repo.GetTopManagers()
	var topmanagers []repository.Employee
	if err == nil {
		for v := rs.Front(); v != nil; v = v.Next() {
			e := v.Value.(repository.Employee)
			es.getEmployeeFull(&e)
			topmanagers = append(topmanagers, e)
		}
		fs, _ := es.repo.GetFreelancers()
		ms, _ := es.repo.GetManagedByUnknown()
		circles, _ := es.GetCirclesOfManagement()
		return &repository.Organisation{Topmanagers: topmanagers, Freelancers: convertListToArray(fs), ManagedByUnkown: convertListToArray(ms), Circles: circles}, nil
	} else {
		return nil, err
	}
}

// GetAllEmployees returns all employee records in the system
func (es *EmployeeService) GetAllEmployees() ([]repository.Employee, error) {
	l, err := es.repo.GetAllEmployees()
	if err == nil {
		return convertListToArray(l), nil
	}
	return nil, err
}

// GetCirclesOfManagement returns the cirles of management if there  (a ->b, b -> c, c->a)
func (es *EmployeeService) GetCirclesOfManagement() ([][]repository.Employee, error) {
	rs, err := es.GetAllEmployees()
	if err == nil {
		var circles [][]repository.Employee
		for _, v := range rs {
			if isEmployeeAlreadyInCircles(circles, v) {
				continue
			}
			// start from this employee, add the manager into a list, stop when reach topmanager or hit the first one in list
			var l []repository.Employee
			if l, b := es.chainOfManagement(rs, l, v); b {
				circles = append(circles, l)
			}
		}
		return circles, nil
	}
	return nil, err
}

func (es *EmployeeService) chainOfManagement(all []repository.Employee, l []repository.Employee, e repository.Employee) ([]repository.Employee, bool) {

	l = append(l, e)

	if e.ManagerID == nil {
		return l, false
	}
	if len(l) > 0 && *e.ManagerID == l[0].ID {
		return l, true
	}
	var next repository.Employee
	for _, v := range all {
		if v.ID == *e.ManagerID {
			next = v
			break
		}
	}

	return es.chainOfManagement(all, l, next)
}

// getEmployeeFull query the employee managed by the given employee and fill in the 'managerOf' array
func (es *EmployeeService) getEmployeeFull(e *repository.Employee) {
	rs, err := es.repo.GetEmployeeManagedByID(strconv.Itoa(e.ID))
	if err == nil {
		for v := rs.Front(); v != nil; v = v.Next() {
			ve := v.Value.(repository.Employee)
			es.getEmployeeFull(&ve)
			e.ManagerOf = append(e.ManagerOf, ve)
		}
	}
}

// AddEmployee insert a new employee record into storage
func (es *EmployeeService) AddEmployee(employee repository.Employee) (*repository.Employee, error) {
	return es.repo.AddEmployee(employee)
}

// DeleteEmployee remove an employee record out of storage
func (es *EmployeeService) DeleteEmployee(id string) error {
	return es.repo.DeleteEmployee(id)
}

func isEmployeeAlreadyInCircles(circles [][]repository.Employee, e repository.Employee) bool {
	for _, a := range circles {
		for _, b := range a {
			if e.ID == b.ID {
				return true
			}
		}
	}
	return false
}

func convertListToArray(l *list.List) []repository.Employee {
	arr := make([]repository.Employee, l.Len())
	i := 0
	for v := l.Front(); v != nil; v = v.Next() {
		e := v.Value.(repository.Employee)
		arr[i] = e
		i = i + 1
	}

	return arr
}
