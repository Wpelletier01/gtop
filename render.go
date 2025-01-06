package main;

import (
    "gtop/module"
    "gtop/core"
    "fmt"
    "math"
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

const (
    
    PBAR_EMPTY = '▱'
    PBAR_FULL = '▰'
    
    BORDER_HORIZONTAL =     '─'
    BORDER_VERTICAL =       '│'
    BORDER_TOP_LEFT =       '┌'
    BORDER_TOP_RIGHT =      '┐'
    BORDER_BOTTOM_LEFT =    '└'
    BORDER_BOTTOM_RIGHT =   '┘'

)

func AddBorder(buf []rune, s core.Size) ([]rune) { 

    // top 
    for i:=1; i<s.Col-1; i++ {
        buf[i] = BORDER_HORIZONTAL
    }
    // bottom
    for i:=(s.Row-1)*s.Col+1; i<(s.Row*s.Col)-2; i++ {
        buf[i] = BORDER_HORIZONTAL
    }
    // left
    for i:=1; i<s.Row-1; i++ {
        buf[i*s.Col] = BORDER_VERTICAL
        buf[(i*s.Col)+(s.Col-2)] = BORDER_VERTICAL
    }
 
    // set corner
    buf[0] = BORDER_TOP_LEFT
    buf[s.Col-2] = BORDER_TOP_RIGHT
    buf[s.Col*(s.Row-1)] = BORDER_BOTTOM_LEFT
    buf[(s.Col*s.Row)-2] = BORDER_BOTTOM_RIGHT

    return buf

}


func CreateProgressBar(length int, percent float64) []rune {
    
    pbar := make([]rune,length)
    
    var nbFull int = int(math.Round(percent * float64(length)))
    for i:=0; i<nbFull; i++ {
        pbar[i] = PBAR_FULL
    }

    for i:=nbFull; i<length; i++ {
        pbar[i] = PBAR_EMPTY
    }
     
    return pbar
}


func RenderScreenFrame(screenSize core.Size) {
    
    // top left
    core.MoveCursor(0,2)
    
    horizontal := make([]rune,screenSize.Col-3)
    
    for i:=0; i < screenSize.Col-3; i++ {
        horizontal[i] = BORDER_HORIZONTAL
    }

    fmt.Printf(
        "%c%s%c",
        BORDER_TOP_LEFT,
        string(horizontal),
        BORDER_TOP_RIGHT,
    )

    // Bottom 
    core.MoveCursor(screenSize.Row,2) // let place for menu bar
    fmt.Printf(
        "%c%s%c",
        BORDER_BOTTOM_LEFT,
        string(horizontal),
        BORDER_BOTTOM_RIGHT,
    )

    // right left
    for i:=2; i<screenSize.Row; i++ {
        core.MoveCursor(i,2)
        fmt.Printf("%c",BORDER_VERTICAL)
        core.MoveCursor(i,screenSize.Col)
        fmt.Printf("%c",BORDER_VERTICAL)
    }

}


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
