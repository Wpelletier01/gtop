package main

import (
	"gtop/core"
	"gtop/module"
	"math"
	//"gtop/module"
	"gtop/module/cpu"
)

type App struct {
    size        core.Size
    event       core.EventRegister
    modules     []module.Model
}  


func (a *App) initModules() {
    
    ts := core.GetTerminalSize()
    
    a.size.Col = int(ts.Col)
    a.size.Row = int(ts.Row)

    a.modules = make([]module.Model, 1, 1)
    

    var cpuModel cpu.Cpu
    

    
    
    quarterHeight := int(math.Floor(float64(a.size.Row) / 4.0))
    halfWidth := int(math.Round(float64(a.size.Col)/2.0))

    cpuModel.Init(core.Position{Col: 0, Row: 0}, core.Size{Col: halfWidth, Row: quarterHeight})
    

    a.modules[0] = cpuModel
    

}


func (a *App) Run() {
    

    // start event listener
    a.event.StartListener()
    
    a.initModules() 
    
    a.drawStatic() // init draw

    main:
    for {
        
        a.event.NeedRedraw = make(chan bool)
        go core.NeedToRedraw(a.event.NeedRedraw)

        select {
            
            case nSize := <-a.event.Resize:

                a.resize(nSize)
                a.drawStatic()
            
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
            case <-a.event.NeedRedraw:
                // We need to draw        
        }
         
        a.drawDynamic()
    }



}


func (a *App) resize(s core.Size) { 
    
    // hardcoded padding
    s.Col -= 1
    s.Row -= 1
    
    a.size = s
    
}

func (a *App) drawStatic() {

    core.ClearScreen()
    core.RenderFrame(core.Position{Col: 0, Row: 0,},a.size)

    for _,module := range a.modules {
        module.DrawStatic()
    }

}

func (a *App) drawDynamic() {

}

func (a *App) draw() {
    
    core.ClearScreen()
    

    
    core.RenderScreenFrame(a.size)
    
    content := make([]core.Renderable,2,2)

    text := core.NewText(
        "This is some text",
        core.GM_RESET,
        core.FG_WHITE,
        core.BG_RED,
        15,
        40,
    )
    
    pbar := core.NewProgressBar(
        0.9,
        core.FG_GREEN,
        core.BG_DEFAULT,
        6,
        30,
        26,
    )


    content[0] = text
    content[1] = pbar
    
   
    for _, obj := range content {
        core.RenderObject(obj)
    }

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

