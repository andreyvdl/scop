package render

import (
	"image"
	"image/draw"
	"os"
)

func loadTexture(filepath string) *image.RGBA {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	return rgba
}

func flipVertically(img *image.RGBA) {
	row := make([]byte, img.Bounds().Dx()*4)
	height := img.Bounds().Dy()
	half_height := height / 2

	for y := 0; y < half_height; y++ {
		top := y * img.Stride
		bottom := (height - 1 - y) * img.Stride

		copy(row, img.Pix[top:top+img.Stride])
		copy(img.Pix[top:top+img.Stride], img.Pix[bottom:bottom+img.Stride])
		copy(img.Pix[bottom:bottom+img.Stride], row)
	}
}
