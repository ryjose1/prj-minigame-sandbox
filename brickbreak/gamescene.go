package brickbreak

import (
	"github.com/ryjose1/minigames/log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	ball   *Ball
	logger *log.BuiltinLogger
}

func NewGameScene(logger *log.BuiltinLogger) *GameScene {
	return &GameScene{
		ball:   NewBall(),
		logger: logger,
	}

}

func (s *GameScene) Update() error {
	s.ball.Update()
	return nil
}

func (s *GameScene) Draw(r *ebiten.Image) {
	s.ball.Draw(r)
	s.logger.Infof("%d %d", s.ball.x, s.ball.y)
}