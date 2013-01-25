package display

/*
#cgo LDFLAGS: -framework IOKit -framework ApplicationServices -fconstant-cfstrings
#include <CoreFoundation/CoreFoundation.h>
#include <ApplicationServices/ApplicationServices.h>
#include <IOKit/graphics/IOGraphicsLib.h>
*/
import "C"

type Display struct {
	id int
}

func MainDisplay() *Display {
	display := Display{int(C.CGMainDisplayID())}

	return &display
}

func (display Display) Brightness() float64 {
	value := C.float(0.0)

	C.IODisplayGetFloatParameter(
		C.CGDisplayIOServicePort(C.CGDirectDisplayID(display.id)),
		C.kNilOptions,
		C.CFStringCreateWithCString(nil, C.CString(C.kIODisplayLinearBrightnessKey), C.kCFStringEncodingMacRoman),
		&value)

	return float64(value)
}

func (display Display) SetBrightness(value float64) {
	C.IODisplaySetFloatParameter(
		C.CGDisplayIOServicePort(C.CGDirectDisplayID(display.id)),
		C.kNilOptions,
		C.CFStringCreateWithCString(nil, C.CString(C.kIODisplayLinearBrightnessKey), C.kCFStringEncodingMacRoman),
		C.float(value))
}
