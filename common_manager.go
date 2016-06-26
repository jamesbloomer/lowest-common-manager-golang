package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Employee struct {
	Name    string
	Manager *Employee
}

func main() {
	fmt.Println(do(os.Stdin))
}

func do(in io.Reader) string {
	i := 0
	var firstSelected string
	var secondSelected string
	var employees []*Employee

	// Create tree
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()

		if isFirstSelected(i) {
			firstSelected = line
		}

		if isSecondSelected(i) {
			secondSelected = line
		}

		if isRelationship(i) {
			employees = addEmployee(line, employees)
		}

		i++
	}

	// Traverse tree
	e1 := getEmployee(firstSelected, employees)
	e2 := getEmployee(secondSelected, employees)

	return searchOuter(e1, e2)
}

func searchOuter(e1 *Employee, e2 *Employee) string {
	result := searchInner(e1, e2)

	if result != "" {
		return result
	}

	return searchOuter(e1.Manager, e2)
}

func searchInner(e1 *Employee, e2 *Employee) string {
	if e1 == nil || e2.Manager == nil {
		return ""
	}

	if e1.Name == e2.Manager.Name {
		return e1.Name
	}

	return searchInner(e1, e2.Manager)
}

func addEmployee(line string, employees []*Employee) []*Employee {
	lineSplit := strings.Split(line, " ")
	manager := getEmployee(lineSplit[0], employees)
	employee := Employee{Name: lineSplit[1], Manager: manager}
	return append(employees, &employee)
}

func getEmployee(name string, employees []*Employee) *Employee {

	for _, e := range employees {
		if e.Name == name {
			return e
		}
	}

	return &Employee{Name: name}
}

func isSize(i int) bool {
	return i == 0
}

func isFirstSelected(i int) bool {
	return i == 1
}

func isSecondSelected(i int) bool {
	return i == 2
}

func isRelationship(i int) bool {
	return i > 2
}
