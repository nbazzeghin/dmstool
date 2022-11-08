/*
Copyright © 2022 Nigel Bazzeghin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package dmstool

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Deg2Dec Converts DMS input string to Decimal Degrees.
// Can accept either single or double positions seperated
// with a comma.
//
// Example Inputs:
//
//	34 deg 58' 59.61" N, 84 deg 20' 8.75" W
//	or
//	110°4'21' W
func Deg2Dec(input string) (result string) {

	inputSplit := strings.Split(input, ",")

	if len(inputSplit) > 1 {
		r1 := Deg2Dec(inputSplit[0])
		r2 := Deg2Dec(inputSplit[1])
		return r1 + " " + r2
	}

	re := regexp.MustCompile(`s|w`)
	swCoord := re.MatchString(strings.ToLower(inputSplit[0]))

	f := 1
	if swCoord {
		f = -1
	}

	reDigits := regexp.MustCompile(`(?m)[^\d*\.?\d*]`)
	//re_sw := regexp.MustCompile(``)

	split := reDigits.Split(input, -1)
	set := []string{}

	for i := range split {
		if split[i] != "" {
			set = append(set, split[i])
		}
	}

	deg, _ := strconv.ParseFloat(set[0], 2)
	min, _ := strconv.ParseFloat(set[1], 2)
	sec, _ := strconv.ParseFloat(set[2], 2)

	var res float64 = 0
	res += deg / float64(f)
	res += min / float64(f*60)
	res += sec / float64(f*3600)
	r := fmt.Sprintf("%f", res)

	return r
}
