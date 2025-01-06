package cpu;

import (
    "gtop/core"
    //"gtop/module"
    
)


type Core struct {
    name    int
}

type Cpu struct {
    size    core.Size
    //view    module.View
} 



func (c *Cpu) Collect() {}

func (c *Cpu) Update() {}

func (c *Cpu) Resize(s core.Size) {
    

}

//func (c *Cpu) GetView() module.View {
//    return c.view
//}
