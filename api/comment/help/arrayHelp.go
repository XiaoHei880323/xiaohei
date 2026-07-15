package helper

type ArrayHelpFunc struct {
}

var ArrayHelpFuncObject ArrayHelpFunc

func (h *ArrayHelpFunc) InArray(val int, arrayInt []int) bool {
	if len(arrayInt) < 1 {
		return false
	}
	for _, v := range arrayInt {
		if v == val {
			return true
		}
	}
	return false
}
