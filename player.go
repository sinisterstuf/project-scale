package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
)

//go:generate ./tools/gen_sprite_tags.sh assets/sprites/Nanobot.json player_anim.go player

const MinJumpTime = 13
const MaxJumpTime = 21

const (
	ActionMoveUp input.Action = iota
	ActionMoveLeft
	ActionMoveDown
	ActionMoveRight
	ActionJump
)

type PlayerAxis int8

const (
	AxisVertical PlayerAxis = iota
	AxisHorizontal
	AxisBoth
)

// Player is the player character in the game
type Player struct {
	*resolv.Object
	Input    *input.Handler
	State    playerAnimationTags
	Sprite   *SpriteSheet
	Frame    int
	Tick     int
	Jumping  bool
	Falling  bool
	Axis     PlayerAxis
	JumpTime int
	WhatTile string
}

func NewPlayer(position []int) *Player {
	object := resolv.NewObject(
		float64(position[0]), float64(position[1]),
		16, 16,
	)
	object.SetShape(resolv.NewRectangle(
		0, 0, // origin
		8, 8,
	))

	return &Player{
		Object: object,
		Sprite: loadSprite("Nanobot"),
	}
}

func (p *Player) Update() {
	p.Tick++
	if p.State == playerJumpingmidair {
		p.JumpTime++
	}
	p.updateMovement()
	p.animate()
	p.Object.Update()
}

func (p *Player) updateMovement() {
	speed := 0.6

	if p.Input.ActionIsJustPressed(ActionJump) {
		p.Jumping = true
		p.State = playerJumpingstart
	}

	if p.Jumping {
		speed = 2.0
		if p.State == playerJumpingmidair {
			p.move(+0, -speed)
		}
	} else if p.Falling {
		switch p.State {
		case playerFallingfreefall:
			speed = -3.0
		case playerFallingstart, playerFallingrecovery:
			speed = -0.5
		}
		p.move(+0, -speed)
	} else {
		p.State = playerIdle
		if p.Input.ActionIsPressed(ActionMoveUp) {
			p.move(+0, -speed)
			p.State = playerClimpingupdown
			p.Axis = AxisVertical
		}
		if p.Input.ActionIsPressed(ActionMoveDown) {
			p.move(+0, +speed)
			p.State = playerClimpingupdown
			p.Axis = AxisVertical
		}
		if p.Input.ActionIsPressed(ActionMoveLeft) {
			p.move(-speed, +0)
			p.State = playerClimbingleftright
			if p.Axis == AxisVertical {
				p.Axis = AxisBoth
			} else {
				p.Axis = AxisHorizontal
			}
		}
		if p.Input.ActionIsPressed(ActionMoveRight) {
			p.move(+speed, +0) // TODO: cancel movement when pressing opposite directions
			p.State = playerClimbingleftright
			if p.Axis == AxisVertical {
				p.Axis = AxisBoth
			} else {
				p.Axis = AxisHorizontal
			}
		}
		if p.Axis == AxisBoth {
			p.State = playerClimbingdiagonally
		}
	}

}

func (p *Player) move(dx, dy float64) {
	if collision := p.Check(dx, 0, TagWall); collision != nil {
		for _, o := range collision.Objects {
			if p.Shape.Intersection(dx, 0, o.Shape) != nil {
				dx = 0
			}
		}
	}
	p.X += dx

	if collision := p.Check(0, dy, TagWall); collision != nil {
		for _, o := range collision.Objects {
			if p.Shape.Intersection(0, dy, o.Shape) != nil {
				dy = 0
				// recover from fall
				if p.Falling {
					p.Falling = false
					p.State = playerFallingrecovery
				}
			}
		}
	}
	p.Y += dy

	if collision := p.Check(0, 0); collision != nil {
		for _, o := range collision.Objects {
			if p.Shape.Intersection(0, 0, o.Shape) != nil {
				p.WhatTile = o.Tags()[0]
			}
		}
	}
}

func (p *Player) animate() {
	if p.Frame == p.Sprite.Meta.FrameTags[p.State].To {
		p.animationBasedStateChanges()
	}
	p.Frame = Animate(p.Frame, p.Tick, p.Sprite.Meta.FrameTags[p.State])
}

// Animation-trigged state changes
func (p *Player) animationBasedStateChanges() {
	switch p.State {

	case playerJumpingstart, playerJumpingmidair:
		if (p.Input.ActionIsPressed(ActionJump) || p.JumpTime < MinJumpTime) && p.JumpTime < MaxJumpTime {
			p.State = playerJumpingmidair
		} else {
			p.State = playerJumpingend
			// Start falling if the jumped ended in a chasm
			if collision := p.Check(0, 0, TagChasm); collision != nil {
				for _, o := range collision.Objects {
					if p.Shape.Intersection(0, 0, o.Shape) != nil {
						p.Jumping = false
						p.JumpTime = 0
						p.State = playerFallingstart
						p.Falling = true
					}
				}
			}
		}

	case playerJumpingend:
		p.State = playerIdle
		p.Jumping = false
		p.JumpTime = 0

	case playerFallingstart:
		p.State = playerFallingfreefall

	case playerFallingrecovery:
		p.State = playerIdle

	}
}

func (p *Player) Draw(camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}

	s := p.Sprite
	frame := s.Sprite[p.Frame]
	img := s.Image.SubImage(image.Rect(
		frame.Position.X,
		frame.Position.Y,
		frame.Position.X+frame.Position.W,
		frame.Position.Y+frame.Position.H,
	)).(*ebiten.Image)

	// Centre sprite
	op.GeoM.Translate(
		float64(-frame.Position.W/4),
		float64(-frame.Position.H/4),
	)

	camera.Surface.DrawImage(img, camera.GetTranslation(op, p.X, p.Y))
}
