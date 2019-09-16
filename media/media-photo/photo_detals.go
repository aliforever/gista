package media_photo

import (
	"fmt"

	"github.com/aliforever/gista/media"
	"github.com/aliforever/gista/utils"
	"github.com/go-errors/errors"
)

type PhotoDetails struct {
	media.Media
	photoType   string
	orientation string
}

func NewPhotoDetails(fileName string) (pd *PhotoDetails, err error) {
	if fileName == "" || !utils.FileOrFolderExists(fileName) {
		err = errors.New(fmt.Sprintf(`The photo file "%s" does not exist on disk.`, fileName))
		return
	}
	var fileSize int64
	fileSize, err = utils.GetFileSize(fileName)
	if err != nil {
		return
	}
	if fileSize < 1 {
		err = errors.New(fmt.Sprintf(`The photo file "%s" is empty.`, fileName))
		return
	}
	var width, height int
	width, height, err = utils.GetImageDimension(fileName)
	if err != nil {
		return
	}
	var orientation string
	orientation, err = utils.GetImageOrientation(fileName)

	var imageType string
	imageType, err = utils.GuessImageFormat(fileName)
	pd = &PhotoDetails{}
	pd.orientation = orientation
	pd.photoType = imageType
	pd.Media = media.NewMedia(fileName, fileSize, width, height)

	return
}

func (pd *PhotoDetails) GetType() string {
	return pd.photoType
}

func (pd *PhotoDetails) HasSwappedAxes() bool {
	arr := []string{"5", "6", "7", "8"}
	for _, orientation := range arr {
		if pd.orientation == orientation {
			return true
		}
	}
	return false
}

func (pd *PhotoDetails) IsHorizontallyFlipped() bool {
	arr := []string{"2", "3", "6", "7"}
	for _, orientation := range arr {
		if pd.orientation == orientation {
			return true
		}
	}
	return false
}

func (pd *PhotoDetails) IsVerticallyFlipped() bool {
	arr := []string{"3", "4", "7", "8"}
	for _, orientation := range arr {
		if pd.orientation == orientation {
			return true
		}
	}
	return false
}

func (pd *PhotoDetails) GetMinAllowedWidth() int {
	return minWidth
}

func (pd *PhotoDetails) GetMaxAllowedWidth() int {
	return maxWidth
}

func (pd *PhotoDetails) Validate(constraints media.Constraints) (err error) {
	err = pd.Media.Validate(constraints)
	if err != nil {
		return
	}
	mediaFileName := pd.GetBaseName()

	imageType := pd.GetType()

	if imageType == "jpeg" {
		err = errors.New(fmt.Sprintf(`The photo file "%s" is not a JPEG file.`, mediaFileName))
	}

	width := pd.GetWidth()

	if width < minWidth || width > maxWidth {
		err = errors.New(fmt.Sprintf(`Instagram only accepts photos that are between %d and %d pixels wide. Your file "%s" is %d pixels wide.`, minWidth, maxWidth, mediaFileName, width))
	}
	return
}
