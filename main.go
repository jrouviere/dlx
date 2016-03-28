package main


import (
	"fmt"
)

// Cell is a "one" in our sparse matrix, it's an element of a 2D doubly linked list
type Cell struct {
	Left, Right, Up, Down *Cell
	Row, Column           *Cell
	data                  interface{}
}

func (cell Cell) String() string {
	return fmt.Sprintf("%v", cell.data)
}

// Matrix defines a sparse matrix, it's the root element of the doubly linked list
// Constraints are columns in our matrix
// Choices are the rows of our matrix
type Matrix struct {
	root *Cell
}

func newMatrix() *Matrix {
	matrix := Matrix{}

	//initially the matrix is completely empty
	matrix.root = new(Cell)
	matrix.root.Left = matrix.root
	matrix.root.Right = matrix.root
	matrix.root.Up = matrix.root
	matrix.root.Down = matrix.root

	matrix.root.data = "root"

	return &matrix
}

func (matrix Matrix) createConstraint(data interface{}) *Cell {
	constraint := Cell{}

	constraint.data = data

	// initially no choice cover this constraint
	constraint.Up = &constraint
	constraint.Down = &constraint
	constraint.Column = &constraint

	// add the constraint to the matrix as the last column
	constraint.Left = matrix.root.Left
	constraint.Right = matrix.root
	constraint.Row = matrix.root

	//TODO: should be constraint.restore
	constraint.Left.Right = &constraint
	constraint.Right.Left = &constraint

	return &constraint
}

func (matrix Matrix) createChoice(data interface{}, constraints []*Cell) *Cell {
	choice := Cell{}

	choice.data = data

	// initially this choice doesn't cover any constraint
	choice.Left = &choice
	choice.Right = &choice
	choice.Row = &choice

	// add the choice to the matrix as the last row
	choice.Up = matrix.root.Up
	choice.Down = matrix.root
	choice.Column = matrix.root

	//TODO: should be row.restore
	choice.Up.Down = &choice
	choice.Down.Up = &choice

	for _, constraint := range constraints {
		cell := Cell{}

		// add it to the choice
		cell.Left = choice.Left
		cell.Right = choice.Left.Right
		cell.Left.Right = &cell
		cell.Right.Left = &cell
		cell.Row = choice.Left.Right

		// add it to the constraint
		cell.Up = constraint.Up
		cell.Down = constraint.Up.Down
		cell.Up.Down = &cell
		cell.Down.Up = &cell
		cell.Column = constraint

		cell.data = fmt.Sprintf("%v: %v", cell.Row.data, cell.Column.data)
	}
	return &choice
}


func (matrix Matrix) coverColumn(column *Cell) {
	column = column.Column
	column.Left.Right = column.Right
	column.Right.Left = column.Left
	for row := column.Down; row != column; row = row.Down {
		for cell := row.Right; cell != row; cell = cell.Right {
			cell.Up.Down = cell.Down
			cell.Down.Up = cell.Up
		}
	}
}

func (matrix Matrix) uncoverColumn(column *Cell) {
	column = column.Column
	for row := column.Up; row != column; row = row.Up {
		for cell := row.Left; cell != row; cell = cell.Left {
			cell.Up.Down = cell
			cell.Down.Up = cell
		}
	}
	column.Left.Right = column
	column.Right.Left = column
}

func (matrix Matrix) solve() (bool, []*Cell) {

	// TODO: use an euristic to take the constraint satisfied by fewer choices

	// choose an unsolved constraint: aka a column
	constraint := matrix.root.Right

	// no more constraint, we are done!
	if constraint == constraint.Right {
		return true, make([]*Cell, 0)
	}

	// cover column
	matrix.coverColumn(constraint)

	for row := constraint.Down; row != constraint; row = row.Down {
		rowHead := row.Row
		for cell := rowHead.Right; cell != rowHead; cell = cell.Right {
			matrix.coverColumn(cell)
		}

		// solve recursively
		found, partial := matrix.solve()
		if found {
			partial = append(partial, rowHead)
			return true, partial
		}

		for cell := rowHead.Left; cell != rowHead; cell = cell.Left {
			matrix.uncoverColumn(cell)
		}

	}

	matrix.uncoverColumn(constraint)

	return false, make([]*Cell, 0)
}

func main() {
	mat := newMatrix()

	/* sample from wikipedia
	A = {1, 4, 7}
	B = {1, 4}
	C = {4, 5, 7}
	D = {3, 5, 6}
	E = {2, 3, 6, 7}
	F = {2, 7}
	*/
	ctr := make([]*Cell, 7)
	for i := 0; i < 7; i++ {
		ctr[i] = mat.createConstraint(i + 1)
	}

	mat.createChoice("A", []*Cell{ctr[0], ctr[3], ctr[6]})
	mat.createChoice("B", []*Cell{ctr[0], ctr[3]})
	mat.createChoice("C", []*Cell{ctr[3], ctr[4], ctr[6]})
	mat.createChoice("D", []*Cell{ctr[2], ctr[4], ctr[5]})
	mat.createChoice("E", []*Cell{ctr[1], ctr[2], ctr[5], ctr[6]})
	mat.createChoice("F", []*Cell{ctr[1], ctr[6]})

	found, result := mat.solve()
	fmt.Println(found, result)
}
