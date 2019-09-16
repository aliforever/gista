package media_constraints

type TvConstraints struct {
}

func (sc *TvConstraints) GetMinAspectRatio() float64 {
	return 0.5
}

func (sc *TvConstraints) GetMaxAspectRatio() float64 {
	return 0.8
}

func (sc *TvConstraints) GetRecommendedRatio() float64 {
	return 0.5625
}

func (sc *TvConstraints) GetRecommendedRatioDeviation() float64 {
	return 0.0025
}

func (sc *TvConstraints) UseRecommendedRatioByDefault() bool {
	return true
}

func (sc *TvConstraints) GetMinDuration() float64 {
	return 15.0
}

func (sc *TvConstraints) GetMaxDuration() float64 {
	return 600.0
}

func (sc *TvConstraints) GetTitle() string {
	return "Tv"
}
