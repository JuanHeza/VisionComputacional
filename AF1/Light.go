package main

import (
	_"fmt"
	"os"
	"log"
	"image"
	"math"
	"image/color"
	"image/jpeg"
)

type Color struct {
    R, G, B float64
}

func SatP(In string,Out string) {
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds()
	imgSet := image.NewRGBA(b)
	
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			Pixel := MakeColor(oldPixel)
			h,s,l := Pixel.Hsl()
			if (s+0.15>0){
				Pixel = Hsl(h,s+.15,l)
			}else{
				Pixel = Hsl(h,s,l)
			}
			R,G,B,_ := Pixel.RGBA()
			r := uint8(R/0x101)
			g := uint8(G/0x101)
			b := uint8(B/0x101)
			pixel := color.RGBA{r,g,b,(100)}
			imgSet.Set(X,Y,pixel)
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SatM(In string,Out string) {
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds()
	imgSet := image.NewRGBA(b)
	
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			Pixel := MakeColor(oldPixel)
			h,s,l := Pixel.Hsl()
			if (s-0.25>0){
				Pixel = Hsl(h,s-.25,l)
			}else{
				Pixel = Hsl(h,s,l)
			}
			R,G,B,_ := Pixel.RGBA()
			r := uint8(R/0x101)
			g := uint8(G/0x101)
			b := uint8(B/0x101)
			pixel := color.RGBA{r,g,b,(100)}
			imgSet.Set(X,Y,pixel)
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func Dark(In string,Out string) {
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds()
	imgSet := image.NewRGBA(b)
	
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			Pixel := MakeColor(oldPixel)
			h,s,l := Pixel.Hsl()
			if (l-0.15>0){
				Pixel = Hsl(h,s,l-.15)
			}else{
				Pixel = Hsl(h,s,l)
			}
			R,G,B,_ := Pixel.RGBA()
			r := uint8(R/0x101)
			g := uint8(G/0x101)
			b := uint8(B/0x101)
			pixel := color.RGBA{r,g,b,(100)}
			imgSet.Set(X,Y,pixel)
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func Light(In string,Out string) {
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	b := img.Bounds()
	imgSet := image.NewRGBA(b)
	
	for X := 0; X < b.Max.X; X++ {
		for Y := 0; Y < b.Max.Y; Y++ {
			oldPixel := img.At(X, Y)
			Pixel := MakeColor(oldPixel)
			h,s,l := Pixel.Hsl()
			if (l+0.15<1.0){
				Pixel = Hsl(h,s,l+.15)
			}else{
				Pixel = Hsl(h,s,l)
			}
			R,G,B,_ := Pixel.RGBA()
			r := uint8(R/0x101)
			g := uint8(G/0x101)
			b := uint8(B/0x101)
			pixel := color.RGBA{r,g,b,(100)}
			imgSet.Set(X,Y,pixel)
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}


func (col Color) Hsl() (h, s, l float64) {
    min := math.Min(math.Min(col.R, col.G), col.B)
    max := math.Max(math.Max(col.R, col.G), col.B)

    l = (max + min) / 2

    if min == max {
        s = 0
        h = 0
    } else {
        if l < 0.5 {
            s = (max - min) / (max + min)
        } else {
            s = (max - min) / (2.0 - max - min)
        }

        if max == col.R {
            h = (col.G - col.B) / (max - min)
        } else if max == col.G {
            h = 2.0 + (col.B-col.R)/(max-min)
        } else {
            h = 4.0 + (col.R-col.G)/(max-min)
        }

        h *= 60

        if h < 0 {
            h += 360
        }
    }

    return
}

// Hsl creates a new Color given a Hue in [0..360], a Saturation [0..1], and a Luminance (lightness) in [0..1]
func Hsl(h, s, l float64) Color {
    if s == 0 {
        return Color{l, l, l}
    }

    var r, g, b float64
    var t1 float64
    var t2 float64
    var tr float64
    var tg float64
    var tb float64

    if l < 0.5 {
        t1 = l * (1.0 + s)
    } else {
        t1 = l + s - l*s
    }

    t2 = 2*l - t1
    h = h / 360
    tr = h + 1.0/3.0
    tg = h
    tb = h - 1.0/3.0

    if tr < 0 {
        tr += 1
    }
    if tr > 1 {
        tr -= 1
    }
    if tg < 0 {
        tg += 1
    }
    if tg > 1 {
        tg -= 1
    }
    if tb < 0 {
        tb += 1
    }
    if tb > 1 {
        tb -= 1
    }

    // Red
    if 6*tr < 1 {
        r = t2 + (t1-t2)*6*tr
    } else if 2*tr < 1 {
        r = t1
    } else if 3*tr < 2 {
        r = t2 + (t1-t2)*(2.0/3.0-tr)*6
    } else {
        r = t2
    }

    // Green
    if 6*tg < 1 {
        g = t2 + (t1-t2)*6*tg
    } else if 2*tg < 1 {
        g = t1
    } else if 3*tg < 2 {
        g = t2 + (t1-t2)*(2.0/3.0-tg)*6
    } else {
        g = t2
    }

    // Blue
    if 6*tb < 1 {
        b = t2 + (t1-t2)*6*tb
    } else if 2*tb < 1 {
        b = t1
    } else if 3*tb < 2 {
        b = t2 + (t1-t2)*(2.0/3.0-tb)*6
    } else {
        b = t2
    }

    return Color{r, g, b}
}

// Implement the Go color.Color interface.
func (col Color) RGBA() (r, g, b, a uint32) {
    r = uint32(col.R*65535.0+0.5)
    g = uint32(col.G*65535.0+0.5)
    b = uint32(col.B*65535.0+0.5)
    a = 0xFFFF
    return
}

func MakeColor(col color.Color) Color {
    r, g, b, a := col.RGBA()

    // Since color.Color is alpha pre-multiplied, we need to divide the
    // RGB values by alpha again in order to get back the original RGB.
    r *= 0xffff
    r /= a
    g *= 0xffff
    g /= a
    b *= 0xffff
    b /= a

    return Color{float64(r)/65535.0, float64(g)/65535.0, float64(b)/65535.0}
}