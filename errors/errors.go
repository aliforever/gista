package errors

import (
	"fmt"

	"github.com/go-errors/errors"
)

var (
	EmptyUsernameOrPassword       = errors.New("Please provide a username and password to login")
	StorageNoUsernameIsSet        = errors.New("No Username is set in storage handler")
	EmptyDeviceString             = errors.New("device string is empty")
	NotLoggedIn                   = errors.New("User not logged in")
	NoDataForStreamCreation       = errors.New("No data is provided for stream creation")
	InvalidRequestOptions         = errors.New("Invalid request options")
	RequestedResourceNotExist     = errors.New("Requested resource does not exist")
	NoResponseFromServer          = errors.New("No response from server. Either a connection or configuration error.")
	UnknownMessageObject          = errors.New("Unknown message object. Expected errors subarray but found something else. Please submit a ticket about needing an Instagram-API library update!")
	UnknownMessageType            = errors.New("Unknown message type. Please submit a ticket about needing an Instagram-API library update!")
	EmptyCookiesFilePath          = errors.New("Empty cookies file path")
	ThrottledResponse             = errors.New("Throttled by Instagram because of too many API requests.")
	RequestHeaderTooLargeResponse = errors.New("The request start-line and/or headers are too large to process.")
)

func InvalidBiography(bio string) error {
	return errors.New(fmt.Sprintf("Invalid biography %s, Please provide a 0 to 150 character string as biography.", bio))
}

func InvalidItem(item interface{}) error {
	return errors.New(fmt.Sprintf("Invalid item: %+v", item))
}

func InvalidHTTPStatus(statusCode int) error {
	return errors.New(fmt.Sprintf("Invalid http status code: %d", statusCode))
}

func UnknownItem(item interface{}) error {
	return errors.New(fmt.Sprintf("Unknown item %+v", item))
}

func UnknownSearchType(t string) error {
	return errors.New(fmt.Sprintf("Unknown search type %s", t))
}

func BadPushChannel(channel string) error {
	return errors.New(fmt.Sprintf("Bad push channel %s", channel))
}

func InvalidAppRefreshInterval(interval int) error {
	return errors.New(fmt.Sprintf("Invalid app refresh interval %d , it should be a positive number.", interval))
}

func TooHighAppRefreshInterval(interval int) error {
	return errors.New(fmt.Sprintf("Instagram's app state refresh interval is NOT allowed to be higher than 6 hours, and the lower the better! , given: %d", interval))
}

func InvalidLoginResponse(message string) error {
	return errors.New(fmt.Sprintf("Invalid login response provided to updateLoginState(). %s", message))
}

func CannotMarshalJSON(structVal interface{}, err string) error {
	return errors.New(fmt.Sprintf("Cannot unmarshal %s to JSON : %s", structVal, err))
}

func CookiesFileNotWritable(filepath, err string) error {
	return errors.New(fmt.Sprintf("Cookies file %s is not writable : %s", filepath, err))
}

func CannotDeleteCookiesFile(filepath, err string) error {
	return errors.New(fmt.Sprintf("Cannot delete cookies file %s : %s", filepath, err))
}

func ParameterMustBeString(param string) error {
	return errors.New(fmt.Sprintf("Parameter is not a string: %s", param))
}

func ErrorBuildingHTTPRequest(err string) error {
	return errors.New(fmt.Sprintf("Error building HTTP request: %s", err))
}

func ErrorGettingHTTPResponse(err string) error {
	return errors.New(fmt.Sprintf("Error getting HTTP response: %s", err))
}

func ErrorReadingHTTPResponseBody(err string) error {
	return errors.New(fmt.Sprintf("Error reading HTTP response body: %s", err))
}

func CannotCreateFormFieldFromFile(filepath, err string) error {
	return errors.New(fmt.Sprintf("Error creating form field from path %s: %s", filepath, err))
}

func CannotOpenFile(filepath, err string) error {
	return errors.New(fmt.Sprintf("Error openning path %s: %s", filepath, err))
}

func NotSupportedApiVersion(version int) error {
	return errors.New(fmt.Sprintf("Api version %d is not supported", version))
}

func NotEnoughDeviceStringResolution(deviceString, minimum string) error {
	return errors.New(fmt.Sprintf("device string %s does not meet the minimum required resolution %s for Instagram", deviceString, minimum))
}

func NotEnoughDeviceStringVersion(deviceString, minimum string) error {
	return errors.New(fmt.Sprintf("device string %s does not meet the minimum required version %s for Instagram", deviceString, minimum))
}

func InvalidDeviceFormat(deviceString string) error {
	return errors.New(fmt.Sprintf("Invalid device format: %s", deviceString))
}

func EmptyParameter(param string) error {
	return errors.New(fmt.Sprintf("Parameter %s cannot be empty", param))
}

func NotValidPersistentKey(key string) error {
	return errors.New(fmt.Sprintf("Key %s is not a valid persistent key", key))
}

func SettingsKeyNotFound(key string) error {
	return errors.New(fmt.Sprintf("Key %s not found in user settings", key))
}

func PathNotExist(path string) error {
	return errors.New(fmt.Sprintf("Path %s does not exist.", path))
}

func PathAlreadyExists(path string) error {
	return errors.New(fmt.Sprintf("Path %s exists already.", path))
}

func CreateFolder(err string) error {
	return errors.New(fmt.Sprintf("Error creating folder: %s", err))
}

func MoveFile(oldPath, newPath, err string) error {
	return errors.New(fmt.Sprintf("Error moving path %s to %s: %s", oldPath, newPath, err))
}

func ReadFile(filePath, err string) error {
	return errors.New(fmt.Sprintf("Error reading storage-file %s: %s", filePath, err))
}

func InvalidStorageVersion(version int) error {
	return errors.New(fmt.Sprintf("Invalid storage version %d", version))
}

func UnknownSettingsStorageType(sType string) error {
	return errors.New(fmt.Sprintf("Unknown Settings Storage Type: %s", sType))
}
