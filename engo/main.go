package main

import (
	"github.com/EngoEngine/engo"
)

const (
	Title string = "Hello Engo";
	Width int = 500;
	Height int = 500; 
)

type myScene struct {};

// Type
func (*myScene) Type() string {
	return "Game"; 
};

// Preload 
func (*myScene) Preload() {
	engo.Files.Load("./assets/img/city.png"); 
}

// Setup 
func (*myScene) Setup(engo.Updater) {

}

func main() {
	opts := engo.RunOptions{
		Title: Title,
		Width: Width,
		Height: Height,  
	};
	
	engo.Run(opts, &myScene{});
} 