package metadata

import "github.com/aliforever/gista/utils"

type Internal struct {
	photoDetails        map[string]string
	videoDetails        map[string]string
	uploadId            string
	videoUploadUrls     string
	videoUploadResponse string
	photoUploadResponse string
	directThreads       string
	directUsers         string
	bestieMedia         bool
}

func NewInternalMetaData(uploadId *string) (md *Internal) {
	md = &Internal{}
	if uploadId != nil {
		md.uploadId = *uploadId
	} else {
		md.uploadId = utils.GenerateUploadId(false)
	}
	md.bestieMedia = false
	return
}

func (i *Internal) GetPhotoDetails() {

}
