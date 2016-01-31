package main

import (
	"github.com/vitkarpov/lorempixum"
	"os"
)

func main() {
	img := lorempixum.GetImage(100, 100)
	lorempixum.StreamImage(os.Stdout, img, "jpeg")
}
