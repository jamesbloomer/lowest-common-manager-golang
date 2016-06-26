package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRunning(t *testing.T) {
	fmt.Println("Tests are running")
}

func TestIsFirstSelected(t *testing.T) {
	assert.True(t, isFirstSelected(1))
	assert.False(t, isFirstSelected(0))
	assert.False(t, isFirstSelected(2))
}

func TestIsSecondSelected(t *testing.T) {
	assert.True(t, isSecondSelected(2))
	assert.False(t, isSecondSelected(1))
	assert.False(t, isSecondSelected(3))
}

func TestIsRelationship(t *testing.T) {
	assert.True(t, isRelationship(3))
	assert.True(t, isRelationship(4))
	assert.False(t, isRelationship(2))
	assert.False(t, isRelationship(1))
}

func TestGetEmployeeDoesNotExist(t *testing.T) {
	var employees []*Employee
	assert.Equal(t, getEmployee("A", employees).Name, "A")
}

func TestGetEmployeeExists(t *testing.T) {
	e := Employee{Name: "A"}
	var employees []*Employee
	employees = append(employees, &e)

	assert.Equal(t, getEmployee("A", employees).Name, e.Name)
}

func TestSearchInnerMatchNoMatch(t *testing.T) {
	m1 := Employee{Name: "m1"}
	e1 := Employee{Name: "e1", Manager: &m1}
	e2 := Employee{Name: "e2"}
	assert.Equal(t, "", searchInner(&e1, &e2))
}

func TestSearchInnerMatchNoRecursion(t *testing.T) {
	e1 := Employee{Name: "e1"}
	e2 := Employee{Name: "e2", Manager: &e1}
	assert.Equal(t, "e1", searchInner(&e1, &e2))
}

func TestSearchInnerMatchRecursion(t *testing.T) {
	e1 := Employee{Name: "e1"}
	m2 := Employee{Name: "m2", Manager: &e1}
	e2 := Employee{Name: "e2", Manager: &m2}
	assert.Equal(t, "e1", searchInner(&e1, &e2))
}

func TestSearchOuterMatchRecursion(t *testing.T) {
	m1 := Employee{Name: "m1"}
	m2 := Employee{Name: "m2", Manager: &m1}
	m3 := Employee{Name: "m3", Manager: &m1}
	e1 := Employee{Name: "e1", Manager: &m2}
	e2 := Employee{Name: "e2", Manager: &m3}
	assert.Equal(t, "m1", searchOuter(&e1, &e2))
}

func Test1(t *testing.T) {
	input := "6\nSarah\nFred\nJune Tom\nTom Nathan\nTom Sarah\nNathan Qing\nNathan Fred"
	in := strings.NewReader(input)

	result := do(in)
	assert.Equal(t, "Tom", result)
}

func Test2(t *testing.T) {
	input := "5\nSarah\nClaudiu\nJune Sarah\nSarah Tom\nTom Simon\nTom Claudiu"
	in := strings.NewReader(input)

	result := do(in)
	assert.Equal(t, "Sarah", result)
}
