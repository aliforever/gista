package media_photo

import "github.com/aliforever/gista/media"

type InstagramPhoto struct {
	details media.Media
}

func (ip *InstagramPhoto) isMod2CanvasRequired() bool {
	return false
}

func (ip *InstagramPhoto) createOutputFile(srcRect Rectangle, dstRect Rectangle, canvas Dimensions) (err error) {
	var outputFile *string
	resource := ip.loadImage()
	return
}

func (ip *InstagramPhoto) loadImage() (err error) {
	switch ip.details.GetType() {
	case "jpeg":

	}
	return
}
