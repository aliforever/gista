package metadata

import (
	media_constraints "github.com/aliforever/gista/media/media-constraints"
	media_photo "github.com/aliforever/gista/media/media-photo"
	"github.com/aliforever/gista/responses"
	"github.com/aliforever/gista/utils"
)

type Internal struct {
	photoDetails        *media_photo.PhotoDetails
	videoDetails        map[string]string
	uploadId            string
	videoUploadUrls     string
	videoUploadResponse string
	photoUploadResponse *responses.UploadPhoto
	directThreads       string
	directUsers         string
	bestieMedia         bool
}

func (i *Internal) SetPhotoUploadResponse(photoUploadResponse *responses.UploadPhoto) {
	i.photoUploadResponse = photoUploadResponse
}

func (i *Internal) GetUploadId() string {
	return i.uploadId
}

func (i *Internal) SetVideoDetails(targetFeed, videoFilename string) {
	//i.videoDetails = videoDetails
}

func (i *Internal) SetPhotoDetails(targetFeed, photoFilename string) (pd *media_photo.PhotoDetails, err error) {
	i.photoDetails, err = media_photo.NewPhotoDetails(photoFilename)
	if err != nil {
		return
	}
	err = i.photoDetails.Validate(media_constraints.ConstraintsFactory{}.CreateFor(targetFeed))
	if err != nil {
		return
	}
	pd = i.photoDetails
	return
}

func (i *Internal) GetVideoDetails() map[string]string {
	return i.videoDetails
}

func (i *Internal) GetPhotoDetails() *media_photo.PhotoDetails {
	return i.photoDetails
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
