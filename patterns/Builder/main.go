package main

// Строитель паттерн проектирования, который позволяет создавать сложные объекты пошагово.
// Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов

// Паттерн позволяет создавать сложные объекты используя простые объекты и поэтапный подход
// Может быть полезен для создания иммутабельных объектов с помощью одного и того же процесса построения объекта

// Из минусов: Увеличение количества кода из-за необходимости создания отдельных строителей для каждого типа объекта

import (
	"builder/director"
	builder "builder/employee"
	"fmt"
)

func main() {
	jBuilder := builder.GetBuilder("Junior")
	mBuilder := builder.GetBuilder("Middle")
	sBuilder := builder.GetBuilder("Senior")

	j := director.NewDirector(&jBuilder)
	m := director.NewDirector(&mBuilder)
	s := director.NewDirector(&sBuilder)

	fmt.Println(j.BuildEmployee("ArtemJ", "ArtemJ"))
	fmt.Println(m.BuildEmployee("ArtemM", "ArtemM"))
	fmt.Println(s.BuildEmployee("ArtemS", "ArtemS"))
}
