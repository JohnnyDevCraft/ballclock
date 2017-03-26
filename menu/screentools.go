package menu

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

//ClearScreen clears the contents of the console window.
func ClearScreen() {
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Windows example it is untested, but I think its working
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	val, ok := clear[runtime.GOOS]
	if ok {
		val()
	}
}

//StringIsInt returns true if the string can be converted to an int.
func StringIsInt(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

//StringToInt converts a string to an integer
func StringToInt(s string) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return 0
}
