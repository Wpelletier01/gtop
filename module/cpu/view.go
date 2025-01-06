package cpu;

import (
    "gtop/module"
    "gtop/core"
)

// core mode
const (
    C_NORMAL = iota
    // no progress bar
    C_SHRINK
)

type CpuView struct { 

} 


type CoreView struct { 
    size    core.Size
    pos     core.Position
}

func (cv *CoreView) build(mode int, core Core) []module.Entity  {
    
    entities := make([]module.Entity, 10)
     

    return entities
} 
