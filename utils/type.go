package utils

func NilCheck(c interface{}) interface{} {
	if c == nil {
		return nil
	} else {
		return c
	}
}
