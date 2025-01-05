package main

import (
    "math"
    "gtop/core"
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


