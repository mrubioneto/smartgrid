package plg_image_thumbnail

import (
	_ "embed"
)

//go:embed dist/jpeg_linux_amd64.bin
var binaryThumbnailJpeg []byte
