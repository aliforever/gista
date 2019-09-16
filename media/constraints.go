package media

type Constraints interface {
	GetTitle() string
	GetMinAspectRatio() float64
	GetMaxAspectRatio() float64
	GetRecommendedRatio() float64
	GetRecommendedRatioDeviation() float64
	UseRecommendedRatioByDefault() bool
	GetMinDuration() float64
	GetMaxDuration() float64
}
