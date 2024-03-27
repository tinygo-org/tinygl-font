// File generated using:
// 	go run ./generate -font=freemono/FreeMono.ttf -size=20 -dpi=72 -package=freemono

package freemono

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      2741
// - glyph metadata:  475
// - glyph mask data: 2065

var Regular20 = font.Make("" +
	"\x00" + // version: 0
	"\x14" + // size:   20
	"\x16" + // height: 22
	"\x10" + // ascent: 16

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xd9\x00" + // "\"" at index 217
	"\xe6\x00" + // "#" at index 230
	"\x07\x01" + // "$" at index 263
	"\x2a\x01" + // "%" at index 298
	"\x4a\x01" + // "&" at index 330
	"\x66\x01" + // "'" at index 358
	"\x6e\x01" + // "(" at index 366
	"\x81\x01" + // ")" at index 385
	"\x93\x01" + // "*" at index 403
	"\xa6\x01" + // "+" at index 422
	"\xb5\x01" + // "," at index 437
	"\xc2\x01" + // "-" at index 450
	"\xca\x01" + // "." at index 458
	"\xd1\x01" + // "/" at index 465
	"\xf4\x01" + // "0" at index 500
	"\x0e\x02" + // "1" at index 526
	"\x20\x02" + // "2" at index 544
	"\x46\x02" + // "3" at index 582
	"\x67\x02" + // "4" at index 615
	"\x83\x02" + // "5" at index 643
	"\x9d\x02" + // "6" at index 669
	"\xbc\x02" + // "7" at index 700
	"\xdc\x02" + // "8" at index 732
	"\xfb\x02" + // "9" at index 763
	"\x19\x03" + // ":" at index 793
	"\x23\x03" + // ";" at index 803
	"\x34\x03" + // "<" at index 820
	"\x4e\x03" + // "=" at index 846
	"\x5e\x03" + // ">" at index 862
	"\x78\x03" + // "?" at index 888
	"\x96\x03" + // "@" at index 918
	"\xbb\x03" + // "A" at index 955
	"\xe7\x03" + // "B" at index 999
	"\x0b\x04" + // "C" at index 1035
	"\x2f\x04" + // "D" at index 1071
	"\x50\x04" + // "E" at index 1104
	"\x71\x04" + // "F" at index 1137
	"\x90\x04" + // "G" at index 1168
	"\xb7\x04" + // "H" at index 1207
	"\xd1\x04" + // "I" at index 1233
	"\xe1\x04" + // "J" at index 1249
	"\x00\x05" + // "K" at index 1280
	"\x26\x05" + // "L" at index 1318
	"\x3d\x05" + // "M" at index 1341
	"\x66\x05" + // "N" at index 1382
	"\x8f\x05" + // "O" at index 1423
	"\xb0\x05" + // "P" at index 1456
	"\xcd\x05" + // "Q" at index 1485
	"\xf4\x05" + // "R" at index 1524
	"\x1b\x06" + // "S" at index 1563
	"\x3b\x06" + // "T" at index 1595
	"\x52\x06" + // "U" at index 1618
	"\x67\x06" + // "V" at index 1639
	"\x93\x06" + // "W" at index 1683
	"\xbf\x06" + // "X" at index 1727
	"\xe5\x06" + // "Y" at index 1765
	"\x04\x07" + // "Z" at index 1796
	"\x24\x07" + // "[" at index 1828
	"\x31\x07" + // "\\" at index 1841
	"\x54\x07" + // "]" at index 1876
	"\x61\x07" + // "^" at index 1889
	"\x72\x07" + // "_" at index 1906
	"\x7b\x07" + // "`" at index 1915
	"\x83\x07" + // "a" at index 1923
	"\xa1\x07" + // "b" at index 1953
	"\xc8\x07" + // "c" at index 1992
	"\xe4\x07" + // "d" at index 2020
	"\x0b\x08" + // "e" at index 2059
	"\x29\x08" + // "f" at index 2089
	"\x41\x08" + // "g" at index 2113
	"\x6a\x08" + // "h" at index 2154
	"\x84\x08" + // "i" at index 2180
	"\x9d\x08" + // "j" at index 2205
	"\xb6\x08" + // "k" at index 2230
	"\xda\x08" + // "l" at index 2266
	"\xe8\x08" + // "m" at index 2280
	"\xff\x08" + // "n" at index 2303
	"\x13\x09" + // "o" at index 2323
	"\x2f\x09" + // "p" at index 2351
	"\x56\x09" + // "q" at index 2390
	"\x7d\x09" + // "r" at index 2429
	"\x92\x09" + // "s" at index 2450
	"\xae\x09" + // "t" at index 2478
	"\xc4\x09" + // "u" at index 2500
	"\xdb\x09" + // "v" at index 2523
	"\xf9\x09" + // "w" at index 2553
	"\x17\x0a" + // "x" at index 2583
	"\x35\x0a" + // "y" at index 2613
	"\x5c\x0a" + // "z" at index 2652
	"\x76\x0a" + // "{" at index 2678
	"\x8c\x0a" + // "|" at index 2700
	"\x95\x0a" + // "}" at index 2709
	"\xa9\x0a" + // "~" at index 2729

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x0c\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x0c\xf4\x00\x05\x07\x2c\x5a\x45\x05\x84\xf2" +

	// glyph 217 for rune '"'
	"\x0c\xf4\xfa\x03\x09\x7c\x7d\x38\x6c\x34\x1c\x04\x01" +

	// glyph 230 for rune '#'
	"\x0c\xf3\x01\x02\x0a\x00\x21\x00\x85\x00\x14\x43\x80\x31\xe0\xff\x0f\x08\x03\x30\x08\xe8\xba\x01\x93\x00\x4c\x41\x80\x14\x00\x22\x00" +

	// glyph 263 for rune '$'
	"\x0c\xf3\x02\x02\x0a\x00\x05\x00\x28\x00\x59\x1e\x0c\x50\x20\x00\xc0\x01\x00\xf8\x02\x00\x34\x00\x40\x22\x00\x88\x02\x0c\xe1\x0b\x00\x05\x14" +

	// glyph 298 for rune '%'
	"\x0c\xf4\x00\x02\x0a\xb0\x07\x60\x30\x80\x80\x00\x49\x03\x90\x02\x02\xe0\x46\x2e\x00\x01\x1a\x00\x87\x01\x08\x08\x30\x24\x40\x3b" +

	// glyph 330 for rune '&'
	"\x0c\xf5\x00\x02\x0a\x00\x01\x00\xab\x00\x06\x40\x80\x00\x40\x07\x40\x32\x1c\x43\x26\x0c\x38\x60\xd0\x00\xab\x1e" +

	// glyph 358 for rune '\''
	"\x0c\xf4\xfa\x05\x07\x7c\x85\x52" +

	// glyph 366 for rune '('
	"\x0c\xf4\x03\x06\x09\xc0\x90\x30\x24\x18\x1c\x0c\x71\x84\x02\x03\x0a\x0c\x04" +

	// glyph 385 for rune ')'
	"\x0c\xf4\x03\x03\x06\x0c\x28\x30\x60\xd0\x01\x17\x34\x24\x0c\x0d\x03\x01" +

	// glyph 403 for rune '*'
	"\x0c\xf4\xfb\x02\x0a\x00\x05\x44\x51\x14\xe4\x1b\x00\x0f\x00\x96\x00\x05\x06" +

	// glyph 422 for rune '+'
	"\x0c\xf6\xff\x01\x0b\x00\x14\x40\x45\xff\xff\x01\x50\x00\x15" +

	// glyph 437 for rune ','
	"\x0c\xfd\x03\x03\x07\xf0\xd1\x83\x07\x0f\x1c\x20\x00" +

	// glyph 450 for rune '-'
	"\x0c\xfa\xfb\x01\x0b\xf4\xff\x1f" +

	// glyph 458 for rune '.'
	"\x0c\xfe\x00\x04\x08\xf4\x05" +

	// glyph 465 for rune '/'
	"\x0c\xf3\x02\x02\x0a\x00\x40\x02\x00\x03\x00\x09\x00\x08\x01\x20\x10\x80\x01\x00\x03\x00\x06\x00\x0c\x00\x14\x00\x30\x00\x50\x00\x40\x00\x00" +

	// glyph 500 for rune '0'
	"\x0c\xf4\x00\x02\x0a\xc0\x3a\xc0\x00\x43\x02\x18\x06\x90\x08\x00\x56\x18\x40\x92\x00\x06\x03\x0c\xb0\x0e" +

	// glyph 526 for rune '1'
	"\x0c\xf4\x00\x02\x0a\x40\x07\x40\x1b\x00\x53\x00\x40\x01\x55\x15\xfd\x7f" +

	// glyph 544 for rune '2'
	"\x0c\xf3\x00\x01\x0a\x00\x04\x00\xb8\x0e\x90\x00\x03\x02\x50\x00\x00\x05\x00\x30\x00\x80\x01\x00\x06\x00\x18\x00\x70\x00\xc0\x01\x00\x03\x80\xf4\xff\x0b" +

	// glyph 582 for rune '3'
	"\x0c\xf3\x01\x02\x0a\x00\x04\x80\xea\x41\x01\x08\x00\x50\x01\x40\x02\xe0\x03\x00\x34\x00\x00\x02\x00\x1c\x02\x60\xb4\x7a\x00\x04\x00" +

	// glyph 615 for rune '4'
	"\x0c\xf4\x00\x02\x0a\x00\x78\x00\x64\x01\x30\x05\x50\x14\xc0\x50\x80\x40\x11\x06\x18\xa8\xfa\x01\x40\x11\x40\x7f" +

	// glyph 643 for rune '5'
	"\x0c\xf4\x00\x02\x0a\xb0\xaa\xc0\x00\x50\xb0\x2f\x80\x01\x03\x00\x20\x00\xc0\x01\x00\xc8\x01\x0c\xb8\x1e" +

	// glyph 669 for rune '6'
	"\x0c\xf3\x00\x03\x0a\x00\x10\x40\xab\x90\x00\x30\x00\x08\x00\x0c\x00\xc8\x2a\x28\x90\x0c\xc0\x0c\x80\x08\xc0\x20\x90\xd0\x2a" +

	// glyph 700 for rune '7'
	"\x0c\xf4\x00\x02\x0a\xac\xea\x22\x00\x05\x00\x08\x00\x30\x00\x90\x00\x80\x00\x00\x03\x00\x08\x00\x18\x00\x30\x00\x80\x00\x40\x01" +

	// glyph 732 for rune '8'
	"\x0c\xf3\x00\x02\x0a\x00\x05\x40\xeb\x01\x03\x0c\x05\x50\x81\x00\x02\xac\x03\x28\x28\x14\x40\x21\x00\x18\x08\x30\xd0\x7a\x00" +

	// glyph 763 for rune '9'
	"\x0c\xf3\x00\x03\x0a\x00\x01\xb0\x1e\x08\x20\x08\x80\x08\xc0\x61\x80\x82\x3a\x02\x00\x03\x40\x02\xc0\x00\x70\xa0\x0e\x00" +

	// glyph 793 for rune ':'
	"\x0c\xf8\x00\x04\x08\xf4\x05\x40\x45\x5f" +

	// glyph 803 for rune ';'
	"\x0c\xf7\x03\x03\x07\x40\x80\x1f\x00\x05\x2f\x3c\xb4\xe0\xc0\x01\x02" +

	// glyph 820 for rune '<'
	"\x0c\xf6\xff\x02\x0a\x00\x80\x03\xd0\x01\x74\x00\x28\x00\x2c\x00\x80\x07\x00\xa0\x00\x00\x28\x00\x40\x03" +

	// glyph 846 for rune '='
	"\x0c\xf8\xfc\x01\x0b\xa4\xaa\x1a\x55\x55\x05\x00\x00\xa0\xaa\xaa" +

	// glyph 862 for rune '>'
	"\x0c\xf6\xff\x02\x0a\x2c\x00\x40\x07\x00\xd0\x01\x00\x38\x00\x80\x03\xd0\x02\xa0\x00\x28\x00\x1c\x00\x00" +

	// glyph 888 for rune '?'
	"\x0c\xf4\x00\x02\x0a\x40\x15\xc0\x96\x43\x01\x24\x00\x80\x00\x40\x02\x40\x03\xc0\x01\x40\x01\x00\x00\x04\xa0\x00\xd0\x07" +

	// glyph 918 for rune '@'
	"\x0c\xf3\x01\x02\x0a\x00\x05\x40\xd7\x01\x03\x0c\x06\x20\x08\x80\x30\xb0\xc3\x30\x08\x83\x20\x0c\x83\x30\xf0\x87\x00\x00\x05\x00\x30\x40\x00\xeb\x02" +

	// glyph 955 for rune 'A'
	"\x0c\xf4\x00\x00\x0c\x00\x15\x00\x00\xea\x03\x00\x40\x19\x00\x00\x83\x00\x00\x08\x03\x00\x18\x14\x00\x30\xc0\x00\xd0\xaa\x03\x80\x01\x24\x00\x02\xc0\x00\x05\x00\x06\xfe\x00\xbf" +

	// glyph 999 for rune 'B'
	"\x0c\xf4\x00\x01\x0b\x50\x15\x00\xae\xaa\x00\x06\x90\x80\x01\x30\x60\x00\x08\x58\x95\x00\xaa\x7a\x80\x01\x60\x60\x00\x60\x60\x00\x18\xff\xff\x01" +

	// glyph 1035 for rune 'C'
	"\x0c\xf4\x01\x01\x0b\x00\x54\x00\xa0\xe5\x06\x06\xc0\x91\x00\x50\x18\x00\x00\x02\x00\x50\x14\x00\x00\x0c\x00\x00\x09\xc0\x01\xac\x0e\x00\x10\x00" +

	// glyph 1071 for rune 'D'
	"\x0c\xf4\x00\x01\x0b\x50\x05\x00\xae\x3e\x00\x03\x30\xc0\x00\x30\x30\x00\x08\x0c\x00\x55\x30\x00\x08\x0c\x40\x02\x03\x34\xf0\xff\x02" +

	// glyph 1104 for rune 'E'
	"\x0c\xf4\x00\x01\x0b\x50\x55\x05\xae\xaa\x03\x06\xc0\x04\x06\x01\x80\xd5\x00\xa0\x3a\x00\x18\x08\x00\x06\x40\x80\x01\x50\xf1\xff\x7f" +

	// glyph 1137 for rune 'F'
	"\x0c\xf4\x00\x01\x0b\x50\x55\x05\xae\xaa\x06\x06\x40\x05\x06\x01\x80\xd5\x00\xa0\x3a\x00\x18\x08\x00\x06\x00\x14\xff\x07\x00" +

	// glyph 1168 for rune 'G'
	"\x0c\xf4\x01\x01\x0b\x00\x54\x00\xa0\xd5\x07\x0a\x80\x81\x00\x00\x18\x00\x00\x02\x00\x10\x02\xf4\x8f\x01\x40\x81\x00\x50\xa0\x00\x14\xd0\xea\x02\x00\x01\x00" +

	// glyph 1207 for rune 'H'
	"\x0c\xf4\x00\x01\x0b\x50\x00\x05\x6c\x90\x07\x05\x90\x14\x64\x55\x02\xa9\xaa\x40\x01\x24\x15\xbe\xe0\x0b" +

	// glyph 1233 for rune 'I'
	"\x0c\xf4\x00\x02\x0a\x50\x55\x90\xaa\x06\x50\x40\x55\x15\xfd\x7f" +

	// glyph 1249 for rune 'J'
	"\x0c\xf4\x01\x01\x0c\x00\x54\x15\x00\xa8\x6b\x00\x00\x02\x55\x14\x00\x02\x51\x00\x0c\xc0\x00\x09\x00\xab\x02\x00\x10\x00\x00" +

	// glyph 1280 for rune 'K'
	"\x0c\xf4\x00\x01\x0b\x50\x01\x05\xae\x90\x07\x06\x30\x80\x41\x02\x60\x24\x00\x98\x03\x00\x6e\x03\x80\x01\x03\x60\x40\x02\x18\xc0\x00\x06\x50\xf0\x0f\xf0" +

	// glyph 1318 for rune 'L'
	"\x0c\xf4\x00\x01\x0b\x50\x05\x00\xe9\x0a\x00\x24\x00\x54\x01\x09\x40\x40\x02\x70\xd1\xff\xff" +

	// glyph 1341 for rune 'M'
	"\x0c\xf4\x00\x00\x0c\x50\x00\x40\xd0\x07\xc0\x07\x26\x80\x09\x98\x00\x23\x60\x09\x86\x80\x31\x0c\x02\x46\x16\x08\x18\x3c\x20\x60\x50\x80\x80\x01\x00\x12\xfe\x00\xbf" +

	// glyph 1382 for rune 'N'
	"\x0c\xf4\x00\x00\x0b\x50\x00\x54\xe4\x02\xb9\xc0\x03\x30\xc0\x08\x30\xc0\x18\x30\xc0\x30\x30\xc0\x90\x30\xc0\x80\x30\xc0\x00\x33\xc0\x00\x36\xc0\x00\x3c\xf0\x0b\x34" +

	// glyph 1423 for rune 'O'
	"\x0c\xf4\x00\x01\x0b\x00\x14\x00\xa0\xa5\x00\x06\x90\x90\x00\x60\x08\x00\x20\x03\x00\x5c\x18\x00\x20\x0c\x00\x03\x0a\xa0\x00\xad\x07" +

	// glyph 1456 for rune 'P'
	"\x0c\xf4\x00\x01\x0a\x50\x05\x80\xab\x2e\x60\x00\x06\x06\xc0\x81\x01\x28\x68\xb9\x80\x56\x00\x18\x00\x14\xff\x07\x00" +

	// glyph 1485 for rune 'Q'
	"\x0c\xf4\x02\x01\x0b\x00\x14\x00\xa0\xa5\x00\x06\x90\x90\x00\x60\x08\x00\x20\x03\x00\x5c\x18\x00\x20\x0c\x00\x03\x0a\xa0\x00\xad\x07\x00\x09\x00\xe0\xbf\x0a" +

	// glyph 1524 for rune 'R'
	"\x0c\xf4\x00\x01\x0c\x50\x05\x00\xb8\xea\x02\x60\x00\x0a\x60\x00\x0c\x81\x01\x0d\x80\xff\x02\x80\x41\x03\x80\x01\x0c\x80\x01\x24\x80\x01\x30\xf0\x0f\xc0\x02" +

	// glyph 1563 for rune 'S'
	"\x0c\xf4\x00\x02\x0a\x00\x05\x80\x96\x4a\x02\x24\x02\x80\x18\x00\x80\x06\x00\x90\x07\x00\x90\x00\x00\x33\x00\xcc\x01\x24\xb6\x2e" +

	// glyph 1595 for rune 'T'
	"\x0c\xf4\x00\x01\x0b\x50\x55\x05\xaa\xaa\x8a\x41\x41\x12\x50\x40\x00\x14\x40\x55\x01\xfe\x0b" +

	// glyph 1618 for rune 'U'
	"\x0c\xf4\x00\x01\x0b\x54\x00\x15\x6e\x90\x0b\x03\xc0\x54\x55\x90\x00\x06\xd0\x7a\x00" +

	// glyph 1639 for rune 'V'
	"\x0c\xf4\x00\x00\x0c\x50\x01\x54\xd0\x0a\xa0\x07\x09\x00\x06\x30\x00\x0c\x80\x01\x14\x00\x08\x20\x00\x30\x80\x00\x40\x81\x01\x00\x0c\x03\x00\x60\x05\x00\x00\x0a\x00\x00\x2c\x00" +

	// glyph 1683 for rune 'W'
	"\x0c\xf4\x00\x00\x0c\x50\x01\x54\xd0\x0a\xa0\x07\x02\x00\x08\x18\x28\x20\x50\xf0\x50\x40\x82\x46\x01\x88\x25\x02\x30\xc2\x08\xc0\x08\x32\x00\x17\xd4\x00\x28\xc0\x02\xe0\x00\x0b" +

	// glyph 1727 for rune 'X'
	"\x0c\xf4\x00\x01\x0b\x50\x00\x05\x2e\x80\x07\x09\x60\x00\x06\x09\x00\xc3\x00\x00\x0f\x00\xc0\x02\x00\xa8\x01\x40\xc2\x01\x30\xc0\x00\x03\xd0\xf0\x07\xfd" +

	// glyph 1765 for rune 'Y'
	"\x0c\xf4\x00\x01\x0b\x50\x00\x05\x2d\x80\x07\x09\x60\x00\x06\x08\x00\xc3\x00\x40\x1a\x00\x80\x02\x00\x50\x00\x15\xe0\xff\x00" +

	// glyph 1796 for rune 'Z'
	"\x0c\xf4\x00\x02\x0a\x50\x55\xa0\xaa\x87\x00\x0c\x02\x18\x00\x24\x00\x30\x00\x30\x00\x60\x00\x90\x00\xc3\x00\x8c\x01\x30\xff\xff" +

	// glyph 1828 for rune '['
	"\x0c\xf4\x03\x05\x09\xb4\x52\x50\x55\x55\x91\x05\x15" +

	// glyph 1841 for rune '\\'
	"\x0c\xf3\x02\x02\x0a\x18\x00\xc0\x00\x00\x06\x00\x20\x00\x01\x08\x10\x40\x02\x00\x0c\x00\x90\x00\x00\x03\x00\x14\x00\xc0\x00\x00\x05\x00\x10" +

	// glyph 1876 for rune ']'
	"\x0c\xf4\x03\x03\x07\xe8\x01\x55\x55\x55\x51\x46\x05" +

	// glyph 1889 for rune '^'
	"\x0c\xf4\xf9\x02\x0a\x00\x0a\x00\x69\x00\x09\x06\x0c\x30\x08\x00\x02" +

	// glyph 1906 for rune '_'
	"\x0c\x01\x03\x00\x0c\xa8\xaa\xaa\x06" +

	// glyph 1915 for rune '`'
	"\x0c\xf3\xf6\x03\x06\x08\x30\xc0" +

	// glyph 1923 for rune 'a'
	"\x0c\xf7\x01\x01\x0b\x00\x69\x00\x64\xa5\x00\x00\x20\x04\xb8\x3a\xc0\x00\x08\x14\x00\x02\x05\xe0\x00\xaa\xe7\x02\x04\x00" +

	// glyph 1953 for rune 'b'
	"\x0c\xf4\x01\x00\x0b\xe4\x00\x00\xc0\x00\x00\x01\x43\x06\x00\x73\x75\x00\x0b\xc0\x00\x03\x40\x02\x03\x00\x06\x1c\x00\x05\x2c\x40\x43\x8f\xaa\x00\x00\x04\x00" +

	// glyph 1992 for rune 'c'
	"\x0c\xf7\x01\x02\x0b\x40\x19\x00\x1a\x3d\x18\x00\xc3\x00\x00\x08\x00\x10\x03\x00\x90\x00\x18\xb4\x7a\x00\x10\x00" +

	// glyph 2020 for rune 'd'
	"\x0c\xf4\x01\x01\x0c\x00\x00\x0e\x00\x00\x0c\x01\x64\x30\x40\x57\x36\xc0\x00\x38\x50\x00\x30\x20\x00\x30\x44\x02\xc0\x00\x07\xe0\x00\xa8\xca\x07\x40\x00\x00" +

	// glyph 2059 for rune 'e'
	"\x0c\xf7\x00\x01\x0b\x00\x69\x00\x74\xa5\x00\x03\x90\x60\x00\x20\xf8\xff\x1f\x02\x00\x40\x02\x00\x80\x01\x20\x80\xab\x07" +

	// glyph 2089 for rune 'f'
	"\x0c\xf4\x00\x02\x0b\x00\xfd\x07\x30\x00\x40\x01\x00\x69\x05\x94\x56\x01\x14\x00\x55\xe1\xff\x07" +

	// glyph 2113 for rune 'g'
	"\x0c\xf7\x04\x01\x0b\x00\x19\x14\x74\x65\x07\x03\xb0\x60\x00\x24\x08\x00\x09\x06\x40\x02\x02\xb0\x80\x02\x26\x40\x2a\x08\x00\x00\x02\x00\x50\x00\x00\x0d\x40\x6a\x00" +

	// glyph 2154 for rune 'h'
	"\x0c\xf4\x00\x01\x0b\x78\x00\x00\x18\x00\x10\x18\x19\x00\x66\x35\x80\x03\x14\x60\x00\x49\x15\xbe\xd0\x0b" +

	// glyph 2180 for rune 'i'
	"\x0c\xf3\x00\x02\x0a\x00\x01\x00\x1c\x00\x10\x00\x00\x00\x50\x01\x40\x19\x00\x50\x40\x55\xfc\xff\x03" +

	// glyph 2205 for rune 'j'
	"\x0c\xf3\x04\x03\x09\x00\x04\x00\x07\x80\x00\x00\x54\x15\x55\x0d\x00\x56\x55\x00\x30\x00\x86\x6a\x00" +

	// glyph 2230 for rune 'k'
	"\x0c\xf4\x00\x01\x0b\xb4\x00\x00\x20\x00\x10\x20\x50\x00\x08\x6d\x00\xc2\x00\x80\x0d\x00\xf0\x02\x00\x4c\x02\x00\x42\x02\x80\x40\x02\x2d\xe0\x0b" +

	// glyph 2266 for rune 'l'
	"\x0c\xf4\x00\x02\x0a\xa0\x07\x00\x14\x50\x55\x15\xff\xff" +

	// glyph 2280 for rune 'm'
	"\x0c\xf7\x00\x00\x0c\x10\x09\x19\xd0\xdb\xda\x01\x0a\x09\x08\x08\x14\x20\x55\xf8\xd1\xc1\x02" +

	// glyph 2303 for rune 'n'
	"\x0c\xf7\x00\x01\x0b\x10\x64\x00\xa9\xd1\x00\x0d\x50\x40\x01\x20\x55\xf8\x02\x2f" +

	// glyph 2323 for rune 'o'
	"\x0c\xf7\x00\x01\x0b\x00\x69\x00\xb0\xe5\x00\x03\xc0\x90\x00\x60\x14\x00\x54\x20\x00\x08\x28\x80\x02\xb8\x2e\x00" +

	// glyph 2351 for rune 'p'
	"\x0c\xf7\x04\x00\x0b\x50\x90\x01\xd0\x1d\x1d\xc0\x02\x30\xc0\x00\x80\xc0\x00\xc0\xc0\x00\x80\xc0\x02\x60\xc0\x09\x28\xc0\xa4\x06\xc0\x00\x00\x45\xaa\x00\x00" +

	// glyph 2390 for rune 'q'
	"\x0c\xf7\x04\x01\x0c\x00\x19\x14\xd0\x95\x6d\x30\x00\x0e\x18\x00\x0c\x08\x00\x0c\x18\x00\x0c\x20\x00\x0d\x90\x80\x0e\x00\x6a\x0c\x00\x00\x0c\x05\x00\xa8\x06" +

	// glyph 2429 for rune 'r'
	"\x0c\xf7\x00\x02\x0b\x54\x40\x41\x8d\xa6\xc0\x06\x00\x1c\x00\xc0\x00\x50\xf1\xff\x03" +

	// glyph 2450 for rune 's'
	"\x0c\xf7\x01\x02\x0a\x40\x1a\xc0\x95\x47\x01\x14\x1c\x00\x80\x2b\x00\x00\x8a\x00\x30\x03\x90\xa8\xba\x00\x10\x00" +

	// glyph 2478 for rune 't'
	"\x0c\xf5\x00\x01\x0a\x80\x00\x10\x75\x15\x90\x57\x05\x20\x00\x54\x01\x03\x10\xd0\xea\x01" +

	// glyph 2500 for rune 'u'
	"\x0c\xf7\x00\x01\x0b\x14\x40\x01\x1d\x90\x02\x06\x80\x54\x60\x00\x09\x24\xd0\x02\xac\xdb\x02" +

	// glyph 2523 for rune 'v'
	"\x0c\xf7\x00\x01\x0b\x54\x00\x15\x5d\x50\x07\x06\x90\x00\x03\x08\x80\x00\x02\x90\x60\x00\x30\x0c\x00\x64\x01\x00\x3c\x00" +

	// glyph 2553 for rune 'w'
	"\x0c\xf7\x00\x01\x0b\x54\x00\x15\x1a\x40\x4a\x02\x80\x81\xb0\x30\x30\x3c\x0c\x48\x15\x02\x35\x5c\x00\x0b\x0e\xc0\x41\x03" +

	// glyph 2583 for rune 'x'
	"\x0c\xf7\x00\x01\x0b\x50\x00\x05\x7c\xd0\x03\x18\x24\x00\x58\x02\x00\x3c\x00\x40\x1a\x00\x28\x28\x80\x02\x28\xf8\x42\x2f" +

	// glyph 2613 for rune 'y'
	"\x0c\xf7\x04\x01\x0b\x54\x00\x15\x1d\x40\x07\x06\x80\x00\x03\x08\x80\x01\x02\xc0\x20\x00\x60\x08\x00\xb0\x00\x00\x28\x00\x00\x02\x10\x80\x01\x80\xaa\x02\x00" +

	// glyph 2652 for rune 'z'
	"\x0c\xf7\x00\x02\x0a\x50\x55\xa0\x55\x47\x01\x06\x00\x09\x00\x09\x00\x0d\x00\x0c\x10\x0c\xc0\xf8\xff\x03" +

	// glyph 2678 for rune '{'
	"\x0c\xf4\x03\x03\x08\x00\x0e\x24\x40\x51\x81\x02\x1d\x00\x02\x50\x14\x90\x00\x1c\x00\x01" +

	// glyph 2700 for rune '|'
	"\x0c\xf4\x02\x05\x07\x54\x55\x55\x55" +

	// glyph 2709 for rune '}'
	"\x0c\xf4\x03\x04\x09\x2c\x00\x06\x50\x50\x01\x0a\xd0\x01\x02\x14\x54\x34\x40\x00" +

	// glyph 2729 for rune '~'
	"\x0c\xf9\xfc\x02\x0a\xe0\x01\x31\x28\x09\x00\x0a" +

	"")
