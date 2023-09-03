package schema

type GameContext struct {
	NextGameFunction GameFunction
	Nickname         string
}

type GameFunction func(g *GameContext) error
