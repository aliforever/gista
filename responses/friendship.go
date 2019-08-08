package responses

import "github.com/aliforever/gista/models"

type Friendship struct {
	Response
	models.FriendshipStatus
}
