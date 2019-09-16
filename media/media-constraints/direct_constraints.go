package media_constraints

type DirectConstraints struct {
}

func (sc *DirectConstraints) GetMinAspectRatio() float64 {
	return 0.8
}

func (sc *DirectConstraints) GetMaxAspectRatio() float64 {
	return 1.91
}

func (sc *DirectConstraints) GetRecommendedRatio() float64 {
	return 1.0
}

func (sc *DirectConstraints) GetRecommendedRatioDeviation() float64 {
	return 0.0
}

func (sc *DirectConstraints) UseRecommendedRatioByDefault() bool {
	return false
}

func (sc *DirectConstraints) GetMinDuration() float64 {
	return 0.1
}

func (sc *DirectConstraints) GetMaxDuration() float64 {
	return 15.0
}

func (sc *DirectConstraints) GetTitle() string {
	return "direct"
}
