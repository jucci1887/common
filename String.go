package common

import (
	"strings"
)

type StrUtil struct{}

var String = new(StrUtil)

// 截取字符串
func (s *StrUtil) Sub(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

// 截取字符串
func (s *StrUtil) SubLen(str string, pos int, length int) string {
	runes := []rune(str)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}

	return string(runes[pos:l])
}

// 首字母大写
func (s *StrUtil) UcFirst(str string) string {
	if len(str) < 1 {
		return ""
	}

	strArray := []rune(str)
	if strArray[0] >= 97 && strArray[0] <= 122 {
		strArray[0] -= 32
	}

	return string(strArray)
}

// 英文字符串全部转大写
func (s *StrUtil) ToUpper(str string) string {
	return strings.ToUpper(str)
}

// 英文字符串全部转小写
func (s *StrUtil) ToLower(str string) string {
	return strings.ToLower(str)
}

// 分割字符串以获取字符串前缀
func (s *StrUtil) SplitGetPrefix(str string, need string) string {
	fileArray := strings.Split(str, need)
	length := len(fileArray)
	strArray := append(fileArray[:length-1], fileArray[(length-1)+1:]...)
	strLength := len(strArray)
	if strLength == 0 {
		return ""
	}

	return strArray[0]
}

// 分割字符串以获取字符串后缀
func (s *StrUtil) SplitGetSuffix(str string, need string) string {
	fileArray := strings.Split(str, need)
	length := len(fileArray)
	return fileArray[length-1]
}

// 分割字符串以获取字符串后缀之前的字符串
func (s *StrUtil) SplitGetBeforeSuffix(str string, need string, joinNeed string) string {
	fileArray := strings.Split(str, need)
	length := len(fileArray)
	newStr := strings.Join(fileArray[:length-1], joinNeed)
	return newStr
}

// 分割换行符以获取字符串
func (s *StrUtil) SplitLineBreak(path string) string {
	return s.SplitString(path, "\n")
}

// 分割字符串
func (s *StrUtil) SplitString(path string, need string) string {
	return s.SubLen(path, 0, strings.LastIndex(path, need))
}

// 反转字符串
func (s *StrUtil) Reverse(str string) string {
	toRune := []rune(str)
	reverse := Array.ReverseRune(toRune)
	return string(reverse)
}
