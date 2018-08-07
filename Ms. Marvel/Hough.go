package main
 
import (
	"fmt"
	"log"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "math"
    "os"
	"encoding/csv"
)

type Dato struct{
	Th, R int
}
var	Mapa = make(map[Dato]int)

func hough(im image.Image, ntx, mry int) draw.Image {
    nimx := im.Bounds().Max.X
    mimy := im.Bounds().Max.Y
    mry = int(mry/2) * 2
    him := image.NewGray(image.Rect(0, 0, ntx, mry))
    draw.Draw(him, him.Bounds(), image.NewUniform(color.White),
        image.ZP, draw.Src)
 
    rmax := math.Hypot(float64(nimx), float64(mimy))
    dr := rmax / float64(mry/2)
    dth := math.Pi / float64(ntx)
 
    for jx := 0; jx < nimx; jx++ {
        for iy := 0; iy < mimy; iy++ {
            col := color.GrayModel.Convert(im.At(jx, iy)).(color.Gray)
            if col.Y == 255 {
                continue
            }
            for jtx := 0; jtx < ntx; jtx++ {
                th := dth * float64(jtx)
				r := float64(jx)*math.Cos(th) + float64(iy)*math.Sin(th)                
                iry := mry/2 - int(math.Floor(r/dr+.5))
                col = him.At(jtx, iry).(color.Gray)
                if col.Y > 0 {
                    col.Y--
                    Mapa[Dato{jtx,iry}]++
                    him.SetGray(jtx, iry, col)
                }
            }
        }
    }
    return him
}
 
func main() {
    f, err := os.Open("Pentagon.png")
    if err != nil {
        fmt.Println(err)
        return
    }
    pent, err := png.Decode(f)
    if err != nil {
        fmt.Println(err)
        return
    }
    if err = f.Close(); err != nil {
        fmt.Println(err)
    }
    h := hough(pent, 460, 360)
    if f, err = os.Create("hough.png"); err != nil {
        fmt.Println(err)
        return
    }
    if err = png.Encode(f, h); err != nil {
        fmt.Println(err)
    }
    if cErr := f.Close(); cErr != nil && err == nil {
        fmt.Println(err)
	}
	GreyMat("hough.png")
}

func GreyMat(In string) {
    var Out = "Matriz.csv"
    var PK []string
    PK = make([]string,2)
	file, err := os.Open(In)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	b := img.Bounds()
	
	Peep := make([][]color.Color, b.Max.X)
    for i := 0; i < b.Max.X; i++ {
        Peep[i] = make([]color.Color, b.Max.Y)
	}
	for x := 0; x < b.Max.X; x++ {
		for y := 0; y < b.Max.Y; y++ {
			oldPixel := img.At(x, y)
			Peep[x][y] = color.GrayModel.Convert(oldPixel)
		}
	}
	Outfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
	defer Outfile.Close()

	writer := csv.NewWriter(Outfile)
	defer writer.Flush()
	for Key, value := range Mapa {
        if value == 1{
            PK[0],PK[1] = fmt.Sprint(Key.Th,",",Key.R),fmt.Sprint(value)
            err = writer.Write(PK)
            if err != nil {
                log.Fatal(err)
            }
        }
	}
}
