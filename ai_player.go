package main

import "math"

type AiBehavior interface {
	moveWith(ball *Ball)
}
type AiPlayer struct {
	object Object
}

func (ai *AiPlayer) Update(ball *Ball) {
	ai.moveWith(ball)
}

// Implement AiBehavior
func (ai *AiPlayer) moveWith(ball *Ball) {
	if (ai.object.y < ball.y && math.Abs(ball.y - ai.object.y) > 10) {
		ai.object.y = math.Min(ai.object.y +6, 240 - ai.object.height )
	} else if (ai.object.y > ball.y && math.Abs(ball.y - ai.object.y) > 10) {
		ai.object.y = math.Min(ai.object.y -6, 240 - ai.object.height )
	}
}
