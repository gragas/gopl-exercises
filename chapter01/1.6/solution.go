// Exercise 1.6: Modify the Lissajous program to produce images in multiple
// colors by adding more values to palette and then displaying them by
// changing the third argument of SetColorIndex in some interesting way.

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

var palette = []color.Color{color.Black}

const (
	ChunkSize   = 50
	NumChunks   = 3
	PaletteSize = NumChunks * ChunkSize
)

func main() {
	fillPalette()
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10    // number of complete x cycles
		res     = 0.01  // resolution of curve
		size    = 100   // 201 x 201 image
		nframes = 128   // 64 total frames
		delay   = 1     // 80ms
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex uint8 = 0
		for t := 0.0; t < 2*math.Pi*cycles; t += res {
			x := math.Cos(t)
			y := math.Sin(t*freq + phase)
			colorIndex = (colorIndex + 1) % PaletteSize
			img.SetColorIndex(
				size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.05
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func fillPalette() {
	var state1, state2, state3 int
	for chunk := 0; chunk < NumChunks; chunk++ {
		if chunk == 0 {
			state1 = 1
			state2 = 0
			state3 = 0
		} else if chunk == 1 {
			state1 = 0
			state2 = 1
			state3 = 0
		} else {
			state1 = 0
			state2 = 0
			state3 = 1
		}
		for i := 0; i <= ChunkSize; i++ {
			addColor := color.RGBA{
				uint8(((state3*i + state1*(ChunkSize-i)) * 0xFF) / ChunkSize),
				uint8(((state1*i + state2*(ChunkSize-i)) * 0xFF) / ChunkSize),
				uint8(((state2*i + state3*(ChunkSize-i)) * 0xFF) / ChunkSize),
				0xFF}
			palette = append(palette, addColor)
		}
	}
}
