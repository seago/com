package img

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"

	"github.com/robfig/graphics-go/graphics"
)

type Img struct {
	reader io.Reader
	img    image.Image
	Ext    string
}

func New(r io.Reader, ext string) (img *Img, err error) {
	// var src image.Image
	// switch ext {
	// case "jpeg":
	// 	fallthrough
	// case "jpg":
	// 	src, err = jpeg.Decode(r)
	// case "png":
	// 	src, err = png.Decode(r)
	// }
	// if err != nil {
	// 	return nil, err
	// }
	src, ext, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	img = &Img{reader: r, img: src, Ext: ext}
	return
}

func (img *Img) Thumbnail(width, height float64) (image.Image, error) {
	bound := img.img.Bounds()
	dx := float64(bound.Dx())
	dy := float64(bound.Dy())
	var rate float64 = 0

	if dx < width && dy < height {
		rate = 1
	}
	rate = math.Min(width/dx, height/dy)
	width = dx * rate
	height = dy * rate

	dst := image.NewNRGBA(image.Rect(0, 0, int(width), int(height))) //根据 等比缩放后的宽高尺寸 新建一画板

	err := graphics.Scale(dst, img.img)

	if err != nil {
		return nil, err
	}
	return dst, nil
}
