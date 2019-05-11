package models

type Thumbnail struct {
	VideoLength            float64  `json:"video_length"`
	ThumbnailWidth         int      `json:"thumbnail_width"`
	ThumbnailHeight        int      `json:"thumbnail_height"`
	ThumbnailDuration      float64  `json:"thumbnail_duration"`
	SpriteUrls             []string `json:"sprite_urls"`
	ThumbnailsPerRow       int      `json:"thumbnails_per_row"`
	MaxThumbnailsPerSprite int      `json:"max_thumbnails_per_sprite"`
	SpriteWidth            int      `json:"sprite_width"`
	SpriteHeight           int      `json:"sprite_height"`
	RenderedWidth          int      `json:"rendered_width"`
}
