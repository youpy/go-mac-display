package display

/*
#cgo LDFLAGS: -framework IOKit -framework ApplicationServices
#include <CoreFoundation/CoreFoundation.h>
#include <ApplicationServices/ApplicationServices.h>
#include <IOKit/graphics/IOGraphicsLib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Display struct {
	id int
}

func MainDisplay() *Display {
	display := Display{int(C.CGMainDisplayID())}

	return &display
}

func DoWithBrightness(callback func(key *C.char)) {
	key := C.CString(C.kIODisplayLinearBrightnessKey)

	callback(key)

	// http://code.google.com/p/go-wiki/wiki/cgo#Go_strings_and_C_strings
	defer C.free(unsafe.Pointer(key))
}

func (display Display) Brightness() float64 {
	value := C.float(0.0)

	DoWithBrightness(func(key *C.char) {
		C.IODisplayGetFloatParameter(
			C.CGDisplayIOServicePort(C.CGDirectDisplayID(display.id)),
			C.kNilOptions,
			C.CFStringCreateWithCString(nil, key, C.kCFStringEncodingMacRoman),
			&value)
	})

	return float64(value)
}

func (display Display) SetBrightness(value float64) {
	DoWithBrightness(func(key *C.char) {
		C.IODisplaySetFloatParameter(
			C.CGDisplayIOServicePort(C.CGDirectDisplayID(display.id)),
			C.kNilOptions,
			C.CFStringCreateWithCString(nil, key, C.kCFStringEncodingMacRoman),
			C.float(value))
	})
}
