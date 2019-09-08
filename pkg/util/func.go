package util

import "time"

// ConvertUUnixtimeToString func
func ConvertUUnixtimeToString(in int64) (out string) {
	second := in
	t := time.Unix(second, 0)
	out = t.Format("2006-01-02 15:04:05")
	return
}
