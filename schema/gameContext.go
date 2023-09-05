package schema

type GameContext struct {
	NextGameFunction GameFunction
	SaveId           int64
	Nickname         string
	Currency         float64
}

type GameFunction func(g *GameContext) error

func (g *GameContext) CallNextGameFunction() error {
	err := g.NextGameFunction(g)
	if err != nil {
		return err
	}
	return nil
}
