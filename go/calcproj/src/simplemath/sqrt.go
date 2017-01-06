package simplemath;

import "math";
func Sqrt(a int) int {
	v := math.Sqrt(float64(a));
	return int(v);
}
