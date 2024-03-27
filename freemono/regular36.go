// File generated using:
// 	go run ./generate -font=freemono/FreeMono.ttf -size=36 -dpi=72 -package=freemono

package freemono

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      6690
// - glyph metadata:  475
// - glyph mask data: 6014

var Regular36 = font.Make("" +
	"\x00" + // version: 0
	"\x24" + // size:   36
	"\x28" + // height: 40
	"\x1d" + // ascent: 29

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xe8\x00" + // "\"" at index 232
	"\x0c\x01" + // "#" at index 268
	"\x57\x01" + // "$" at index 343
	"\xad\x01" + // "%" at index 429
	"\x0b\x02" + // "&" at index 523
	"\x53\x02" + // "'" at index 595
	"\x66\x02" + // "(" at index 614
	"\x99\x02" + // ")" at index 665
	"\xc4\x02" + // "*" at index 708
	"\xf3\x02" + // "+" at index 755
	"\x15\x03" + // "," at index 789
	"\x2e\x03" + // "-" at index 814
	"\x3c\x03" + // "." at index 828
	"\x4a\x03" + // "/" at index 842
	"\xb5\x03" + // "0" at index 949
	"\x03\x04" + // "1" at index 1027
	"\x2d\x04" + // "2" at index 1069
	"\x85\x04" + // "3" at index 1157
	"\xdf\x04" + // "4" at index 1247
	"\x28\x05" + // "5" at index 1320
	"\x6f\x05" + // "6" at index 1391
	"\xcb\x05" + // "7" at index 1483
	"\x1d\x06" + // "8" at index 1565
	"\x75\x06" + // "9" at index 1653
	"\xd1\x06" + // ":" at index 1745
	"\xeb\x06" + // ";" at index 1771
	"\x15\x07" + // "<" at index 1813
	"\x63\x07" + // "=" at index 1891
	"\x81\x07" + // ">" at index 1921
	"\xd3\x07" + // "?" at index 2003
	"\x1c\x08" + // "@" at index 2076
	"\x76\x08" + // "A" at index 2166
	"\xef\x08" + // "B" at index 2287
	"\x54\x09" + // "C" at index 2388
	"\xa7\x09" + // "D" at index 2471
	"\xfa\x09" + // "E" at index 2554
	"\x44\x0a" + // "F" at index 2628
	"\x8e\x0a" + // "G" at index 2702
	"\xf3\x0a" + // "H" at index 2803
	"\x26\x0b" + // "I" at index 2854
	"\x42\x0b" + // "J" at index 2882
	"\x7a\x0b" + // "K" at index 2938
	"\xee\x0b" + // "L" at index 3054
	"\x18\x0c" + // "M" at index 3096
	"\x81\x0c" + // "N" at index 3201
	"\xef\x0c" + // "O" at index 3311
	"\x4f\x0d" + // "P" at index 3407
	"\x99\x0d" + // "Q" at index 3481
	"\x08\x0e" + // "R" at index 3592
	"\x7c\x0e" + // "S" at index 3708
	"\xd9\x0e" + // "T" at index 3801
	"\x06\x0f" + // "U" at index 3846
	"\x45\x0f" + // "V" at index 3909
	"\xbe\x0f" + // "W" at index 4030
	"\x32\x10" + // "X" at index 4146
	"\xa0\x10" + // "Y" at index 4256
	"\xee\x10" + // "Z" at index 4334
	"\x47\x11" + // "[" at index 4423
	"\x5b\x11" + // "\\" at index 4443
	"\xc6\x11" + // "]" at index 4550
	"\xdb\x11" + // "^" at index 4571
	"\x02\x12" + // "_" at index 4610
	"\x19\x12" + // "`" at index 4633
	"\x28\x12" + // "a" at index 4648
	"\x7e\x12" + // "b" at index 4734
	"\xe8\x12" + // "c" at index 4840
	"\x2e\x13" + // "d" at index 4910
	"\x98\x13" + // "e" at index 5016
	"\xea\x13" + // "f" at index 5098
	"\x25\x14" + // "g" at index 5157
	"\x93\x14" + // "h" at index 5267
	"\xcf\x14" + // "i" at index 5327
	"\xfc\x14" + // "j" at index 5372
	"\x2d\x15" + // "k" at index 5421
	"\x89\x15" + // "l" at index 5513
	"\xa7\x15" + // "m" at index 5543
	"\xda\x15" + // "n" at index 5594
	"\x07\x16" + // "o" at index 5639
	"\x59\x16" + // "p" at index 5721
	"\xba\x16" + // "q" at index 5818
	"\x1b\x17" + // "r" at index 5915
	"\x4b\x17" + // "s" at index 5963
	"\x94\x17" + // "t" at index 6036
	"\xc9\x17" + // "u" at index 6089
	"\xf8\x17" + // "v" at index 6136
	"\x4c\x18" + // "w" at index 6220
	"\xa0\x18" + // "x" at index 6304
	"\xed\x18" + // "y" at index 6381
	"\x5b\x19" + // "z" at index 6491
	"\x99\x19" + // "{" at index 6553
	"\xc9\x19" + // "|" at index 6601
	"\xd7\x19" + // "}" at index 6615
	"\x05\x1a" + // "~" at index 6661

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x16\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x16\xea\x01\x08\x0d\xd0\x03\x3e\x05\x3d\x15\xf0\x00\x4b\x05\x1c\x40\x00\x00\x15\xf8\xd1\x3f\xfc\x02\x06" +

	// glyph 232 for rune '"'
	"\x16\xea\xf5\x05\x10\xf8\x0b\xff\xd1\x2f\xfc\xd3\x1f\xfc\xc3\x1f\xf8\xc3\x0f\xf8\xc2\x0f\xf4\xc2\x0f\xf4\x81\x0f\xf0\x40\x0b\xf0\x00\x02\x50\x00" +

	// glyph 268 for rune '#'
	"\x16\xe9\x02\x03\x12\x00\xd0\x80\x02\x00\xe0\xc0\x02\x01\x80\x02\x07\x00\xc0\x02\x07\x14\x00\x1c\x30\x00\x50\x6d\x79\x05\xfc\xff\xff\x3f\x00\x1d\x38\x00\x00\x0d\x38\x40\x41\xe5\xd5\x57\xe0\xff\xff\xff\x01\xe0\xc0\x02\x04\x80\x02\x07\x00\xc0\x02\x07\x10\x00\x0b\x0c\x00\x00\x07\x0d\x40\x00\x0c\x34\x00" +

	// glyph 343 for rune '$'
	"\x16\xe8\x03\x04\x12\x00\x40\x00\x00\x00\xb0\x00\x10\x00\xf9\x06\x00\xf0\xeb\xab\x00\x1f\x00\x3f\xe0\x00\x00\x0f\x2c\x00\x40\x01\x07\x00\x00\xc0\x02\x00\x00\xe0\x02\x00\x00\xe0\x1b\x00\x00\x40\xff\x06\x00\x00\x94\x0f\x00\x00\x00\x0f\x00\x00\x00\x17\x07\x00\x00\xc7\x03\x00\xe0\xf0\x03\x00\x2d\xec\x5b\xf9\x02\x41\xfe\x0b\x00\x00\x2c\x00\x54\x01" +

	// glyph 429 for rune '%'
	"\x16\xea\x01\x03\x12\x00\xfd\x01\x00\x80\x9b\x0b\x00\xd0\x01\x1c\x00\xa0\x00\x38\x00\xb0\x00\x34\x00\xa0\x00\x38\x00\xd0\x01\x1d\x00\x80\x9b\x0b\x00\x00\xf9\x01\x90\x00\x00\x40\xbe\x00\x00\xfd\x06\x00\xf4\x1b\x00\xe0\x6f\x00\x00\xbc\x01\x50\x00\x00\x00\xfe\x0b\x00\x80\x07\x2d\x00\xc0\x01\x34\x00\xd0\x00\x70\x01\x00\x07\xd0\x00\x00\x1e\xb4\x00\x00\xf8\x2f\x00\x00\x40\x01\x00" +

	// glyph 523 for rune '&'
	"\x16\xed\x01\x04\x11\x00\xa0\x11\x00\xe0\xff\x07\x40\x0b\x18\x00\x3c\x00\x00\xc0\x02\x00\x10\xd0\x00\x00\x00\x2c\x00\x00\x40\x03\x00\x00\xfe\x00\x00\x7c\x1d\xd0\xe3\x80\x03\x1e\x0b\xb0\x70\x70\x00\x5d\x03\x07\x80\x2f\xb0\x00\xf0\x00\x0e\x40\x0f\xc0\x07\xfd\x17\xf0\xbf\xf0\x03\x54\x01\x00" +

	// glyph 595 for rune '\''
	"\x16\xea\xf5\x08\x0d\xf4\x1f\xfc\xc3\x6f\xf0\x07\x7e\xe0\x03\x3d\xc0\x03\x14" +

	// glyph 614 for rune '('
	"\x16\xea\x05\x0a\x11\x00\x20\x00\x74\x00\x3c\x00\x1e\x00\x0f\x80\x07\xc0\x03\xd0\x02\xe0\x01\xf0\x00\xd1\x03\x54\xf0\x00\xc1\x07\x80\x0b\x40\x0f\x00\x0f\x00\x2e\x00\x3c\x00\x78\x00\xf0\x00\xd0\x01\x40\x00" +

	// glyph 665 for rune ')'
	"\x16\xea\x04\x05\x0b\x14\x00\x1e\x00\x0f\x80\x07\xc0\x03\xe0\x01\xf4\x00\x3c\x00\x1e\x40\x0b\xd0\x03\xf0\x55\x01\xb4\x00\x1e\xc0\x03\xb0\x00\x0e\xc0\x03\x78\x00\x0b\xe0\x00" +

	// glyph 708 for rune '*'
	"\x16\xea\xf7\x04\x12\x00\x80\x00\x00\x00\xb0\x00\x50\x00\xc0\x02\x04\xbf\xb1\xe4\x07\xf9\xff\x1f\x00\xd0\x1f\x00\x00\xf0\x0b\x00\x00\x4f\x07\x00\xe0\x81\x03\x00\x2c\xc0\x03\x40\x02\xd0\x00" +

	// glyph 755 for rune '+'
	"\x16\xed\xff\x03\x13\x00\x00\x07\x00\x00\x00\x2c\x00\x50\x55\xfc\xff\xff\xff\xa2\xaa\xbe\xaa\x06\x00\xb0\x00\x40\x55\x01\x00\x18\x00\x00" +

	// glyph 789 for rune ','
	"\x16\xfb\x05\x05\x0c\x80\xff\xc0\x7f\xd0\x2f\xe0\x0f\xf0\x0b\xf0\x03\xf4\x02\xf8\x00\x7c\x00\x3c\x00" +

	// glyph 814 for rune '-'
	"\x16\xf5\xf7\x03\x13\xf8\xff\xff\xff\xa2\xaa\xaa\xaa\x0a" +

	// glyph 828 for rune '.'
	"\x16\xfc\x01\x08\x0e\xf0\x0b\xfe\xc7\xff\xd1\x3f\x80\x01" +

	// glyph 842 for rune '/'
	"\x16\xe8\x03\x04\x12\x00\x00\x00\x1c\x00\x00\x40\x03\x00\x00\xf0\x00\x00\x00\x1d\x00\x00\xc0\x03\x00\x00\x74\x00\x00\x00\x0f\x00\x00\xc0\x01\x00\x00\x38\x00\x00\x00\x07\x00\x00\xe0\x00\x00\x00\x2c\x00\x00\x80\x03\x00\x00\xb0\x00\x00\x00\x0e\x00\x00\xc0\x02\x00\x00\x34\x00\x00\x00\x0b\x00\x00\xd0\x01\x00\x00\x3c\x00\x00\x40\x07\x00\x00\xf0\x00\x00\x00\x1d\x00\x00\x80\x03\x00\x00\x70\x00\x00\x00\x0e\x00\x00\x80\x01\x00\x00\x00" +

	// glyph 949 for rune '0'
	"\x16\xea\x01\x04\x12\x00\xf8\x1f\x00\xc0\x57\x3e\x00\x3c\x00\x2d\x80\x03\x00\x0e\xb0\x00\x00\x0b\x1d\x00\x80\x83\x03\x00\xd0\xf1\x00\x00\x70\x2c\x00\x00\x1c\x0b\x00\x00\x5b\xb1\x00\x00\x70\xe1\x00\x00\x70\x34\x00\x00\x0e\x2c\x00\xc0\x03\x0f\x00\x74\x40\x0b\x00\x0f\x40\x0b\xf4\x00\x40\xff\x0f\x00\x00\x55\x00\x00" +

	// glyph 1027 for rune '1'
	"\x16\xea\x00\x04\x12\x00\xf0\x02\x00\x00\xbf\x00\x00\xf0\x2d\x00\x00\x1f\x0b\x00\xf0\xc1\x02\x00\x1d\xb0\x00\x00\x00\x2c\x00\x54\x55\x55\x45\x55\x7d\x55\xe0\xff\xff\x7f" +

	// glyph 1069 for rune '2'
	"\x16\xea\x00\x03\x11\x00\xf8\x7f\x00\xc0\x5b\xb9\x00\x3c\x00\xb4\x80\x03\x00\x74\x70\x00\x00\x28\x08\x00\x00\x0d\x00\x00\x80\x03\x00\x00\xf0\x00\x00\x00\x1d\x00\x00\xd0\x02\x00\x00\x3c\x00\x00\xc0\x03\x00\x00\x3c\x00\x00\xd0\x03\x00\x00\x2d\x00\x00\xd0\x02\x00\x00\x2d\x00\x00\xd0\x02\x00\x00\x1e\x00\x00\xe2\x01\x00\xd0\x7c\x55\x55\x39\xff\xff\xff\x0f" +

	// glyph 1157 for rune '3'
	"\x16\xea\x01\x03\x12\x00\xf8\xbf\x00\x40\x6f\xe5\x07\xd0\x02\x00\x0e\x80\x00\x00\x3c\x00\x00\x00\x34\x01\x00\x00\xe0\x00\x00\x00\xb0\x00\x00\x00\x3d\x00\x00\xf9\x0b\x00\x00\xfe\x03\x00\x00\x00\x2e\x00\x00\x00\xb4\x00\x00\x00\xd0\x01\x00\x00\xc0\x02\x00\x00\x80\x07\x00\x00\x00\x0b\x00\x00\x40\x47\x07\x00\xd0\x02\x7e\x00\xbd\x00\xe4\xff\x1f\x00\x00\x65\x00\x00" +

	// glyph 1247 for rune '4'
	"\x16\xea\x00\x04\x11\x00\x00\xfc\x00\x00\xe0\x0f\x00\x40\xd3\x00\x00\x2c\x0d\x00\xe0\xd0\x00\x00\x07\x0d\x00\x38\xd0\x00\xd0\x01\x0d\x00\x0b\xd0\x00\x34\x00\x0d\xc0\x02\xd0\x00\x0e\x00\x0d\x70\x00\xd0\xc0\x03\x00\x0d\xac\xaa\xfa\xc6\xff\xff\xff\x00\x00\xd0\x50\x01\x00\x95\x17\x00\xf0\xff\x03" +

	// glyph 1320 for rune '5'
	"\x16\xea\x01\x03\x12\x40\xff\xff\x0b\x80\xab\xaa\x0a\x80\x03\x00\x00\x55\x01\x9e\xff\x07\x00\xfe\x56\x2e\x00\x09\x00\xb4\x00\x00\x00\xe0\x00\x00\x00\xc0\x02\x00\x00\x80\x17\x00\x00\x00\x3c\x00\x00\x00\x2c\x04\x00\x00\x0e\x2d\x00\x40\x0b\xf4\x02\xf4\x02\x80\xff\x7f\x00\x00\x90\x05\x00" +

	// glyph 1391 for rune '6'
	"\x16\xea\x01\x05\x13\x00\x40\xff\x0f\x00\xfc\x56\x02\xd0\x07\x00\x00\x2d\x00\x00\xc0\x03\x00\x00\x38\x00\x00\x00\x07\x00\x00\xd0\x00\x00\x00\x28\x00\x00\x00\x0b\xf9\x0b\xc0\xe1\x9b\x0f\x70\x1d\x00\x0f\xdc\x01\x00\x0b\x2f\x00\x80\xc3\x03\x00\xd0\xa0\x00\x00\x70\x34\x00\x00\x0d\x1c\x00\x40\x03\x0e\x00\xf0\x00\x0b\x00\x0e\x40\x1f\xe0\x02\x40\xff\x1f\x00\x00\x55\x00\x00" +

	// glyph 1483 for rune '7'
	"\x16\xea\x00\x04\x11\xfc\xff\xff\xcf\xab\xaa\xfa\x0c\x00\x00\xcd\x00\x00\xe0\x00\x00\x00\x0b\x00\x00\x70\x00\x00\x80\x03\x00\x00\x2c\x00\x00\xc0\x01\x00\x00\x0d\x00\x00\xb0\x00\x00\x00\x07\x00\x00\x34\x00\x00\x80\x03\x00\x00\x1c\x00\x00\xd0\x00\x00\x00\x0e\x00\x00\xb0\x00\x00\x40\x03\x00\x00\x38\x00\x00\xc0\x02\x00\x00\x0c\x00" +

	// glyph 1565 for rune '8'
	"\x16\xea\x01\x04\x12\x00\xfd\x2f\x00\xe0\x57\x3e\x00\x1d\x00\x3c\xc0\x02\x00\x2c\x34\x00\x00\x0e\x0e\x00\x40\x43\x03\x00\xd0\xc0\x01\x00\x3c\xe0\x00\x80\x03\xe0\x02\x7d\x00\xd0\xff\x03\x00\xbc\xea\x02\xd0\x02\xc0\x03\x2c\x00\xc0\x82\x03\x00\xd0\xb0\x00\x00\x70\xe1\x00\x00\x70\x34\x00\x00\x0e\x3c\x00\xd0\x01\x7c\x00\x2e\x00\xf8\xff\x01\x00\x50\x01\x00" +

	// glyph 1653 for rune '9'
	"\x16\xea\x01\x05\x13\x00\xfe\x0b\x00\xf0\x56\x1f\x00\x0e\x00\x0e\xd0\x01\x00\x0b\x38\x00\x40\x03\x0b\x00\xc0\xc2\x01\x00\xe0\xb0\x00\x00\x38\x2c\x00\x00\x0f\x0e\x00\xb0\x07\x0f\x00\xce\x01\x1f\xf4\x70\x00\xff\x0b\x0c\x00\x14\x40\x03\x00\x00\xe0\x00\x00\x00\x1c\x00\x00\x80\x03\x00\x00\x74\x00\x00\x40\x0f\x00\x00\xf8\x00\x04\xd0\x0b\x00\xff\x7f\x00\x00\x55\x00\x00\x00" +

	// glyph 1745 for rune ':'
	"\x16\xf1\x01\x08\x0e\xe0\x07\xfe\xc7\xff\xe1\x3f\x90\x02\x00\x50\x15\xfc\x82\xff\xf1\x7f\xf4\x0f\x60\x00" +

	// glyph 1771 for rune ';'
	"\x16\xf1\x05\x05\x0d\x00\xbe\x00\xfe\x07\xf8\x2f\xd0\x7f\x00\x69\x00\x00\x50\x05\xf8\x0f\xf0\x1f\xc0\x3f\x40\x7f\x00\xfe\x00\xfc\x01\xf4\x02\xe0\x03\xc0\x0b\x00\x0f\x00" +

	// glyph 1813 for rune '<'
	"\x16\xed\xfe\x03\x13\x00\x00\x00\x40\x01\x00\x00\xd0\x0b\x00\x00\xe0\x07\x00\x00\xf8\x01\x00\x00\x7d\x00\x00\x40\x2f\x00\x00\xd0\x0b\x00\x00\xe0\x07\x00\x00\xf8\x01\x00\x00\xd0\x0b\x00\x00\x00\xf8\x01\x00\x00\x00\x7e\x00\x00\x00\x80\x1f\x00\x00\x00\xf4\x02\x00\x00\x00\xbd\x00\x00\x00\x80\x1f\x00\x00\x00\xe0\x02" +

	// glyph 1891 for rune '='
	"\x16\xf2\xf9\x02\x14\x54\x55\x55\x55\x05\xff\xff\xff\xff\x0b\x00\x00\x00\x00\x14\x54\x55\x55\x55\xc1\xff\xff\xff\xff\x02" +

	// glyph 1921 for rune '>'
	"\x16\xed\xff\x03\x13\x08\x00\x00\x00\xf0\x02\x00\x00\x00\xbd\x00\x00\x00\x80\x1f\x00\x00\x00\xe0\x07\x00\x00\x00\xbd\x00\x00\x00\x40\x2f\x00\x00\x00\xd0\x0b\x00\x00\x00\xf8\x01\x00\x00\xf0\x02\x00\x00\xf8\x01\x00\x00\x7e\x00\x00\x40\x2f\x00\x00\xd0\x0b\x00\x00\xf4\x02\x00\x00\xf8\x01\x00\x00\x7c\x00\x00\x00\x10\x00\x00\x00\x00" +

	// glyph 2003 for rune '?'
	"\x16\xeb\x01\x05\x12\x00\xa9\x06\x40\xfe\xff\x07\x7c\x00\xf4\xc1\x01\x00\x3c\x1c\x00\x00\x47\x00\x00\x70\x00\x00\x00\x07\x00\x00\x38\x00\x00\xd0\x02\x00\x80\x0b\x00\x80\x2f\x00\x00\x2f\x00\x00\xb0\x00\x10\x00\x04\x00\x00\x00\x00\x40\x01\xf0\x0b\x00\x80\xff\x01\x00\xf8\x0f\x00\x00\x29\x00\x00" +

	// glyph 2076 for rune '@'
	"\x16\xe9\x02\x04\x11\x00\x50\x01\x00\xd0\xff\x02\x80\x1b\xf4\x00\x2d\x00\x2c\xb0\x00\x40\x43\x07\x00\x70\x38\x00\x00\xcb\x02\x00\xb0\x1c\x00\xe8\xcb\x01\xf4\xba\x1c\xd0\x02\xcb\x00\x0b\xb0\x0c\x70\x00\x1b\x03\x2c\xc0\x72\x80\x07\x2c\x07\xe0\xff\x73\x00\xa0\x2e\x0b\x00\x00\xe0\x00\x00\x00\x1d\x00\x00\xc0\x03\x00\x00\xb4\x00\x80\x00\xbd\x95\x0f\x40\xff\x1f\x00" +

	// glyph 2166 for rune 'A'
	"\x16\xeb\x00\x00\x15\x00\x50\x55\x00\x00\x00\xd0\xff\x3f\x00\x00\x00\x00\x74\x0b\x00\x00\x00\x80\xe3\x00\x00\x00\x00\x1c\x1d\x00\x00\x00\xd0\xc0\x02\x00\x00\x00\x0a\x38\x00\x00\x00\x70\x00\x07\x00\x00\x40\x03\xf0\x00\x00\x00\x2c\x00\x0d\x00\x00\xc0\x00\xc0\x02\x00\x00\x0e\x00\x38\x00\x00\xb0\x00\x40\x07\x00\x40\xff\xff\xbf\x00\x00\x38\x00\x00\x0e\x00\xc0\x01\x00\xc0\x01\x00\x0d\x00\x00\x38\x00\xb0\x00\x00\x40\x03\x00\x07\x00\x00\xb0\x00\x79\x05\x00\x54\x5f\xf4\xff\x02\xc0\xff\x0f" +

	// glyph 2287 for rune 'B'
	"\x16\xeb\x00\x01\x14\x40\x55\x55\x00\x00\xf4\xff\xff\x2f\x00\x00\x0e\x00\xf5\x01\x00\x0e\x00\x80\x03\x00\x0e\x00\x00\x0b\x00\x0e\x00\x00\x0e\x01\x38\x00\x00\x2c\x00\x38\x00\x00\x0f\x00\x78\x55\xe5\x03\x00\xf8\xff\xff\x01\x00\x38\x00\x94\x1f\x00\x38\x00\x00\x78\x00\x38\x00\x00\xe0\x00\x38\x00\x00\xc0\x00\x38\x00\x00\xc0\x05\xe0\x00\x00\x80\x03\xe0\x00\x00\xe0\x01\xf5\x55\x55\x7e\x40\xff\xff\xff\x0b\x00" +

	// glyph 2388 for rune 'C'
	"\x16\xeb\x01\x02\x13\x00\x40\xaa\x01\x00\x40\xff\xff\xa2\x00\x6e\x00\xf8\x0b\xb8\x00\x00\xbd\xc0\x02\x00\x40\x0b\x0f\x00\x00\xb0\x70\x00\x00\x00\x45\x03\x00\x00\x00\x38\x00\x00\x00\x50\x15\x0d\x00\x00\x00\xc0\x02\x00\x00\x00\x38\x00\x00\x00\x01\x0f\x00\x00\x38\xc0\x03\x00\xe0\x01\xf0\x02\xe0\x07\x00\xf8\xff\x0b\x00\x00\x94\x05\x00" +

	// glyph 2471 for rune 'D'
	"\x16\xeb\x00\x01\x13\x40\x55\x05\x00\x00\xfd\xff\xff\x02\x00\x74\x00\xe4\x07\x00\x0d\x00\x80\x07\x40\x03\x00\x80\x03\xd0\x00\x00\xc0\x02\x34\x00\x00\xe0\x00\x0d\x00\x00\x70\x40\x03\x00\x00\x6c\x55\x00\x03\x00\x00\x1c\xc0\x00\x00\x80\x03\x34\x00\x00\xb4\x00\x0c\x00\x00\x0f\x40\x03\x00\xf4\x00\xe4\x56\xe5\x0f\x40\xff\xff\x6f\x00\x00" +

	// glyph 2554 for rune 'E'
	"\x16\xeb\x00\x01\x13\x40\x55\x55\x55\x05\xfd\xff\xff\xff\x03\xe0\x00\x00\xe0\x54\x00\x0e\x00\x00\x04\x80\x03\x60\x00\x00\xe0\x00\x2c\x00\x00\x78\x55\x0b\x00\x00\xfe\xff\x02\x00\x80\x03\xb0\x00\x10\x80\x03\x50\x00\x00\xe0\x00\x00\x40\x01\x38\x00\x00\xa0\x15\xd4\x57\x55\x55\x4b\xff\xff\xff\xff\x02" +

	// glyph 2628 for rune 'F'
	"\x16\xeb\x00\x01\x13\x40\x55\x55\x55\x05\xfd\xff\xff\xff\x0b\xe0\x00\x00\xc0\x02\x38\x00\x00\xa0\x05\xe0\x00\x00\x40\x01\x38\x00\x06\x00\x00\x0e\xc0\x02\x00\x80\x57\xb5\x00\x00\xe0\xff\x2f\x00\x00\x38\x00\x0b\x00\x01\x38\x00\x05\x00\x00\x0e\x00\x00\x40\x15\xd4\x57\x15\x00\x40\xff\xff\x0b\x00\x00" +

	// glyph 2702 for rune 'G'
	"\x16\xeb\x01\x02\x14\x00\x40\xaa\x01\x00\x00\xfd\xff\x8b\x02\xe0\x06\x40\xbf\x00\x1e\x00\x00\x2e\xd0\x02\x00\x00\x0b\x3c\x00\x00\x80\x02\x0b\x00\x00\x00\xd0\x00\x00\x00\x00\x38\x00\x00\x00\x40\x28\x00\x00\x00\x40\x38\x00\xe0\xff\x3f\x0e\x00\x50\xd5\x47\x03\x00\x00\xa0\xc0\x01\x00\x00\x28\xf0\x00\x00\x00\x0a\xb4\x00\x00\x80\x02\xb8\x00\x00\xb0\x00\xf4\x02\x80\x2f\x00\xe0\xff\xbf\x00\x00\x40\x69\x00\x00" +

	// glyph 2803 for rune 'H'
	"\x16\xeb\x00\x02\x14\x40\x15\x00\x54\x01\xfc\x3f\xc0\xff\x03\x38\x00\x00\x1d\x00\x0d\x00\x40\x03\x55\x01\x5d\x55\x55\x03\x40\xff\xff\xff\x00\xd0\x00\x00\x34\x50\x55\x51\x5e\x01\x94\x1b\xfc\xbf\x00\xff\x2f" +

	// glyph 2854 for rune 'I'
	"\x16\xeb\x00\x04\x12\x50\x55\x55\x01\xff\xff\xff\x07\x00\x2c\x00\x54\x55\x55\x55\x51\x55\x5b\x15\xf8\xff\xff\x1f" +

	// glyph 2882 for rune 'J'
	"\x16\xeb\x01\x03\x15\x00\x00\x55\x55\x05\x00\xf4\xff\xff\x0f\x00\x00\xc0\x02\x54\x55\xc5\x01\x00\xc0\x02\x14\x07\x00\x00\x07\x10\x0b\x00\x80\x03\x80\x0b\x00\x74\x00\x80\x1f\x80\x0b\x00\x00\xff\x7f\x00\x00\x00\x64\x01\x00\x00" +

	// glyph 2938 for rune 'K'
	"\x16\xeb\x00\x01\x15\x40\x55\x00\x40\x15\xd0\xff\x1f\xd0\xff\x02\xe0\x00\x00\xb8\x00\x80\x03\x00\xb4\x00\x00\x0e\x00\xb8\x00\x00\x38\x00\x78\x00\x00\xe0\x00\x78\x00\x00\x80\x03\x3c\x00\x00\x00\x0e\x3c\x00\x00\x00\x38\xbc\x00\x00\x00\xe0\xad\x1f\x00\x00\x80\x2f\xf4\x01\x00\x00\x2e\x00\x0f\x00\x00\x38\x00\xf0\x00\x00\xe0\x00\x40\x07\x00\x80\x03\x00\x3c\x00\x00\x0e\x00\xd0\x01\x00\x38\x00\x00\x0f\x00\xe0\x00\x00\x34\x00\xd4\x57\x00\xc0\x16\xf4\xff\x07\x00\xfe\x01" +

	// glyph 3054 for rune 'L'
	"\x16\xeb\x00\x02\x14\x50\x55\x05\x00\x00\xfe\xff\x1f\x00\x00\x40\x07\x00\x00\x00\xd0\x00\x00\x00\x55\x55\x01\xd0\x00\x00\x70\x55\x41\xe5\x56\x55\x75\xf8\xff\xff\xff\x1f" +

	// glyph 3096 for rune 'M'
	"\x16\xeb\x00\x00\x15\x40\x05\x00\x00\x50\x01\xff\x02\x00\xc0\xff\x80\x3b\x00\x00\xed\x01\x34\x07\x00\xf0\x0d\x40\xe3\x00\x00\xd7\x00\x34\x1c\x00\x38\x0d\x40\x83\x03\xc0\xd1\x00\x34\x74\x00\x0e\x0d\x40\x03\x0f\xb0\xd0\x00\x34\xd0\x40\x03\x0d\x40\x03\x2c\x2c\xd0\x00\x34\x40\xd3\x01\x0d\x40\x03\xb0\x0e\xd0\x00\x34\x00\x7e\x00\x0d\x40\x03\xc0\x03\xd0\x00\x34\x00\x00\x00\x0d\x15\xe4\x16\x00\x50\xb9\xd1\xff\x0b\x00\xff\x3f" +

	// glyph 3201 for rune 'N'
	"\x16\xeb\x00\x01\x14\x50\x01\x00\x50\x15\xfc\x0f\x00\xfc\xff\x40\x2f\x00\x00\x0e\x40\x7b\x00\x00\x0d\x40\xf3\x00\x00\x0d\x40\xc3\x01\x00\x0d\x40\x83\x03\x00\x0d\x40\x03\x0b\x00\x0d\x40\x03\x1d\x00\x0d\x40\x03\x3c\x00\x0d\x40\x03\x70\x00\x0d\x40\x03\xe0\x00\x0d\x40\x03\xc0\x02\x0d\x40\x03\x40\x07\x0d\x40\x03\x00\x0f\x0d\x40\x03\x00\x2c\x0d\x40\x03\x00\x38\x0d\x40\x03\x00\xb0\x0d\x40\x03\x00\xd0\x0f\x90\x5b\x00\xc0\x0f\xf4\xff\x02\x00\x0f" +

	// glyph 3311 for rune 'O'
	"\x16\xeb\x01\x02\x14\x00\x40\x6a\x00\x00\x00\xfd\xff\x02\x00\xe0\x02\xd0\x07\x00\x1e\x00\xc0\x03\xd0\x02\x00\xc0\x03\x3c\x00\x00\xc0\x42\x07\x00\x00\xe0\xe0\x00\x00\x00\x34\x2c\x00\x00\x00\x2c\x07\x00\x00\x00\x5b\xb1\x00\x00\x00\xb0\x38\x00\x00\x00\x0d\x1d\x00\x00\x80\x03\x0f\x00\x00\x74\x00\x0b\x00\x00\x0f\x80\x0b\x00\xf0\x00\x40\x1f\x80\x0f\x00\x40\xff\xbf\x00\x00\x00\x64\x01\x00" +

	// glyph 3407 for rune 'P'
	"\x16\xeb\x00\x01\x12\x40\x55\x15\x00\x40\xff\xff\xff\x01\x00\x0e\x00\xf9\x00\xe0\x00\x00\x3d\x00\x0e\x00\x40\x07\xe0\x00\x00\xb0\x00\x0e\x00\x00\x1e\x80\x03\x00\xc0\x02\x38\x00\x00\x0f\x80\x03\x00\x3d\x00\xf8\xff\xbf\x00\x80\xab\x6a\x00\x00\x38\x00\x00\x40\x55\x50\x5f\x55\x00\x40\xff\xff\x0b\x00" +

	// glyph 3481 for rune 'Q'
	"\x16\xeb\x04\x02\x14\x00\x40\x6a\x00\x00\x00\xfd\xff\x02\x00\xe0\x02\xd0\x07\x00\x1e\x00\xc0\x03\xd0\x02\x00\xc0\x03\x3c\x00\x00\xc0\x42\x07\x00\x00\xe0\xe0\x00\x00\x00\x74\x2c\x00\x00\x00\x2c\x07\x00\x00\x00\x5b\xb1\x00\x00\x00\xb0\x38\x00\x00\x00\x0d\x1c\x00\x00\x80\x03\x0e\x00\x00\x74\x00\x0b\x00\x00\x0f\x40\x0b\x00\xf0\x00\x80\x1f\x80\x0f\x00\x40\xff\xbf\x00\x00\x00\x7d\x01\x00\x00\xe0\x02\x00\x04\x00\xff\xff\xd6\x07\xe0\x5b\xe4\x7f\x00" +

	// glyph 3592 for rune 'R'
	"\x16\xeb\x00\x01\x15\x40\x55\x15\x00\x00\xd0\xff\xff\x7f\x00\x00\xe0\x00\x90\x0f\x00\x80\x03\x00\xf0\x00\x00\x0e\x00\x00\x07\x00\x38\x00\x00\x2c\x00\xe0\x00\x00\xe0\x00\x80\x03\x00\xc0\x02\x00\x0e\x00\x80\x07\x00\x38\x00\x90\x07\x00\xe0\xaa\xfa\x07\x00\x80\xff\xff\x01\x00\x00\x0e\x40\x0f\x00\x00\x38\x00\xf0\x00\x00\xe0\x00\x00\x0f\x00\x80\x03\x00\xb0\x00\x00\x0e\x00\x40\x07\x00\x38\x00\x00\x3c\x00\xe0\x00\x00\xc0\x02\xd4\x57\x00\x00\x5e\xf4\xff\x07\x00\xf0\x03" +

	// glyph 3708 for rune 'S'
	"\x16\xeb\x01\x03\x12\x00\x90\x2a\x00\x00\xfd\xff\x72\x40\x0b\x40\xbf\xc0\x02\x00\xbc\xe0\x00\x00\xb4\xf0\x00\x00\x70\xf0\x00\x00\x00\xe0\x00\x00\x00\xc0\x07\x00\x00\x00\xbf\x01\x00\x00\xe4\xbf\x01\x00\x00\xe4\x1f\x00\x00\x00\x7c\x00\x00\x00\xf0\x00\x00\x00\xd0\x24\x00\x00\xd0\x38\x00\x00\xe0\x78\x00\x00\xf0\xf8\x01\x00\x78\xb8\x1b\x40\x2f\x28\xfd\xff\x07\x00\x50\x16\x00" +

	// glyph 3801 for rune 'T'
	"\x16\xeb\x00\x02\x13\x50\x55\x55\x55\x45\xff\xff\xff\xff\x74\x00\x2c\x00\x4e\x03\xc0\x02\xd0\x81\x00\xb0\x00\x20\x00\x00\x0b\x00\x54\x55\x55\x01\x50\xf5\x55\x00\x40\xff\xff\x0f\x00" +

	// glyph 3846 for rune 'U'
	"\x16\xeb\x01\x01\x14\x40\x55\x00\x50\x15\xf4\xff\x02\xfc\xff\x40\x07\x00\x00\x0e\x40\x03\x00\x00\x0e\x55\x55\x15\xc0\x01\x00\x80\x03\xc0\x02\x00\xc0\x02\x80\x03\x00\xe0\x00\x00\x0f\x00\xb4\x00\x00\x7d\x00\x3e\x00\x00\xe0\xff\x0b\x00\x00\x00\x59\x00\x00" +

	// glyph 3909 for rune 'V'
	"\x16\xeb\x00\x00\x15\x50\x55\x00\x00\x55\x81\xff\x2f\x00\xfc\xff\x80\x07\x00\x00\xe0\x00\x70\x00\x00\x00\x0b\x00\x0e\x00\x00\x70\x00\xd0\x00\x00\x80\x03\x00\x2c\x00\x00\x1c\x00\x80\x03\x00\xd0\x00\x00\x70\x00\x00\x0b\x00\x00\x0f\x00\x70\x00\x00\xd0\x00\x80\x03\x00\x00\x2c\x00\x1c\x00\x00\x80\x03\xd0\x00\x00\x00\x74\x00\x0b\x00\x00\x00\x0b\x70\x00\x00\x00\xe0\x80\x03\x00\x00\x00\x1c\x1c\x00\x00\x00\x80\xd3\x00\x00\x00\x00\x74\x0b\x00\x00\x00\x00\x7f\x00\x00\x00\x00\xe0\x03\x00\x00" +

	// glyph 4030 for rune 'W'
	"\x16\xeb\x00\x01\x15\x50\x15\x00\x40\x55\xf0\xff\x03\xd0\xff\x0f\x0f\x00\x00\x40\x07\x38\x00\x00\x00\x0c\xe0\x00\x00\x00\x34\x40\x03\xd0\x03\xd0\x00\x0d\x80\x1f\x80\x03\x70\x00\xbb\x00\x0e\xc0\x01\x8c\x03\x2c\x00\x0b\x34\x0d\xb0\x00\x2c\xb0\x70\xc0\x01\xe0\xc0\xc1\x02\x07\x80\x43\x03\x0e\x0d\x00\x0d\x0e\x74\x34\x00\x34\x2c\xc0\xd2\x00\xd0\x31\x00\x8e\x03\x00\xd7\x00\x34\x0a\x00\x9c\x02\xc0\x2d\x00\xb0\x07\x00\x7b\x00\xc0\x0f\x00\xf8\x01\x00\x3e\x00\xd0\x03\x00" +

	// glyph 4146 for rune 'X'
	"\x16\xeb\x00\x01\x14\x40\x15\x00\x40\x15\xf0\xbf\x00\xf0\xbf\x40\x0b\x00\x40\x0f\x00\x0e\x00\x80\x03\x00\x3c\x00\xd0\x01\x00\xb0\x00\xb0\x00\x00\xd0\x02\x3c\x00\x00\x80\x07\x0e\x00\x00\x00\x4f\x07\x00\x00\x00\xfc\x02\x00\x00\x00\xf4\x00\x00\x00\x00\xfc\x01\x00\x00\x00\x8e\x03\x00\x00\x40\x07\x0f\x00\x00\xc0\x02\x2c\x00\x00\xf0\x00\x74\x00\x00\x38\x00\xe0\x01\x00\x1d\x00\xc0\x03\x00\x0b\x00\x00\x0f\xd0\x5b\x00\x50\x6f\xf4\xff\x00\xf4\xff" +

	// glyph 4256 for rune 'Y'
	"\x16\xeb\x00\x02\x14\x50\x05\x00\x50\x05\xff\x0b\x00\xff\x0b\x3d\x00\x00\x2d\x00\x1d\x00\x80\x03\x00\x0f\x00\x70\x00\x00\x0b\x00\x0f\x00\x40\x07\xe0\x00\x00\xc0\x03\x1c\x00\x00\xc0\xc2\x03\x00\x00\xd0\x39\x00\x00\x00\xf0\x07\x00\x00\x00\xf0\x00\x00\x00\x00\x2c\x00\x40\x55\x01\x50\xf5\x55\x00\x00\xfd\xff\x3f\x00" +

	// glyph 4334 for rune 'Z'
	"\x16\xeb\x00\x03\x12\x40\x55\x55\x05\xe0\xff\xff\x3f\xe0\x00\x00\x3c\xe0\x00\x00\x1d\xe0\x00\x00\x0b\xe0\x00\x80\x03\xe0\x00\xd0\x01\x00\x00\xb0\x00\x00\x00\x3c\x00\x00\x00\x0e\x00\x00\x40\x07\x00\x00\xc0\x02\x00\x00\xe0\x00\x00\x00\x74\x00\x40\x00\x2c\x00\xe0\x00\x0f\x00\xe0\x80\x03\x00\xe0\xd0\x01\x00\xe0\xb0\x00\x00\xe0\xb4\x55\x55\xf5\xf4\xff\xff\xff" +

	// glyph 4423 for rune '['
	"\x16\xea\x05\x0a\x10\xf8\x2f\xaf\xca\x02\x54\x55\x55\x55\x55\x15\xff\x4f\x55\x01" +

	// glyph 4443 for rune '\\'
	"\x16\xe8\x03\x04\x12\x1c\x00\x00\x00\x0e\x00\x00\x00\x07\x00\x00\x80\x03\x00\x00\xc0\x01\x00\x00\xe0\x00\x00\x00\x74\x00\x00\x00\x3c\x00\x00\x00\x1d\x00\x00\x00\x0f\x00\x00\x40\x03\x00\x00\xc0\x02\x00\x00\xd0\x00\x00\x00\xb0\x00\x00\x00\x38\x00\x00\x00\x2c\x00\x00\x00\x0e\x00\x00\x00\x0b\x00\x00\x80\x03\x00\x00\xc0\x01\x00\x00\xe0\x00\x00\x00\x70\x00\x00\x00\x3c\x00\x00\x00\x1d\x00\x00\x00\x0f\x00\x00\x40\x07\x00\x00\x80\x01" +

	// glyph 4550 for rune ']'
	"\x16\xea\x05\x05\x0c\xa0\x6a\xa0\xba\x00\xb0\x55\x55\x55\x55\x55\x45\xff\x0b\x55\x01" +

	// glyph 4571 for rune '^'
	"\x16\xea\xf3\x04\x12\x00\x80\x01\x00\x00\xf8\x00\x00\x40\xfb\x00\x00\xb0\xb0\x00\x00\x0f\x74\x00\xe0\x00\x78\x00\x1d\x00\x38\xd0\x02\x00\x3c\x2c\x00\x00\x1c" +

	// glyph 4610 for rune '_'
	"\x16\x02\x05\x00\x16\x54\x55\x55\x55\x55\x05\xff\xff\xff\xff\xff\x8b\xaa\xaa\xaa\xaa\xaa\x01" +

	// glyph 4633 for rune '`'
	"\x16\xe9\xee\x05\x0c\x74\x00\xf0\x01\x80\x07\x00\x2e\x00\x74" +

	// glyph 4648 for rune 'a'
	"\x16\xf0\x01\x02\x14\x00\x00\x55\x00\x00\x40\xff\xff\x01\x00\xf4\x01\xe0\x02\x00\x00\x00\xd0\x00\x00\x00\x00\x70\x00\x00\x00\x00\x2c\x00\x00\x40\x00\x0b\x00\xe0\xff\xef\x02\x80\x6f\x41\xb9\x00\xb8\x00\x00\x2c\x00\x0b\x00\x00\x0b\xd0\x00\x00\xc0\x02\x34\x00\x00\xb8\x00\x2c\x00\x80\x2f\x00\x2d\x00\x3e\x5f\x00\xfd\xff\xc1\x7f\x00\x94\x01\x00\x00" +

	// glyph 4734 for rune 'b'
	"\x16\xea\x01\x01\x14\xa8\x02\x00\x00\x00\xe8\x03\x00\x00\x00\x40\x03\x00\x00\x00\x15\xd0\x00\x54\x01\x00\xd0\x80\xff\x1f\x00\xd0\xf4\x01\xf9\x00\xd0\x2d\x00\xd0\x02\xd0\x0b\x00\x40\x07\xd0\x03\x00\x00\x0f\xd0\x02\x00\x00\x0d\xd0\x01\x00\x00\x1c\xd0\x00\x00\x00\x1c\xd0\x01\x00\x00\x1c\xd0\x02\x00\x00\x0d\xd0\x03\x00\x00\x0e\xd0\x0b\x00\x40\x07\xd0\x2d\x00\xd0\x02\xe5\xf4\x01\xb9\x00\xff\xc0\xff\x2f\x00\x00\x00\x64\x01\x00" +

	// glyph 4840 for rune 'c'
	"\x16\xf0\x01\x03\x13\x00\x40\x15\x00\x00\xe0\xff\x1f\x03\xf4\x01\xe4\x1f\xb0\x00\x00\x7d\xf0\x00\x00\xd0\xd1\x01\x00\x00\x87\x03\x00\x00\x00\x0b\x00\x00\x00\xc5\x03\x00\x00\x00\x0d\x00\x00\x00\xf0\x00\x00\x00\x42\x0b\x00\x40\x0f\xf4\x01\xe0\x0b\x40\xff\xff\x06\x00\x40\x1a\x00\x00" +

	// glyph 4910 for rune 'd'
	"\x16\xea\x01\x02\x15\x00\x00\x00\xa8\x02\x00\x00\x00\xe8\x03\x00\x00\x00\x40\x03\x15\x00\x50\x05\xd0\x00\x00\xfe\xbf\xd0\x00\xc0\x1b\xd0\xd3\x00\xf0\x01\x00\xde\x00\x38\x00\x00\xfc\x00\x2c\x00\x00\xf0\x00\x0d\x00\x00\xf0\x00\x0e\x00\x00\xe0\x00\x0e\x00\x00\xd0\x00\x0e\x00\x00\xe0\x00\x0d\x00\x00\xf0\x00\x2c\x00\x00\xf0\x00\x38\x00\x00\xfc\x00\xf0\x00\x00\xee\x00\xc0\x0b\xd0\xd3\x16\x00\xfe\xbf\xd0\x3f\x00\x90\x06\x00\x00" +

	// glyph 5016 for rune 'e'
	"\x16\xf0\x01\x02\x13\x00\x40\x55\x00\x00\x80\xff\xbf\x00\x00\x6f\x40\x3e\x00\x3c\x00\x00\x0f\xe0\x00\x00\xc0\x02\x07\x00\x00\x34\x38\x00\x00\x00\x87\xab\xaa\xaa\xba\xf8\xff\xff\xff\x8b\x03\x00\x00\x00\x74\x00\x00\x00\x00\x0b\x00\x00\x00\xd0\x01\x00\x00\x00\x78\x00\x00\x74\x00\x7e\x00\xf5\x03\x40\xff\xff\x06\x00\x00\x69\x01\x00" +

	// glyph 5098 for rune 'f'
	"\x16\xea\x00\x04\x14\x00\x00\xa9\x56\x00\x00\xff\xff\x07\x00\x1f\x00\x00\x00\x0d\x00\x00\x00\x38\x00\x00\x00\xa0\x00\x00\x00\x80\x03\x00\x00\xfe\xff\xff\x07\x50\x7d\x55\x05\x00\xa0\x00\x00\x50\x55\x55\x54\x7d\x55\x05\xf0\xff\xff\x3f\x00" +

	// glyph 5157 for rune 'g'
	"\x16\xf0\x07\x02\x14\x00\x40\x05\x00\x00\x00\xfe\x7f\xf0\x0f\xf4\x01\xb9\x6c\x01\x0f\x00\xb4\x07\xf0\x00\x00\xf4\x01\x1d\x00\x00\x7c\x80\x03\x00\x00\x1d\x84\x03\x00\x00\x1c\xe0\x00\x00\x40\x07\x74\x00\x00\xe0\x01\x3c\x00\x00\x7d\x00\x2d\x00\xc0\x1e\x00\x7d\x00\x2e\x07\x00\xfd\xff\xc1\x01\x00\x94\x06\x70\x00\x00\x00\x00\x1c\x00\x00\x00\x40\x03\x00\x00\x00\xe0\x00\x00\x00\x00\x2c\x00\x00\x00\xd0\x03\x00\xd0\xff\x2f\x00\x00\xa4\x6a\x01\x00" +

	// glyph 5267 for rune 'h'
	"\x16\xea\x00\x02\x14\xa8\x02\x00\x00\x00\xfa\x00\x00\x00\x00\x38\x00\x00\x00\x54\x80\x03\x54\x00\x00\xe0\xe0\xff\x02\x00\x38\x1f\xd0\x03\x00\xbe\x00\xc0\x02\x80\x0b\x00\xe0\x00\xe0\x00\x00\x34\x50\x55\x45\x7d\x01\x40\x6e\xf0\xff\x01\xf8\xbf" +

	// glyph 5327 for rune 'i'
	"\x16\xe9\x00\x03\x12\x00\x40\x01\x00\x00\x80\x07\x00\x05\x00\x10\x00\x00\x00\x00\x00\x50\xc0\xff\x0b\x00\x40\x55\x0b\x00\x00\x00\x0b\x00\x55\x55\x05\x55\xb5\x55\x85\xff\xff\xff\x0f" +

	// glyph 5372 for rune 'j'
	"\x16\xe9\x07\x05\x11\x00\x00\x14\x00\x00\xf0\x50\x00\x00\x14\x00\x00\x00\x50\xf4\xff\xff\x41\x55\x55\x07\x00\x00\x5c\x55\x55\x55\x01\x00\x40\x03\x00\x00\x0f\x00\x00\x1e\xf4\xff\x1f\x90\xaa\x06\x00" +

	// glyph 5421 for rune 'k'
	"\x16\xea\x00\x02\x14\xa4\x06\x00\x00\x00\xe9\x02\x00\x00\x00\xb0\x00\x00\x00\x54\x01\x2c\x80\xff\x0b\x00\x0b\x80\x5f\x00\xc0\x02\xf4\x00\x00\xb0\x80\x0b\x00\x00\x2c\x78\x00\x00\x00\xcb\x07\x00\x00\xc0\xff\x00\x00\x00\xf0\xb7\x00\x00\x00\x3c\xb4\x00\x00\x00\x0b\xf4\x00\x00\xc0\x02\xf4\x00\x00\xb0\x00\xf4\x00\x00\x2c\x00\xf4\x00\x50\x0b\x00\xf9\x05\xfe\x02\xc0\xff\x07" +

	// glyph 5513 for rune 'l'
	"\x16\xea\x00\x03\x12\x80\xaa\x06\x00\x80\xaa\x0b\x00\x00\x00\x0b\x00\x55\x55\x55\x55\x41\x55\x6d\x55\xe1\xff\xff\xff\x03" +

	// glyph 5543 for rune 'm'
	"\x16\xf0\x00\x00\x15\x00\x40\x05\x40\x05\x40\x3f\xfe\x03\xff\x02\x90\x7b\xf0\x3d\xb4\x00\xf4\x00\xbc\x00\x0e\x40\x07\xc0\x03\xd0\x00\x34\x00\x2c\x00\x0d\x55\x55\x90\x17\xc0\x07\xd0\x46\xff\x03\xfc\x02\xfd" +

	// glyph 5594 for rune 'n'
	"\x16\xf0\x00\x02\x14\x00\x00\x54\x00\x00\xfd\xe0\xff\x03\x00\x39\x1f\xd0\x03\x00\xbd\x00\xc0\x03\x40\x0b\x00\xd0\x00\xd0\x00\x00\x70\x50\x55\x45\x79\x01\x40\x6e\xf0\xff\x01\xf0\x7f" +

	// glyph 5639 for rune 'o'
	"\x16\xf0\x01\x02\x13\x00\x00\x55\x00\x00\x40\xff\xbf\x00\x00\x7e\x40\x3e\x00\xb8\x00\x00\x0f\xd0\x02\x00\xc0\x03\x0f\x00\x00\x74\x70\x00\x00\x00\x4b\x07\x00\x00\xe0\x34\x00\x00\x00\x4e\x07\x00\x00\xe0\x70\x00\x00\x00\x0f\x0f\x00\x00\x74\xd0\x02\x00\xc0\x03\x78\x00\x00\x0f\x00\x7e\x00\x7d\x00\x40\xff\xbf\x00\x00\x00\x59\x00\x00" +

	// glyph 5721 for rune 'p'
	"\x16\xf0\x07\x01\x14\x00\x00\x50\x05\x00\xfc\x43\xff\xbf\x00\x94\xd3\x07\xe4\x03\x40\x7b\x00\x40\x0f\x40\x2f\x00\x00\x2d\x40\x0f\x00\x00\x38\x40\x0b\x00\x00\x74\x40\x07\x00\x00\x70\x01\x1d\x00\x00\xd0\x01\x3d\x00\x00\xe0\x00\x7d\x00\x00\xb0\x00\xed\x01\x00\x3c\x00\x8d\x0b\x80\x0f\x00\x0d\xfd\xff\x02\x00\x0d\x50\x1a\x00\x00\x0d\x00\x00\x00\x54\xe1\xff\x1f\x00\x00\xa0\xaa\x1a\x00\x00\x00" +

	// glyph 5818 for rune 'q'
	"\x16\xf0\x07\x02\x15\x00\x40\x15\x00\x00\x00\xf8\xff\x42\xff\x40\x2f\x40\x5f\x5b\xc0\x03\x00\xb8\x03\xf0\x00\x00\xe0\x03\x74\x00\x00\xc0\x03\x38\x00\x00\x80\x03\xe1\x00\x00\x00\x0d\xe0\x00\x00\x00\x0e\xd0\x01\x00\x00\x0f\xc0\x03\x00\x80\x0f\x00\x0b\x00\xe0\x0e\x00\x7d\x00\x7d\x0d\x00\xf0\xff\x1f\x0d\x00\x00\x69\x01\x0d\x00\x00\x00\x00\x0d\x54\x01\x00\x00\xfe\xff\x02\x00\x00\xa9\xaa\x02" +

	// glyph 5915 for rune 'r'
	"\x16\xf0\x00\x03\x14\x00\x00\x00\x14\x00\xff\x03\xf8\x1f\x50\x3d\xf4\xd1\x07\x80\xe3\x02\x10\x00\xf8\x07\x00\x00\x80\x0f\x00\x00\x00\x38\x00\x00\x50\x55\x51\xf5\x55\x15\x00\xff\xff\xff\x03\x00" +

	// glyph 5963 for rune 's'
	"\x16\xf0\x01\x03\x12\x00\x50\x05\x00\x00\xfd\xff\x35\x80\x1f\x40\x3f\xc0\x02\x00\x3c\xd0\x00\x00\x38\xd0\x01\x00\x10\x80\x1b\x00\x00\x00\xfd\x5b\x00\x00\x40\xfa\x0b\x00\x00\x00\x3d\x00\x00\x00\xb0\x30\x00\x00\xf0\x74\x00\x00\xb0\xf4\x00\x00\x78\xf4\x0b\x40\x2f\x30\xfe\xff\x07\x00\x50\x16\x00" +

	// glyph 6036 for rune 't'
	"\x16\xec\x01\x01\x12\x00\x28\x00\x00\x00\xc0\x02\x00\x00\x15\xfd\xff\xff\x1f\x40\xf5\x55\x55\x00\x00\x0b\x00\x00\x54\x55\x01\xe0\x00\x00\x00\x00\x0d\x00\x00\x01\xc0\x07\x40\x2f\x00\xf0\xff\x7f\x00\x00\x90\x16\x00" +

	// glyph 6089 for rune 'u'
	"\x16\xf1\x01\x01\x14\xf4\x0f\x00\xfe\x07\x50\x0f\x00\x94\x07\x00\x0e\x00\x00\x07\x55\x55\x00\x0d\x00\xc0\x07\x00\x1c\x00\xf0\x07\x00\x78\x40\x2e\x1b\x00\xe0\xff\x07\x7f\x00\x40\x19\x00\x00" +

	// glyph 6136 for rune 'v'
	"\x16\xf1\x00\x01\x15\xf8\xff\x01\xf8\xff\x51\x6e\x01\x50\x7d\x01\xb0\x00\x00\xb0\x00\x80\x03\x00\xd0\x00\x00\x1c\x00\xc0\x02\x00\xe0\x00\x00\x07\x00\x00\x07\x00\x0e\x00\x00\x3c\x00\x1c\x00\x00\xd0\x01\x38\x00\x00\x00\x0b\xb0\x00\x00\x00\x34\xd0\x00\x00\x00\xc0\xc2\x02\x00\x00\x00\x4e\x03\x00\x00\x00\xb0\x0f\x00\x00\x00\x80\x1f\x00\x00" +

	// glyph 6220 for rune 'w'
	"\x16\xf1\x00\x01\x15\xf8\x3f\x00\xd0\xff\x91\x5b\x00\x00\xf5\x01\x1c\x00\x00\xc0\x02\xb0\x00\x00\x00\x07\x80\x03\xf4\x00\x0d\x00\x0d\xe0\x03\x38\x00\x70\xc0\x2e\xe0\x00\xc0\x41\xe3\xc0\x02\x00\x0a\x0f\x07\x07\x00\x34\x1c\x2c\x0d\x00\xd0\x34\xe0\x38\x00\x00\xf7\x00\xb7\x00\x00\xec\x01\xec\x01\x00\xe0\x03\xd0\x03\x00\x40\x0f\x00\x0f\x00" +

	// glyph 6304 for rune 'x'
	"\x16\xf1\x00\x02\x14\xf0\x7f\x00\xff\x0f\xf4\x06\x40\xbd\x01\xf0\x00\x40\x0b\x00\xf0\x00\xb4\x00\x00\xf0\x40\x0b\x00\x00\xf0\xb8\x00\x00\x00\xf0\x07\x00\x00\x00\xf8\x01\x00\x00\x80\xf7\x01\x00\x00\x78\xf0\x01\x00\x80\x07\xe0\x01\x00\x78\x00\xe0\x01\xc0\x07\x00\xe0\x01\xbd\x05\x40\xf5\xc2\xff\x07\xe0\xff\x02" +

	// glyph 6381 for rune 'y'
	"\x16\xf1\x07\x02\x14\xfc\x2f\x00\xfc\x2f\xb9\x05\x00\xe5\x02\x2c\x00\x00\x38\x00\x0e\x00\x00\x07\x00\x0b\x00\xe0\x00\x80\x03\x00\x1c\x00\xc0\x02\x80\x03\x00\xe0\x00\x70\x00\x00\xb0\x00\x0e\x00\x00\x38\xc0\x01\x00\x00\x2c\x38\x00\x00\x00\x0e\x07\x00\x00\x00\xeb\x00\x00\x00\x80\x1f\x00\x00\x00\xc0\x03\x00\x00\x00\x70\x00\x00\x00\x00\x0e\x00\x00\x00\xc0\x01\x00\x00\x00\x38\x00\x00\x00\x00\x07\x00\x00\xf8\xff\x2f\x00\x00\xa9\xaa\x0a\x00\x00" +

	// glyph 6491 for rune 'z'
	"\x16\xf1\x00\x04\x12\xf8\xff\xff\x0f\x5e\x55\xe5\x83\x03\x00\x3c\xd0\x00\xc0\x03\x00\x00\x38\x00\x00\x80\x07\x00\x00\x74\x00\x00\x40\x0b\x00\x00\xb4\x00\x00\x00\x0b\x00\x00\xf0\x00\x40\x00\x0f\x00\xb0\xf0\x00\x00\x2c\x6e\x55\x55\xcb\xff\xff\xff\x02" +

	// glyph 6553 for rune '{'
	"\x16\xea\x05\x06\x0f\x00\x40\x06\x00\x7f\x00\x78\x00\xc0\x02\x55\x05\xc0\x01\x00\x0e\x40\x7e\x00\xfc\x01\x00\xb8\x00\x00\x0e\x00\xc0\x01\x00\x2c\x50\x15\x00\x0e\x00\xd0\x02\x00\xf8\x02\x00\x04" +

	// glyph 6601 for rune '|'
	"\x16\xea\x05\x0a\x0c\x08\x5b\x55\x55\x55\x55\x55\x71\x04" +

	// glyph 6615 for rune '}'
	"\x16\xea\x05\x07\x10\x68\x00\x80\x2f\x00\x80\x03\x00\x70\x00\x55\x01\x2c\x40\x00\x1d\x00\x80\x5f\x00\xe0\x0f\xc0\x07\x00\x0e\x00\xb0\x00\x00\x07\x50\x55\xc0\x03\xc0\x1f\x00\x14\x00\x00" +

	// glyph 6661 for rune '~'
	"\x16\xf3\xf9\x03\x12\x00\x15\x00\x00\x80\xff\x00\x80\xe0\xc1\x07\xf0\x74\x00\x3e\x3d\x14\x00\xf4\x0f\x00\x00\x40\x00" +

	"")
