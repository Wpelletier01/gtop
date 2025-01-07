package module;

import "gtop/core"

type Model interface {
    
    Init()
    Collect()
    // Refresh data
    Update()
    // The terminal has been resized
    Resize(s core.Size)
}



// TODO: they should have an renderable interface
type Renderable interface {}



type Text struct {
    Data                string
    GraphicMode         uint
    ForegroundColor     uint
    BackgroundColor     uint
    relativePosition    core.Position
}

func NewText(data string, gmode uint, fclr uint, bclr uint, row int, col int) Text {
    return Text{
        Data: data,
        GraphicMode: gmode,
        ForegroundColor: fclr,
        BackgroundColor: bclr,
        relativePosition: core.Position{ Row: row, Col: col},
    }
}


type ProgressBar struct {
    Percent     float32
    FullColor   uint
    EmptyColor  uint
    Length      uint
}

