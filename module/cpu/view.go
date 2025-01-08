package cpu;

import (
    //"gtop/module"
    "gtop/core"
    //"fmt"
)






func generateStaticView(pos core.Position) []core.Renderable {

    var obj []core.Renderable
    
    
    moduleTitle := core.NewText(
        "CPU",
        core.GM_BOLD,
        core.FG_DEFAULT,
        core.BG_DEFAULT,
        2,
        2,
    )
    
    obj = append(obj,moduleTitle)

    return obj

}


// 
