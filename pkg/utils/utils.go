package utils

import (
	"strings"
)

func SliceToMap(s []string) map[string]string {
	m := make(map[string]string)
	for _, v := range s {
		k := strings.Split(v, "=")
		if len(k) != 2 {
			m[k[0]] = ""
		} else {
			m[k[0]] = k[1]
		}
	}
	return m
}
