package builder

type SeniorBuilder struct {
	name     string
	lastName string
	position string
	salary   int
}

func NewSeniorBuilder() *SeniorBuilder {
	return &SeniorBuilder{}
}

func (s *SeniorBuilder) SetName(name string) {
	s.name = name
}

func (s *SeniorBuilder) SetLastName(lastName string) {
	s.lastName = lastName
}

func (s *SeniorBuilder) SetPosition() {
	s.position = "Senior"
}

func (s *SeniorBuilder) SetSalary(salary int) {
	s.salary = salary
}

func (j *SeniorBuilder) GetEmployee() Employee {
	return Employee{
		name:     j.name,
		lastName: j.lastName,
		position: j.position,
		salary:   j.salary,
	}
}
