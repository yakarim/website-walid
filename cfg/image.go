package cfg

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"

	"github.com/nfnt/resize"
)

// Image ...
func (c Cfg) Image(file *multipart.FileHeader, width, height uint) (string, []byte, string, int64) {

	f, err := file.Open()
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	name := file.Filename
	size1 := file.Size
	buffer := make([]byte, size1)

	f.Read(buffer)

	types := http.DetectContentType(buffer)
	img, size := size(buffer, width, height, types)
	return name, img, types, size
}

func size(imgByte []byte, width, height uint, types string) ([]byte, int64) {
	img, _, _ := image.Decode(bytes.NewReader(imgByte))
	m := resize.Resize(width, height, img, resize.Lanczos3)
	buf := new(bytes.Buffer)

	switch types {
	case "image/png":
		png.Encode(buf, m)
	case "image/gif":
		gif.Encode(buf, m, nil)
	default:
		jpeg.Encode(buf, m, nil)
	}

	size := int64(len(buf.Bytes()))

	return buf.Bytes(), size
}
