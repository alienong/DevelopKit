/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/17 14:29
 */

package misc

import "strconv"

/**
@description: String To Float64
@param s: string
@return: Float64
*/
func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

/**
@description:Float64 To String
@param f: Float64
@return: string
*/
func Float64ToString(f float64) string {
	s := strconv.FormatFloat(f, 'E', -1, 64)
	return s
}

/**
@description:Add Pad At The End of String Util The Length of String Equals length
@param str: The Original String
@param pad: The Added String
@param length: total length of string
@return: string
*/
func StringPadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

/**
@description:Add Pad At The Start of String Util The Length of String Equals length
@param str: The Original String
@param pad: The Added String
@param length: total length of string
@return: string
*/
func StringPadLeft(str, pad string, length int) string {
	for {
		if len(str) >= length {
			return str[0:length]
		}
		str = pad + str
	}
}
