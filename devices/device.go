package devices

import (
	"strconv"
	"strings"

	"github.com/mcuadros/go-version"

	good_devices "github.com/aliforever/gista/devices/good-devices"
	"github.com/aliforever/gista/errors"
)

const requiredAndroidVersion = "2.2"

type Device struct {
	appVersion     string
	versionCode    string
	userLocale     string
	deviceString   string
	userAgent      string
	fbUserAgents   []string
	androidVersion string
	androidRelease string
	dpi            string
	resolution     string
	manufacturer   string
	brand          *string
	model          string
	device         string
	cpu            string
}

func (d *Device) SetDeviceString(device string) {
	d.deviceString = device
}

func (d *Device) GetDeviceString() string {
	return d.deviceString
}

func (d *Device) initFromDeviceString(deviceString *string) error {
	if deviceString == nil || *deviceString == "" {
		return errors.EmptyDeviceString
	}
	parts := strings.Split(*deviceString, ";")
	if len(parts) != 7 {
		return errors.InvalidDeviceFormat(*deviceString)
	}
	androidOS := strings.Split(parts[0], "/")
	if version.Compare(androidOS[1], requiredAndroidVersion, "<") {
		return errors.NotEnoughDeviceStringVersion(*deviceString, requiredAndroidVersion)
	}
	resolution := strings.SplitN(parts[2], "x", 2)
	x, _ := strconv.Atoi(strings.TrimSpace(resolution[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(resolution[1]))
	pixelCount := x * y
	if pixelCount < 2073600 {
		return errors.NotEnoughDeviceStringResolution(*deviceString, "1920x1080")
	}
	manufacturerAndBrand := strings.SplitN(parts[3], "/", 2)
	d.deviceString = *deviceString
	d.androidVersion = androidOS[0]
	d.androidRelease = androidOS[1]
	d.dpi = parts[1]
	d.resolution = parts[2]
	d.manufacturer = manufacturerAndBrand[0]
	if len(manufacturerAndBrand) > 1 {
		d.brand = &manufacturerAndBrand[1]
	} else {
		d.brand = nil
	}
	d.model = parts[4]
	d.device = parts[5]
	d.cpu = parts[6]
	d.userAgent = BuildUserAgent(d.appVersion, d.userLocale, d)
	return nil
}

func (d *Device) GetManufacturer() string {
	return d.manufacturer
}

func (d *Device) GetBrand() *string {
	return d.brand
}

func (d *Device) GetAndroidVersion() string {
	return d.androidVersion
}

func (d *Device) GetAndroidRelease() string {
	return d.androidRelease
}

func (d *Device) GetDPI() string {
	return d.dpi
}

func (d *Device) GetResolution() string {
	return d.resolution
}

func (d *Device) GetModel() string {
	return d.model
}

func (d *Device) GetDevice() string {
	return d.device
}

func (d *Device) GetCPU() string {
	return d.cpu
}

func (d *Device) GetUserAgent() string {
	return d.userAgent
}

func NewDevice(appVersion, versionCode, userLocale string, deviceString *string, autoFallback bool /*true*/) *Device {
	d := Device{}
	d.appVersion = appVersion
	d.versionCode = versionCode
	d.userLocale = userLocale
	if autoFallback && (deviceString == nil || !good_devices.IsGoodDevice(*deviceString)) {
		random := good_devices.GetRandomGoodDevice()
		deviceString = &random
	}
	d.initFromDeviceString(deviceString)
	return &d
}
