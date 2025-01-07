package cpu;

import (
    "gtop/core"
    //"gtop/module"
    
)


type Core struct {
    name    int
    percent float64
    temp    int
}

type Cpu struct {
    name        string
    clockspeed  string
    cores       []Core
    
} 


func (c *Cpu) Init() {}

func (c *Cpu) Collect() {}

func (c *Cpu) Update() {}

func (c *Cpu) Resize(s core.Size) {
    

}

