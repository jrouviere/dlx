package main

// Sparse matrix implementation with 2D doubly linked list
type Cell struct {
	Left, Right, Up, Down *Cell
	data interface{}
}


// // Constraints are columns in our matrix
// type Constraint struct {
// 	Cell
// 	description string
// }

// // Choices are the rows of our matrix
// type Choice struct {
// 	Cell
// 	description string
// }

// The problem matrix, simply defined as the root element of the doubly linked list
type Matrix struct {
	root Cell
}


func newMatrix() *Matrix {
	matrix := Matrix{}
	
	//initially the matrix is completely empty
	matrix.root.Left = &matrix.root
	matrix.root.Right = &matrix.root
	matrix.root.Up = &matrix.root
	matrix.root.Down = &matrix.root
	
	return &matrix
}

func (matrix Matrix) createConstraint(data interface{}) *Cell {
	constraint := Cell{}
	
	// initially no choice cover this constraint
	constraint.Up = &constraint
	constraint.Down = &constraint
	
	constraint.data = data
	
	// add the constraint to the matrix as the last column
	constraint.Left = matrix.root.Left
	constraint.Right = &matrix.root
	
	//TODO: should be constraint.restore
	constraint.Left.Right = &constraint
	constraint.Right.Left = &constraint
	
	return &constraint
}

func (matrix Matrix) createChoice(data interface{}, constraints []*Cell) {
	choice := Cell{}
	
	// initially this choice doesn't cover any constraint
	choice.Left = &choice
	choice.Right = &choice
	
	choice.data = data
	
	// add the choice to the matrix as the last row 
	choice.Up = matrix.root.Up
	choice.Down = &matrix.root
	
	//TODO: should be row.restore
	choice.Up.Down = &choice
	choice.Down.Up = &choice
}

/*
A = {1, 4, 7}
B = {1, 4}
C = {4, 5, 7}
D = {3, 5, 6}
E = {2, 3, 6, 7}
F = {2, 7}
*/
func main() {
	mat := newMatrix()
	
	ctr := make([]*Cell, 7)
	for i := 0; i < 7; i++ {
		ctr[i] = mat.createConstraint(i+1)
	}
	
	mat.createChoice("A", []*Cell{ctr[0], ctr[3], ctr[6]})
	mat.createChoice("B", []*Cell{ctr[0], ctr[3]})
	mat.createChoice("C", []*Cell{ctr[3], ctr[4], ctr[6]})
	mat.createChoice("D", []*Cell{ctr[2], ctr[4], ctr[5]})
	mat.createChoice("E", []*Cell{ctr[1], ctr[2], ctr[5], ctr[6]})
	mat.createChoice("F", []*Cell{ctr[1], ctr[6]})
}