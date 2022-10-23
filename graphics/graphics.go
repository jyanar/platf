package graphics

import (
	"image"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var Atlas *ebiten.Image
var Quads []*ebiten.Image

// In the atlas, numbers are left -> right
var Tile *ebiten.Image        // 0
var ToggleFloor *ebiten.Image // 1
var Empty *ebiten.Image       // 2
var Symbol *ebiten.Image      // 3
var Spikes *ebiten.Image      // 4
var LeverOff *ebiten.Image    // 5
var LeverOn *ebiten.Image     // 6
var Player *ebiten.Image      // 8
var Enemy *ebiten.Image       // 20

func Load() error {
	var err error

	Quads, err = parseAtlas("atlas.png", 16, 16)
	if err != nil {
		return err
	}

	Tile = Quads[0]
	ToggleFloor = Quads[1]
	Empty = Quads[2]
	Symbol = Quads[3]
	Spikes = Quads[4]
	LeverOff = Quads[5]
	LeverOn = Quads[6]
	Player = Quads[8]
	Enemy = Quads[21]

	return nil
}

type Animation struct {
	t        float64
	start    int
	len      int
	duration float64
}

func NewAnimation(start, len int, duration float64) *Animation {
	return &Animation{
		t:        0,
		start:    start,
		len:      len,
		duration: duration,
	}
}

func (a *Animation) Init(start int, len int, duration float64) {
	a.t = 0
	a.start = start
	a.len = len
	a.duration = duration
}

func (a *Animation) Update() {
	// Assuming dt is 1/60 seconds
	a.t += 1.0 / 60.0
}

func (a Animation) GetFrame() int {
	return a.start + int(math.Floor(float64(a.len)*math.Mod(a.t, a.duration)/a.duration))
}

// func (a Animation) Draw(screen *ebiten.Image) {
// 	screen.DrawImage(Quads[a.GetFrame()])
// }

// https://stackoverflow.com/questions/49594259/reading-image-in-go
func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func getEbitenImage(filepath string) (*ebiten.Image, error) {
	img, err := getImageFromFilePath("atlas.png")
	if err != nil {
		return &ebiten.Image{}, err
	}
	return ebiten.NewImageFromImage(img), nil
}

func parseAtlas(filepath string, quadWidth, quadHeight int) ([]*ebiten.Image, error) {
	Quads = []*ebiten.Image{}
	var err error
	Atlas, err = getEbitenImage(filepath)
	if err != nil {
		return nil, err
		// log.Println(err)
	}

	atlasWidth, atlasHeight := Atlas.Size()
	ncol, nrow := atlasWidth/quadWidth, atlasHeight/quadHeight
	for irow := 0; irow < nrow; irow++ {
		for icol := 0; icol < ncol; icol++ {
			x0 := icol * quadWidth
			y0 := irow * quadHeight
			x1 := x0 + quadWidth
			y1 := y0 + quadHeight
			Quads = append(Quads, Atlas.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image))
		}
	}
	return Quads, nil
}

// func getEbitenSubImageAt(atlas *ebiten.Image, x, y, width, height int) (*ebiten.Image, error) {
// 	return atlas.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image), nil
// }
