// Colornames: https://pkg.go.dev/golang.org/x/image/colornames 

package main

import (
	"image"
	"os"
	_"math" 
	"time" 
	_ "image/png" 
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	Title string = "Dummy Witch - Is She Cute ?";  
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

	witchPic, err := loadPicture("./assets/img/witch.png");  
	if err != nil {
		panic(err); 
	}

	witchSprite := pixel.NewSprite(witchPic, witchPic.Bounds()); 

	
	// Sprite Management. 
	
	angle := 0.0;
	lastTime := time.Now(); 
	
	// Put all transform in the loop to make the animation. 
	for !window.Closed() {
		// Delta Time. 
		deltaTime := time.Since(lastTime).Seconds();
		lastTime = time.Now(); 
		
		// Clear Window every loop. 
		window.Clear(colornames.Lightblue);  
		
		// Assign Sprite and Center Alignment. 
		witch := pixel.IM;  
		centerAlign := window.Bounds().Center(); 

		// Framerate. 
		angle += deltaTime * 2;   

		// 2D Transformation. 
		witch = witch.Rotated(pixel.ZV, angle);     
		// witch = witch.Scaled(pixel.ZV, 3);     
		witch = witch.Moved(centerAlign);  
		witchSprite.Draw(window, witch);    

		window.Update(); 
	}
}

func main() {
	pixelgl.Run(run); 
} 