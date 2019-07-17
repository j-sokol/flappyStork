# FlappyStork

Reimplementation of Flappy Bird in console. This was little project for me to pick up Go language. Library tcell was used for rendering graphics.


## How to run
You need to have Go compiler downloaded. Then run
```bash
go build
```
and you can run it using
```bash
./flappyStork
```

## Screenshots

### Graphics kudos

Big thanks to guys below who provided textures for this game.

Stork is from [this](https://opengameart.org/content/baby-bounce-nes) set, and blocks from [this](https://opengameart.org/content/micro-world-old-tileset) one.


### TODO
Here are some things that need to be improved:

- better hitbox
- show score
    - draw lower bar
- send score to some centralized location
- animate player
- menu at the end to restart
- save highscore to file?
- make gravity not linear