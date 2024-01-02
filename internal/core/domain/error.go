package domain

type Error struct {
	Err     error
	Message string
	Code    int
	Data    any
}
