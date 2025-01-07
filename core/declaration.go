package core;

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

// Character
const (
    
    PBAR_EMPTY = '▢'
    PBAR_FULL = '■'
    
    BORDER_HORIZONTAL =     '─'
    BORDER_VERTICAL =       '│'
    BORDER_TOP_LEFT =       '┌'
    BORDER_TOP_RIGHT =      '┐'
    BORDER_BOTTOM_LEFT =    '└'
    BORDER_BOTTOM_RIGHT =   '┘'
    BORDER_MIDDLE_TOP =     '┬'
    BORDER_MIDDLE_BOTTOM =  '┴'
    BORDER_MIDDLE_CENTER =  '┼'
)


