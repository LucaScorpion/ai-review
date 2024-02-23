package util

func Must[T any](val T, err error) T {
	PanicIfError(err)
	return val
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
