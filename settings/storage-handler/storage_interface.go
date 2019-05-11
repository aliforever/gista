package storage_handler

type StorageInterface interface {
	/*SetSettings(key, value string) error
	SetUsername(username string) error
	SetActiveUser(username string) error*/
	OpenLocation(config map[string]string) error
	HasUser(username string) bool
	MoveUser(oldUsername, newUsername string) error
	DeleteUser(username string) error
	OpenUser(username string) error
	LoadUserSettings() (map[string]string, error)
	SaveUserSettings(userSettings *map[string]string, key string) error
	HasUserCookies() bool
	GetUserCookiesFilePath() string
	LoadUserCookies() (*string, error)
	SaveUserCookies(rawData string) error
	CloseUser() error
	CloseLocation() error
}

type StorageHandlerVars struct {
	username     *string
	userSettings *map[string]string
}
