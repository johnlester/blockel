package blockel

//Top-level API
//BestSplit(Block, []Splits, Criteria)


import (
	"image"
	"image/color"
	"image/jpeg"
	// "fmt"
	// "runtime"
	//	"log"
)

var (
	// numCores = runtime.NumCPU()
)

const (
)



//Block of color
type Block struct {
	top float64
	bottom float64
	right float64
	left float64
	area float64
	color Color
	totalColorDelta float64
	averageColorDelta float64
	//Score		// Value to sort on; customizable score to pick which block to divide
}

//Color interface--to be implmented by RBG, LAB, etc.
type Color interface {

}

//Represents a way of dividing up a single block into 2 or more blocks
type Split struct {
	xSplits []float64	//should be 0 < x < 1
	xSplits []float64	//should be 0 < x < 1
}


//Collection of blocks
//Can be sorted by color deltas or other characteristics
var Blocks = make([]Blocks, 1, 100)




//Function that returns slice of possible Splits to try
//


//Color functions
//Average slice of colors
//Median slice of colors
//Mode slice of colors

//Drawing functions for Blocks
//Optional borders around each block
//Can output to SVG and PNG
