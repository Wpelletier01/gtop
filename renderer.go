package main;

import (
    "gtop/module"
    "fmt"
)

const (
    ESC = '\x1b'
)

// Graphics Mode
const (
    GM_RESET = iota
    GM_BOLD
    GM_DIM
    GM_ITALIC
    GM_UNDERLINE
    GM_BLINKING
    GM_NULL // not used
    GM_REVERSE
    GM_STRIKE_THROUGH

)

// COLOR
const (
    FG_BLACK = iota + 30
    FG_RED
    FG_GREEN
    FG_YELLOW
    FG_BLUE
    FG_MAGENTA
    FG_CYAN
    FG_WHITE
    FG_EMPTY
    FG_DEFAULT // not used
    BG_BLACK
    BG_RED
    BG_GREEN
    BG_YELLOW
    BG_BLUE
    BG_MAGENTA
    BG_CYAN
    BG_WHITE
    BG_EMPTY // not used
    BG_DEFAULT
)



func RenderEntity(e module.Entity) string {
    
    return fmt.Sprintf(
        "%c[%d;%d;%dm%s%c[%dm",
        ESC,
        e.GraphicMode,
        e.ForegroundColor,
        e.BackgroundColor,
        e.Data,
        ESC,
        GM_RESET,
    )
}
