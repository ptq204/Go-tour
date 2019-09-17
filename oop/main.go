package main

import "fmt"

func main() {
	bob := Student{Human{"Bob", 17}, "16CTT", 12}
	mike := Student{Human{"Mike", 20}, "16CLC", 5}
	jane := Teacher{Human{"Jane", 25}, 422.13, 6}
	john := Teacher{Human{"Jonh", 32}, 500, 7}

	bob.talk()
	fmt.Printf("Number of courses that %s registered is %d\n", mike.name, mike.getNumberOfCourses())

	jane.walk()
	fmt.Printf("Salary of John is: %f\n", john.getSalary())
}
