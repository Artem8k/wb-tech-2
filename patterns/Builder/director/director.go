package director

import builder "builder/employee"

type Director struct {
	builder builder.EmployeeBuilder
}

func (d *Director) newDirector(b *builder.EmployeeBuilder) *Director {
	return &Director{
		builder: *b,
	}
}

func (d *Director) buildEmployee(name, lastName string, salary int) builder.Employee {
	d.builder.SetName(name)
	d.builder.SetLastName(lastName)
	d.builder.SetPosition()
	d.builder.SetSalary(salary)
	return d.builder.GetEmployee()
}
