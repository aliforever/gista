package geometry

import (
	"fmt"
	"math"

	"github.com/go-errors/errors"
)

type Rectangle struct {
	x           int
	y           int
	width       int
	height      int
	aspectRatio float64
}

func NewRectangle(x, y, width, height int) (r Rectangle) {
	r = Rectangle{x: x, y: y, width: width, height: height}
	r.aspectRatio = float64(width / height)
	return
}

func (r *Rectangle) GetX() int {
	return r.x
}

func (r *Rectangle) GetY() int {
	return r.y
}

func (r *Rectangle) GetX1() int {
	return r.x
}

func (r *Rectangle) GetY1() int {
	return r.y
}

func (r *Rectangle) GetX2() int {
	return r.x + r.width
}

func (r *Rectangle) GetY2() int {
	return r.y + r.height
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) GetAspectRatio() float64 {
	return r.aspectRatio
}

func (r *Rectangle) WithSwappedAxes() (nr Rectangle) {
	return NewRectangle(r.y, r.x, r.height, r.width)
}

func (r *Rectangle) WithRescaling(newScale *float64, roundingFunc *string) (nr *Rectangle, err error) {
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

	newWidth := scale * float64(r.width)
	newHeight := scale * float64(r.height)
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
	rec := NewRectangle(r.x, r.y, int(newWidth), int(newHeight))
	nr = &rec
	return
}
