package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
//"fmt"
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
        //cycles  = 1.5  	// number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 100    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )
    //freq := rand.Float64() * 3.0 // relative frequency of y oscillator
    anim := gif.GIF{LoopCount: nframes}
    //phase := 0 // phase difference
    //x := 0
    //y := 50
    sumA := math.Pi/90
    //empieza := 0 
    angulo := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
	
        //for t := empieza; t < cycles*math.Pi; t += res {
        //    x := math.Cos(t)
        //    y := math.Sin(t)
	//    img.SetColorIndex(size +int(x*size/2), size+int(y*size/2+0.5), blackIndex)
	//    img.SetColorIndex(i+int(t)*100,i+int(t)*100, blackIndex)
	//fmt.Println("\n")
	//fmt.Println(x)
	//f/mt.Println("\n")

			//if t>=math.Pi {
			//	img.SetColorIndex(size + size/2+int(x*size/2) + phase, size+int(y*size/2+0.5), blackIndex)
			//}else if t<=-math.Pi {
			//	img.SetColorIndex( size/2+int(x*size/2) - 2*size + phase, size+int(y*size/2+0.5), blackIndex)
			//}else if t<=0 {
			//	img.SetColorIndex( size/2+int(x*size/2) - size + phase, size+int(y*size/2+0.5), blackIndex)
			//}else{
			//	img.SetColorIndex(size/2+int(x*size/2) + phase, size+int(y*size/2+0.5), blackIndex)
			//}
	for a := angulo; a < angulo+math.Pi*1.5; a += res {
	    y := math.Sin(a)
	    x := math.Cos(a)
	    if a< math.Pi/2{
		img.SetColorIndex(size +int((x-math.Pi)*size/2), size+int(y*size/2+0.5), blackIndex)
	    }else if a < math.Pi{
		img.SetColorIndex(size +int(x*size/2), size+int(y*size/2+0.5), blackIndex)
	    }
 	    else {
		img.SetColorIndex(size +int(x*size/2), size+int(y*size/2+0.5), blackIndex)
	    }
	    //img.SetColorIndex(size +int(x*size/2), size+int(y*size/2+0.5), blackIndex)
        }

	angulo = angulo + sumA
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
