package module;

import "gtop/core"

type Model interface {
    
    Collect()
    // Refresh data
    Update()
    // The terminal has been resized
    Resize(s core.Size)
    // return the view of the module 
    getEntity() []Entity
}



type Entity struct {
    Data            string
    GraphicMode     uint
    ForegroundColor uint
    BackgroundColor uint
}


