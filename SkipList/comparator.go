package SkipList

/********************************************************
* @author: Ihc
* @date: 2022/4/20 0020 10:19
* @version: 1.0
* @description: 内置类型比较大小
*********************************************************/

// BuiltinTypeComparator 内置类型比较大小; m和n为内置类型
// 若m > n,则返回 1; 若m < n,则返回-1; 若m = n,则返回0
func BuiltinTypeComparator(m, n interface{}) int {
	if m == n {
		return 0
	}
	switch m.(type) {
	case uint8, uint16, uint32, uint64, uint, int8, int16, int32, int64, int, uintptr:
		return compareInt(m, n)
	case float32:
		if m.(float32) < n.(float32) {
			return -1
		}
	case float64:
		if m.(float64) < n.(float64) {
			return -1
		}
	case bool:
		if m.(bool) == false && n.(bool) == true {
			return -1
		}
	case string:
		if m.(string) < n.(string) {
			return -1
		}
	case complex64:
		return compareComplex64(m.(complex64), n.(complex64))
	case complex128:
		return compareComplex128(m.(complex128), n.(complex128))
	}
	return 1
}

// compareComplex64 比较复数
func compareComplex64(m complex64, n complex64) int {
	if real(m) < real(n) {
		return -1
	}
	if real(m) == real(n) && imag(m) < imag(n) {
		return -1
	}
	return 1
}

// compareComplex128 比较复数
func compareComplex128(m complex128, n complex128) int {
	if real(m) < real(n) {
		return -1
	}
	if real(m) == real(n) && imag(m) < imag(n) {
		return -1
	}
	return 1
}

// compareInt 比较int类型
func compareInt(m, n interface{}) int {
	return -1
}
