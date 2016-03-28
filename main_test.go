package main

import (
	"testing"
	"reflect"
)

func TestBasic2x2(t *testing.T) {
	mat := newMatrix()

	c0 := mat.createConstraint("C0")
	c1 := mat.createConstraint("C1")
	r0 := mat.createChoice("R0", []*Cell{c0})
	r1 := mat.createChoice("R1", []*Cell{c1})

	found, result := mat.solve()

	if found == false {
		t.Error("Should have found a solution")
	}
	if !reflect.DeepEqual(result, []*Cell{r1, r0}) {
		t.Error("Wrong result")
	}
}



func TestBasic3x3(t *testing.T) {
	mat := newMatrix()

	c0 := mat.createConstraint("C0")
	c1 := mat.createConstraint("C1")
	c2 := mat.createConstraint("C2")

	mat.createChoice("R0", []*Cell{c0})
	r1 := mat.createChoice("R1", []*Cell{c1})
	r2 := mat.createChoice("R2", []*Cell{c0, c2})

	found, result := mat.solve()

	if found == false {
		t.Error("Should have found a solution")
	}
	if !reflect.DeepEqual(result, []*Cell{r1, r2}) {
		t.Error("Wrong result")
	}
}


func TestBasic4x4(t *testing.T) {
	mat := newMatrix()

	c0 := mat.createConstraint("C0")
	c1 := mat.createConstraint("C1")
	c2 := mat.createConstraint("C2")
	c3 := mat.createConstraint("C3")

	r0 := mat.createChoice("R0", []*Cell{c0})
	r1 := mat.createChoice("R1", []*Cell{c1})
	mat.createChoice("R2", []*Cell{c1, c2})
	r3 := mat.createChoice("R3", []*Cell{c2, c3})

	found, result := mat.solve()

	if found == false {
		t.Error("Should have found a solution")
	}
	if !reflect.DeepEqual(result, []*Cell{r3, r1, r0}) {
		t.Error("Wrong result")
	}
}