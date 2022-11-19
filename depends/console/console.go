package Console

import (
	"fmt"
	"runtime"
)

// 普通前景色
const (
	FgBlack   = iota + 30 // 30
	FgRed                 // 31
	FgGreen               // 32
	FgYellow              // 33
	FgBlue                // 34
	Fgmagenta             // 35
	FgCyan                // 36
	FgWhite               // 37
)

// 高亮度的前景色
const (
	FgHiBlack = iota + 90 // 90
	FgHiRed               // 91
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHimagenta
	FgHiCyan
	FgHiWhite
)

// 普通背景色
const (
	BgBlack   = iota + 40 // 40 黑色
	BgRed                 // 41 红色
	BgGreen               // 42 绿色
	BgYellow              // 43 黄色
	BgBlue                // 44 蓝色
	Bgmagenta             // 45 紫色
	BgCyan                // 46 青色
	BgWhite               // 47 白色
)

// 高亮度的背景色
const (
	BgHiBlack   = iota + 100 // 100
	BgHiRed                  // 101
	BgHiGreen                // 102
	BgHiYellow               // 103
	BgHiBlue                 // 104
	BgHimagenta              // 105
	BgHiCyan                 // 106
	BgHiWhite                // 107
)

func Highlight(msg string, bg int, fg int) string {
	if "windows" == runtime.GOOS {
		return msg
	}
	return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", bg, fg, msg)
}

func Error(a ...interface{}) string {
	return Highlight(fmt.Sprint(a...), BgBlack, FgRed)
}
func Succeed(a ...interface{}) string {
	return Highlight(fmt.Sprint(a...), BgBlack, FgGreen)
}
func Info(a ...interface{}) string {
	return Highlight(fmt.Sprint(a...), BgBlack, FgGreen)
}

func LogError(a ...interface{}) {
	fmt.Print(Highlight(fmt.Sprint(a...), BgBlack, FgRed))
}
func LogSucceed(a ...interface{}) {
	fmt.Print(Highlight(fmt.Sprint(a...), BgBlack, FgGreen))
}
func LogInfo(a ...interface{}) {
	fmt.Print(Highlight(fmt.Sprint(a...), BgBlack, FgWhite))
}
func Log(a ...interface{}) {
	fmt.Print(fmt.Sprint(a...))
}

func LogErrorln(a ...interface{}) {
	fmt.Println(Highlight(fmt.Sprint(a...), BgBlack, FgRed))
}
func LogSucceedln(a ...interface{}) {
	fmt.Println(Highlight(fmt.Sprint(a...), BgBlack, FgGreen))
}
func LogInfoln(a ...interface{}) {
	fmt.Println(Highlight(fmt.Sprint(a...), BgBlack, FgWhite))
}
func Logln(a ...interface{}) {
	fmt.Println(fmt.Sprint(a...))
}
