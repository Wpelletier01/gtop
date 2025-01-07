package cpu;

import (
    "gtop/module"
    "gtop/core"
    "fmt"
)

// core mode
const (
    C_NORMAL = iota
    // no progress bar
    C_SHRINK
)



func coreBuilder(c Core, pos core.Position, mode int) []module.Renderable {

    object := make([]module.Renderable,10,10)
    
    object[0] = module.NewText(
        fmt.Sprintf("C%d",c.name),
        core.GM_NULL,
        core.FG_WHITE

    )
    
    

    return object
}
