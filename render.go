package main;

import (
//    "gtop/module"
    "gtop/core"
    "fmt"
)


func RenderScreenFrame(screenSize core.Size) {
    
    // top left
    core.MoveCursor(0,2)
    
    horizontal := make([]rune,screenSize.Col-3)
    
    for i:=0; i < screenSize.Col-3; i++ {
        horizontal[i] = core.BORDER_HORIZONTAL
    }

    fmt.Printf(
        "%c%s%c",
        core.BORDER_TOP_LEFT,
        string(horizontal),
        core.BORDER_TOP_RIGHT,
    )

    // Bottom 
    core.MoveCursor(screenSize.Row,2) // let place for menu bar
    fmt.Printf(
        "%c%s%c",
        core.BORDER_BOTTOM_LEFT,
        string(horizontal),
        core.BORDER_BOTTOM_RIGHT,
    )

    // right left
    for i:=2; i<screenSize.Row; i++ {
        core.MoveCursor(i,2)
        fmt.Printf("%c",core.BORDER_VERTICAL)
        core.MoveCursor(i,screenSize.Col)
        fmt.Printf("%c",core.BORDER_VERTICAL)
    }

}


//func RenderEntity(e module.Entity, p core.Position) {
//    
//    core.MoveCursor(p.Row+1,p.Col+1) // hardcoded padding
//
//    fmt.Printf(
//        "%c[%d;%d;%dm%s%c[%dm",
//        ESC,
//        e.GraphicMode,
//        e.ForegroundColor,
//        e.BackgroundColor,
//        e.Data,
//        ESC,
//        GM_RESET,
//    )
//}
