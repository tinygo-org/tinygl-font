// File generated using:
// 	go run ./generate -font=proggy/ProggyTiny.ttf -size=40 -dpi=72 -package=proggy -weight=Tiny

package proggy

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      2932
// - glyph metadata:  475
// - glyph mask data: 2256

var Tiny40 = font.Make("" +
	"\x00" + // version: 0
	"\x28" + // size:   40
	"\x19" + // height: 25
	"\x13" + // ascent: 19

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xdc\x00" + // "\"" at index 220
	"\xe7\x00" + // "#" at index 231
	"\x07\x01" + // "$" at index 263
	"\x35\x01" + // "%" at index 309
	"\x60\x01" + // "&" at index 352
	"\x8c\x01" + // "'" at index 396
	"\x95\x01" + // "(" at index 405
	"\xb0\x01" + // ")" at index 432
	"\xcb\x01" + // "*" at index 459
	"\xee\x01" + // "+" at index 494
	"\x07\x02" + // "," at index 519
	"\x14\x02" + // "-" at index 532
	"\x21\x02" + // "." at index 545
	"\x29\x02" + // "/" at index 553
	"\x48\x02" + // "0" at index 584
	"\x6c\x02" + // "1" at index 620
	"\x84\x02" + // "2" at index 644
	"\xb2\x02" + // "3" at index 690
	"\xe0\x02" + // "4" at index 736
	"\x07\x03" + // "5" at index 775
	"\x2e\x03" + // "6" at index 814
	"\x58\x03" + // "7" at index 856
	"\x7f\x03" + // "8" at index 895
	"\xa3\x03" + // "9" at index 931
	"\xca\x03" + // ":" at index 970
	"\xd6\x03" + // ";" at index 982
	"\xe8\x03" + // "<" at index 1000
	"\x0b\x04" + // "=" at index 1035
	"\x23\x04" + // ">" at index 1059
	"\x46\x04" + // "?" at index 1094
	"\x74\x04" + // "@" at index 1140
	"\xa3\x04" + // "A" at index 1187
	"\xc4\x04" + // "B" at index 1220
	"\xe8\x04" + // "C" at index 1256
	"\x0c\x05" + // "D" at index 1292
	"\x30\x05" + // "E" at index 1328
	"\x54\x05" + // "F" at index 1364
	"\x71\x05" + // "G" at index 1393
	"\x9b\x05" + // "H" at index 1435
	"\xb5\x05" + // "I" at index 1461
	"\xc9\x05" + // "J" at index 1481
	"\xe3\x05" + // "K" at index 1507
	"\x11\x06" + // "L" at index 1553
	"\x28\x06" + // "M" at index 1576
	"\x4c\x06" + // "N" at index 1612
	"\x66\x06" + // "O" at index 1638
	"\x80\x06" + // "P" at index 1664
	"\x9d\x06" + // "Q" at index 1693
	"\xba\x06" + // "R" at index 1722
	"\xd7\x06" + // "S" at index 1751
	"\x05\x07" + // "T" at index 1797
	"\x19\x07" + // "U" at index 1817
	"\x30\x07" + // "V" at index 1840
	"\x47\x07" + // "W" at index 1863
	"\x61\x07" + // "X" at index 1889
	"\x85\x07" + // "Y" at index 1925
	"\x9f\x07" + // "Z" at index 1951
	"\xcd\x07" + // "[" at index 1997
	"\xe2\x07" + // "\\" at index 2018
	"\x01\x08" + // "]" at index 2049
	"\x16\x08" + // "^" at index 2070
	"\x2e\x08" + // "_" at index 2094
	"\x3c\x08" + // "`" at index 2108
	"\x48\x08" + // "a" at index 2120
	"\x6b\x08" + // "b" at index 2155
	"\x89\x08" + // "c" at index 2185
	"\xa2\x08" + // "d" at index 2210
	"\xc0\x08" + // "e" at index 2240
	"\xe3\x08" + // "f" at index 2275
	"\x01\x09" + // "g" at index 2305
	"\x25\x09" + // "h" at index 2341
	"\x3c\x09" + // "i" at index 2364
	"\x4f\x09" + // "j" at index 2383
	"\x6e\x09" + // "k" at index 2414
	"\x91\x09" + // "l" at index 2449
	"\xa0\x09" + // "m" at index 2464
	"\xb2\x09" + // "n" at index 2482
	"\xc4\x09" + // "o" at index 2500
	"\xdd\x09" + // "p" at index 2525
	"\xfa\x09" + // "q" at index 2554
	"\x17\x0a" + // "r" at index 2583
	"\x30\x0a" + // "s" at index 2608
	"\x53\x0a" + // "t" at index 2643
	"\x70\x0a" + // "u" at index 2672
	"\x86\x0a" + // "v" at index 2694
	"\xa2\x0a" + // "w" at index 2722
	"\xb8\x0a" + // "x" at index 2744
	"\xdb\x0a" + // "y" at index 2779
	"\xfc\x0a" + // "z" at index 2812
	"\x1f\x0b" + // "{" at index 2847
	"\x3a\x0b" + // "|" at index 2874
	"\x47\x0b" + // "}" at index 2887
	"\x62\x0b" + // "~" at index 2914

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x0f\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x0f\xee\x00\x05\x08\x68\xbc\x55\x55\x15\x40\x68\xbc\x01" +

	// glyph 220 for rune '"'
	"\x0f\xec\xf4\x02\x0a\xf8\xe0\x57\x15\x29\xa4" +

	// glyph 231 for rune '#'
	"\x0f\xf1\x00\x00\x0d\x80\x0f\x3e\x10\xfa\xeb\xaf\xf1\xff\xff\x6f\x80\x0f\x3e\x50\xc5\xff\xff\xbf\xa1\xbf\xfe\x1a\xe0\x83\x0f\x04" +

	// glyph 263 for rune '$'
	"\x0f\xee\x00\x00\x0d\x00\xa0\x01\x00\x00\x2f\x00\x01\xfe\xff\x6f\xa8\xfa\xab\xc6\x0b\x2f\x00\x01\xfe\xff\x40\x40\xfa\xab\x06\x00\x2f\xbc\xf1\xff\xff\x40\xa8\xfa\x2b\x00\x00\x2f\x00\x01" +

	// glyph 309 for rune '%'
	"\x0f\xf1\x00\x00\x0f\x80\x0f\x00\xf8\xa1\xaa\x06\xaa\xf2\xc2\x0b\x2f\x04\xf8\xe0\x03\x10\x90\xaa\xaa\x01\x00\xbc\xf0\x42\x80\x0f\x3e\xf8\xa1\x2a\xa4\xaa\xf2\x02\x00\x2f\x04" +

	// glyph 352 for rune '&'
	"\x0f\xee\x00\x00\x0f\x40\xaa\x01\x00\x80\xff\x02\x00\xf1\x02\xf8\x00\x54\x01\xfe\x0b\x00\x84\xaa\xaa\x42\xca\x0b\xe0\x83\x1f\x2f\x00\xf0\x42\xa8\xaa\xaa\xaa\x80\xff\x3f\xf8\x01" +

	// glyph 396 for rune '\''
	"\x0f\xec\xf4\x05\x08\xbc\x55\x85\x06" +

	// glyph 405 for rune '('
	"\x0f\xec\x03\x02\x0a\x00\xe0\x07\xa0\x2a\xc0\x0b\xe1\x03\x50\x55\x55\x91\x6a\x00\xf0\x42\x00\xe0\x07\x00\x29" +

	// glyph 432 for rune ')'
	"\x0f\xec\x03\x02\x0a\xf8\x00\x44\xaa\x01\xc0\x0b\x01\x80\x5f\x55\x55\x01\xa8\x0a\xf0\x42\xf8\x00\x44\x0a\x00" +

	// glyph 459 for rune '*'
	"\x0f\xf1\xfe\x00\x0d\x00\xf0\x02\x10\x1a\xbc\xa0\xf1\xc2\x0b\x6f\x80\xff\x3f\x10\xaa\xfe\xaa\xf1\xc2\x0b\x6f\x00\xf0\x02\x10\x00\x68\x00\x00" +

	// glyph 494 for rune '+'
	"\x0f\xf1\xfe\x00\x0d\x00\xf0\x02\x50\xc5\xff\xff\xbf\xa1\xea\xaf\x1a\x00\xbc\x00\x54\x00\xa0\x01\x00" +

	// glyph 519 for rune ','
	"\x0f\xfd\x03\x02\x08\x00\x1a\xc0\x1b\x3e\x10\x29\x00" +

	// glyph 532 for rune '-'
	"\x0f\xf6\xf9\x00\x0d\xfc\xff\xff\x1b\xaa\xaa\xaa\x01" +

	// glyph 545 for rune '.'
	"\x0f\xfd\x00\x05\x08\x68\xbc\x01" +

	// glyph 553 for rune '/'
	"\x0f\xec\x03\x00\x0d\x00\x00\xc0\x5b\x05\x00\xe0\x03\x55\x00\xf0\x02\x50\x05\xf8\x00\x00\x55\xbc\x00\x00\x10\x1a\x00\x00\x00" +

	// glyph 584 for rune '0'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x15\x2f\xbc\xf0\xc6\x0b\x1a\xbc\xbc\x00\xc0\x5b\xa1\xaa\xaa\x1a\xe0\xff\x0f\x04" +

	// glyph 620 for rune '1'
	"\x0f\xee\x00\x02\x0a\x00\x1a\x00\xbc\x10\xfe\x0b\x91\xbe\x00\xf0\x42\x55\x15\xe9\xaf\xf8\xff\x07" +

	// glyph 644 for rune '2'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x68\x00\xc0\x0b\x00\x00\xbc\x01\xc0\xff\x40\x40\xaa\x2a\x00\xf8\x00\x00\xf1\x02\x00\x40\xfc\xaa\xaa\xc6\xff\xff\xbf\x01" +

	// glyph 690 for rune '3'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x68\x00\xc0\x0b\x00\x00\xbc\x01\xc0\xff\x40\x00\xa0\xaa\x06\x00\x00\xbc\xf1\x02\x00\x6f\xa8\xaa\xaa\x06\xf8\xff\x03\x01" +

	// glyph 736 for rune '4'
	"\x0f\xee\x00\x00\x0d\x00\x00\x29\x00\x00\xe0\x03\x01\xc0\xff\x40\x40\xaa\x3f\x00\xf8\xe0\x03\xf1\x02\xf8\x40\xfc\xaa\xbf\xc6\xff\xff\xbf\x01\x00\xf8\x40\x15" +

	// glyph 775 for rune '5'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\xf1\x02\x00\x40\x15\xff\xff\x0f\x84\xaa\xaa\x6a\x00\x00\xc0\x1b\x2f\x00\xf0\x86\xaa\xaa\x6a\x80\xff\x3f\x10" +

	// glyph 814 for rune '6'
	"\x0f\xee\x00\x00\x0d\x00\xa0\x2a\x00\x00\xff\x03\x01\x3e\x00\x40\xa8\x0a\x00\xc0\x0b\x00\x00\xf1\xff\xff\x40\xfc\xaa\xaa\xc6\x0b\x00\xbc\x15\xaa\xaa\xaa\x01\xfe\xff\x40" +

	// glyph 856 for rune '7'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\x01\x00\x00\x6f\x00\x00\xa9\x06\x00\xe0\x03\x15\x00\xa8\x0a\x00\xc0\x0b\x40\x05\xa4\x1a\x00\x80\x0f\x00\x10" +

	// glyph 895 for rune '8'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x15\xe0\xff\x0f\x84\xaa\xaa\x6a\xbc\x00\xc0\x5b\xa1\xaa\xaa\x1a\xe0\xff\x0f\x04" +

	// glyph 931 for rune '9'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x15\xe0\xff\xff\x06\xa4\xaa\xbe\x00\x00\xc0\x1b\x00\x80\x0f\x04\xa4\xaa\x02\x80\xff\x02\x10" +

	// glyph 970 for rune ':'
	"\x0f\xf3\xfe\x05\x08\x68\xbc\x01\x54\xf1\x86\x06" +

	// glyph 982 for rune ';'
	"\x0f\xf3\x00\x02\x08\x00\x1a\xc0\x1b\x00\x50\x05\xf0\x46\xaa\xe1\x03\x01" +

	// glyph 1000 for rune '<'
	"\x0f\xf1\xfe\x00\x0d\x00\x00\xc0\x1b\x00\xa8\xaa\x01\xc0\xff\x40\xfc\x0f\x00\x10\xaa\xaa\x0a\x00\xc0\xff\x40\x00\x00\xc0\x1b\x00\x00\xa0\x01" +

	// glyph 1035 for rune '='
	"\x0f\xf3\xfb\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\x01\x00\x00\x40\xa8\xaa\xaa\xc6\xff\xff\xbf\x01" +

	// glyph 1059 for rune '>'
	"\x0f\xf1\xfe\x00\x0d\xbc\x00\x00\x10\xaa\x6a\x00\x00\xfe\x0b\x40\x00\x00\xfe\x1b\x90\xaa\xaa\x01\xfe\x0b\x40\xbc\x00\x00\x10\x1a\x00\x00\x00" +

	// glyph 1094 for rune '?'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x68\x00\xc0\x0b\x00\x00\xbc\x01\x00\xf8\x40\x00\xa0\x2a\x00\x00\x2f\x00\x01\x00\x00\x40\x00\xa0\x01\x00\x00\x2f\x00\x01" +

	// glyph 1140 for rune '@'
	"\x0f\xee\x00\x00\x0f\x40\xaa\xaa\x06\x80\xff\xff\x0b\xf1\x02\x00\xe0\xc7\x0b\xaa\x82\xcf\x0b\xff\x83\x5f\xf1\xc2\xff\xfa\xf3\xc2\xff\xff\xc7\x0b\x00\x00\x10\xaa\xaa\xaa\x2a\xe0\xff\xff\x7f" +

	// glyph 1187 for rune 'A'
	"\x0f\xee\x00\x00\x0d\x00\xa0\x01\x00\x00\x2f\x00\x15\x90\xaa\x0a\x00\x3e\xf8\x40\x05\xf8\xfa\x03\x80\xff\x3f\x10\x2f\x00\xf0\x56\x01" +

	// glyph 1220 for rune 'B'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x15\xff\xff\x0f\xc4\xaf\xaa\x6a\xbc\x00\xc0\x5b\xf1\xab\xaa\x1a\xff\xff\x0f\x04" +

	// glyph 1256 for rune 'C'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\xbc\x00\x80\xc6\x0b\x00\x00\x55\xc5\x0b\x00\xbc\xa1\xaa\xaa\x1a\xe0\xff\x0f\x04" +

	// glyph 1292 for rune 'D'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\x01\xc0\xff\x2f\x00\xf1\x02\xf8\x40\xbc\x00\xa9\xc6\x0b\x00\xbc\x55\xc5\x0b\xe0\x03\xf1\xab\xaa\x00\xff\xbf\x00\x04" +

	// glyph 1328 for rune 'E'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\xf1\x02\x00\x40\x15\xff\xff\x0f\xc4\xaf\xaa\x02\xbc\x00\x00\x50\xf1\xab\xaa\x1a\xff\xff\xff\x06" +

	// glyph 1364 for rune 'F'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\xf1\x02\x00\x40\x15\xff\xff\x0f\xc4\xaf\xaa\x02\xbc\x00\x00\x50\x55" +

	// glyph 1393 for rune 'G'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\xbc\x00\x80\xc6\x0b\x00\x00\xf1\x02\xf8\x6f\xbc\x00\xe9\xcb\x0b\x00\xbc\x15\xaa\xaa\xfa\x02\xfe\xff\x6f" +

	// glyph 1435 for rune 'H'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x55\xc5\xff\xff\xbf\xf1\xab\xaa\x2f\x2f\x00\xf0\x56\x15" +

	// glyph 1461 for rune 'I'
	"\x0f\xee\x00\x02\x0a\xa4\xaa\xe2\xff\x1f\xc0\x0b\x55\x55\x15\xe9\xaf\xf8\xff\x07" +

	// glyph 1481 for rune 'J'
	"\x0f\xee\x00\x00\x0d\x00\xa0\xaa\x06\x00\xff\xbf\x01\x00\x00\x6f\x55\x55\x85\xaa\xaa\x6a\xfc\xff\x3f\x10" +

	// glyph 1507 for rune 'K'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\xf1\x02\xf8\x40\xbc\xa0\x2a\xc0\x0b\x2f\x00\xf1\x3f\x00\x40\xfc\xaa\x01\xc0\x0b\x2f\x00\xf1\x02\xf8\x40\xbc\x00\xa9\xc6\x0b\x00\xbc\x01" +

	// glyph 1553 for rune 'L'
	"\x0f\xee\x00\x00\x0d\x68\x00\x00\xc0\x0b\x00\x00\x55\x55\x55\xf1\xab\xaa\x1a\xff\xff\xff\x06" +

	// glyph 1576 for rune 'M'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x15\xbf\x42\xfa\xf2\x3f\xf8\x6f\xc5\xaf\xaa\xbe\xbc\xf0\xc2\x5b\xf1\x82\x06\x2f\x2f\x00\xf0\x06" +

	// glyph 1612 for rune 'N'
	"\x0f\xee\x00\x00\x0d\xa8\x0a\x80\xc6\xff\x00\xbc\x15\xbf\x6a\xf0\xf2\xc2\x0b\x6f\x55\xf1\x02\xf8\x6f\x15" +

	// glyph 1638 for rune 'O'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x55\x55\x85\xaa\xaa\x6a\x80\xff\x3f\x10" +

	// glyph 1664 for rune 'P'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x15\xff\xff\x0f\xc4\xaf\xaa\x02\xbc\x00\x00\x50\x55" +

	// glyph 1693 for rune 'Q'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x55\x55\xbc\x00\xfe\x1b\xaa\xea\xff\x02\xfe\xff\x6f" +

	// glyph 1722 for rune 'R'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x15\xff\xff\x0f\xc4\xaf\xaa\x6a\xbc\x00\xc0\x5b\x55" +

	// glyph 1751 for rune 'S'
	"\x0f\xee\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\xbc\x00\x80\xc6\x0b\x00\x00\x01\xfe\xff\x40\x40\xaa\xaa\x06\x00\x00\xbc\xf1\x02\x00\x6f\xa8\xaa\xaa\x06\xf8\xff\x03\x01" +

	// glyph 1797 for rune 'T'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\x01\xc0\x0b\x40\x55\x55\x55\x01" +

	// glyph 1817 for rune 'U'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x55\x55\x55\xa1\xaa\xaa\x1a\xe0\xff\x0f\x04" +

	// glyph 1840 for rune 'V'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x55\x05\xf8\xe0\x03\x55\x00\xf0\x02\x50\x05" +

	// glyph 1863 for rune 'W'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\xf1\xc2\x0b\x6f\x55\xa1\xaa\xaa\x1a\xe0\x83\x0f\x54\x15" +

	// glyph 1889 for rune 'X'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x15\xaa\x42\xaa\x01\x3e\xf8\x40\x00\xf0\x02\x10\x90\xaa\x0a\x00\x3e\xf8\x40\xbc\x00\xc0\x5b\x05" +

	// glyph 1925 for rune 'Y'
	"\x0f\xee\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x55\x05\xf8\xff\x03\x01\xe9\xaf\x00\x00\xbc\x00\x54\x15" +

	// glyph 1951 for rune 'Z'
	"\x0f\xee\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\x01\x00\x00\x6f\x00\x00\xa9\x06\x00\xe0\x03\x01\xc0\x0b\x40\x40\xaa\x01\x00\xf8\x00\x00\xf1\x02\x00\x40\xfc\xaa\xaa\xc6\xff\xff\xbf\x01" +

	// glyph 1997 for rune '['
	"\x0f\xec\x03\x02\x0a\xf8\xff\x87\xaf\x2a\x3e\x00\x55\x55\x55\x55\xf8\xff\x47\xaa\x2a" +

	// glyph 2018 for rune '\\'
	"\x0f\xec\x03\x00\x0d\xbc\x00\x00\x50\x05\xf8\x00\x00\x55\x00\xf0\x02\x50\x05\x00\xe0\x03\x55\x00\x00\xc0\x1b\x00\x00\xa0\x01" +

	// glyph 2049 for rune ']'
	"\x0f\xec\x03\x02\x0a\xf8\xff\x47\xaa\x3f\x00\xf8\x55\x55\x55\x55\xf8\xff\x47\xaa\x2a" +

	// glyph 2070 for rune '^'
	"\x0f\xee\xf6\x00\x0d\x00\xa0\x01\x00\x00\x2f\x00\x01\x3e\xf8\x40\xa8\x0a\xa9\xc6\x0b\x00\xbc\x01" +

	// glyph 2094 for rune '_'
	"\x0f\x00\x03\x00\x0f\xfc\xff\xff\xff\xa1\xaa\xaa\xaa\x02" +

	// glyph 2108 for rune '`'
	"\x0f\xee\xf4\x05\x0a\x68\xc0\x0b\x01\x7e\x40\x0a" +

	// glyph 2120 for rune 'a'
	"\x0f\xf3\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\x01\x00\x00\x6f\x40\xaa\xea\x0b\xf8\xff\xbf\xf1\x02\x00\x6f\xa8\xaa\xea\x0b\xf8\xff\xbf\x01" +

	// glyph 2155 for rune 'b'
	"\x0f\xec\x00\x00\x0d\xbc\x00\x00\x50\x55\xfc\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x55\xf1\xab\xaa\x1a\xff\xff\x0f\x04" +

	// glyph 2185 for rune 'c'
	"\x0f\xf3\x00\x00\x0d\x40\xaa\xaa\x06\xf8\xff\xbf\xf1\x02\x00\x40\x55\xa1\xaa\xaa\x1a\xe0\xff\xff\x06" +

	// glyph 2210 for rune 'd'
	"\x0f\xec\x00\x00\x0d\x00\x00\xc0\x5b\x55\x40\xaa\xea\x0b\xf8\xff\xbf\xf1\x02\x00\x6f\x55\xa1\xaa\xaa\x2f\xe0\xff\xff\x06" +

	// glyph 2240 for rune 'e'
	"\x0f\xf3\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\xfc\xaa\xea\xcb\xff\xff\xbf\xf1\x02\x00\x40\xa8\xaa\xaa\x06\xf8\xff\xbf\x01" +

	// glyph 2275 for rune 'f'
	"\x0f\xec\x00\x00\x0d\x00\xf0\xff\x1b\x90\xaa\xaa\x01\x3e\x00\x40\x85\xfe\xaa\x02\xfc\xff\x3f\x10\xe0\x03\x00\x54\x55\x05" +

	// glyph 2305 for rune 'g'
	"\x0f\xf3\x05\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x55\xa1\xaa\xaa\x2f\xe0\xff\xff\x06\x00\x00\xbc\x01\xa9\xaa\x1a\xe0\xff\x0f\x04" +

	// glyph 2341 for rune 'h'
	"\x0f\xec\x00\x00\x0d\xbc\x00\x00\x50\x55\xfc\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x55\x55" +

	// glyph 2364 for rune 'i'
	"\x0f\xee\x00\x02\x08\x00\x1a\xc0\x1b\x00\x10\xa9\x86\xff\x06\xf0\x56\x55\x05" +

	// glyph 2383 for rune 'j'
	"\x0f\xee\x05\x00\x0a\x00\x00\x29\x00\x80\x1f\x00\x00\x10\x00\xa8\x0a\x00\xff\x07\x00\xe0\x57\x55\x55\xa8\xaa\x2a\xff\xbf\x10" +

	// glyph 2414 for rune 'k'
	"\x0f\xec\x00\x02\x0d\xf8\x00\x00\x55\x85\x0f\x80\x86\x0f\xc0\x1b\x3e\xf8\x40\xf8\xaa\x02\xf8\x2f\x00\xe1\x83\x0f\x84\x0f\xa9\x86\x0f\xc0\x1b" +

	// glyph 2449 for rune 'l'
	"\x0f\xec\x00\x02\x08\xf8\x6f\xa4\x2f\xc0\x5b\x55\x55\x55\x05" +

	// glyph 2464 for rune 'm'
	"\x0f\xf3\x00\x00\x0d\xa8\x0a\x29\xc0\xff\xe0\x03\xf1\xc2\x0b\x6f\x55\x55" +

	// glyph 2482 for rune 'n'
	"\x0f\xf3\x00\x00\x0d\xa8\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x55\x55" +

	// glyph 2500 for rune 'o'
	"\x0f\xf3\x00\x00\x0d\x40\xaa\x2a\x00\xf8\xff\x03\xf1\x02\x00\x6f\x55\xa1\xaa\xaa\x1a\xe0\xff\x0f\x04" +

	// glyph 2525 for rune 'p'
	"\x0f\xf3\x05\x00\x0d\xa8\xaa\x2a\xc0\xff\xff\x03\xf1\x02\x00\x6f\x55\xf1\xab\xaa\x1a\xff\xff\x0f\xc4\x0b\x00\x00\x55" +

	// glyph 2554 for rune 'q'
	"\x0f\xf3\x05\x00\x0d\x40\xaa\xaa\x06\xf8\xff\xbf\xf1\x02\x00\x6f\x55\xa1\xaa\xaa\x2f\xe0\xff\xff\x06\x00\x00\xbc\x55" +

	// glyph 2583 for rune 'r'
	"\x0f\xf3\x00\x00\x0d\x68\xa0\x2a\xc0\x0b\xff\x03\xf1\x3f\x00\x6f\xfc\x0a\x80\xc6\x0b\x00\x00\x55\x05" +

	// glyph 2608 for rune 's'
	"\x0f\xf3\x00\x00\x0d\x40\xaa\xaa\x06\xf8\xff\xbf\xf1\x02\x00\x40\xa8\xaa\x2a\x00\xf8\xff\x03\x01\x00\x00\x6f\xa8\xaa\xaa\xc6\xff\xff\x03\x01" +

	// glyph 2643 for rune 't'
	"\x0f\xee\x00\x02\x0d\xa4\x00\x00\xf8\x00\x00\x15\xbe\xaa\x00\xfe\xff\x40\xf8\x00\x00\x55\x45\xaa\xaa\x06\xf0\xff\x1b" +

	// glyph 2672 for rune 'u'
	"\x0f\xf3\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x55\x55\xa8\xaa\xea\x0b\xf8\xff\xbf\x01" +

	// glyph 2694 for rune 'v'
	"\x0f\xf3\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x15\xaa\x42\xaa\x01\x3e\xf8\x40\x05\xa4\xaa\x02\x00\xf0\x02\x10" +

	// glyph 2722 for rune 'w'
	"\x0f\xf3\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\xf1\xc2\x0b\x6f\x15\xe0\x83\x0f\x54\x01" +

	// glyph 2744 for rune 'x'
	"\x0f\xf3\x00\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x01\x3e\xf8\x40\x40\xaa\x2a\x00\x00\x2f\x00\x01\x3e\xf8\x40\xa8\x0a\xa9\xc6\x0b\x00\xbc\x01" +

	// glyph 2779 for rune 'y'
	"\x0f\xf3\x05\x00\x0d\x68\x00\x80\xc6\x0b\x00\xbc\x55\x55\xa8\xaa\xea\x0b\xf8\xff\xbf\x01\x00\x00\x6f\x40\xaa\xaa\x06\xf8\xff\x03\x01" +

	// glyph 2812 for rune 'z'
	"\x0f\xf3\x00\x00\x0d\xa8\xaa\xaa\xc6\xff\xff\xbf\x01\x00\xf8\x40\x00\xa0\x2a\x00\x00\x2f\x00\x01\x3e\x00\x40\xe8\xaf\xaa\xc6\xff\xff\xbf\x01" +

	// glyph 2847 for rune '{'
	"\x0f\xec\x03\x02\x0a\x00\xe0\x07\xa0\x2a\xc0\x0b\x55\x85\x0f\x40\xa4\x1a\x00\xbc\x50\x55\x00\xe0\x07\x00\x29" +

	// glyph 2874 for rune '|'
	"\x0f\xec\x03\x05\x08\xbc\x55\x55\x55\x55\x55\xa1\x01" +

	// glyph 2887 for rune '}'
	"\x0f\xec\x03\x02\x0a\xf8\x00\x44\xaa\x01\xc0\x0b\x55\x05\x00\x7e\x00\xaa\x02\xbc\x50\x55\xf8\x00\x44\x0a\x00" +

	// glyph 2914 for rune '~'
	"\x0f\xf6\xfb\x00\x0f\x80\xff\x02\xf8\xa1\xaa\xaa\xaa\xf2\x02\xf8\x2f\x04" +

	"")
