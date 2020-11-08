package imgconv

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func Test_ConvertAndSaveAsBmp(t *testing.T) {
	img := getTestImage()
	gray := ToGrayScale(img)
	gray.SaveAsBmp("./test/out_gray.bmp")

	binary1 := ToBinary(img, 127, false)
	binary1.SaveAsBmp("./test/out_binary_false.bmp")

	binary2 := ToBinary(img, 127, true)
	binary2.SaveAsBmp("./test/out_binary_true.bmp")
}

func Test_ConvertAndSaveAsPng(t *testing.T) {
	img := getTestImage()
	gray := ToGrayScale(img)
	gray.SaveAsPng("./test/out_gray.png")

	binary1 := ToBinary(img, 127, false)
	binary1.SaveAsPng("./test/out_binary_false.png")

	binary2 := ToBinary(img, 127, true)
	binary2.SaveAsPng("./test/out_binary_true.png")
}

// private functions

func getTestImage() image.Image {
	testImgPath := "./test/test.png"

	imgFile, _ := os.Open(testImgPath)
	defer imgFile.Close()
	img, _ := png.Decode(imgFile)

	return img
}
