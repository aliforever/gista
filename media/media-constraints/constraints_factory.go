package media_constraints

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/media"
)

type ConstraintsFactory struct {
}

func (cf ConstraintsFactory) CreateFor(targetFeed string) (c media.Constraints) {
	switch targetFeed {
	case constants.FeedStory:
		c = &StoryConstraints{}
	case constants.FeedDirect:
		c = &DirectConstraints{}
	case constants.FeedDirectStory:
		c = &DirectStoryConstraints{}
		// TODO: FEED_TIMELINE_ALBUM, FEED_TIMELINE
	case constants.FeedTv:
		c = &TvConstraints{}
	case constants.FeedTimelineAlbum:
		c = &AlbumConstraints{}
	//case constants.FeedTimeline:
	default:
		c = &TimelineConstraints{}
	}
	return
}
