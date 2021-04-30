package main

import "fmt"

func main() {
	// We used an array for student, you can use any db for this operation
	var students []Student
	students = CreateStudent(10, "Alex", "De Souza", "Fenerbahçe", students)
	students = CreateStudent(11, "Neymar", "Jr", "PSG", students)
	students = CreateStudent(8, "Kevin", "De Bruyne", "M.City", students)
	fmt.Println(students)

	// selectedStudent := GetStudentById(10, students)
	// fmt.Println("Selected students information is ", selectedStudent)

	// var s1 Student
	// s1.Id = 7
	// s1.FirstName = "Cristiano"
	// s1.LastName = "Ronaldo"
	// s1.ClassName = "Juventus"

	// students = UpdateStudent(10, s1, students)
	// fmt.Println(students)

	students = DeleteStudent(10, students)
	fmt.Println(students)

}

type Student struct {
	Id        int
	FirstName string
	LastName  string
	ClassName string
}

func GetStudentById(studentId int, array []Student) Student {
	var s1 Student
	for _, item := range array {
		if item.Id == studentId {
			s1 = item
		}
	}
	return s1
}

func CreateStudent(id int, FirstName string, LastName string, ClassName string, array []Student) []Student {
	var s1 Student
	s1.Id = id
	s1.FirstName = FirstName
	s1.LastName = LastName
	s1.ClassName = ClassName

	array = append(array, s1)
	return array
}

func UpdateStudent(studentId int, updatedStudent Student, array []Student) []Student {
	var students []Student
	for _, item := range array {
		if item.Id == studentId {
			students = append(students, updatedStudent)
		} else {
			students = append(students, item)
		}
	}
	return students
}

func DeleteStudent(studentId int, array []Student) []Student {
	var students []Student
	for _, item := range array {
		if item.Id != studentId {
			students = append(students, item)
		}
	}
	return students
}
