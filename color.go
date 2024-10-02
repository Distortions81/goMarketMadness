package main

import "image/color"

var tiColor []color.NRGBA = []color.NRGBA{
	{A: 0},                           //0 transparent
	{R: 255, G: 255, B: 255, A: 255}, //1 black
	{R: 79, G: 176, B: 69, A: 255},   //2 light green
	{R: 192, G: 202, B: 119, A: 255}, //3 med green
	{R: 95, G: 81, B: 237, A: 255},   //4 dark blue
	{R: 192, G: 116, B: 255, A: 255}, //5 light blue
	{R: 173, G: 101, B: 77, A: 255},  //6 dark red
	{R: 103, G: 195, B: 228, A: 255}, //7 cyan
	{R: 204, G: 110, B: 80, A: 255},  //8 med red
	{R: 240, G: 146, B: 116, A: 255}, //9 light red
	{R: 193, G: 202, B: 81, A: 255},  //10 dark yellow
	{R: 209, G: 215, B: 129, A: 255}, //11 light yellow
	{R: 72, G: 156, B: 59, A: 255},   //12 dark green
	{R: 176, G: 104, B: 190, A: 255}, //13 magenta
	{R: 204, G: 204, B: 204, A: 255}, //14 gray
	{R: 255, G: 255, B: 255, A: 255}, //15 white
}
