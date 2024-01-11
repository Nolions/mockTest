package game

type Game interface {
	GetType() string
	Start()
	End()
}
