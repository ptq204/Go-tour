package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int64
}

type Men interface {
	talk()
	walk()
}

type Student struct {
	Human
	class      string
	numCourses int64
}

type Teacher struct {
	Human
	salary      float64
	lessonHours int64
}

func (t Teacher) talk() {
	fmt.Printf("I'm a teacher and My name is %s\n", t.name)
}

func (t Teacher) walk() {
	fmt.Println("Walk like a teacher")
}

func (t Teacher) getSalary() float64 {
	return t.salary
}

func (s Student) talk() {
	fmt.Printf("I'm a student and My name is %s\n", s.name)
}

func (s Student) walk() {
	fmt.Println("Walk like a student")
}

func (s Student) getNumberOfCourses() int64 {
	return s.numCourses
}
