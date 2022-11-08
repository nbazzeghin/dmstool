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

package cmd

import (
	"fmt"
	"github.com/nbazzeghin/dmstool/pkg/dmstool"
	"github.com/spf13/cobra"
)

var mapLinks bool

// deg2decCmd represents the deg2dec command
var deg2decCmd = &cobra.Command{
	Use:     "deg2dec \"coordinate\"",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"dd"},
	Short:   "Convert from Degrees Minutes Seconds to Decimal Degrees",
	Long: `Converts a coordinate from Degrees Minutes Seconds to Decimal Degrees.
Single and double coordinate are supported so long as they are seperated by a comma.

Example Inputs:
  "34 deg 58' 59.61\" N, 84 deg 20' 8.75\" W"
  or
  "110°4'21' W"

NOTE: Check your escape quotes.`,
	Run: func(cmd *cobra.Command, args []string) {
		res := dmstool.Deg2Dec(args[0], mapLinks)
		fmt.Println(res)
	},
}

func init() {
	deg2decCmd.Flags().BoolVarP(&mapLinks, "maplinks", "m", false, "Return google map links (if two coordinates have been passed)")
	rootCmd.AddCommand(deg2decCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deg2decCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deg2decCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
