package media

type Details interface {
	GetWidth() int
	GetHeight() int
	GetType() string
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
