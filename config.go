package main

import "time"

// Globals
var jumpImpulse = false // Init val
var jumpTime = 0        // Init val
var tickTime = time.Duration(20)

// Player
var playerImgPath = "assets/player/3.png"
var constGravity = 0.7
var constJump = 1.3
var impulseTime = 15
var playerOffset = 10

// Barier block
var blockImgPath = "assets/player/5.png"
var blockWidth = 14
var verticalDistance = 7

// Scoreboard
var scoreboardImgPath = "assets/player/8.png"
var scoreboardHeight = 2
