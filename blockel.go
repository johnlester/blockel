package blockel

//Top-level API
//BestSplit(Block, []Splits, Criteria)


import (
	"image"
	"image/color"
	_ "image/jpeg"
	// "fmt"
	"os"
	"log"
	// "github.com/ajstarks/svgo"

)

var (
	// numCores = runtime.NumCPU()
)

const (
	BlockSetReservedSize = 40
)

type Dimensions struct {
	// Full new picture Dimensions are 0.0, 0.0 to width, height of original image
	left, right, top, bottom float64
}

//Block of color
type Block struct {
	Dim Dimensions
	BlockColor color.Color
	Score float64 		// Value to sort on; customizable score to pick which block to divide
}

func NewBlock(d Dimensions, c color.Color) Block {
	newBlock := Block{Dim: d, BlockColor: c}
	return newBlock
}

type BlockSet struct {
	Blocks []Block
	ReferenceImage *image.Image	
}

func NewBlockSet(refImageFilePath string) *BlockSet {
	// Open the reference image file
	file, err := os.Open(refImageFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// Decode the image
	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	
	//Create first Block
	bounds := m.Bounds()
	imgLeft := bounds.Min.X
	imgRight := bounds.Max.X
	imgTop := bounds.Min.Y
	imgBottom := bounds.Max.Y
	imgWidth := imgRight - imgLeft
	imgHeight := imgBottom - imgTop
	dim := Dimensions{left: 0.0, right: float64(imgWidth), top: 0.0,  bottom: float64(imgHeight)}
	aveColor := averageRGB(m, dim)
	firstBlock := NewBlock(dim, aveColor)

	
	
	//Create a BlockSet
	blocksInSet := make([]block, 1, BlockSetReservedSize)
	newBlockSet = BlockSet{blocks: blocksinSet, referenceImage: &m}
	
	return newBlockSet
}

func (bs *BlockSet) AddBlock(b *Block) {
	//Calculate score for Block
	b.Score = CalculateScore(bs.ReferenceImage, *b)
}

func CalculateScore(m *image.Image, b Block) float64 {
	rect := dimensionsOnImage(m, b.Dim)
	minx := rect.Min.X
	maxx := rect.Max.X
	miny := rect.Min.Y
	maxy := rect.Max.Y
	
	var delta uint64 = 0
	for y := miny; y < maxy; y++ {
		for x := minx; x < maxx; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			//sum of squares
			delta += 0
		}
	}
	return float64(delta)
}

func dimensionsOnImage(m *image.Image, dimensions Dimensions) image.Rect {
	bounds := m.Bounds()
	left := bounds.Min.X
	top := bounds.Min.Y
	
	newLeft := left + int(dimensions.left)
	newRight := left + int(dimensions.right)
	newTop := top + int(dimensions.top)
	newBottom := top + int(dimensions.bottom)
	
	return image.Rect(newLeft, newTop, newRight, newBottom)
}

func averageRGB(m image.Image, dimensions) color.Color {
	//Calculate average color


	var rt, gt, bt, at uint64
	for y := miny; y < maxy; y++ {
		for x := minx; x < maxx; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			rt += uint64(r)
			gt += uint64(g)
			bt += uint64(b)
			at += uint64(a)
		}
	}
	pixelCount := (uint64(maxy - miny) * uint64(maxx - minx))	
	ra := uint8(rt / (256 * pixelCount))
	ga := uint8(gt / (256 * pixelCount))
	ba := uint8(bt / (256 * pixelCount))
	aa := uint8(at / (256 * pixelCount))
	aveColor := color.RGBA{ra, ga, ba, aa}
	return aveColor
}


//Criteria: takes 2 or more block-sets covering the same area and rank orders them
// type Criteria interface {
// 	Score(BlockSet)
// }




//Represents a way of dividing up a single block into 2 or more blocks
// type Split struct {
// 	xSplits []float64	//should be 0 < x < 1
// 	xSplits []float64	//should be 0 < x < 1
// }


//Collection of blocks

//Function that returns slice of possible Splits to try
//


//Color functions
//Average slice of colors
//Median slice of colors
//Mode slice of colors

//Drawing functions for Blocks
//Optional borders around each block
//Can output to SVG and PNG
