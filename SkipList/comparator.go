package SkipList

/********************************************************
* @author: Ihc
* @date: 2022/4/20 0020 10:19
* @version: 1.0
* @description: 内置类型比较大小
*********************************************************/

// BuiltinTypeComparator 内置类型比较大小; m和n为内置类型并且m和n必须为同类型
// 非同类型数据比较会报错
// 若m > n,则返回 1; 若m < n,则返回-1; 若m = n,则返回0
func BuiltinTypeComparator(m, n interface{}) int {
	// 若m与n类型不同会报错
	// 可以获取其真实类型进行比较:reflect.TypeOf(m).Kind() == reflect.TypeOf(n).Kind()
	// 这里不做比较,使用此接口要注意
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
	switch m.(type) {
	case uint8:
		if m.(uint8) < n.(uint8) {
			return -1
		}
	case uint16:
		if m.(uint16) < n.(uint16) {
			return -1
		}
	case uint32:
		if m.(uint32) < n.(uint32) {
			return -1
		}
	case uint64:
		if m.(uint64) < n.(uint64) {
			return -1
		}
	case uint:
		if m.(uint) < n.(uint) {
			return -1
		}
	case int8:
		if m.(int8) < n.(int8) {
			return -1
		}
	case int16:
		if m.(int16) < n.(int16) {
			return -1
		}
	case int32:
		if m.(int32) < n.(int32) {
			return -1
		}
	case int64:
		if m.(int64) < n.(int64) {
			return -1
		}
	case int:
		if m.(int) < n.(int) {
			return -1
		}
	case uintptr:
		if m.(uintptr) < n.(uintptr) {
			return -1
		}
	}
	return 1
}
