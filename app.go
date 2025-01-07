package main

import (
    "gtop/core"
    //"gtop/module"
)

type App struct {
    size        core.Size
    event       core.EventRegister             
}  



func (a *App) Run() {
    

    // start event listener
    a.event.StartListener()

    main:
    for {
        
        select {
            
            case nSize := <-a.event.Resize:

                a.resize(nSize)
            
            case input := <-a.event.Input:
                if input[0] == 27 { // maybe an ansi esc seq 
                    // ESC key press hard quit
                    if input[1] == 0 {
                        break main
                    }
                    // rest other esc sequence
			        // arrow key
                    // [A up
                    // [B down
                    // [C right
                    // [D left
                }
                
        }
         

        
        a.draw()
    }



}


func (a *App) resize(s core.Size) { 
    
    // hardcoded padding
    s.Col -= 1
    s.Row -= 1
    
    a.size = s
    
}



func (a *App) draw() {
    
    core.ClearScreen()
    

    
    RenderScreenFrame(a.size)
//    fmt.Printf("%s",sb.String())
//    
//    e := module.Entity {
//        Data: "This is a test",
//        ForegroundColor: FG_WHITE,
//        BackgroundColor: BG_RED,
//        GraphicMode: GM_NULL,
//    }
//
//    RenderEntity(e,core.Position{Col:a.size.Col-15,Row:10,})

}

