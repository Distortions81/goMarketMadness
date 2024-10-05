package main

import "image/color"

var barColors []color.NRGBA = []color.NRGBA{
	tiColor[6],
	tiColor[3],
	tiColor[1],
	tiColor[11],
	tiColor[12],
	tiColor[13],
	tiColor[15],
	tiColor[4],
	tiColor[2],
	tiColor[13],
	tiColor[8],
	tiColor[14],
	tiColor[5],
	tiColor[9],
	tiColor[10],
	tiColor[6],
}

var tiColor []color.NRGBA = []color.NRGBA{
	{A: 0},                              //0 transparent
	{R: 0, G: 0, B: 0, A: 255},          //1 black
	{R: 0x21, G: 0xc8, B: 0x42, A: 255}, //2 med green
	{R: 0x5e, G: 0xdc, B: 0x78, A: 255}, //3 light green
	{R: 0x54, G: 0x55, B: 0xed, A: 255}, //4 dark blue
	{R: 0x7d, G: 0x76, B: 0xfc, A: 255}, //5 light blue
	{R: 0xd4, G: 0x52, B: 0x4d, A: 255}, //6 dark red
	{R: 0x42, G: 0xeb, B: 0xf5, A: 255}, //7 cyan
	{R: 0xfc, G: 0x55, B: 0x54, A: 255}, //8 med red
	{R: 0xff, G: 0x79, B: 0x78, A: 255}, //9 light red
	{R: 0xd4, G: 0xc1, B: 0x54, A: 255}, //10 dark yellow
	{R: 0xe6, G: 0xce, B: 0x80, A: 255}, //11 light yellow
	{R: 0x21, G: 0xb0, B: 0x3b, A: 255}, //12 dark green
	{R: 0xc9, G: 0x5b, B: 0xba, A: 255}, //13 magenta
	{R: 0xcc, G: 0xcc, B: 0xcc, A: 255}, //14 gray
	{R: 0xff, G: 0xff, B: 0xff, A: 255}, //15 white

}

const (
	C_CLEAR = iota
	C_BLACK
	C_MGREEN
	C_LGREEN
	C_DBLUE
	C_LBLUE
	C_DRED
	C_CYAN
	C_MRED
	C_LRED
	C_DYELLOW
	C_LYELLOW
	C_DGREEN
	C_MAGENTA
	C_GRAY
	C_WHITE
)
