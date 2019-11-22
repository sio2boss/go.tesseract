package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	leptonica "github.com/sio2boss/go.tesseract/pkg/leptonica"
	tesseract "github.com/sio2boss/go.tesseract/pkg/tesseract"
	"log"
	"os"
)

var usage = `go-tesseract, sample app that uses tesseract via golang instead of via file and process execs.
Usage:
  go-tesseract <imagefile>
  go-tesseract -h | --help
  go-tesseract --version

Arguments:
  <imagefile>     Path to image to run tesseract OCR against.

Options:
  -h --help       Show this screen
  --version       Show version
`

func main() {

	// get the image to try
	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], tesseract.Version())
	image, _ := arguments["<imagefile>"].(string)

	// Create new tess instance and point it to the tessdata location. Set language to english.
	t, err := tesseract.New("eng")
	if err != nil {
		log.Fatalf("Error while initializing Tess: %s\n", err)
	}
	defer t.Close()

	// Open a new Pix from file with leptonica
	pix, err := leptonica.NewPixFromFile(image)
	if err != nil {
		log.Fatalf("Error while getting pix from file: %s\n", err)
	}
	defer pix.Close() // remember to cleanup

	// Set the page seg mode to autodetect
	t.SetPageSegMode(tesseract.PSM_AUTO_OSD)

	// Set the image to the tesseract instance
	t.SetImagePix(pix)

	// rRtrieve text from the tesseract instance
	fmt.Println(t.Text())

}
