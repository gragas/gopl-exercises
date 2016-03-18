// Exercise 1.5: Change the Lissajous program's color palette to green on
// black, for added authenticity. To create the web color #RRGGBB, use
// color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair fo hexadecimal
// digits represents the intensity of the red, green, or blue component of
// the pixel.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles   = 10      // number of complete x cycles
		res      = 0.001  // resolution of curve
		size     = 100    // 201 x 201 image
		nframes  = 128     // 64 total frames
		delay    = 1      // 80ms
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < 2*math.Pi*cycles; t += res {
			x := math.Cos(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size + int(x * size + 0.5), size + int(y * size + 0.5),
				greenIndex)
		}
		phase += 0.001
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
