package main

import (
	"os"
	"crypto/md5"
	"image"
	"image/color"
	"image/png"
)

type coloringStrategy func(int, int) color.Color

func main() {
	identifier := os.Args[1]
	checksum := calculate_checksum(identifier)
	img := draw_image(get_coloring_strategy_2(checksum))
	save_image(img, "1.png")
}

func draw_image(strategy coloringStrategy) *image.RGBA {
	img := image.NewRGBA(image.Rect(0,0, 255, 255))
	for x := 0; x < 255; x++ {
		for y := 0; y < 255; y++ {
			img.Set(x, y, strategy(x, y))
		}		
	}
	return img
}

func calculate_checksum(data string) [16]byte {
	return md5.Sum([]byte(data))
}

func get_coloring_strategy_1(checksum [16]byte) coloringStrategy {

	return func(x int, y int) color.Color {
		row := x / 64
		column := y / 64
		hue := uint8(checksum[row * 4 + column])
		return color.RGBA{hue, hue, hue, 255}
	}
}

func get_coloring_strategy_2(checksum [16]byte) coloringStrategy {
	foreground := color.RGBA{
		uint8(checksum[0]),
		uint8(checksum[7]),
		uint8(checksum[15]),
		255,
	}
	background := color.RGBA{255,255,255,255}
	return func(x int, y int) color.Color {
		row := x / 51
		if row > 2 {
			row = 4 - row
		}
		column := y / 51
		result := background;
		if int(checksum[row * 5 + column]) % 2 == 0 {
			result = foreground
		}
		return result;
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func save_image(img *image.RGBA, filename string) {
	f, err := os.Create(filename)
	check(err)
	png.Encode(f, img)
}