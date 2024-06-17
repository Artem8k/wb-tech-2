package builder

type EmployeeBuilder interface {
	SetName(name string)
	SetLastName(name string)
	SetPosition()
	SetSalary(salary int)
	GetEmployee() Employee
}

func GetBuilder(employeePosition string) EmployeeBuilder {
	switch len(employeePosition) > 0 {
	case employeePosition == "Senior":
		return NewSeniorBuilder()

	case employeePosition == "Middle":
		return NewMiddleBuilder()

	case employeePosition == "Junior":
		return NewJuniorBuilder()
	}
	return nil
}
