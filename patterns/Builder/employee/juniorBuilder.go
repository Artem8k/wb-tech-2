package builder

type JuniorBuilder struct {
	name     string
	lastName string
	position string
	salary   int
}

func NewJuniorBuilder() *JuniorBuilder {
	return &JuniorBuilder{}
}

func (j *JuniorBuilder) SetName(name string) {
	j.name = name
}

func (j *JuniorBuilder) SetLastName(lastName string) {
	j.lastName = lastName
}

func (j *JuniorBuilder) SetPosition() {
	j.position = "Junior"
}

func (j *JuniorBuilder) SetSalary() {
	j.salary = 80000
}

func (j *JuniorBuilder) GetEmployee() Employee {
	return Employee{
		name:     j.name,
		lastName: j.lastName,
		position: j.position,
		salary:   j.salary,
	}
}
