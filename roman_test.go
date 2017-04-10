package roman

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"strconv"
)

func TestTo(t *testing.T) {

	/*
		iM = 1000
	iD = 500
	iC = 100
	iL = 50
	iX = 10
	iV = 5
	iI = 1

	sM = "M"
	sD = "D"
	sC = "C"
	sL = "L"
	sX = "X"
	sV = "V"
	sI = "I"
	 */

	for input, want := range map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IIII",
		5: "V",
		6: "VI",
		9: "VIIII",
		10: "X",
		11: "XI",
		15: "XV",
		50: "L",
		51: "LI",
		100: "C",
		500: "D",
		1000: "M",



	} {
		Convey(fmt.Sprintf("A value of %s", strconv.Itoa(input)), t, func() {
			Convey(fmt.Sprintf("Will return %s", want), func() {
				r := To(input)
				So(r, ShouldEqual, want)
			})
		})
	}
}
