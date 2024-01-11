package util

import "github.com/sundonghui/stool/consts"

func Max[T consts.Int | consts.Uint | consts.Float | rune | byte | uintptr](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T consts.Int | consts.Uint | consts.Float | rune | byte | uintptr](a, b T) T {
	if a > b {
		return b
	}
	return a
}
