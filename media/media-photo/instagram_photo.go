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

	"github.com/aliforever/gista/utils"

	"github.com/disintegration/imaging"

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
	var result image.Image

	result, err = ip.processResource(img, srcRect, dstRect, canvas)
	if err != nil {
		return
	}

	p := "IMG"
	var f *os.File

	f, err = utils.CreateTempFile(ip.Parent.TmpPath, &p)
	if err != nil {
		return
	}

	err = jpeg.Encode(f, result, &jpeg.Options{Quality: JPGQuality})
	if err != nil {
		if utils.FileOrFolderExists(f.Name()) {
			os.Remove(f.Name())
		}
		return
	}
	outputFile = f.Name()
	return
}

func (ip *InstagramPhoto) processResource(img image.Image, srcRect, dstRect geometry.Rectangle, canvas geometry.Dimensions) (result image.Image, err error) {
	if ip.details.HasSwappedAxes() {
		srcRect = srcRect.WithSwappedAxes()
		dstRect = dstRect.WithSwappedAxes()
		canvas = canvas.WithSwappedAxes()
	}

	// Prepare Image to place
	preImg := image.NewRGBA(image.Rect(0, 0, srcRect.GetWidth(), srcRect.GetHeight()))
	draw.Draw(preImg, preImg.Bounds(), img, image.Point{X: srcRect.GetX(), Y: srcRect.GetY()}, draw.Src)
	preImgRescaled := imaging.Resize(preImg, dstRect.GetWidth(), dstRect.GetHeight(), imaging.Lanczos)

	// White Canvas Background
	background := image.NewRGBA(image.Rect(0, 0, canvas.GetWidth(), canvas.GetHeight()))
	draw.Draw(background, background.Bounds(), &image.Uniform{C: ip.Parent.BgColor}, image.ZP, draw.Src)

	// Draw Image to place on White Canvas on new image

	dimensions := image.Rect(dstRect.GetX(), dstRect.GetY(), dstRect.GetWidth()+dstRect.GetX(), dstRect.GetHeight()+dstRect.GetY())
	draw.Draw(background, background.Bounds(), background, image.Point{}, draw.Src)
	draw.Draw(background, dimensions, preImgRescaled, image.Point{}, draw.Src)
	result, err = ip.rotateResource(background, ip.Parent.BgColor)
	return
}

func (ip *InstagramPhoto) rotateResource(original image.Image, bgColor color.Color) (img image.Image, err error) {
	var angle float64
	var flip *int
	if ip.details.HasSwappedAxes() {
		if ip.details.IsHorizontallyFlipped() && ip.details.IsVerticallyFlipped() {
			angle = -90
			f := ImgFlipHorizontal
			flip = &f
		} else if ip.details.IsHorizontallyFlipped() {
			angle = -90
		} else if ip.details.IsVerticallyFlipped() {
			angle = 90
		} else {
			angle = -90
			f := ImgFlipVertical
			flip = &f
		}
	} else {
		if ip.details.IsHorizontallyFlipped() && ip.details.IsVerticallyFlipped() {
			f := ImgFlipBoth
			flip = &f
		} else if ip.details.IsHorizontallyFlipped() {
			f := ImgFlipHorizontal
			flip = &f
		} else if ip.details.IsVerticallyFlipped() {
			f := ImgFlipVertical
			flip = &f
		} else {
			// Do nothing.
		}
	}
	if flip != nil {
		if *flip == ImgFlipHorizontal {
			img = imaging.FlipH(original)
		} else if *flip == ImgFlipVertical {
			img = imaging.FlipV(original)
		} else {
			img = imaging.FlipV(original)
			img = imaging.FlipH(img)
		}
	} else {
		img = original
	}

	// Return original resource if no rotation is needed.
	if angle == 0 {
		return
	}
	img = imaging.Rotate(img, angle, bgColor)
	return
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
