package algorithms

import (
	"testing"
)

func testReverseStrings(t *testing.T) {
	strs := make(map[string]string)
	strs["macbooks!"] = "!skoobcam"
	strs["face the facts"] = "stcaf eht ecaf"
	strs["shaq was the best actor to have ever lived"] = "devil reve evah ot rotca tseb eht saw qahs"

	for input, expected := range strs {
		actual, err := reverse(input)
		if err != nil {
			t.Errorf("This function failed: %s", err)
		}
		if actual != expected {
			t.Errorf("Expected [%s]; got [%s]", expected, actual)
		}
	}
}
