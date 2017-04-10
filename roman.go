package roman

//import "fmt"

// allow conversion from int to roman numeral string
// allow conversion from roman numeral string to int
// fully documented
// fully tested
// must go up to integer 100
// bonus: allow mathmatical functions on roman numerals

/*
Symbol	I	V	X	L	C	D	M
Value	1	5	10	50	100	500	1,000

In a few specific cases, to avoid four characters being repeated in succession (such as IIII or XXXX),
subtractive notation is used: as in this table:

Number		4	9	40	90	400	900
Notation	IV	IX	XL	XC	CD	CM
 */

const (
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
)

// TODO: turn string into rune
func printNumerals(char string, freq int) string {
	//fmt.Println(fmt.Sprintf("printing %s %d times", char, freq))

	out := ""
	for i := 0; i < freq; i++ {
		out += char
	}
	return out
}

func To(in int) string {
	//fmt.Println(fmt.Sprintf("processing %s", in))

	if m := in / iM; m > 0 {
		return printNumerals(sM, m) + To(in % iM)
	} else if m := in / iD; m > 0 {
		return printNumerals(sD, m) + To(in % iD)
	} else if m := in / iC; m > 0 {
		return printNumerals(sC, m) + To(in % iC)
	} else if m := in / iL; m > 0 {
		return printNumerals(sL, m) + To(in % iL)
	} else if m := in / iX; m > 0 {
		return printNumerals(sX, m) + To(in % iX)
	} else if m := in / iV; m > 0 {
		return printNumerals(sV, m) + To(in % iV)
	} else if m := in / iI; m > 0 {
		return printNumerals(sI, m) + To(in % iI)
	}

	return ""
}

func From(in string) int {
	out := 0
	return out
}
