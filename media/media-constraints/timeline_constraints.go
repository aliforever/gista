package media_constraints

type TimelineConstraints struct {
}

func (sc *TimelineConstraints) GetMinAspectRatio() float64 {
	return 0.8
}

func (sc *TimelineConstraints) GetMaxAspectRatio() float64 {
	return 1.91
}

func (sc *TimelineConstraints) GetRecommendedRatio() float64 {
	return 1.0
}

func (sc *TimelineConstraints) GetRecommendedRatioDeviation() float64 {
	return 0.0
}

func (sc *TimelineConstraints) GetMinDuration() float64 {
	return 3.0
}

func (sc *TimelineConstraints) GetMaxDuration() float64 {
	return 60.0
}

func (sc *TimelineConstraints) GetTitle() string {
	return "timeline"
}

func (sc *TimelineConstraints) UseRecommendedRatioByDefault() bool {
	return false
}
