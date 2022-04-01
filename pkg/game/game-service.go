package game

type GameServicer interface {
}

type GameAccessor interface{}

type gameService struct {
	accessor GameAccessor
}

func NewGameService(gameAccessor GameAccessor) GameServicer {
	return &gameService{
		accessor: gameAccessor,
	}
}
