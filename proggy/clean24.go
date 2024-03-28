// File generated using:
// 	go run ./generate -font=proggy/ProggyClean.ttf -size=24 -dpi=72 -package=proggy -weight=Clean

package proggy

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      2282
// - glyph metadata:  475
// - glyph mask data: 1606

var Clean24 = font.Make("" +
	"\x00" + // version: 0
	"\x18" + // size:   24
	"\x14" + // height: 20
	"\x0f" + // ascent: 15

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xd8\x00" + // "\"" at index 216
	"\xe1\x00" + // "#" at index 225
	"\xfd\x00" + // "$" at index 253
	"\x1c\x01" + // "%" at index 284
	"\x40\x01" + // "&" at index 320
	"\x61\x01" + // "'" at index 353
	"\x69\x01" + // "(" at index 361
	"\x7d\x01" + // ")" at index 381
	"\x91\x01" + // "*" at index 401
	"\xa8\x01" + // "+" at index 424
	"\xb9\x01" + // "," at index 441
	"\xc3\x01" + // "-" at index 451
	"\xcd\x01" + // "." at index 461
	"\xd4\x01" + // "/" at index 468
	"\xf3\x01" + // "0" at index 499
	"\x0d\x02" + // "1" at index 525
	"\x25\x02" + // "2" at index 549
	"\x45\x02" + // "3" at index 581
	"\x61\x02" + // "4" at index 609
	"\x82\x02" + // "5" at index 642
	"\x9c\x02" + // "6" at index 668
	"\xb6\x02" + // "7" at index 694
	"\xca\x02" + // "8" at index 714
	"\xe2\x02" + // "9" at index 738
	"\xfc\x02" + // ":" at index 764
	"\x05\x03" + // ";" at index 773
	"\x12\x03" + // "<" at index 786
	"\x2b\x03" + // "=" at index 811
	"\x3e\x03" + // ">" at index 830
	"\x59\x03" + // "?" at index 857
	"\x75\x03" + // "@" at index 885
	"\x9c\x03" + // "A" at index 924
	"\xb1\x03" + // "B" at index 945
	"\xcd\x03" + // "C" at index 973
	"\xe7\x03" + // "D" at index 999
	"\x01\x04" + // "E" at index 1025
	"\x19\x04" + // "F" at index 1049
	"\x2d\x04" + // "G" at index 1069
	"\x4e\x04" + // "H" at index 1102
	"\x60\x04" + // "I" at index 1120
	"\x6f\x04" + // "J" at index 1135
	"\x80\x04" + // "K" at index 1152
	"\xa6\x04" + // "L" at index 1190
	"\xb4\x04" + // "M" at index 1204
	"\xc7\x04" + // "N" at index 1223
	"\xd9\x04" + // "O" at index 1241
	"\xf3\x04" + // "P" at index 1267
	"\x07\x05" + // "Q" at index 1287
	"\x26\x05" + // "R" at index 1318
	"\x45\x05" + // "S" at index 1349
	"\x6b\x05" + // "T" at index 1387
	"\x7c\x05" + // "U" at index 1404
	"\x8c\x05" + // "V" at index 1420
	"\x9f\x05" + // "W" at index 1439
	"\xbd\x05" + // "X" at index 1469
	"\xd7\x05" + // "Y" at index 1495
	"\xed\x05" + // "Z" at index 1517
	"\x13\x06" + // "[" at index 1555
	"\x23\x06" + // "\\" at index 1571
	"\x42\x06" + // "]" at index 1602
	"\x52\x06" + // "^" at index 1618
	"\x68\x06" + // "_" at index 1640
	"\x73\x06" + // "`" at index 1651
	"\x7c\x06" + // "a" at index 1660
	"\x94\x06" + // "b" at index 1684
	"\xab\x06" + // "c" at index 1707
	"\xc1\x06" + // "d" at index 1729
	"\xd8\x06" + // "e" at index 1752
	"\xf2\x06" + // "f" at index 1778
	"\x07\x07" + // "g" at index 1799
	"\x20\x07" + // "h" at index 1824
	"\x33\x07" + // "i" at index 1843
	"\x40\x07" + // "j" at index 1856
	"\x58\x07" + // "k" at index 1880
	"\x77\x07" + // "l" at index 1911
	"\x82\x07" + // "m" at index 1922
	"\x92\x07" + // "n" at index 1938
	"\xa0\x07" + // "o" at index 1952
	"\xb2\x07" + // "p" at index 1970
	"\xc9\x07" + // "q" at index 1993
	"\xe0\x07" + // "r" at index 2016
	"\xf0\x07" + // "s" at index 2032
	"\x0a\x08" + // "t" at index 2058
	"\x1b\x08" + // "u" at index 2075
	"\x29\x08" + // "v" at index 2089
	"\x37\x08" + // "w" at index 2103
	"\x55\x08" + // "x" at index 2133
	"\x6b\x08" + // "y" at index 2155
	"\x80\x08" + // "z" at index 2176
	"\x9a\x08" + // "{" at index 2202
	"\xb4\x08" + // "|" at index 2228
	"\xbf\x08" + // "}" at index 2239
	"\xd9\x08" + // "~" at index 2265

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x0b\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x0b\xf4\x00\x04\x06\x78\x55\x15\x40\xe2" +

	// glyph 216 for rune '"'
	"\x0b\xf2\xf7\x03\x08\x18\xc6\xb2\x15" +

	// glyph 225 for rune '#'
	"\x0b\xf4\x00\x00\x0b\x00\x38\x0e\x05\xfe\xff\x0b\xf9\xbe\x06\xb0\x2c\x10\xea\xfb\x02\xff\xff\x03\x38\x0e\x40\x01" +

	// glyph 253 for rune '$'
	"\x0b\xf4\x02\x01\x09\x00\x0e\x00\xbe\x0a\xfc\x3f\x8e\x03\x05\xfc\x0b\xe0\xab\x00\x8e\x47\xfa\x2a\xfe\x2f\x00\x0e\x00\x24\x00" +

	// glyph 284 for rune '%'
	"\x0b\xf4\x00\x00\x0b\xe0\x00\x0e\xa8\x86\x0a\x2c\xcb\x02\xa1\xaa\x06\x80\xe3\x00\x00\xe0\x38\x00\xa8\xaa\x01\x2c\xcb\x06\x69\xa8\x06\x0e\xe0\x00" +

	// glyph 320 for rune '&'
	"\x0b\xf4\x00\x01\x0b\xc0\x0f\x00\xa9\x1a\x80\x03\x0b\x44\xaa\x86\x01\x3f\xb0\x38\xb0\x2c\x0e\xa8\x86\x03\x38\x44\xaa\xaa\x01\xff\xb2" +

	// glyph 353 for rune '\''
	"\x0b\xf2\xf7\x04\x06\x24\x5e\x01" +

	// glyph 361 for rune '('
	"\x0b\xf2\x03\x03\x08\x00\x06\xb0\xe0\x50\x2c\x50\x55\xa8\x00\x0e\x41\x1a\xc0\x02" +

	// glyph 381 for rune ')'
	"\x0b\xf2\x03\x03\x08\x18\xc0\x02\xe0\x50\x00\x5b\x55\x90\x06\x0e\xa1\x02\x0b\x00" +

	// glyph 401 for rune '*'
	"\x0b\xf7\xff\x01\x09\x00\x0e\x90\x38\x89\xe3\x38\xf0\x2f\xa4\xaf\xe2\x38\x0e\xe0\x00\x40\x02" +

	// glyph 424 for rune '+'
	"\x0b\xf7\xff\x01\x09\x00\x0e\x14\xfe\xff\xa4\xaf\x02\x38\x10\x40\x02" +

	// glyph 441 for rune ','
	"\x0b\xfd\x03\x01\x05\xc0\x56\xa4\xe1\x00" +

	// glyph 451 for rune '-'
	"\x0b\xfa\xfc\x01\x09\xf8\xff\x93\xaa\x0a" +

	// glyph 461 for rune '.'
	"\x0b\xfd\x00\x03\x05\x6c\x01" +

	// glyph 468 for rune '/'
	"\x0b\xf2\x02\x01\x09\x00\x40\x02\x00\x1e\x00\xa8\x00\xb0\x04\x90\x06\x80\x03\x01\x2a\x00\x2c\x40\xa4\x01\xe0\x00\x10\x09\x00" +

	// glyph 499 for rune '0'
	"\x0b\xf4\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x78\x38\x89\xe3\x38\x1e\x4e\xe2\x38\x80\x47\xaa\x2a\xf0\x2f" +

	// glyph 525 for rune '1'
	"\x0b\xf4\x00\x01\x09\x00\x0e\x00\x3e\x00\xfc\x00\x8e\x03\x24\x0e\x00\x38\x50\x45\xfa\x2a\xfe\xff" +

	// glyph 549 for rune '2'
	"\x0b\xf4\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x38\x00\xe0\x00\xa0\x02\xc0\x02\xe0\x00\xa0\x02\xc0\x02\xe0\x00\x80\xab\x2a\xfe\xff" +

	// glyph 581 for rune '3'
	"\x0b\xf4\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x38\x00\xe0\x00\xa9\x02\xf8\x02\x00\x78\xe1\x00\x4e\xaa\x2a\xf0\x2f" +

	// glyph 609 for rune '4'
	"\x0b\xf4\x00\x01\x0b\x00\x80\x03\x00\xf8\x00\x00\x3f\x00\x38\x0e\x80\x8a\x03\xb0\xe0\x80\x03\x38\xe0\xaa\x6f\xf8\xff\x2f\x00\xe0\x50" +

	// glyph 642 for rune '5'
	"\x0b\xf4\x00\x01\x09\xf8\xff\xe3\xaa\x8a\x03\x40\xb8\x6a\xe0\xff\x02\x00\x78\xe1\x00\x4e\xaa\x2a\xf0\x2f" +

	// glyph 668 for rune '6'
	"\x0b\xf4\x00\x01\x09\x00\xbe\x00\xaa\x01\x2c\x00\x0e\x00\xb8\x6a\xe0\xff\x82\x03\x78\x45\xaa\x2a\xf0\x2f" +

	// glyph 694 for rune '7'
	"\x0b\xf4\x00\x01\x09\xf8\xff\x93\xaa\x0f\x00\x38\x00\x2c\x05\xe0\x40\x01\x0b\x50" +

	// glyph 714 for rune '8'
	"\x0b\xf4\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x78\xa4\xaa\x02\xff\x82\x03\x78\x45\xaa\x2a\xf0\x2f" +

	// glyph 738 for rune '9'
	"\x0b\xf4\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x78\x05\xfc\x3f\xa0\xfa\x00\x80\x03\xc0\x02\xa8\x06\xf0\x03" +

	// glyph 764 for rune ':'
	"\x0b\xf7\x00\x04\x06\x78\x01\x85\x17" +

	// glyph 773 for rune ';'
	"\x0b\xf7\x03\x01\x05\xc0\x16\x00\x05\x6c\x45\x1a\x0e" +

	// glyph 786 for rune '<'
	"\x0b\xf7\xff\x00\x09\x00\xc0\x0f\xa0\xaa\x00\x3f\xc0\x0f\x00\xa8\x2a\x00\xf0\x03\x00\xc0\x0f\x00\xa8" +

	// glyph 811 for rune '='
	"\x0b\xf8\xfd\x01\x0b\xa4\xaa\x1a\xfe\xff\x0b\x00\x00\x90\xaa\x6a\xf8\xff\x2f" +

	// glyph 830 for rune '>'
	"\x0b\xf7\xff\x01\x0b\xf8\x02\x00\xa9\x1a\x00\xe0\x0b\x00\x00\xbe\x00\xa9\x1a\x80\x2f\x80\x2f\x00\x90\x06\x00" +

	// glyph 857 for rune '?'
	"\x0b\xf4\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x38\x00\xe0\x00\xa0\x02\xc0\x02\xe0\x40\x01\x00\x00\x90\x00\x80\x03" +

	// glyph 885 for rune '@'
	"\x0b\xf4\x00\x00\x0b\x00\xff\x02\x90\xaa\x0a\xe0\x00\x0e\x2c\xf8\xb2\x2c\xea\xb2\x2c\xcb\xb2\xb1\xa8\xaf\xb1\xe0\x3f\x80\x03\x00\x40\xaa\x2a\x00\xfc\x3f\x00" +

	// glyph 924 for rune 'A'
	"\x0b\xf4\x00\x01\x0b\x00\xbe\x40\x01\x0b\x0e\x05\xfc\x3f\x90\xaa\x6a\x38\x00\x6c\x05" +

	// glyph 945 for rune 'B'
	"\x0b\xf4\x00\x01\x0b\xf8\xbf\x00\xae\xaa\x80\x03\x38\x84\xab\x3e\xe0\xff\x0f\x38\x00\x6c\x85\xab\xaa\xe1\xff\x0f" +

	// glyph 973 for rune 'C'
	"\x0b\xf4\x00\x01\x0b\x00\xfe\x03\xa0\xaa\x06\x2c\xc0\xe2\x00\x00\x55\x01\x0b\xb0\x80\xaa\x1a\x80\xff\x00" +

	// glyph 999 for rune 'D'
	"\x0b\xf4\x00\x01\x0b\xf8\xbf\x00\xae\xaa\x80\x03\x38\xe0\x00\xb0\x55\xe1\x00\x0e\xb8\xaa\x02\xfe\x2f\x00" +

	// glyph 1025 for rune 'E'
	"\x0b\xf4\x00\x01\x09\xf8\xff\xe3\xaa\x8a\x03\x40\xb8\x6a\xe0\xff\x82\x03\x40\x85\xab\x2a\xfe\xff" +

	// glyph 1049 for rune 'F'
	"\x0b\xf4\x00\x01\x09\xf8\xff\xe3\xaa\x8a\x03\x40\xb8\x6a\xe0\xff\x82\x03\x40\x55" +

	// glyph 1069 for rune 'G'
	"\x0b\xf4\x00\x01\x0b\x00\xfe\x03\xa0\xaa\x06\x2c\xc0\xe2\x00\x00\x85\x03\xff\xe2\x80\xba\x38\x00\x2c\xb0\x00\x0b\xa8\xaa\x01\xf8\x0f" +

	// glyph 1102 for rune 'H'
	"\x0b\xf4\x00\x01\x0b\x38\x00\x6c\x85\xab\xea\xe2\xff\xbf\x38\x00\x6c\x55" +

	// glyph 1120 for rune 'I'
	"\x0b\xf4\x00\x03\x08\xfc\x8b\x6f\xe0\x50\x55\xe1\x1b\xff\x02" +

	// glyph 1135 for rune 'J'
	"\x0b\xf4\x00\x01\x08\xc0\xbf\x80\xba\x00\xb0\x55\x15\xa9\x1a\xfe\x03" +

	// glyph 1152 for rune 'K'
	"\x0b\xf4\x00\x01\x0b\x38\x00\x2c\x0e\x90\x86\x03\x38\xe0\xc0\x02\x38\x69\x00\x8e\x03\x80\xff\x00\xe0\xaa\x01\x38\xb0\x00\x0e\xe0\x80\x03\xa4\xe1\x00\xb0" +

	// glyph 1190 for rune 'L'
	"\x0b\xf4\x00\x01\x09\x38\x00\x54\x55\x85\xab\x2a\xfe\xff" +

	// glyph 1204 for rune 'M'
	"\x0b\xf4\x00\x00\x0b\xfc\x00\xbe\xc5\xb2\x2c\x5b\x2c\x38\xb0\xc5\x02\x00\x5b" +

	// glyph 1223 for rune 'N'
	"\x0b\xf4\x00\x01\x0b\xf8\x02\x6c\xe1\x38\xb0\x85\x03\xcb\x16\x0e\xe0\x5b" +

	// glyph 1241 for rune 'O'
	"\x0b\xf4\x00\x01\x0b\x00\xbe\x00\xa0\xaa\x00\x2c\x38\xe0\x00\xb0\x55\x01\x0b\x0e\x80\xaa\x02\x80\x2f\x00" +

	// glyph 1267 for rune 'P'
	"\x0b\xf4\x00\x01\x09\xf8\xbf\xe0\xaa\x8a\x03\x78\x85\xff\x0b\xae\x1a\x38\x00\x54" +

	// glyph 1287 for rune 'Q'
	"\x0b\xf4\x02\x01\x0b\x00\xbe\x00\xa0\xaa\x00\x2c\x38\xe0\x00\xb0\x55\x01\x0b\x0e\x80\xaa\x1a\x80\x2f\x0b\x00\xc0\x02\x00\x60" +

	// glyph 1318 for rune 'R'
	"\x0b\xf4\x00\x01\x0b\xf8\xbf\x00\xae\xaa\x80\x03\x38\x54\xf8\xbf\x00\xae\x2e\x80\x03\x0b\xe0\x00\x0e\x38\x40\x1a\x0e\x00\x0b" +

	// glyph 1349 for rune 'S'
	"\x0b\xf4\x00\x01\x0b\xc0\xff\x03\xa9\xaa\x86\x03\xc0\xe2\x00\x00\xa4\x0a\x00\xf0\x03\x00\x00\x3f\x00\x80\x6a\x00\x00\x2c\x0e\x00\x4b\xaa\xaa\x01\xff\x0f" +

	// glyph 1387 for rune 'T'
	"\x0b\xf4\x00\x00\x0b\xfc\xff\xbf\xa8\xbe\x6a\x00\x38\x00\x55\x55\x01" +

	// glyph 1404 for rune 'U'
	"\x0b\xf4\x00\x01\x0b\x38\x00\x6c\x55\x55\xa4\xaa\x1a\xf0\xff\x00" +

	// glyph 1420 for rune 'V'
	"\x0b\xf4\x00\x00\x0b\x2c\x00\xb0\x05\x0e\xe0\x50\x00\xcb\x02\x05\x80\x03\x50" +

	// glyph 1439 for rune 'W'
	"\x0b\xf4\x00\x00\x0b\x2c\x00\xb0\x2c\x24\xb0\x2c\x38\xb0\xb1\xa8\xc6\xb2\x2c\xcb\x86\xba\xbc\x06\xbe\xfc\x00\x0e\xe0\x50" +

	// glyph 1469 for rune 'X'
	"\x0b\xf4\x00\x01\x0b\x38\x00\x6c\x01\x0b\x0e\x80\xaa\x02\x80\x2f\x10\xa0\xaa\x00\x2c\x38\xe0\x00\xb0\x05" +

	// glyph 1495 for rune 'Y'
	"\x0b\xf4\x00\x00\x0b\x2c\x00\xb0\x05\x0e\xe0\x00\x69\xa8\x00\xb0\x2c\x00\x80\x03\x50\x15" +

	// glyph 1517 for rune 'Z'
	"\x0b\xf4\x00\x01\x0b\xf8\xff\x2f\xa9\xaa\x0b\x00\xc0\x02\x00\x0e\x00\xa0\x02\x00\x2c\x00\xe0\x00\x00\x2a\x00\xc0\x02\x00\x0e\x00\x80\xab\xaa\xe1\xff\xbf" +

	// glyph 1555 for rune '['
	"\x0b\xf2\x03\x03\x08\xa8\xc6\xbf\x2c\x50\x55\x55\xc5\x6b\xfc\x0b" +

	// glyph 1571 for rune '\\'
	"\x0b\xf2\x02\x01\x09\x24\x00\xe0\x00\x10\x69\x00\xc0\x02\x04\xa8\x00\x80\x03\x01\xa4\x01\x00\x4b\x00\xa0\x02\x00\x1e\x00\x90" +

	// glyph 1602 for rune ']'
	"\x0b\xf2\x03\x03\x08\xa8\xc6\xbf\x00\x5b\x55\x55\x85\xba\xfc\x0b" +

	// glyph 1618 for rune '^'
	"\x0b\xf2\xfc\x01\x09\x00\x09\x00\x38\x10\xa0\x1a\xc0\xb2\x44\x1a\x2a\x0e\xe0\x91\x00\x09" +

	// glyph 1640 for rune '_'
	"\x0b\x00\x02\x00\x0b\xfc\xff\xbf\xa8\xaa\x6a" +

	// glyph 1651 for rune '`'
	"\x0b\xf2\xf6\x03\x06\x18\x2c\xe0\x90" +

	// glyph 1660 for rune 'a'
	"\x0b\xf7\x00\x01\x09\xc0\xbf\x00\xaa\x0a\x00\x38\xf0\xff\xa4\xea\xe3\x00\x1e\xa9\xfa\xc0\xff\x03" +

	// glyph 1684 for rune 'b'
	"\x0b\xf2\x00\x01\x09\x24\x00\xe0\x00\x50\xe1\xff\x82\xab\x2a\x0e\xe0\x55\xb8\xaa\xe2\xff\x02" +

	// glyph 1707 for rune 'c'
	"\x0b\xf7\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x38\x0e\x00\x85\x03\x38\xa9\xaa\xc0\xbf\x00" +

	// glyph 1729 for rune 'd'
	"\x0b\xf2\x00\x01\x09\x00\x40\x02\x00\x5e\x01\xff\x4f\xaa\x3e\x0e\xe0\x55\xa4\xea\x03\xff\x0f" +

	// glyph 1752 for rune 'e'
	"\x0b\xf7\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x38\xfe\xff\xb8\xaa\xe2\x00\x80\x03\x38\xa9\xaa\xc0\xbf\x00" +

	// glyph 1778 for rune 'f'
	"\x0b\xf2\x00\x01\x09\x00\xa9\x02\xf8\x0f\x2c\x40\xe1\xff\x42\xbe\x06\xb0\x00\x55\x05" +

	// glyph 1799 for rune 'g'
	"\x0b\xf7\x05\x01\x09\xc0\xff\x93\xaa\x8f\x03\x78\x15\xa9\xfa\xc0\xff\x03\x00\x5e\xc0\xbf\x00\xaa\x01" +

	// glyph 1824 for rune 'h'
	"\x0b\xf2\x00\x01\x09\x24\x00\xe0\x00\x50\xe1\xff\x82\xab\x2a\x0e\xe0\x55\x05" +

	// glyph 1843 for rune 'i'
	"\x0b\xf2\x00\x03\x06\x90\xe0\x00\xc5\x8f\x0f\x5e\x55" +

	// glyph 1856 for rune 'j'
	"\x0b\xf2\x03\x01\x08\x00\x60\x00\xb0\x00\x00\x05\xe0\x0b\x90\x0b\x00\x5b\x55\x91\xaa\xe1\x3f\x00" +

	// glyph 1880 for rune 'k'
	"\x0b\xf2\x00\x01\x09\x24\x00\xe0\x00\x50\xe1\x00\x8e\x03\x2a\x0e\x2c\x38\x0e\xe0\x3e\x80\xff\x00\x0e\x2c\x38\xa0\xe2\x00\x0e" +

	// glyph 1911 for rune 'l'
	"\x0b\xf2\x00\x03\x06\xa8\xfc\xe0\x55\x55\x15" +

	// glyph 1922 for rune 'm'
	"\x0b\xf7\x00\x00\x0b\xfc\xcb\x0f\xbc\xaa\x6a\x2c\x38\xb0\x55\x05" +

	// glyph 1938 for rune 'n'
	"\x0b\xf7\x00\x01\x09\xf8\xbf\xe0\xaa\x8a\x03\x78\x55\x01" +

	// glyph 1952 for rune 'o'
	"\x0b\xf7\x00\x01\x09\xc0\xbf\x90\xaa\x8a\x03\x78\x15\xa9\xaa\xc0\xbf\x00" +

	// glyph 1970 for rune 'p'
	"\x0b\xf7\x05\x01\x09\xf8\xbf\xe0\xaa\x8a\x03\x78\x15\xae\xaa\xf8\xbf\xe0\x00\x50\x91\x00\x00" +

	// glyph 1993 for rune 'q'
	"\x0b\xf7\x05\x01\x09\xc0\xff\x93\xaa\x8f\x03\x78\x15\xa9\xfa\xc0\xff\x03\x00\x5e\x01\x00\x09" +

	// glyph 2016 for rune 'r'
	"\x0b\xf7\x00\x01\x09\x38\xbe\xe0\xaa\x8a\x2f\x38\x0e\x00\x55\x01" +

	// glyph 2032 for rune 's'
	"\x0b\xf7\x00\x01\x09\xc0\xff\x93\xaa\x8a\x03\x00\xf0\x03\x80\x6a\x00\xc0\x02\x00\x38\xa9\xaa\xf8\xbf\x00" +

	// glyph 2058 for rune 't'
	"\x0b\xf4\x00\x03\x09\x2c\x40\xf1\xff\xbc\x2a\x0b\x50\x85\xaa\x82\xff" +

	// glyph 2075 for rune 'u'
	"\x0b\xf7\x00\x01\x09\x38\x80\x57\x15\xa9\xfa\xc0\xff\x03" +

	// glyph 2089 for rune 'v'
	"\x0b\xf7\x00\x01\x09\x38\x80\x17\xb0\x2c\x05\xe0\x40\x01" +

	// glyph 2103 for rune 'w'
	"\x0b\xf7\x00\x00\x0b\x2c\x00\xb0\x2c\x24\xb0\x2c\x38\xb0\xb1\xa8\xc6\xb2\x2c\xcb\x82\x2f\x3f\x80\x1b\x3e\x80\x03\x38\x00" +

	// glyph 2133 for rune 'x'
	"\x0b\xf7\x00\x01\x09\x38\x80\x93\x86\x0a\x2c\x0b\x80\x03\x05\x2c\x0b\x69\xa8\x38\x80\x03" +

	// glyph 2155 for rune 'y'
	"\x0b\xf7\x05\x01\x09\x38\x80\x57\x15\xa9\xfa\xc0\xff\x03\x00\x5e\xc0\xbf\x00\xaa\x01" +

	// glyph 2176 for rune 'z'
	"\x0b\xf7\x00\x01\x09\xf8\xff\x93\xaa\x0f\x00\x38\x00\x2c\x00\x69\x00\x38\x00\x2c\x00\xf9\xaa\xf8\xff\x03" +

	// glyph 2202 for rune '{'
	"\x0b\xf2\x03\x01\x09\x00\xa0\x02\xc0\x0f\xe0\x40\x55\xf8\x02\x90\x2a\x00\xe0\x40\x15\x40\xaa\x00\xf0\x03" +

	// glyph 2228 for rune '|'
	"\x0b\xf2\x03\x04\x06\x24\x5e\x55\x55\x55\x01" +

	// glyph 2239 for rune '}'
	"\x0b\xf2\x03\x01\x09\xa4\x01\xe0\x0b\x00\xe0\x40\x55\x00\xf0\x03\xa4\x0a\xe0\x40\x15\xa9\x02\xf8\x02\x00" +

	// glyph 2265 for rune '~'
	"\x0b\xf8\xfc\x00\x0b\x90\x2a\x60\xe0\x3f\xb0\x2c\xf8\x0f\x18\xa4\x0a" +

	"")
