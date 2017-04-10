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
		//1: "I",
		//2: "II",
		//3: "III",
		//4: "IIII",  // no reverse
		4: "IV",  // reverse
		//5: "V",
		//6: "VI",
		//9: "VIIII", 	// no reverse
		9: "IX", 	// reverse
		//10: "X",
		//11: "XI",
		//15: "XV",
		40: "XL",
		//50: "L",
		//51: "LI",
		90: "XC",
		//100: "C",
		400: "CD",
		//500: "D",
		900: "CM",
		//1000: "M",
		//1050: "ML",
		//1051: "MLI",
		//1100: "MC",
		//1600: "MDC",
		//2000: "MM",




	} {
		Convey(fmt.Sprintf("A value of %s", strconv.Itoa(input)), t, func() {
			Convey(fmt.Sprintf("Will return %s", want), func() {
				r := To(input)
				So(r, ShouldEqual, want)
			})
		})
	}
}
