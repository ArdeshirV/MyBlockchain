package main

const (
	Normal      = "\033[0m"
	Bold        = "\033[1m"
	Red         = "\033[0;31m"
	Teal        = "\033[0;36m"
	White       = "\033[0;37m"
	Blue        = "\033[0;34m"
	Green       = "\033[0;32m"
	Yellow      = "\033[0;33m"
	Magenta     = "\033[0;35m"
	WhiteBold   = "\033[1;0m"
	RedBold     = "\033[1;31m"
	BlueBold    = "\033[1;34m"
	TealBold    = "\033[1;36m"
	GreenBold   = "\033[1;32m"
	YellowBold  = "\033[1;33m"
	MagentaBold = "\033[1;35m"
)

func NormalText(text string) string {
	return Normal + text + Normal
}

func WhiteText(text string) string {
	return NormalText(text)
}

func NormalBoldText(text string) string {
	return Bold + text + Normal
}

func WhiteBoldText(text string) string {
	return NormalBoldText(text)
}

func RedText(text string) string {
	return RedBold + text + Normal
}

func BlueText(text string) string {
	return Blue + text + Normal
}

func TealText(text string) string {
	return Teal + text + Normal
}

func GreenText(text string) string {
	return Green + text + Normal
}

func YellowText(text string) string {
	return Yellow + text + Normal
}

func MagentaText(text string) string {
	return Magenta + text + Normal
}

func BoldText(text string) string {
	return Bold + text + Normal
}

func RedBoldText(text string) string {
	return RedBold + text + Normal
}

func BlueBoldText(text string) string {
	return BlueBold + text + Normal
}

func TealBoldText(text string) string {
	return TealBold + text + Normal
}

func GreenBoldText(text string) string {
	return GreenBold + text + Normal
}

func YellowBoldText(text string) string {
	return YellowBold + text + Normal
}

func MagentaBoldText(text string) string {
	return MagentaBold + text + Normal
}

func Prompt(text string) string {
	return Bold + text + Normal
}

func Out(text string) string {
	return MagentaBold + text + Normal
}

func In(text string) string {
	return GreenBold + text + Normal
}
