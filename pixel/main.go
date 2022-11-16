// Colornames: https://pkg.go.dev/golang.org/x/image/colornames 

package main

import (
	"image"
	"os"
	_ "image/png" 
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	Title string = "Hello Pixel";
	Width float64 = 800;
	Height float64 = 600;  
)

// Load Picture. 
func loadPicture(imgPath string) (pixel.Picture, error) {
	file, err := os.Open(imgPath);
	if err != nil {
		return nil, err; 
	}
	defer file.Close(); 

	img, _, err := image.Decode(file);
	if err != nil {
		return nil, err; 
	}

	return pixel.PictureDataFromImage(img), nil; 
}

func run() {
	config := pixelgl.WindowConfig{
		Title: Title,
		Bounds: pixel.R(0, 0, Width, Height), 
	}

	window, err := pixelgl.NewWindow(config); 
	if err != nil {
		panic(err); 
	}

	pic, err := loadPicture("./assets/img/witch.png");  
	if err != nil {
		panic(err); 
	}

	sprite := pixel.NewSprite(pic, pic.Bounds()); 

	window.Clear(colornames.Coral); 
	sprite.Draw(window, pixel.IM.Moved(window.Bounds().Center()));   

	for !window.Closed() {
		window.Update(); 
	}
}

func main() {
	pixelgl.Run(run); 
} 