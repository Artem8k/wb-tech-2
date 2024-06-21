package director

import builder "builder/employee"

type Director struct {
	builder builder.EmployeeBuilder
}

func NewDirector(b *builder.EmployeeBuilder) *Director {
	return &Director{
		builder: *b,
	}
}

func (d *Director) BuildEmployee(name, lastName string) builder.Employee {
	d.builder.SetName(name)
	d.builder.SetLastName(lastName)
	d.builder.SetPosition()
	d.builder.SetSalary()
	return d.builder.GetEmployee()
}
