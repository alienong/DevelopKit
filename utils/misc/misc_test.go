/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/5/29 14:37
 */

package misc

import (
	"fmt"
	"testing"
)

func TestSliceArray(t *testing.T) {
	test1 := []int{1, 2, 3, 45, 6, 7, 8, 9}
	err, arr := SliceArray(3, &test1, make([][]int, 0))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(arr.([][]int))
}
