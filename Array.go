package common

import (
	"strings"
)

type ArrayUtil struct{}

var Array = new(ArrayUtil)

// []string数组转字符串
func (a *ArrayUtil) StringArrayOffsetToString(arr []string, need string, offset int) string {
	if offset != 0 {
		arrLen := len(arr)
		newArr := arr[:arrLen-offset]
		newStr := strings.Join(newArr, need)
		return newStr
	}

	return ""
}

// 反转[]rune数组
func (a *ArrayUtil) ReverseRune(runes []rune) []rune {
	for front, back := 0, len(runes)-1; front < back; front, back = front+1, back-1 {
		runes[front], runes[back] = runes[back], runes[front]
	}
	return runes
}

// 反转[]string数组
func (a *ArrayUtil) ReverseStringArray(arr []string) []string {
	for front, back := 0, len(arr)-1; front < back; front, back = front+1, back-1 {
		arr[front], arr[back] = arr[back], arr[front]
	}
	return arr
}

// 比较int切片是否相等
func (a *ArrayUtil) EqIntSlice(source, target []int) bool {
	if (source == nil) != (target == nil) {
		return false
	}

	if len(source) != len(target) {
		return false
	}

	for i := range source {
		if source[i] != target[i] {
			return false
		}
	}

	return true
}

// 比较string切片是否相等
func (a *ArrayUtil) EqStringSlice(source, target []string) bool {
	if (source == nil) != (target == nil) {
		return false
	}

	if len(source) != len(target) {
		return false
	}

	for i := range source {
		if source[i] != target[i] {
			return false
		}
	}

	return true
}
