package core

type Mode int

const (
	MANUAL Mode = iota
	AUTO
)

type Robot struct {
	Mode Mode	
}
