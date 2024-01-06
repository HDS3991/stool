package util

import "github.com/sundonghui/stool/consts"

func Contains[T consts.Number | string](list []T, element T) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}
