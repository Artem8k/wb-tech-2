package main

import (
	builder "builder/employee"
	"fmt"
)

func main() {
	jBuilder := builder.GetBuilder("Junior")
	mBuilder := builder.GetBuilder("Middle")
	sBuilder := builder.GetBuilder("Senior")

	jBuilder.SetName("ArtemJ")
	jBuilder.SetLastName("ArtemJ")
	jBuilder.SetPosition()
	jBuilder.SetSalary(80000)

	mBuilder.SetName("ArtemM")
	mBuilder.SetLastName("ArtemM")
	mBuilder.SetPosition()
	mBuilder.SetSalary(130000)

	sBuilder.SetName("ArtemM")
	sBuilder.SetLastName("ArtemM")
	sBuilder.SetPosition()
	sBuilder.SetSalary(200000)

	fmt.Println(jBuilder.GetEmployee())
	fmt.Println(mBuilder.GetEmployee())
	fmt.Println(sBuilder.GetEmployee())
}
