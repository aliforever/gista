package responses

import "github.com/aliforever/gista/models"

type FriendshipsShow struct {
	Response
	models.FriendshipStatus
}
