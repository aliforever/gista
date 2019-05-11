package devices

type DeviceInterface interface {
	GetDeviceString() string
	SetDeviceString(device string)
	GetManufacturer() string
	GetBrand() *string
	GetAndroidVersion() string
	GetAndroidRelease() string
	/*GetFBUserAgent() string*/
	GetDPI() string
	GetResolution() string
	GetModel() string
	GetDevice() string
	GetCPU() string
	GetUserAgent() string
}
