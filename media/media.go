package media

import (
	"fmt"
	"path"

	"github.com/go-errors/errors"
)

type Media struct {
	fileSize int64
	fileName string
	width    int
	height   int
	Details
}

func NewMedia(fileName string, fileSize int64, width, height int) (m Media) {
	m = Media{fileName: fileName, fileSize: fileSize, width: width, height: height}
	return
}

func (m *Media) GetWidth() int {
	if m.HasSwappedAxes() {
		return m.height
	}
	return m.width
}

func (m *Media) GetHeight() int {
	if m.HasSwappedAxes() {
		return m.width
	}
	return m.height
}

func (m *Media) GetAspectRatio() float64 {
	return float64(m.height) / float64(m.width)
}

func (m *Media) GetFileName() string {
	return m.fileName
}

func (m *Media) GetBaseName() string {
	return path.Base(m.fileName)
}

func (m *Media) Validate(constraints Constraints) (err error) {
	fileName := m.GetBaseName()
	if m.HasSwappedAxes() || m.IsVerticallyFlipped() || m.IsHorizontallyFlipped() {
		err = errors.New(fmt.Sprintf("Instagram only accepts non-rotated media. Your file '%s' is either rotated or flipped or both.", fileName))
	}
	aspectRatio := m.GetAspectRatio()
	minAspectRatio := constraints.GetMinAspectRatio()
	maxAspectRatio := constraints.GetMaxAspectRatio()
	if aspectRatio < minAspectRatio || aspectRatio > maxAspectRatio {
		err = errors.New(fmt.Sprintf("Instagram only accepts %s media with aspect ratios between %.3f and %.3f. Your file '%s' has a %.4f aspect ratio.", constraints.GetTitle(), minAspectRatio, maxAspectRatio, fileName, aspectRatio))
	}
	return
}
