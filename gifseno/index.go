package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    //"math/rand"
    "os"
)

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 1  	// number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 100    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )
    //freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    phase := 0 // phase difference
    //x := 0
    //y := 50
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := math.Pi*-2; t < cycles*2*math.Pi; t += res {
            x := math.Cos(t)
            y := math.Sin(t)

			if t>=math.Pi {
				img.SetColorIndex(size + size/2+int(x*size/2) + phase, size+int(y*size/2+0.5), blackIndex)
			}else if t<=-math.Pi {
				img.SetColorIndex( size/2+int(x*size/2) - 2*size + phase, size+int(y*size/2+0.5), blackIndex)
			}else if t<=0 {
				img.SetColorIndex( size/2+int(x*size/2) - size + phase, size+int(y*size/2+0.5), blackIndex)
			}else{
				img.SetColorIndex(size/2+int(x*size/2) + phase, size+int(y*size/2+0.5), blackIndex)
			}
        }
		//for pixX:=-4; pixX<4; pixX++{
		//	for  pixY:=-4; pixY<4; pixY++{
		//	img.SetColorIndex(int(x)+pixX + phase, int(y)+pixY, blackIndex)
		//	}
		//}
        phase += 2
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
