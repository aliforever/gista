package media_photo

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/kr/pretty"

	"github.com/go-errors/errors"

	Instagram_media "github.com/aliforever/gista/media/Instagram-media"
	"github.com/aliforever/gista/media/geometry"
)

type InstagramPhoto struct {
	details *PhotoDetails
	Parent  *Instagram_media.InstagramMedia
}

func NewInstagramPhoto(inputFile string, options map[string]interface{}) (ip *InstagramPhoto, err error) {
	ip = &InstagramPhoto{}
	ip.Parent, err = Instagram_media.NewInstagramMedia(ip.createOutputFile, inputFile, options)
	if err != nil {
		return
	}
	ip.details, err = NewPhotoDetails(inputFile)
	if err != nil {
		return
	}
	ip.Parent.Details = ip.details.Media
	ip.Parent.Details.Details = ip.details
	return
}

func (ip *InstagramPhoto) isMod2CanvasRequired() bool {
	return false
}

func (ip *InstagramPhoto) createOutputFile(srcRect geometry.Rectangle, dstRect geometry.Rectangle, canvas geometry.Dimensions) (outputFile string, err error) {
	//var outputFile *string
	var img image.Image
	img, err = ip.loadImage()
	if err != nil {
		os.Remove(ip.Parent.InputFile)
		return
	}
	ip.processResource(img, srcRect, dstRect, canvas)
	return
}

func (ip *InstagramPhoto) processResource(img image.Image, srcRect, dstRect geometry.Rectangle, canvas geometry.Dimensions) {
	if ip.details.HasSwappedAxes() {
		srcRect = srcRect.WithSwappedAxes()
		dstRect = dstRect.WithSwappedAxes()
		canvas = canvas.WithSwappedAxes()
	}
	output := image.NewRGBA(image.Rect(0, 0, canvas.GetWidth(), canvas.GetHeight()))
	colors := color.RGBA{R: ip.Parent.BgColor[0], G: ip.Parent.BgColor[1], B: ip.Parent.BgColor[2], A: 255}
	draw.Draw(output, output.Bounds(), &image.Uniform{C: colors}, image.ZP, draw.Src)
	draw.Draw(output, output.Bounds(), img, image.Point{X: 90, Y: srcRect.GetY()}, draw.Src)
	pretty.Println("Source X:", srcRect.GetX(), "Source Y: ", srcRect.GetY(), "Dest X: ", dstRect.GetX(), "Dest Y: ", dstRect.GetY(), "Source Width: ", srcRect.GetWidth(), "Source Height: ", srcRect.GetHeight())
	save, err := os.Create("test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = jpeg.Encode(save, output, nil)
	if err != nil {
		fmt.Println(err)
	}
	//imagecopyresampled
}

func (ip *InstagramPhoto) loadImage() (img image.Image, err error) {
	switch ip.details.GetType() {
	case "jpeg", "png", "gif":
		var file *os.File
		file, err = os.Open(ip.Parent.InputFile)
		if err != nil {
			return
		}
		defer file.Close()
		img, _, err = image.Decode(file)
	default:
		err = errors.New(fmt.Sprintf("Unsupported Image Type: %s", ip.Parent.InputFile))
	}
	return
}
