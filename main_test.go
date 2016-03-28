package main

import (
	"testing"
	"reflect"
)

func TestSimple(t *testing.T) {
	mat := newMatrix()

	c0 := mat.createConstraint("C0")
	c1 := mat.createConstraint("C1")
	r0 := mat.createChoice("R0", []*Cell{c0})
	r1 := mat.createChoice("R1", []*Cell{c1})

	found, result := mat.solve()

	if found == false {
		t.Fail()
	}
	if !reflect.DeepEqual(result, []*Cell{r1, r0}) {
		t.Fail()
	}
}