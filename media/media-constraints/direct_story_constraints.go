package media_constraints

type DirectStoryConstraints struct {
}

func (sc *DirectStoryConstraints) GetMinAspectRatio() float64 {
	return 0.56
}

func (sc *DirectStoryConstraints) GetMaxAspectRatio() float64 {
	return 0.67
}

func (sc *DirectStoryConstraints) GetRecommendedRatio() float64 {
	return 0.5625
}

func (sc *DirectStoryConstraints) GetRecommendedRatioDeviation() float64 {
	return 0.0025
}

func (sc *DirectStoryConstraints) UseRecommendedRatioByDefault() bool {
	return true
}

func (sc *DirectStoryConstraints) GetMinDuration() float64 {
	return 0.1
}

func (sc *DirectStoryConstraints) GetMaxDuration() float64 {
	return 15.0
}

func (sc *DirectStoryConstraints) GetTitle() string {
	return "direct story"
}
