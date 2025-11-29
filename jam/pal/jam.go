package pal

import "image/color"

var jamColors = []color.RGBA{
	{0x3d, 0x17, 0x2f, 0xff}, // Color 0
	{0x58, 0x26, 0x38, 0xff}, // Color 1
	{0x77, 0x39, 0x41, 0xff}, // Color 2
	{0x9f, 0x54, 0x4b, 0xff}, // Color 3
	{0xbf, 0x76, 0x51, 0xff}, // Color 4
	{0xf0, 0xa0, 0x5a, 0xff}, // Color 5
	{0xf6, 0xcd, 0xa0, 0xff}, // Color 6
	{0xb6, 0xd6, 0x2f, 0xff}, // Color 7
	{0x6d, 0xa9, 0x23, 0xff}, // Color 8
	{0x2a, 0x78, 0x39, 0xff}, // Color 9
	{0x2f, 0x52, 0x4f, 0xff}, // Color 10
	{0x1c, 0x14, 0x19, 0xff}, // Color 11
	{0x53, 0x4b, 0x4b, 0xff}, // Color 12
	{0x7a, 0x6c, 0x6d, 0xff}, // Color 13
	{0x9f, 0x91, 0x8c, 0xff}, // Color 14
	{0xd0, 0xc6, 0xb8, 0xff}, // Color 15
	{0xe0, 0xf8, 0xf7, 0xff}, // Color 16
	{0x88, 0xed, 0xf3, 0xff}, // Color 17
	{0x0a, 0xcd, 0xe0, 0xff}, // Color 18
	{0x2a, 0x74, 0xa7, 0xff}, // Color 19
	{0x2a, 0x3d, 0x74, 0xff}, // Color 20
	{0x2a, 0x1e, 0x42, 0xff}, // Color 21
	{0x4f, 0x34, 0x5f, 0xff}, // Color 22
	{0x8c, 0x3d, 0x8a, 0xff}, // Color 23
	{0xce, 0x5a, 0x91, 0xff}, // Color 24
	{0xff, 0xad, 0xb6, 0xff}, // Color 25
	{0xf6, 0xd3, 0x25, 0xff}, // Color 26 - Inserted yellow.
	{0xf6, 0xb4, 0x00, 0xff}, // Color 27
	{0xf6, 0x7b, 0x00, 0xff}, // Color 28
	{0xe7, 0x3e, 0x19, 0xff}, // Color 29
	{0xa8, 0x2d, 0x2d, 0xff}, // Color 30
	{0x7f, 0x6c, 0x92, 0xff}, // Color 31
}

// Zughy32 but with contrast dialed up a bit and yellow instead of bonus gray.
var Jam = struct {
	Colors                                                 []color.RGBA
	Brown0, Brown1, Brown2, Brown3, Brown4, Brown5, Brown6 color.RGBA
	Green3, Green2, Green1, Green0                         color.RGBA
	Gray0, Gray1, Gray2, Gray3, Gray4, Gray5               color.RGBA
	Blue4, Blue3, Blue2, Blue1, Blue0                      color.RGBA
	Purple0, Purple1, Purple2, Purple3                     color.RGBA
	Warm4, Warm3, Warm2, Warm1, Warm0                      color.RGBA
	Purple1B                                               color.RGBA
}{
	Colors:   jamColors,
	Brown0:   jamColors[0],
	Brown1:   jamColors[1],
	Brown2:   jamColors[2],
	Brown3:   jamColors[3],
	Brown4:   jamColors[4],
	Brown5:   jamColors[5],
	Brown6:   jamColors[6],
	Green3:   jamColors[7],
	Green2:   jamColors[8],
	Green1:   jamColors[9],
	Green0:   jamColors[10],
	Gray0:    jamColors[11],
	Gray1:    jamColors[12],
	Gray2:    jamColors[13],
	Gray3:    jamColors[14],
	Gray4:    jamColors[15],
	Gray5:    jamColors[16],
	Blue4:    jamColors[17],
	Blue3:    jamColors[18],
	Blue2:    jamColors[19],
	Blue1:    jamColors[20],
	Blue0:    jamColors[21],
	Purple0:  jamColors[22],
	Purple1:  jamColors[23],
	Purple2:  jamColors[24],
	Purple3:  jamColors[25],
	Warm4:    jamColors[26],
	Warm3:    jamColors[27],
	Warm2:    jamColors[28],
	Warm1:    jamColors[29],
	Warm0:    jamColors[30],
	Purple1B: jamColors[31],
}
