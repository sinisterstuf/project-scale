package main

// DO NOT EDIT
// Generated by: ./tools/gen_sprite_tags.sh

type playerAnimationTags uint8

const (
	playerIdle playerAnimationTags = iota
	playerClimpingupdown
	playerClimbingleftright
	playerClimbingdiagonally
	playerJumpingstart
	playerJumpingmidair
	playerJumpingend
	playerSlippingstart
	playerSlippingscrambling
	playerSlippingrecovery
	playerFallingstart
	playerFallingfreefall
	playerFallingrecovery
)