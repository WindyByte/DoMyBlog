package utils

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	return i, err
}
func StringToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return i, err
}
