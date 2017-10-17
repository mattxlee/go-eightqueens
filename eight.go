package main

import (
	"fmt"
)

// Dir is a direct struct
type Dir struct {
	dx, dy int
}

// Check if the direct with distance is valid or not
func (dir *Dir) Check(ox, oy, s int) (int, int, bool) {
	x := s*dir.dx + ox
	y := s*dir.dy + oy
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return x, y, false
	}
	return x, y, true
}

// DIRS contains all directs
var DIRS = []Dir{
	{0, 1},
	{1, 0},
	{1, 1},
	{-1, -1},
	{0, -1},
	{-1, 0},
	{-1, 1},
	{1, -1},
}

// Data can store and access board
type Data struct {
	data  [8][8]int
	count int
	total int
}

// Initialize data
func (d *Data) Initialize() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			d.data[i][j] = 0
		}
	}
	d.count = 0
	d.total = 0
}

// Get board data
func (d *Data) Get(x, y int) int {
	return d.data[x][y]
}

// Add board data
func (d *Data) Add(x, y int) {
	d.data[x][y] = 1
	d.count++
}

// Remove board data
func (d *Data) Remove(x, y int) {
	d.data[x][y] = 0
	d.count--
}

// Check the board position available
func (d *Data) Check(x, y int) bool {
	if d.data[x][y] != 0 {
		return false
	}
	for di := 0; di < len(DIRS); di++ {
		dir := DIRS[di]
		for s := 1; s <= 8; s++ {
			if x, y, valid := dir.Check(x, y, s); valid {
				if d.data[x][y] != 0 {
					return false
				}
			}
		}
	}
	return true
}

// IsDone returns true if board is solved
func (d *Data) IsDone() bool {
	return d.count == 8
}

// Print board status
func (d *Data) Print() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if d.data[i][j] == 0 {
				fmt.Print(" . ")
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println()
	}
}

// Search in depth
func Search(dp int, d *Data) {
	if d.IsDone() {
		d.total++
		fmt.Printf("==== Total: %v ====\n", d.total)
		d.Print()
	}
	for p := dp; p < 64; p++ {
		sx := p % 8
		sy := p / 8
		if d.Check(sx, sy) {
			d.Add(sx, sy)
			Search(p+1, d)
			d.Remove(sx, sy)
		}
	}
}

func main() {
	var d Data
	d.Initialize()
	Search(0, &d)
}
