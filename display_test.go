package display

import "testing"

func TestGetDisplayBrightness(t *testing.T) {

	display := MainDisplay()
	brightness := display.Brightness()

	if !(brightness >= 0.0 && brightness <= 1.0) {
		t.Fatalf("brightness is invalid")
	}
}

func TestSetDisplayBrightness(t *testing.T) {
	display := MainDisplay()
	display.SetBrightness(1.0)

	brightness := display.Brightness()

	if brightness != 1.0 {
		t.Fatalf("brightness is invalid: %f", brightness)
	}
}
