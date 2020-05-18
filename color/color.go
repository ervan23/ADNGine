package color

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColor(r, g, b, a uint8) Color { return Color{R: r, G: g, B: b, A: a} }

var (
	// LightGray ...
	LightGray = NewColor(200, 200, 200, 255)
	//Gray ...
	Gray = NewColor(130, 130, 130, 255)
	//DarkGray ...
	DarkGray = NewColor(80, 80, 80, 255)
	//Yellow ...
	Yellow = NewColor(253, 249, 0, 255) // Yellow
	//Gold ...
	Gold = NewColor(255, 203, 0, 255) // Gold
	//Orange ...
	Orange = NewColor(255, 161, 0, 255) // Orange
	//Pink ...
	Pink = NewColor(255, 109, 194, 255) // Pink
	//Red ...
	Red = NewColor(230, 41, 55, 255) // Red
	//Maroon ...
	Maroon = NewColor(190, 33, 55, 255) // Maroon
	//Green ...
	Green = NewColor(0, 228, 48, 255) // Green
	//Lime ...
	Lime = NewColor(0, 158, 47, 255) // Lime
	//DarkGreen ...
	DarkGreen = NewColor(0, 117, 44, 255) // Dark Green
	//SkyBlue ...
	SkyBlue = NewColor(102, 191, 255, 255) // Sky Blue
	//Blue ...
	Blue = NewColor(0, 121, 241, 255) // Blue
	//DarkBlue ...
	DarkBlue = NewColor(0, 82, 172, 255) // Dark Blue
	//Purple ...
	Purple = NewColor(200, 122, 255, 255) // Purple
	//Violet ...
	Violet = NewColor(135, 60, 190, 255) // Violet
	//DarkPurple ...
	DarkPurple = NewColor(112, 31, 126, 255) // Dark Purple
	//Beige ...
	Beige = NewColor(211, 176, 131, 255) // Beige
	//Brown ...
	Brown = NewColor(127, 106, 79, 255) // Brown
	//DarkBrown ...
	DarkBrown = NewColor(76, 63, 47, 255) // Dark Brown
	//White ...
	White = NewColor(255, 255, 255, 255) // White
	//Black ...
	Black = NewColor(0, 0, 0, 255) // Black
	//Blank (Transparent)
	Blank = NewColor(0, 0, 0, 0) // Blank (Transparent)
	//Magenta ...
	Magenta = NewColor(255, 0, 255, 255) // Magenta
	//RayWhite - Off White
	RayWhite = NewColor(245, 245, 245, 255) // My own White (raylib logo)
	//Aqua
	Aqua = NewColor(0, 162, 156, 255)
	//Gopher Blue
	GopherBlue = NewColor(1, 173, 216, 255)
	//Transparent
	Transparent = NewColor(0, 0, 0, 0)
)

