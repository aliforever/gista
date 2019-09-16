package media_constraints

type StoryConstraints struct {
}

func (sc *StoryConstraints) GetMinAspectRatio() float64 {
	return 0.56
}

func (sc *StoryConstraints) GetMaxAspectRatio() float64 {
	return 0.67
}

func (sc *StoryConstraints) GetRecommendedRatio() float64 {
	return 0.5625
}

func (sc *StoryConstraints) GetRecommendedRatioDeviation() float64 {
	return 0.0025
}

func (sc *StoryConstraints) GetMinDuration() float64 {
	return 1.0
}

func (sc *StoryConstraints) GetMaxDuration() float64 {
	return 15.0
}

func (sc *StoryConstraints) GetTitle() string {
	return "story"
}

func (sc *StoryConstraints) UseRecommendedRatioByDefault() bool {
	return true
}
