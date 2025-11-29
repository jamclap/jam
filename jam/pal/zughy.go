package pal

import "image/color"

var zughyColors = []color.RGBA{
	{0x47, 0x2d, 0x3c, 0xff}, // Color 0
	{0x5e, 0x36, 0x43, 0xff}, // Color 1
	{0x7a, 0x44, 0x4a, 0xff}, // Color 2
	{0xa0, 0x5b, 0x53, 0xff}, // Color 3
	{0xbf, 0x79, 0x58, 0xff}, // Color 4
	{0xee, 0xa1, 0x60, 0xff}, // Color 5
	{0xf4, 0xcc, 0xa1, 0xff}, // Color 6
	{0xb6, 0xd5, 0x3c, 0xff}, // Color 7
	{0x71, 0xaa, 0x34, 0xff}, // Color 8
	{0x39, 0x7b, 0x44, 0xff}, // Color 9
	{0x3c, 0x59, 0x56, 0xff}, // Color 10
	{0x30, 0x2c, 0x2e, 0xff}, // Color 11
	{0x5a, 0x53, 0x53, 0xff}, // Color 12
	{0x7d, 0x70, 0x71, 0xff}, // Color 13
	{0xa0, 0x93, 0x8e, 0xff}, // Color 14
	{0xcf, 0xc6, 0xb8, 0xff}, // Color 15
	{0xdf, 0xf6, 0xf5, 0xff}, // Color 16
	{0x8a, 0xeb, 0xf1, 0xff}, // Color 17
	{0x28, 0xcc, 0xdf, 0xff}, // Color 18
	{0x39, 0x78, 0xa8, 0xff}, // Color 19
	{0x39, 0x47, 0x78, 0xff}, // Color 20
	{0x39, 0x31, 0x4b, 0xff}, // Color 21
	{0x56, 0x40, 0x64, 0xff}, // Color 22
	{0x8e, 0x47, 0x8c, 0xff}, // Color 23
	{0xcd, 0x60, 0x93, 0xff}, // Color 24
	{0xff, 0xae, 0xb6, 0xff}, // Color 25
	{0xf4, 0xb4, 0x1b, 0xff}, // Color 26
	{0xf4, 0x7e, 0x1b, 0xff}, // Color 27
	{0xe6, 0x48, 0x2e, 0xff}, // Color 28
	{0xa9, 0x3b, 0x3b, 0xff}, // Color 29
	{0x82, 0x70, 0x94, 0xff}, // Color 30
	{0x4f, 0x54, 0x6b, 0xff}, // Color 31
}

// The original Zughy 32 palette: https://lospec.com/palette-list/zughy-32
var Zughy = struct {
	Colors                                                 []color.RGBA
	Brown0, Brown1, Brown2, Brown3, Brown4, Brown5, Brown6 color.RGBA
	Green3, Green2, Green1, Green0                         color.RGBA
	Gray0, Gray1, Gray2, Gray3, Gray4, Gray5               color.RGBA
	Blue4, Blue3, Blue2, Blue1, Blue0                      color.RGBA
	Purple0, Purple1, Purple2, Purple3                     color.RGBA
	Warm3, Warm2, Warm1, Warm0                             color.RGBA
	Purple1B, ZughyGrayB                                   color.RGBA
}{
	Colors:     zughyColors,
	Brown0:     zughyColors[0],
	Brown1:     zughyColors[1],
	Brown2:     zughyColors[2],
	Brown3:     zughyColors[3],
	Brown4:     zughyColors[4],
	Brown5:     zughyColors[5],
	Brown6:     zughyColors[6],
	Green3:     zughyColors[7],
	Green2:     zughyColors[8],
	Green1:     zughyColors[9],
	Green0:     zughyColors[10],
	Gray0:      zughyColors[11],
	Gray1:      zughyColors[12],
	Gray2:      zughyColors[13],
	Gray3:      zughyColors[14],
	Gray4:      zughyColors[15],
	Gray5:      zughyColors[16],
	Blue4:      zughyColors[17],
	Blue3:      zughyColors[18],
	Blue2:      zughyColors[19],
	Blue1:      zughyColors[20],
	Blue0:      zughyColors[21],
	Purple0:    zughyColors[22],
	Purple1:    zughyColors[23],
	Purple2:    zughyColors[24],
	Purple3:    zughyColors[25],
	Warm3:      zughyColors[26],
	Warm2:      zughyColors[27],
	Warm1:      zughyColors[28],
	Warm0:      zughyColors[29],
	Purple1B:   zughyColors[30],
	ZughyGrayB: zughyColors[31],
}
