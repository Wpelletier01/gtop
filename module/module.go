package module;

import "gtop/core"

type Model interface {
    
    Init()
    Collect()
    // Refresh data
    Update()
    // The terminal has been resized
    Resize(s core.Size)
}




