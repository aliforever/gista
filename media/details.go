package media

type Details interface {
	GetWidth() int
	GetHeight() int
	GetAspectRatio() float64
	GetFileName() string
	GetBaseName() string
	GetFileSize() int
	GetMinAllowedWidth() int
	GetMaxAllowedWidth() int
	HasSwappedAxes() bool
	IsHorizontallyFlipped() bool
	IsVerticallyFlipped() bool
	Validate(constraints Constraints) error
}
