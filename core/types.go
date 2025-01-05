package core


type Size struct {
    Col int
    Row int
}

func (s Size) Area() int {
    return s.Col * s.Row
}


type Position struct {
    X int
    Y int
}
