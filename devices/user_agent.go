package devices

import (
	"fmt"

	"github.com/aliforever/gista/constants"
)

const userAgentFormat string = "Instagram %s Android (%s/%s; %s; %s; %s; %s; %s; %s; %s; %s)"

func BuildUserAgent(appVersion, userLocale string, device DeviceInterface) string {
	manufacturerWithBrand := device.GetManufacturer()
	if device.GetBrand() != nil {
		manufacturerWithBrand += "/" + *device.GetBrand()
	}
	return fmt.Sprintf(userAgentFormat,
		appVersion, device.GetAndroidVersion(),
		device.GetAndroidRelease(),
		device.GetDPI(),
		device.GetResolution(),
		manufacturerWithBrand,
		device.GetModel(),
		device.GetDevice(),
		device.GetCPU(),
		userLocale,
		constants.VersionCode,
	)
}
