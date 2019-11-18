# Employee Hierarchical Relationship 

Transform the tabular employee data of a small company to a hierarchical structure follow this rule

- CEO of the company doesn't have a manager.
- Some of employees may have no manager and also don't manage anyone
- Some managerId are invalid employeeId 

# The Solution

Find the top managers, who has no manager and should manage someone.
Use top managers as the root, traverse back to the fill up the tree structure of `managers -> [employees]`.
Can have multiple top managers (in case company need to replace the big boss)

Search the data for employees that belong to the following groups

- 'Freelancers', who have no manager and don't manage anyone or who are self-managed (managerId is the same as employeeId)
- 'ManagedByUnkown', who has a manager that is not an employee (managerId is not a valid employeeId)
- Circles of management, those make a circle as below

```
b managed a, c managed b, a managed c
```

# Tech Stack

An API for Employee management

```
GET `/api/v1/employees/` returns all employee records
```
```
POST `/api/v1/employees/` for adding new employee
```
```
DELETE `/api/v1/employees/:id` for remove employee(with id)
```
```
GET `/api/v1/employees/org` returns the structure of the company,
```

A simple web page to consume the API and display company Employee Directory

## React for front-end

React is used for front-end for fast rendering

## Go with gin & SQLite3 for API

RDBMS is chosen to model the relationship between employee.
Sqlite is chosen for the scope of this project.
A graph database (DGraph) has been tried but failed to store the invalid relationship (managerId is not a valid employeeId)

# Build Docker

You will need docker to build and run

```
docker build -t tiny-hrm .
```

# Run

```
docker run -it -p 3000:3000 tiny-hrm
```

# Test

Open http://localhost:3000 to view the employee directory and hierarchical view

Add more employee
```
curl -d '{"id": 530, "name": "test", "managerId": 150}' http://localhost:3000/api/v1/employees/
```
Delete a selected employee record
```
curl -X DELETE http://localhost:3000/api/v1/employees/:id 
```

Refresh page to view the changes on organisation structure.