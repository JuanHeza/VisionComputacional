package main

import(
    "os"
    "log"
    "image/jpeg"
    "image"
    "image/color"
)

func SNO(In string,Out string){
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
            BY := b.Max.Y
            BX := b.Max.X
            var Pixel color.Color            
            if (X!=BX-1 && Y!=BY-1){
                Pixel = img.At(X+1,Y+1)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SNE(In string,Out string){
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
            BY := b.Max.Y
            var Pixel color.Color            
            if (X!=0 && Y!=BY-1){
                Pixel = img.At(X-1,Y+1)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SSE(In string,Out string){
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
            var Pixel color.Color            
            if (X!=0 && Y!=0){
                Pixel = img.At(X-1,Y-1)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SSO(In string,Out string){
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
            BX := b.Max.X
            var Pixel color.Color            
            if (X!=BX-1 && Y!=0){
                Pixel = img.At(X+1,Y-1)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SN(In string,Out string){
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
            BY := b.Max.Y
            var Pixel color.Color            
            if (Y!=BY-1){
                Pixel = img.At(X,Y+1)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SS(In string,Out string){
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
            var Pixel color.Color            
            if (Y!=0){
                Pixel = img.At(X,Y-1)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SE(In string,Out string){
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
            var Pixel color.Color            
            if (X!=0){
                Pixel = img.At(X-1,Y)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}

func SO(In string,Out string){
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
            BX := b.Max.X
            var Pixel color.Color            
            if (X!=BX-1){
                Pixel = img.At(X+1,Y)                
            }else{
                Pixel = img.At(X,Y)
            }
			imgSet.Set(X,Y,Pixel)		
		}
	}
		outputfile, err := os.Create(Out)
		if err != nil {
			log.Fatal(err)
		}
			defer outputfile.Close()
		jpeg.Encode(outputfile,imgSet,nil)
}