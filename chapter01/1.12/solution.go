// Exercise 1.12: Modify the Lissajous server to read parameter values from
// the URL. For example, you might arrange it so that a URL like
// http://loclahost:8000/?cycles=20 sets the number of cycles to 20 instead of
// the default 5. Use the strconv.Atoi function to convert the string
// parameter into an integer. You can see its documentation with go doc
// strconv.Atoi.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	blackIndex = 0
	greenIndex = 1
	default_cycles = 10
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		params := r.Form["cycles"]
		cycles := default_cycles
		if len(params) > 0 {
			s := params[0]
			cycles, _ = strconv.Atoi(s)
		}
		lissajous(w, float64(cycles))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res        = 0.001 // resolution of curve
		size       = 100   // 201 x 201 image
		nframes    = 128   // 64 total frames
		delay      = 1     // 80ms
		phase_diff = 0.001
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
				size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += phase_diff
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
