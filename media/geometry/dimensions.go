package geometry

import (
	"fmt"
	"math"

	"github.com/go-errors/errors"
)

type Dimensions struct {
	width       int
	height      int
	aspectRatio float64
}

func (d *Dimensions) GetAspectRatio() float64 {
	return d.aspectRatio
}

func (d *Dimensions) GetWidth() int {
	return d.width
}

func (d *Dimensions) GetHeight() int {
	return d.height
}

func (d *Dimensions) WithSwappedAxes() (nd Dimensions) {
	nd = NewDimensions(d.height, d.width)
	return
}

func (d *Dimensions) WithRescaling(newScale *float64, roundingFunc *string) (nd Dimensions, err error) {
	scale := 1.0
	round := "round"
	if newScale != nil {
		scale = *newScale
	}
	if roundingFunc != nil {
		if *roundingFunc != "round" && *roundingFunc != "floor" && *roundingFunc != "ceil" {
			err = errors.New(fmt.Sprintf(`Invalid rounding function "%s".`, *roundingFunc))
			return
		}
		round = *roundingFunc
	}

	newWidth := scale * float64(d.width)
	newHeight := scale * float64(d.width)
	if round == "round" {
		newWidth = math.Round(newWidth)
		newHeight = math.Round(newHeight)
	} else if round == "floor" {
		newWidth = math.Floor(newWidth)
		newHeight = math.Floor(newHeight)
	} else if round == "ceil" {
		newWidth = math.Ceil(newWidth)
		newHeight = math.Ceil(newHeight)
	}
	nd = NewDimensions(int(newWidth), int(newHeight))
	return
}

func NewDimensions(width, height int) (d Dimensions) {
	d = Dimensions{width: width, height: height}
	d.aspectRatio = float64(width) / float64(height)
	return
}
