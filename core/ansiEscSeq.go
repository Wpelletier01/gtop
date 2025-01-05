package core;

import "fmt"

func HideCursor() {
    fmt.Print("\033[?25l")
}

func ShowCursor() {
    fmt.Print("\033[?25h")
}

// go to a buffer for keeping intact the current buffer where gtop
// is executed
func EnableAltBuffer() {        
    fmt.Print("\033[?1049h")
}

// return to the buffer where gtop was executed 
func DisableAltBuffer() {
    fmt.Print("\033[?1049l")
}

// move the cursor to the home position (usually top left corner of the terminal) 
func GoHome() {
    fmt.Print("\033[H")
}

// clear the full visible terminal screen
func ClearScreen() {
    fmt.Print("\033[3J")
}
