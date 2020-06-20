# ttftoiconbitmap

TTF font to icon bitmap converter. This takes an input TTF font file such as FontAwesome or really any reasonable TTF font, and renders it into a template bitmap.

```
Usage of ./ttftoiconbitmap:
  -chars string
        Characters to extract into bitmaps
  -color string
        Font color (default "#ffffff")
  -outdir string
        Output directory (default ".")
  -outprefix string
        Output filename prefix (default "char")
  -size int
        Font size (-1 = autosize) (default -1)
  -template string
        Template bitmap filename
  -ttf string
        TTF filename
  -yfix
        Employ Y-fix for lowercase characters
  -yoffset int
        Y offset to drawing characters (default -10)
```

Example command line:

```
 ./ttftoiconbitmap --ttf otfs/FontAwesome5Free-Solid.ttf --template icon_ring_template.png --outdir fa --outprefix fa
```
