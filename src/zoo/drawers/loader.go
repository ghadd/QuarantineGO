package drawers

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"

	"github.com/faiface/pixel"
)

var filenames = [...]string{
	"zoo/images/eagle.png",
	"zoo/images/parrot.png",
	"zoo/images/whale.png",
	"zoo/images/shark.png",
	"zoo/images/giraffe.png",
	"zoo/images/lion.png",
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// GetSprites return all image sprites
func GetSprites() map[string]*pixel.Sprite {
	images := len(filenames)

	pics := make([]pixel.Picture, images)
	errs := make([]error, images)
	pics[0], errs[0] = loadPicture(filenames[0])
	pics[1], errs[1] = loadPicture(filenames[1])
	pics[2], errs[2] = loadPicture(filenames[2])
	pics[3], errs[3] = loadPicture(filenames[3])
	pics[4], errs[4] = loadPicture(filenames[4])
	pics[5], errs[5] = loadPicture(filenames[5])

	for i, v := range errs {
		if v != nil {
			fmt.Print(pics[i])
			log.Fatal(v)
		}
	}

	sprites := make(map[string]*pixel.Sprite)
	for i := range pics {
		key := filenames[i][strings.Index(filenames[i][4:], "/")+5 : strings.Index(filenames[i], ".")]
		value := pixel.NewSprite(pics[i], pics[i].Bounds())
		sprites[key] = value
	}

	return sprites
}
