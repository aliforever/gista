package responses

type ResumableOffset struct {
	Response
	Offset *int `json:"offset"`
}

func (ro *ResumableOffset) IsOk() bool {
	if ro.Offset != nil && *ro.Offset >= 0 {
		return true
	} else {
		if ro.GetMessage() == "" {
			ro.Message = "Offset for resumable uploader is missing or invalid."
		}
	}
	return false
}
