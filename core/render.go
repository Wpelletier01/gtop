package core

import (
	//    "gtop/module"
	"fmt"
	"math"
    "strings"
)

func generateRenderMode(gm uint, fg uint, bg uint) string {
    return fmt.Sprintf("%c[%d;%d;%dm",ESC, gm, fg, bg)
} 

type Renderable interface {
    getRaw() string
    getPos() Position
}

type Text struct {
    Data                string
    GraphicMode         uint
    ForegroundColor     uint
    BackgroundColor     uint
    pos                 Position
}

func NewText(data string, gmode uint, fg uint, bg uint, row int, col int) Text {
    return Text{
        Data: data,
        GraphicMode: gmode,
        ForegroundColor: fg,
        BackgroundColor: bg,
        pos: Position{ Row: row, Col: col},
    }
}

func (t Text) getRaw() string {

    return fmt.Sprintf(
        "%s%s%c[%dm",
        generateRenderMode(t.GraphicMode, t.ForegroundColor, t.BackgroundColor),
        t.Data,
        ESC,
        GM_RESET,
    )

}

func (t Text) getPos() Position {
    return t.pos
}


type ProgressBar struct {
    Percent     float64
    FullColor   uint
    EmptyColor  uint
    Length      uint
    pos         Position
}

func NewProgressBar(
    percent float64, 
    fullColor uint, 
    emptyColor uint, 
    length uint,
    row int,
    col int) ProgressBar {
    return ProgressBar{
        Percent: percent,
        FullColor: fullColor,
        EmptyColor: emptyColor,
        Length: length,
        pos:    Position{ Col: col, Row: row, },
    } 
}

func (p ProgressBar) getRaw() string {
    
    
    full := int(math.Round(p.Percent * float64(p.Length)))
    
    var builder strings.Builder
    

    // Write full part
    builder.WriteString(generateRenderMode(GM_RESET,p.FullColor,BG_DEFAULT))

    for i:=0; i<full; i++ {
        builder.WriteRune(PBAR_FULL)
    }
    
    // Write empty part
    builder.WriteString(generateRenderMode(GM_RESET,p.EmptyColor,BG_DEFAULT))
    for i:=0; i<int(p.Length)-full; i++ {
        builder.WriteRune(PBAR_EMPTY)
        
    }
    //return "Bruh"
    return builder.String()

}

func (p ProgressBar) getPos() Position {
    return p.pos
}



func RenderScreenFrame(screenSize Size) {
     
    // top left
    MoveCursor(0,2)
    
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
    MoveCursor(screenSize.Row,2) // let place for menu bar
    fmt.Printf(
        "%c%s%c",
        BORDER_BOTTOM_LEFT,
        string(horizontal),
        BORDER_BOTTOM_RIGHT,
    )
    
    
    // TODO: maybe a better way to render the middle border
    halfScreen := int(float64(screenSize.Col) / 2.0)+1
    
    
    MoveCursor(1,halfScreen)
    fmt.Printf("%c",BORDER_MIDDLE_TOP)
        
    // right and left side
    for i:=2; i<screenSize.Row; i++ {
        MoveCursor(i,2)
        fmt.Printf("%c",BORDER_VERTICAL)
        MoveCursor(i,halfScreen)
        fmt.Printf("%c",BORDER_VERTICAL)
        MoveCursor(i,screenSize.Col)
        fmt.Printf("%c",BORDER_VERTICAL)
    }
    
    MoveCursor(screenSize.Row,halfScreen)
    fmt.Printf("%c",BORDER_MIDDLE_BOTTOM)

}


func RenderObject(obj Renderable) {
    
    pos := obj.getPos()
    MoveCursor(pos.Row,pos.Col)
    fmt.Print(obj.getRaw())

}
