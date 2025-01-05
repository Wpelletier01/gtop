package main

import (
    "fmt"
    "strings"
    "gtop/core"
)

type App struct {
    size        core.Size
    buffer      []rune
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
        
        a.addToBuffer(
            CreateProgressBar(7,0.5),
            core.Position{ X: 44, Y: 20},
            core.Size{ Col: 7, Row: 1}, 
        )

        a.draw()
    }




}


func (a *App) resize(s core.Size) { 
    
    // hardcoded padding
    s.Col -= 1
    s.Row -= 1
    
    a.size = s
    
    a.buildBuffer()
}

func (a *App) buildBuffer() {
     
    a.buffer = CreateEmptyBuffer(a.size)
    // add terminal border 
    a.buffer = AddBorder(a.buffer,a.size)
    
}


func (a *App) draw() {
    
    core.ClearScreen()
    
    var sb strings.Builder
    for i:=0; i<a.size.Row; i++ {
        sb.WriteString(" ") // for padding
        sb.WriteString(string(a.buffer[i*a.size.Col:i*a.size.Col + a.size.Col-1])) 
        sb.WriteString("\r\n") // in Raw mode, a new line is indicated by this
    }
    // start to paint at the top of the string
    core.GoHome()
    
    fmt.Printf("%s",sb.String())
}

func (a *App) addToBuffer(b []rune, pos core.Position, size core.Size) {
    
    // not draw on border
    pos.X++ 
    pos.Y++
    
    for i:=0; i<size.Row; i++ {
        start:=(pos.Y-1)*a.size.Col + (pos.X-1)
        for j:=0; j<size.Col; j++ {
            a.buffer[start+j] = b[i*size.Col+j]
        }

    }



}
