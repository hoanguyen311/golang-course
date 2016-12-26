package main

import "fmt"

func main() {
	students := make([][]string, 0)

	student1 := make([]string, 4)

	student1[0] = "Thai Hoa"
	student1[1] = "IT"
	student1[2] = "8"
	student1[3] = "9"

	students = append(students, student1)

	fmt.Println(students)

	student2 := make([]string, 4)

	student2[0] = "Thi Cuc"
	student2[1] = "German"
	student2[2] = "10"
	student2[3] = "10"

	students = append(students, student2)

	fmt.Println(students)

	for _, s := range students {
		fmt.Println(s[0])
	}
}
