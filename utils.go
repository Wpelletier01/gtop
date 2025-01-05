package main;

import (
    "gtop/core"
)

func CreateEmptyBuffer(s core.Size) ([]rune) {
    
    area := s.Area()

    nBuf := make([]rune, area, area)

    for i:=0; i<s.Col*s.Row; i++ {
        nBuf[i] = ' '
    } 

    return nBuf

}


