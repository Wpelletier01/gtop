package main;

import (
    "os"
    "fmt"
    "gtop/core"
)

func main() {
    
    // init
    var app App

    //enter raw mode
    revertFunc, err := core.RawMode()
    
    if err != nil {
        fmt.Fprintf(os.Stderr,"Termios error: %s",err)
        os.Exit(1)
    }
    
    defer revertFunc()
    
    // enter new buffer
    core.EnableAltBuffer()
    core.HideCursor()
    
    
    // ... logic
    app.Run()
    
    // exit safely
    core.ShowCursor()
    core.DisableAltBuffer()

}
