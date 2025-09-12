package pal

import "image/color"

const (
	ZughyBrown0   int = iota // Color 0
	ZughyBrown1              // Color 1
	ZughyBrown2              // Color 2
	ZughyBrown3              // Color 3
	ZughyBrown4              // Color 4
	ZughyBrown5              // Color 5
	ZughyBrown6              // Color 6
	ZughyGreen3              // Color 7
	ZughyGreen2              // Color 8
	ZughyGreen1              // Color 9
	ZughyGreen0              // Color 10
	ZughyGray0               // Color 11
	ZughyGray1               // Color 12
	ZughyGray2               // Color 13
	ZughyGray3               // Color 14
	ZughyGray4               // Color 15
	ZughyGray5               // Color 16
	ZughyBlue4               // Color 17
	ZughyBlue3               // Color 18
	ZughyBlue2               // Color 19
	ZughyBlue1               // Color 20
	ZughyBlue0               // Color 21
	ZughyPurple0             // Color 22
	ZughyPurple1             // Color 23
	ZughyPurple2             // Color 24
	ZughyPurple3             // Color 25
	ZughyWarm3               // Color 26
	ZughyWarm2               // Color 27
	ZughyWarm1               // Color 28
	ZughyWarm0               // Color 29
	ZughyPurple1B            // Color 30
	ZughyGrayB               // Color 31
)

// The original Zughy 32 palette: https://lospec.com/palette-list/zughy-32
var Zughy = []color.RGBA{
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
