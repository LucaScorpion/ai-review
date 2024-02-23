package ansi

const esc = "\x1B"
const reset = "0"

func mode(mode string) string {
	return esc + "[" + mode + "m"
}

func Reset() string {
	return mode(reset)
}
