package cpu;

import (
    "gtop/core"
    //"gtop/module"
    
)


type Core struct {
    usage   float64
    temp    int
}

type Cpu struct {
    name        string
    clockspeed  string
    cores       []Core
    size        core.Size
    pos         core.Position
} 


func (c *Cpu) Init(pos core.Position,  s core.Size) {
    
    c.name = "Ryzen 5 3500u"
    c.clockspeed = "2.1Ghz"
    
    c.cores = make([]Core,1,1)
    
    c.cores[0] = Core{ usage: 20.0, temp: 23 }
    
    c.pos = pos // never change
    c.size = s
}

func (c Cpu) DrawStatic() {
    
    // 1 draw border
    core.RenderFrame(c.pos,c.size)
    
    for _, obj := range generateStaticView(c.pos) {
        core.RenderObject(obj)
    }

}


func (c *Cpu) Collect() {}

func (c *Cpu) Update() {}

func (c *Cpu) Resize(s core.Size) {
    

}

