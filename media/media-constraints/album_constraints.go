package media_constraints

type AlbumConstraints struct {
}

func (ac *AlbumConstraints) GetMinAspectRatio() float64 {
	return 0.8
}

func (ac *AlbumConstraints) GetMaxAspectRatio() float64 {
	return 1.91
}

func (ac *AlbumConstraints) GetRecommendedRatio() float64 {
	return 1.0
}

func (ac *AlbumConstraints) GetRecommendedRatioDeviation() float64 {
	return 0.0
}

func (ac *AlbumConstraints) GetMinDuration() float64 {
	return 3.0
}

func (ac *AlbumConstraints) GetMaxDuration() float64 {
	return 60.0
}

func (ac *AlbumConstraints) GetTitle() string {
	return "timeline"
}

func (ac *AlbumConstraints) UseRecommendedRatioByDefault() bool {
	return false
}
