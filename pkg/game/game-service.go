package game

type GameServicer interface{}

type GameAccessor interface{}

type gameService struct {
	repo GameAccessor
}

func NewGameServicer(gameRepo GameAccessor) *gameService {
	return &gameService{
		repo: gameRepo,
	}
}
