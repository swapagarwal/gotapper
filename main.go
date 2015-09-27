package main

import (
    "fmt"
    tl "github.com/JoelOtter/termloop"
    "math"
    "math/rand"
    "time"
)

var (
    TileWidth = 10
    TileHeight = 7
    BorderWidth = 1
    BorderHeight = 1
    Time = 10.0
    X = 10
    Y = 15
    TilePos [4]int
)

type Tile struct {
    r *tl.Rectangle
    px int
    py int
}

func (r *Tile) Draw(s *tl.Screen)    { r.r.Draw(s) }
func (r *Tile) Size() (int, int)     { return r.r.Size() }
func (r *Tile) Position() (int, int) { return r.r.Position() }

func (r *Tile) Tick(ev tl.Event) {
    if ev.Type == tl.EventKey {
        r.px, r.py = r.r.Position()
        switch ev.Key {
        case tl.KeyArrowRight:
            //r.r.SetPosition(r.px + 1, r.py)
            break
        case tl.KeyArrowLeft:
            //r.r.SetPosition(r.px - 1, r.py)
            break
        case tl.KeyArrowUp:
            //r.r.SetPosition(r.px, r.py - 1)
            break
        case tl.KeyArrowDown:
            //r.r.SetPosition(r.px, r.py + 1)
            break
        }
        if ty := r.py + TileHeight + BorderHeight; ty > Y {
            TilePos[0] = TilePos[1]
            TilePos[1] = TilePos[2]
            TilePos[2] = TilePos[3]
            TilePos[3] = rand.Intn(4)
            r.r.SetPosition(X + TilePos[3] * (TileWidth + BorderWidth), Y - 3 * (TileHeight + BorderHeight))
        } else {
            r.r.SetPosition(r.px, ty)
        }
    }
}

type RemainingTime struct {
    r *tl.Text
    t float64
}

func (r *RemainingTime) Draw(s *tl.Screen) {
    r.t = math.Max(r.t - s.TimeDelta(), 0)
    r.r.SetText(fmt.Sprintf("%.3f", r.t))
    r.r.Draw(s)
}

func (r *RemainingTime) Tick(ev tl.Event) {}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    game := tl.NewGame()
    level := tl.NewBaseLevel(tl.Cell{
        Bg: tl.ColorWhite,
    })
    for i := 0; i < 4; i++ {
        TilePos[i] = rand.Intn(4)
        level.AddEntity(&Tile{
            r:    tl.NewRectangle(X + TilePos[i] * (TileWidth + BorderWidth), Y - i * (TileHeight + BorderHeight), TileWidth, TileHeight, tl.ColorBlack),
        })
    }
    level.AddEntity(tl.NewText(X + TileWidth / 2 - 1, Y + TileHeight, "←", tl.ColorBlack, tl.ColorWhite))
    level.AddEntity(tl.NewText(X + (TileWidth + BorderWidth) + TileWidth / 2 - 1, Y + TileHeight, "↓", tl.ColorBlack, tl.ColorWhite))
    level.AddEntity(tl.NewText(X + 2 * (TileWidth + BorderWidth) + TileWidth / 2 - 1, Y + TileHeight, "↑", tl.ColorBlack, tl.ColorWhite))
    level.AddEntity(tl.NewText(X + 3 * (TileWidth + BorderWidth) + TileWidth / 2 - 1, Y + TileHeight, "→", tl.ColorBlack, tl.ColorWhite))
    level.AddEntity(&RemainingTime{
        r:    tl.NewText(0, 0, fmt.Sprintf("%.3f", Time), tl.ColorRed, tl.ColorDefault),
        t:    Time,
    })
    game.Screen().SetLevel(level)
    game.Start()
}
