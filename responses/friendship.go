package responses

import "github.com/aliforever/gista/models"

type Friendship struct {
	Response
	FriendshipStatus models.FriendshipStatus `json:"friendship_status"`
}
