package main

// DO NOT EDIT
// Generated by: ./tools/gen_sprite_tags.sh

type playerAnimationTags uint8

const (
	playerIdle playerAnimationTags = iota
	playerClimb
	playerLeanstart
	playerLeanloop
	playerLeanend
	playerPushwallstart
	playerPushwallloop
	playerPushwallend
	playerNogrip
	playerJumpstart
	playerJumploop
	playerJumpendfloor
	playerJumpendmantle
	playerJumpendwall
	playerSlipstart
	playerSliploop
	playerSlipend
	playerFallstart
	playerFallloop
	playerFallendfloor
	playerFallendwall
	playerStand
	playerWalkright
	playerWalkleft
	playerSwitchtotopview
	playerGrapplestart
	playerGrappleloop
	playerGrappleEnd
)

var playerAnimationNames = []string{
	"Idle",
	"Climb",
	"Lean start",
	"Lean loop",
	"Lean end",
	"Push wall start",
	"Push wall loop",
	"Push wall end",
	"No grip",
	"Jump start",
	"Jump loop",
	"Jump end floor",
	"Jump end mantle",
	"Jump end wall",
	"Slip start",
	"Slip loop",
	"Slip end",
	"Fall start",
	"Fall loop",
	"Fall end floor",
	"Fall end wall",
	"Stand",
	"Walk right",
	"Walk left",
	"Switch to topview",
	"Grapple start",
	"Grapple loop",
	"Grapple End",
}
