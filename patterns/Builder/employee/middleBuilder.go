package builder

type MiddleBuilder struct {
	name     string
	lastName string
	position string
	salary   int
}

func NewMiddleBuilder() *MiddleBuilder {
	return &MiddleBuilder{}
}

func (m *MiddleBuilder) SetName(name string) {
	m.name = name
}

func (m *MiddleBuilder) SetLastName(lastName string) {
	m.lastName = lastName
}

func (m *MiddleBuilder) SetPosition() {
	m.position = "Middle"
}

func (m *MiddleBuilder) SetSalary() {
	m.salary = 130000
}

func (j *MiddleBuilder) GetEmployee() Employee {
	return Employee{
		name:     j.name,
		lastName: j.lastName,
		position: j.position,
		salary:   j.salary,
	}
}
