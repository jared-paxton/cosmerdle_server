package game

type GameServicer interface{
    
}

type GameAccessor interface{}

type gameService struct {
	accessor GameAccessor
}

func NewGameServicer(gameAccessor GameAccessor) *gameService {
	return &gameService{
		accessor: gameAccessor,
	}
}
