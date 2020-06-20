package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var ttfFilename = flag.String("ttf", "", "TTF filename")
var templateBitmapFilename = flag.String("template", "", "Template bitmap filename")
var fontSize = flag.Int("size", -1, "Font size (-1 = autosize)")
var fontYfix = flag.Bool("yfix", false, "Employ Y-fix for lowercase characters")
var fontChars = flag.String("chars", "", "Characters to extract into bitmaps")
var fontColor = flag.String("color", "#ffffff", "Font color")
var outDir = flag.String("ourdir", ".", "Output directory")

func renderFontChar(outFileName string, tplImage image.Image, f *truetype.Font, color color.RGBA, size int, yFix bool, ch rune) {

	img := image.NewRGBA(tplImage.Bounds())
	draw.Draw(img, img.Bounds(), tplImage, image.ZP, draw.Src)

	dpi := 72
	if size == -1 {
		size = (img.Bounds().Dy() / 3) * 2
	}

	d := &font.Drawer{
		Dst: img,
		Src: image.NewUniform(color),
		Face: truetype.NewFace(f, &truetype.Options{
			Size:    float64(size),
			DPI:     float64(dpi),
			Hinting: font.HintingFull,
		}),
	}

	b, _ := d.BoundString(string(ch))

	y := (img.Bounds().Dy() + size) / 2

	//fmt.Println("Rune dimensions:", b.Max.X.Ceil(), b.Max.Y.Ceil())

	if yFix {
		y = (img.Bounds().Dy()+size)/2 + b.Max.Y.Ceil()
	}

	d.Dot = fixed.Point26_6{
		X: (fixed.I(img.Bounds().Dx()) - d.MeasureString(string(ch))) / 2,
		Y: fixed.I(y),
	}

	d.DrawString(string(ch))

	outFile, err := os.Create(outFileName)
	if err != nil {
		fmt.Println("Cannot create file", outFileName, err)
		os.Exit(1)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		fmt.Println("Cannot encode png", err)
		os.Exit(1)
	}
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

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Cannot read template bitmap filename:", err)
		return
	}
	f.Close()
	//fmt.Println("Template bitmap format:", format)

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

	color, err := ParseHexColor(*fontColor)
	if err != nil {
		fmt.Println("Error parsing font color:", err)
		return
	}

	if *fontChars != "" {
		for _, ch := range *fontChars {
			outFileName := fmt.Sprintf("%s/char_%x.png", *outDir, int(ch))
			renderFontChar(outFileName, img, font, color, *fontSize, *fontYfix, rune(ch))
		}
	}

}
