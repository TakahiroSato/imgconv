package imgconv

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/bmp"
)

// ConvertedImg : 変換済み画像データインターフェース
type ConvertedImg interface {
	SaveAsBmp(string)
	SaveAsPng(string)
}

type convertedImg struct {
	img image.Image
}

// SaveAsBmp : bmpエンコードして指定パスに保存
func (d convertedImg) SaveAsBmp(path string) {
	destImg, err := os.Create(path)
	defer destImg.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bmp.Encode(destImg, d.img)
}

func (d convertedImg) SaveAsPng(path string) {
	destImg, err := os.Create(path)
	defer destImg.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(destImg, d.img)
}

// ToGrayScale : グレースケールに変換
// 参考 : https://qiita.com/tenntenn/items/0471e5f494df82c3e825
// args
// img : 変換対象Image
func ToGrayScale(img image.Image) ConvertedImg {
	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	imageManipulation(bounds, func(x, y int) {
		c := color.Gray16Model.Convert(img.At(x, y))
		gray, _ := c.(color.Gray16)
		dest.Set(x, y, gray)
	})

	return convertedImg{dest}
}

// ToBinary : 二値画像に変換
// args
// img : 変換対象Image
// threshold : 白黒判別閾値(0~255)。R, G, Bどれか一つでもこの値未満だった場合黒にする(reverse=trueの場合は白)
// reverse : 白黒反転フラグ false=閾値以上だと白、未満だと黒 true=閾値より以上だと黒、未満だと白
func ToBinary(img image.Image, threshold uint8, reverse bool) ConvertedImg {
	bounds := img.Bounds()
	dest := image.NewRGBA(bounds)
	imageManipulation(bounds, func(x, y int) {
		r32, g32, b32, a32 := img.At(x, y).RGBA()
		r := uint8(r32 >> 8)
		g := uint8(g32 >> 8)
		b := uint8(b32 >> 8)
		a := uint8(a32 >> 8)
		var val uint8
		if reverse {
			val = uint8(0)
		} else {
			val = uint8(255)
		}

		th := threshold
		if r < th || g < th || b < th {
			if reverse {
				val = uint8(255)
			} else {
				val = uint8(0)
			}
		}
		dest.SetRGBA(x, y, color.RGBA{val, val, val, a})
	})

	return convertedImg{dest}
}

// private functions

func imageManipulation(bounds image.Rectangle, operation func(x, y int)) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			operation(x, y)
		}
	}
}
