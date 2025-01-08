package core;

import (
    "os"
    "os/signal"
    "syscall"
    "fmt"
    "time"
)


type EventRegister struct {
    
    Resize      chan Size
    Input       chan []byte
    NeedRedraw  chan bool    

}

func (e *EventRegister) StartListener() {
    
    // init channel
    e.Resize = make(chan Size)
    e.Input = make(chan []byte)

    go listenForResize(e.Resize)
    go listenForInput(e.Input)
    
}

func listenForResize(ec chan Size ) {
    
    var size Size

    // Send one for initialisatio
    // TODO: look for errors
    ts := GetTerminalSize()
        
    size.Col = int(ts.Col)
    size.Row = int(ts.Row)
    
    ec<-size

    for {
        c := make(chan os.Signal, 1)
        signal.Notify(c, syscall.SIGWINCH)
        
        <-c 

        // TODO: look for errors
        ts := GetTerminalSize()
        
        size.Col = int(ts.Col)
        size.Row = int(ts.Row)

        ec <- size

    }


}

func listenForInput(ec chan []byte) {


    for {
        
        buf := make([]byte, 3)
        _, err := os.Stdin.Read(buf)

        if err != nil {
            // TODO: better handling of course
			fmt.Println("Error reading input:", err)
		    continue
        }
        
        ec<-buf
        

    }

}

func NeedToRedraw(ec chan bool) {
    
    time.Sleep(DRAW_MAX_TIME * time.Second)
    ec <-true

}
 
