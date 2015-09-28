package main

import (
    "fmt"
    tl "github.com/JoelOtter/termloop"
    "math"
    "math/rand"
    "strconv"
    "time"
)

const (
    TileWidth = 10
    TileHeight = 7
    BorderWidth = 1
    BorderHeight = 1
    Time = 10.0
    X = 10
    Y = 15
)

var (
    Response = 0
    Status = 1
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
        if Status == 1 {
            r.px, r.py = r.r.Position()
            if ty := r.py + TileHeight + BorderHeight; ty > Y {
                Response = 2
                switch ev.Key {
                case tl.KeyArrowLeft:
                    if TilePos[0] == 0 {
                        Response = 1
                    }
                    break
                case tl.KeyArrowDown:
                    if TilePos[0] == 1 {
                        Response = 1
                    }
                    break
                case tl.KeyArrowUp:
                    if TilePos[0] == 2 {
                        Response = 1
                    }
                    break
                case tl.KeyArrowRight:
                    if TilePos[0] == 3 {
                        Response = 1
                    }
                    break
                }
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
}

type RemainingTime struct {
    r *tl.Text
    s *tl.Text
    t float64
    m *tl.Text
    e *tl.Text
}

func (r *RemainingTime) Draw(s *tl.Screen) {
    if Status == 1 {
        r.t = math.Max(r.t - s.TimeDelta(), 0)
        if r.t == 0 {
            Status = 0
            r.e.SetText("Time up!")
        } else {
            if Response == 1 {
                s, _ := strconv.Atoi(r.s.Text())
                s = s + 1
                if s % 5 == 0 {
                    r.t = r.t + 1
                }
                r.s.SetText(strconv.Itoa(s))
                switch s {
                    case 10:
                        r.m.SetText("You've got it!")
                    case 20:
                        r.m.SetText("Keep going!")
                    case 30:
                        r.m.SetText("You're doing great!")
                    case 40:
                        r.m.SetText("You rock!")
                    case 50:
                        r.m.SetText("Don't stop!")
                    case 60:
                        r.m.SetText("I like your style!")
                    case 70:
                        r.m.SetText("Awesome!")
                    case 80:
                        r.m.SetText("How do you do that?")
                    case 90:
                        r.m.SetText("Don't ever stop!!")
                    case 100:
                        r.m.SetText("I'm really impressed.")
                    case 150:
                        r.m.SetText("You're really still here?")
                    case 200:
                        r.m.SetText("That's incredible!")
                }
            } else if Response == 2 {
                Status = 0
                r.e.SetText("Game Over")
            }
        }
        Response = 0
        r.r.SetText(fmt.Sprintf("%.3f", r.t))
    }
    r.r.Draw(s)
    r.s.Draw(s)
    r.m.Draw(s)
    r.e.Draw(s)
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
        r:    tl.NewText(X + 4 * (TileWidth + BorderWidth), 0, fmt.Sprintf("%.3f", Time), tl.ColorRed, tl.ColorDefault),
        s:    tl.NewText(0, 0, "0", tl.ColorRed, tl.ColorDefault),
        t:    Time,
        m:    tl.NewText(0, Y + TileHeight + 1, "", tl.ColorRed, tl.ColorDefault),
        e:    tl.NewText(X + 4 * (TileWidth + BorderWidth), Y + TileHeight + 1, "", tl.ColorRed, tl.ColorDefault),
    })
    game.Screen().SetLevel(level)
    game.Start()
}
