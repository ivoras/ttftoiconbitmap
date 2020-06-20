package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
)

var ttfFilename = flag.String("ttf", "", "TTF filename")
var templateBitmapFilename = flag.String("template", "", "Template bitmap filename")
var fontSize = flag.Int("size", -1, "Font size (-1 = autosize)")
var fontYfix = flag.Int("yoffset", 0, "Y offset to add to font position")
var fontChars = flag.String("chars", "", "Characters to extract into bitmaps")
var outDir = flag.String("ourdir", ".", "Output directory")

func renderFontChar(outFileName string, font *truetype.Font, img image.Image, ch rune) {

}

func main() {
	flag.Parse()

	if *ttfFilename == "" {
		fmt.Println("TTF filename is a required argument")
		return
	}
	if *templateBitmapFilename == "" {
		fmt.Println("Template bitmap filename is a required argument")
		return
	}
	f, err := os.Open(*templateBitmapFilename)
	if err != nil {
		fmt.Println("Cannot open template bitmap file:", err)
		return
	}

	img, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("Cannot read template bitmap filename:", err)
		return
	}
	fmt.Println("Template bitmap format:", format)
	f.Close()

	f, err = os.Open(*ttfFilename)
	if err != nil {
		fmt.Println("Cannot open font file:", err)
		return
	}
	fontBytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Cannot read font file:", err)
		return
	}

	font, err := truetype.Parse(fontBytes)

	if st, err := os.Stat(*outDir); err != nil || !st.IsDir() {
		fmt.Println("Output directory not found:", *outDir)
		return
	}

	if *fontChars != "" {
		for _, ch := range *fontChars {
			outFileName := fmt.Sprintf("%s/char_%x.png", *outDir, int(ch))
			renderFontChar(outFileName, font, img, rune(ch))
		}
	}

}
