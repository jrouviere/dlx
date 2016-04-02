package main

import (
	"testing"
)

func cellInSlice(cell *Cell, slice []*Cell) bool {
	for _, c := range(slice) {
		if cell == c {
			return true
		}
	}
	return false
}

func checkSolution(result, expected []*Cell, t *testing.T) {
	//we don't care about the order
	for _, cell := range(result) {
		if ! cellInSlice(cell, expected) {
			t.Errorf("Result contains unexpected value {%v}, expected: %v", cell, expected)
		}
	}

	for _, cell := range(expected) {
		if ! cellInSlice(cell, result) {
			t.Errorf("Expected value {%v} not found in result: %v", cell, result)
		}
	}

	if len(result) != len(expected) {
		t.Errorf("Unexpected number of results %v != %v", len(result), len(expected))
	}
}

func TestBasic2x2(t *testing.T) {
	mat := newMatrix()

	c0 := mat.createConstraint("C0")
	c1 := mat.createConstraint("C1")
	r0 := mat.createChoice("R0", []*Cell{c0})
	r1 := mat.createChoice("R1", []*Cell{c1})

	found, result := mat.solve()

	if !found {
		t.Error("Should have found a solution")
	}
	checkSolution(result, []*Cell{r0, r1}, t)
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

	if !found {
		t.Error("Should have found a solution")
	}
	checkSolution(result, []*Cell{r1, r2}, t)
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

	if !found {
		t.Error("Should have found a solution")
	}
	checkSolution(result, []*Cell{r0, r1, r3}, t)
}

func TestNoSolution(t *testing.T) {
	mat := newMatrix()

	c0 := mat.createConstraint("C0")
	mat.createConstraint("C1")
	mat.createChoice("R0", []*Cell{c0})
	mat.createChoice("R1", []*Cell{c0})

	found, result := mat.solve()

	if found {
		t.Error("Solution should have not been found")
	}
	checkSolution(result, []*Cell{}, t)
}

func TestCover(t *testing.T) {

	/* sample from wikipedia
	A = {1, 4, 7}
	B = {1, 4}
	C = {4, 5, 7}
	D = {3, 5, 6}
	E = {2, 3, 6, 7}
	F = {2, 7}
	=> solution is B, D, F
	*/

	mat := newMatrix()

	ctr := make([]*Cell, 7)
	for i := 0; i < 7; i++ {
		ctr[i] = mat.createConstraint(i + 1)
	}

	mat.createChoice("A", []*Cell{ctr[0], ctr[3], ctr[6]})
	B := mat.createChoice("B", []*Cell{ctr[0], ctr[3]})
	mat.createChoice("C", []*Cell{ctr[3], ctr[4], ctr[6]})
	D := mat.createChoice("D", []*Cell{ctr[2], ctr[4], ctr[5]})
	mat.createChoice("E", []*Cell{ctr[1], ctr[2], ctr[5], ctr[6]})
	F := mat.createChoice("F", []*Cell{ctr[1], ctr[6]})

	found, result := mat.solve()
	if !found {
		t.Error("Should have found a solution")
	}
	checkSolution(result, []*Cell{B, D, F}, t)
}