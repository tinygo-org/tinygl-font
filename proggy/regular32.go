// File generated using:
// 	go run ./generate -font=proggy/ProggyVector-Regular.ttf -size=32 -dpi=72 -package=proggy -weight=Regular

package proggy

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      6306
// - glyph metadata:  475
// - glyph mask data: 5630

var Regular32 = font.Make("" +
	"\x00" + // version: 0
	"\x20" + // size:   32
	"\x2d" + // height: 45
	"\x1e" + // ascent: 30

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xe6\x00" + // "\"" at index 230
	"\xf2\x00" + // "#" at index 242
	"\x65\x01" + // "$" at index 357
	"\xc1\x01" + // "%" at index 449
	"\x34\x02" + // "&" at index 564
	"\xa7\x02" + // "'" at index 679
	"\xaf\x02" + // "(" at index 687
	"\xe0\x02" + // ")" at index 736
	"\x18\x03" + // "*" at index 792
	"\x54\x03" + // "+" at index 852
	"\x77\x03" + // "," at index 887
	"\x8e\x03" + // "-" at index 910
	"\x9c\x03" + // "." at index 924
	"\xa9\x03" + // "/" at index 937
	"\x12\x04" + // "0" at index 1042
	"\x77\x04" + // "1" at index 1143
	"\xa0\x04" + // "2" at index 1184
	"\xfa\x04" + // "3" at index 1274
	"\x5c\x05" + // "4" at index 1372
	"\xaf\x05" + // "5" at index 1455
	"\xfa\x05" + // "6" at index 1530
	"\x58\x06" + // "7" at index 1624
	"\xb5\x06" + // "8" at index 1717
	"\x13\x07" + // "9" at index 1811
	"\x74\x07" + // ":" at index 1908
	"\x88\x07" + // ";" at index 1928
	"\xa4\x07" + // "<" at index 1956
	"\xe4\x07" + // "=" at index 2020
	"\x01\x08" + // ">" at index 2049
	"\x45\x08" + // "?" at index 2117
	"\x93\x08" + // "@" at index 2195
	"\x17\x09" + // "A" at index 2327
	"\x8a\x09" + // "B" at index 2442
	"\xe5\x09" + // "C" at index 2533
	"\x3b\x0a" + // "D" at index 2619
	"\x8a\x0a" + // "E" at index 2698
	"\xb4\x0a" + // "F" at index 2740
	"\xd9\x0a" + // "G" at index 2777
	"\x38\x0b" + // "H" at index 2872
	"\x52\x0b" + // "I" at index 2898
	"\x6d\x0b" + // "J" at index 2925
	"\x99\x0b" + // "K" at index 2969
	"\x01\x0c" + // "L" at index 3073
	"\x1b\x0c" + // "M" at index 3099
	"\x6a\x0c" + // "N" at index 3178
	"\xc7\x0c" + // "O" at index 3271
	"\x16\x0d" + // "P" at index 3350
	"\x51\x0d" + // "Q" at index 3409
	"\xad\x0d" + // "R" at index 3501
	"\x10\x0e" + // "S" at index 3600
	"\x66\x0e" + // "T" at index 3686
	"\x84\x0e" + // "U" at index 3716
	"\xaf\x0e" + // "V" at index 3759
	"\x17\x0f" + // "W" at index 3863
	"\x8a\x0f" + // "X" at index 3978
	"\xfd\x0f" + // "Y" at index 4093
	"\x42\x10" + // "Z" at index 4162
	"\xa2\x10" + // "[" at index 4258
	"\xbc\x10" + // "\\" at index 4284
	"\x29\x11" + // "]" at index 4393
	"\x47\x11" + // "^" at index 4423
	"\x70\x11" + // "_" at index 4464
	"\x86\x11" + // "`" at index 4486
	"\x9b\x11" + // "a" at index 4507
	"\xe8\x11" + // "b" at index 4584
	"\x37\x12" + // "c" at index 4663
	"\x75\x12" + // "d" at index 4725
	"\xc0\x12" + // "e" at index 4800
	"\x12\x13" + // "f" at index 4882
	"\x3f\x13" + // "g" at index 4927
	"\x9d\x13" + // "h" at index 5021
	"\xc8\x13" + // "i" at index 5064
	"\xde\x13" + // "j" at index 5086
	"\x0f\x14" + // "k" at index 5135
	"\x61\x14" + // "l" at index 5217
	"\x72\x14" + // "m" at index 5234
	"\x99\x14" + // "n" at index 5273
	"\xbb\x14" + // "o" at index 5307
	"\x08\x15" + // "p" at index 5384
	"\x57\x15" + // "q" at index 5463
	"\x9f\x15" + // "r" at index 5535
	"\xbc\x15" + // "s" at index 5564
	"\xfa\x15" + // "t" at index 5626
	"\x1e\x16" + // "u" at index 5662
	"\x4a\x16" + // "v" at index 5706
	"\x98\x16" + // "w" at index 5784
	"\xf2\x16" + // "x" at index 5874
	"\x44\x17" + // "y" at index 5956
	"\x84\x17" + // "z" at index 6020
	"\xc6\x17" + // "{" at index 6086
	"\x17\x18" + // "|" at index 6167
	"\x26\x18" + // "}" at index 6182
	"\x79\x18" + // "~" at index 6265

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x13\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x13\xe9\x00\x07\x0c\xf0\x57\x05\x3f\x55\x15\xf8\x40\x0a\x00\x14\x10\xc0\x1f\xfe\xd3\x3f\xf8\x00" +

	// glyph 230 for rune '"'
	"\x13\xea\xf2\x05\x0e\xf8\x80\x8f\x0f\xfc\x55\x05" +

	// glyph 242 for rune '#'
	"\x13\xea\x00\x00\x13\x00\x00\x68\x40\x0a\x00\x00\x7c\xc0\x0b\x00\x00\x3d\xc0\x07\x00\x00\x3e\xd0\x03\x00\x00\x2f\xe0\x03\x00\x00\x1f\xf0\x02\x80\xea\xaf\xfa\xab\xc0\xff\xff\xff\xff\x80\xea\xff\xff\xaa\x00\xd0\x07\xbc\x00\x00\xe0\x03\x7c\x00\x00\xf0\x02\x3d\x00\x00\xf0\x01\x3e\x00\x54\xf9\x55\x7f\x05\xfc\xff\xff\xff\x1f\xf8\xff\xff\xff\x1f\x00\x7c\xc0\x0b\x00\x00\x3d\xc0\x07\x00\x00\x3e\xd0\x03\x00\x00\x2f\xe0\x03\x00\x00\x1f\xf0\x02\x00\x40\x0f\xf0\x01\x00" +

	// glyph 357 for rune '$'
	"\x13\xe9\x01\x03\x11\x00\x80\x02\x00\x00\xf4\x00\x00\x40\xbf\x05\x00\xfe\xff\x2f\xf0\xef\xeb\x0f\xbe\xf4\x40\xc1\x0f\x3d\x00\xf0\x42\x0f\x00\xfc\xd0\x03\x00\xbe\xf4\x00\x00\xff\x7f\x00\x40\xff\xff\x02\x00\xf9\xff\x07\x00\xf4\xf9\x07\x00\x3d\xf0\x03\x40\x0f\xf4\x01\x40\x0f\xf8\x1c\xd0\x03\x2f\xbf\xf9\xfa\x83\xff\xff\x2f\x00\xe5\xbf\x00\x00\xd0\x03\x00\x00\xa0\x00\x00" +

	// glyph 449 for rune '%'
	"\x13\xea\x00\x00\x13\x00\x14\x00\x40\x05\x80\xff\x01\xc0\x0b\xf0\xff\x07\xf0\x03\xf0\x80\x0f\xf4\x02\xf4\x00\x0f\xfc\x00\xf0\x40\x0f\x7e\x00\xf0\xe7\x0b\x3f\x00\xc0\xff\xc3\x0f\x00\x00\x69\xd0\x0b\x00\x00\x00\xf0\x03\x00\x00\x00\xf8\x01\x00\x00\x00\xfc\x00\x00\x00\x00\x3e\x00\x00\x00\x40\x2f\x90\x01\x00\xc0\x0f\xfe\x1f\x00\xe0\x47\xaf\x3f\x00\xf0\x83\x0b\xbc\x00\xf8\xc1\x07\xb8\x00\xbd\x80\x0b\x7c\x00\x3f\x40\xaf\x3f\x40\x2f\x00\xfd\x0f\x40\x05\x00\x50\x01" +

	// glyph 564 for rune '&'
	"\x13\xe9\x01\x01\x13\x00\x50\x05\x00\x00\x80\xff\x1f\x00\x00\xf8\xff\x0f\x00\x00\x3f\xd0\x07\x00\xd0\x07\xf0\x01\x40\xd0\x0f\xf8\x00\x00\xf0\x8b\x2f\x00\x00\xf0\xff\x02\x00\x00\xf0\x1f\x00\x00\x00\xff\x07\x00\x00\xf0\xfb\x03\x00\x00\x7e\xf4\x03\xf0\xd2\x07\xf8\x02\xbc\xf8\x00\xfc\x01\x2f\x2f\x00\xfc\xc0\xc7\x0b\x00\xfd\xf0\xf0\x03\x00\xbe\x3f\xf8\x01\x00\xff\x07\xfd\x00\x00\xbf\x00\xfe\x01\xf4\x3f\x00\xfe\xff\xbf\x3f\x00\xfd\xff\x82\x2f\x00\x40\x01\x00\x00" +

	// glyph 679 for rune '\''
	"\x13\xea\xf2\x08\x0b\xf8\x55\x15" +

	// glyph 687 for rune '('
	"\x13\xe8\x03\x05\x0c\x00\xf8\x00\x7c\x00\x3e\x40\x1f\xc0\x0f\xc0\x0b\xe0\x07\xf0\x03\xd1\x0b\xe0\x0b\xe0\x07\x54\xd1\x0b\xc0\x0f\x04\x3e\x00\xbd\x00\xfc\x00\xf4\x01\xf0\x03\xd0\x07\x80\x0f\x00\x0a" +

	// glyph 736 for rune ')'
	"\x13\xe7\x03\x07\x0e\x14\x00\x7c\x00\xf8\x00\xf0\x02\xe0\x03\xc0\x0b\xc0\x0f\x40\x1f\x00\x2f\x00\x3f\x01\xf8\x05\xe0\x0b\xd0\x0b\xe0\x17\x80\x0f\xc0\x0f\xc0\x0b\xd0\x07\xf0\x03\xf0\x02\xf8\x00\xbc\x00\x3e\x00\x1f\x00\x05\x00" +

	// glyph 792 for rune '*'
	"\x13\xed\xfd\x01\x12\x00\x00\x3f\x00\x50\xe0\x01\x3f\xd0\x43\xff\xf0\x83\xbf\xd0\xbf\xbf\xff\x02\xe0\xff\xff\x02\x00\xe0\xff\x03\x00\x40\xff\x7f\x00\x40\xff\xff\x7f\x00\xfe\xf7\xf7\x3f\xf0\x07\x3f\xf4\x07\x0a\xf0\x03\x28\x00\x00\x3f\x00\x50" +

	// glyph 852 for rune '+'
	"\x13\xed\xfd\x01\x12\x00\x00\x29\x00\x00\x00\xf0\x03\x00\x55\xe1\xff\xff\xff\x7f\x54\x55\x7f\x55\x05\x00\xf0\x03\x00\x55\x00\x00\x3e\x00\x00" +

	// glyph 887 for rune ','
	"\x13\xfb\x05\x06\x0c\x00\x01\xf0\x07\xfe\x43\xff\x80\x3f\xc0\x07\xf8\x40\x1f\xf4\x02\x2c\x00" +

	// glyph 910 for rune '-'
	"\x13\xf4\xf7\x02\x11\xf4\xff\xff\xbf\x41\x55\x55\x55\x01" +

	// glyph 924 for rune '.'
	"\x13\xfb\x00\x07\x0c\x40\x00\x7f\xf8\x4f\xff\xe0\x03" +

	// glyph 937 for rune '/'
	"\x13\xe9\x02\x02\x11\x00\x00\x00\xfc\x00\x00\x00\x7e\x00\x00\x00\x3f\x00\x00\x80\x1f\x00\x00\xc0\x0f\x00\x00\xe0\x07\x00\x00\xf0\x03\x00\x00\xf8\x01\x00\x00\xfc\x00\x00\x00\x7e\x00\x00\x00\x3f\x00\x00\x80\x1f\x00\x00\xc0\x0f\x00\x00\xe0\x07\x00\x00\xf0\x03\x00\x00\xf8\x01\x00\x00\xfc\x00\x00\x00\x7d\x00\x00\x00\x3f\x00\x00\x40\x1f\x00\x00\xc0\x0f\x00\x00\xd0\x07\x00\x00\xf0\x03\x00\x00\xf4\x01\x00\x00\xa4\x00\x00\x00" +

	// glyph 1042 for rune '0'
	"\x13\xe9\x01\x02\x11\x00\x00\x05\x00\x00\xf8\xff\x00\x00\xff\xff\x07\xc0\x7f\xe5\x0f\xe0\x0f\x80\x3f\xf0\x03\x00\x3f\xf0\x02\x80\xbf\xf4\x01\xe0\xff\xf8\x01\xf4\xff\xfc\x00\xfc\xfe\xfc\x00\xbf\xf8\xfc\x80\x3f\xf8\xfc\xe0\x0f\xf8\xfc\xf4\x07\xf8\xfc\xfd\x01\xfc\xf8\xbf\x00\xfc\xf8\x3f\x00\xfc\xf4\x0f\x00\xbd\xf0\x07\x00\x7f\xe0\x0f\x40\x3f\xc0\x7f\xe5\x1f\x00\xff\xff\x07\x00\xf8\xff\x01\x00\x40\x05\x00" +

	// glyph 1143 for rune '1'
	"\x13\xea\x00\x04\x10\x00\xfc\x01\x00\xfd\x07\x00\xfe\x1f\x00\xff\x7f\x00\xfc\xf5\x02\xf0\xd0\x0b\x80\x40\x2f\x00\x00\xbd\x00\x55\x55\x15\xa9\xfe\xaa\xfc\xff\xff\x07" +

	// glyph 1184 for rune '2'
	"\x13\xe9\x00\x02\x11\x00\x40\x01\x00\x40\xfe\xbf\x00\xf0\xff\xff\x07\xf4\x1b\xf4\x1f\x74\x00\xc0\x3f\x00\x00\x00\x3f\x00\x00\x00\x7f\x01\x00\x00\xfc\x00\x00\x00\xfd\x00\x00\x00\x7f\x00\x00\xc0\x2f\x00\x00\xe0\x0b\x00\x00\xf8\x02\x00\x00\xfe\x00\x00\x80\x3f\x00\x00\xe0\x0f\x00\x00\xf8\x02\x00\x00\xbe\x00\x00\x80\x2f\x00\x00\xd0\xaf\xaa\xaa\xe2\xff\xff\xff\x07" +

	// glyph 1274 for rune '3'
	"\x13\xe9\x01\x02\x11\x00\x40\x01\x00\x90\xff\xbf\x01\xf0\xff\xff\x0b\xf0\x16\xe4\x1f\x00\x00\x40\x3f\x00\x00\x00\x7f\x00\x00\x00\x7e\x00\x00\x00\x7f\x00\x00\x40\x3f\x00\x00\xe0\x1f\x00\xf8\xff\x07\x00\xf8\xff\x01\x00\xa4\xfa\x0f\x00\x00\x80\x3f\x00\x00\x00\xbe\x00\x00\x00\xfd\x00\x00\x00\xfc\x01\x00\x00\xf4\x13\x00\x00\xfc\xe2\x17\x90\xff\xe0\xff\xff\x3f\x90\xff\xff\x0b\x00\x50\x15\x00\x00" +

	// glyph 1372 for rune '4'
	"\x13\xea\x00\x01\x12\x00\x00\xf0\x1f\x00\x00\xc0\xff\x01\x00\x00\xfe\x1f\x00\x00\xf0\xfb\x01\x00\xc0\x2f\x1f\x00\x00\xfd\xf0\x01\x00\xf0\x07\x1f\x00\x80\x2f\xf0\x01\x00\xfd\x00\x1f\x00\xf0\x07\xf0\x01\x80\x2f\x00\x1f\x00\xfd\x00\xf0\x01\xf0\x07\x00\x1f\x80\x3f\x00\xf0\x01\xf8\xab\xaa\xbf\x8a\xff\xff\xff\xff\x01\x00\x00\x7e\x40\x15" +

	// glyph 1455 for rune '5'
	"\x13\xea\x01\x02\x11\xe0\xff\xff\x1f\xf0\xff\xff\x1f\xf0\xab\xaa\x0a\xf0\x03\x00\x00\x15\xfc\x55\x01\x00\xfc\xff\x7f\x00\xfc\xff\xff\x02\x68\x00\xfd\x07\x00\x00\xe0\x0f\x00\x00\xc0\x1f\x00\x00\x80\x6f\x05\x00\x00\xf0\x47\x00\x00\xfc\xc3\x1b\x50\xff\xc1\xff\xff\x7f\x40\xff\xff\x0b\x00\x40\x15\x00\x00" +

	// glyph 1530 for rune '6'
	"\x13\xea\x01\x02\x11\x00\x40\xff\x0b\x00\xf8\xff\x1f\x00\xfe\x5b\x1e\x80\x3f\x00\x00\xd0\x0f\x00\x00\xf0\x07\x00\x00\xf4\x03\x00\x00\xf4\x41\x15\x00\xf8\xf5\xff\x03\xfc\xfd\xff\x1f\xfc\x2f\x80\x3f\xfc\x07\x00\xbf\xfc\x03\x00\xfd\xfc\x02\x00\xfc\xf8\x02\x00\xfc\xd1\x0b\x00\xf0\xc3\x0f\x00\xf4\x83\x2f\x00\xfc\x02\xbf\x40\xff\x00\xfd\xff\x3f\x00\xe0\xff\x0b\x00\x00\x14\x00\x00" +

	// glyph 1624 for rune '7'
	"\x13\xea\x00\x02\x11\xf8\xff\xff\xff\xfc\xff\xff\xff\xa4\xaa\xaa\x7f\x00\x00\x00\x3f\x00\x00\x40\x2f\x00\x00\xc0\x0f\x00\x00\xd0\x0f\x00\x00\xe0\x07\x00\x00\xf0\x03\x00\x00\xf8\x02\x00\x00\xfc\x00\x00\x00\xfd\x00\x00\x00\x7f\x00\x00\x00\x3f\x00\x00\x80\x2f\x00\x00\xc0\x0f\x00\x00\xd0\x0b\x00\x00\xf0\x07\x00\x00\xf4\x03\x00\x00\xf8\x02\x00\x00\xfc\x00\x00\x00\xbe\x00\x00" +

	// glyph 1717 for rune '8'
	"\x13\xe9\x01\x02\x11\x00\x40\x01\x00\x00\xfd\xff\x01\x80\xff\xff\x0f\xe0\x1f\xd0\x3f\xf0\x07\x00\x7f\xf0\x03\x00\xbe\xf4\x03\x00\xbe\xc1\x0f\x00\xfc\x41\x3f\x00\xfe\x00\xfd\xff\x2f\x00\xf4\xff\x0b\x00\xff\xeb\x7f\xc0\x2f\x00\xfd\xd1\x0f\x00\xf8\xe3\x07\x00\xf0\xf3\x07\x00\xf0\xc7\x2f\x00\xc0\x8f\x3f\x00\xf0\x0f\xff\x01\xfc\x07\xfc\xff\xff\x01\xe0\xff\x2f\x00\x00\x54\x00\x00" +

	// glyph 1811 for rune '9'
	"\x13\xe9\x00\x02\x11\x00\x40\x01\x00\x00\xfe\xbf\x00\xc0\xff\xff\x07\xf0\x1f\xe0\x0f\xf4\x03\x40\x3f\xf8\x02\x00\x3f\xfc\x01\x00\x7e\xfc\x00\x00\xbd\xfc\x00\x00\xfd\xfc\x00\x00\xfe\xfc\x01\x00\xff\xf4\x03\x40\xff\xf0\x0b\xd0\xff\xd0\xff\xfe\xfe\x00\xff\xbf\xfc\x00\x90\x06\xbd\x00\x00\x00\x7e\x00\x00\x00\x3f\x00\x00\x80\x2f\x00\x00\xd0\x0f\x80\x01\xf9\x07\xd0\xff\xff\x01\xc0\xff\x2f\x00" +

	// glyph 1908 for rune ':'
	"\x13\xee\xfd\x07\x0c\xe0\x47\xff\xc1\x2f\x10\x00\x40\x15\xb4\xd0\x7f\xf0\x0b\x14" +

	// glyph 1928 for rune ';'
	"\x13\xee\x02\x07\x0c\xe0\x47\xff\xc1\x2f\x10\x00\x40\x15\xb4\xd0\x7f\xf0\x0f\xfc\xc0\x07\x3e\xf8\xc0\x03\x08\x00" +

	// glyph 1956 for rune '<'
	"\x13\xee\xfc\x01\x12\x00\x00\x00\x40\x0f\x00\x00\x90\xff\x00\x00\xe4\xff\x07\x00\xf8\xff\x02\x00\xfd\xbf\x00\x00\xfe\x6f\x00\x00\xf8\x1b\x00\x00\x10\xf9\xbf\x00\x00\x00\xf4\xff\x02\x00\x00\xe0\xff\x07\x00\x00\x90\xff\x1f\x00\x00\x40\xff\x03\x00\x00\x00\x3e" +

	// glyph 2020 for rune '='
	"\x13\xf0\xfa\x01\x12\xa4\xaa\xaa\xaa\x8a\xff\xff\xff\xff\x01\x00\x00\x00\x40\x85\xff\xff\xff\xff\x91\xaa\xaa\xaa\x2a" +

	// glyph 2049 for rune '>'
	"\x13\xee\xfc\x01\x12\x78\x00\x00\x00\x80\xbf\x01\x00\x00\xf4\xff\x02\x00\x00\xd0\xff\x07\x00\x00\x80\xff\x1f\x00\x00\x40\xfe\x6f\x00\x00\x00\xf9\x0f\x00\x00\x80\xff\x00\x00\xd0\xff\x07\x00\xe4\xff\x02\x00\xf9\xbf\x01\x00\xfe\x7f\x00\x00\xf8\x2f\x00\x00\x80\x1b\x00\x00\x00" +

	// glyph 2117 for rune '?'
	"\x13\xe9\x00\x03\x11\x00\xe8\x6f\x00\xe0\xff\xff\x01\xff\xaa\xff\xf0\x0f\x00\xbd\xbc\x00\x00\x3f\x08\x00\xc0\x0f\x00\x00\xf4\x02\x00\x40\x3f\x00\x00\xf8\x07\x00\xc0\x7f\x00\x00\xfc\x03\x00\x80\x2f\x00\x00\xf0\x03\x40\x01\x00\x00\x00\x05\x00\x04\x00\x00\xc0\x1f\x00\x00\xf8\x0f\x00\x00\xfd\x03\x00\x00\x3e\x00\x00" +

	// glyph 2195 for rune '@'
	"\x13\xe9\x01\x00\x15\x00\x00\xfd\xaf\x01\x00\x00\xfe\xff\xff\x01\x00\xfc\x5b\xe5\x7f\x00\xf0\x0b\x00\xd0\x1f\xc0\x2f\x00\x00\xf0\x03\xbe\x00\xb9\x01\x3e\xf0\x03\xf8\xff\xe0\x43\x2f\xf0\xff\xbf\x3f\xf8\x40\x3f\xf0\xff\xc0\x0f\xfc\x00\xfc\x02\xfc\xc0\x0b\x80\x0f\xc0\x0b\xbc\x00\xf8\x00\xf1\x03\x3f\x00\x3f\x00\x3f\xf0\x03\xf4\x03\xe0\x07\xfd\xd1\x3f\x00\xfd\x80\xff\xff\x03\xc0\x1f\xd0\xff\x3d\x00\xf4\x03\x90\x01\x00\x00\xfe\x01\x00\x00\x0e\x80\xbf\x00\x00\xfd\x03\xe0\xff\xaa\xff\x0f\x00\xf4\xff\xff\x1f\x00\x00\x90\xff\x1b\x00" +

	// glyph 2327 for rune 'A'
	"\x13\xea\x00\x00\x13\x00\x00\xfd\x02\x00\x00\x00\xfe\x03\x00\x00\x00\xff\x07\x00\x00\x40\xef\x0b\x00\x00\x80\xcf\x0f\x00\x00\xc0\x8f\x1f\x00\x00\xd0\x0b\x2f\x00\x00\xe0\x03\x3f\x00\x00\xf0\x03\x7e\x00\x00\xf4\x02\xbd\x00\x00\xf8\x01\xfc\x00\x00\xfc\x00\xfc\x01\x00\xbd\x00\xf4\x02\x00\x7e\x00\xf0\x03\x00\xff\xff\xff\x03\x40\xff\xff\xff\x0b\x80\x6f\x55\xd5\x0f\xc0\x0f\x00\xc0\x0f\xd0\x0f\x00\x80\x2f\xe0\x0b\x00\x40\x3f\xf0\x07\x00\x00\x3f\xf4\x03\x00\x00\xbe" +

	// glyph 2442 for rune 'B'
	"\x13\xea\x00\x02\x12\xf4\xff\x6b\x00\xe0\xff\xff\x1f\x80\xaf\xaa\xff\x00\x7e\x00\xd0\x0f\xf8\x01\x00\x3e\xe0\x07\x00\xf4\x11\x7e\x00\xc0\x0f\xf8\x01\x80\x2f\xe0\xff\xff\x2f\x80\xff\xff\x2f\x00\xfe\xaa\xfa\x0b\xf8\x01\x00\xbd\xe0\x07\x00\xd0\x87\x1f\x00\x00\x6f\xf8\x01\x00\xf0\xe3\x07\x00\xd0\x8b\x1f\x00\xc0\x1f\xbe\x55\xe9\x2f\xf8\xff\xff\x2f\xe0\xff\xff\x06\x00" +

	// glyph 2533 for rune 'C'
	"\x13\xe9\x01\x02\x11\x00\x00\x50\x00\x00\xd0\xff\x2f\x00\xfc\xff\xff\x00\xff\x01\xf9\xc0\x2f\x00\x80\xd0\x0f\x00\x00\xf0\x07\x00\x00\xf0\x03\x00\x00\xf4\x03\x00\x00\xf8\x02\x00\x00\xf1\x07\x00\x00\x84\x2f\x00\x00\x10\xfd\x00\x00\x40\xf0\x07\x00\x00\xe0\x0f\x00\x00\xc0\x2f\x00\x40\x40\xff\x01\xf8\x00\xfd\xff\xff\x00\xe0\xff\x6f\x00\x00\x54\x01" +

	// glyph 2619 for rune 'D'
	"\x13\xea\x00\x02\x12\xfc\xff\x2b\x00\xf0\xff\xff\x0f\xc0\xaf\xea\xff\x00\x7f\x00\xf4\x0f\xfc\x01\x40\x7f\xf0\x07\x00\xf8\xc2\x1f\x00\xd0\x0f\x7f\x00\x00\x3f\xf1\x07\x00\xf0\x57\xf1\x07\x00\xf0\x13\x7f\x00\x40\x3f\xfc\x01\x00\xbe\xf0\x07\x00\xfd\xc0\x1f\x00\xfd\x02\xbf\x95\xff\x02\xfc\xff\xff\x02\xf0\xff\x6f\x00\x00" +

	// glyph 2698 for rune 'E'
	"\x13\xea\x00\x03\x11\xfc\xff\xff\x7f\xfc\xaa\xaa\x2a\x3f\x00\x00\x50\x15\xff\xff\xff\xc2\xff\xff\xff\xf0\xab\xaa\x1a\xfc\x00\x00\x40\x55\xf1\xab\xaa\xaa\xfc\xff\xff\x7f" +

	// glyph 2740 for rune 'F'
	"\x13\xea\x00\x03\x12\xf4\xff\xff\x7f\xd1\x5f\x55\x55\xd0\x0b\x00\x00\x54\x45\xbf\xaa\x1a\x40\xff\xff\x1f\x40\xbf\xaa\x1a\x40\x2f\x00\x00\x50\x55\x15" +

	// glyph 2777 for rune 'G'
	"\x13\xe9\x01\x01\x11\x00\x00\x50\x00\x00\x00\xfe\xbf\x01\x80\xff\xff\x1f\x80\xbf\x01\x7d\x40\x3f\x00\x80\x01\x7f\x00\x00\x00\xfd\x00\x00\x00\xfc\x01\x00\x00\xf0\x03\x00\x00\x44\x3f\x00\x00\x40\xf4\x03\xf0\xff\x07\x3f\x00\x00\x3f\xfc\x00\x00\xf8\xf0\x07\x00\xe0\x83\x3f\x00\x80\x0f\xfc\x00\x00\x3e\xe0\x0f\x00\xf8\x00\xff\x01\xf4\x03\xf0\xff\xff\x0b\x00\xfd\xff\x07\x00\x00\x55\x00" +

	// glyph 2872 for rune 'H'
	"\x13\xea\x00\x02\x11\xfc\x01\x00\xfc\x55\x55\xfc\xff\xff\xff\xf1\xab\xaa\xfa\xf3\x07\x00\xf0\x57\x55\x05" +

	// glyph 2898 for rune 'I'
	"\x13\xea\x00\x03\x10\xfc\xff\xff\x1f\xa9\xfe\xaa\x02\xc0\x0f\x40\x55\x55\x55\x45\xaa\xbf\xaa\xfc\xff\xff\x1f" +

	// glyph 2925 for rune 'J'
	"\x13\xea\x01\x02\x0f\x00\xfc\xff\x1f\x00\xaa\xfa\x03\x00\x00\x7f\x55\x55\x05\x00\x00\xfd\x01\x00\x80\x2f\x00\x00\xfc\x01\x00\xe4\x0f\xff\xff\x7f\xf0\xff\xff\x01\x55\x15\x00\x00" +

	// glyph 2969 for rune 'K'
	"\x13\xea\x00\x02\x13\xfc\x01\x00\xf4\xc7\x1f\x00\xd0\x1f\xfc\x01\x40\x7f\xc0\x1f\x00\xfd\x01\xfc\x01\xf4\x07\xc0\x1f\xd0\x1f\x00\xfc\x41\x7f\x00\xc0\x1f\xfd\x01\x00\xfc\xf5\x07\x00\xc0\xef\x3f\x00\x00\xfc\xff\x0f\x00\xc0\xff\xfd\x02\x00\xfc\x47\x7f\x00\xc0\x1f\xe0\x0f\x00\xfc\x01\xfc\x02\xc0\x1f\x00\x7f\x00\xfc\x01\xd0\x0f\xc0\x1f\x00\xfc\x03\xfc\x01\x00\xbf\xc0\x1f\x00\xd0\x1f\xfc\x01\x00\xf8\xc3\x1f\x00\x00\xff" +

	// glyph 3073 for rune 'L'
	"\x13\xea\x00\x03\x12\xf8\x01\x00\x00\xf8\x02\x00\x00\x55\x55\x55\x55\xe1\xaf\xaa\xaa\xe1\xff\xff\xff\x07" +

	// glyph 3099 for rune 'M'
	"\x13\xea\x00\x01\x12\xf8\x0b\x00\xf4\xcf\xff\x00\x80\xff\xfc\x1f\x00\xfc\xcf\xff\x02\xd0\xfb\xfc\x3d\x00\xaf\xcf\xcf\x0b\xf4\xf9\xfc\xf8\x80\x8f\xcf\x0f\x1f\xbc\xf8\xfc\xf0\xd2\x83\xcf\x0f\x3d\x2f\xf8\xfc\xc0\xfb\x81\xcf\x0f\xf8\x0f\xf8\xfc\x40\xbf\x80\xcf\x0f\xf0\x03\xf8\xfc\x00\x14\x80\xcf\x0f\x00\x00\xf8\x55\x05" +

	// glyph 3178 for rune 'N'
	"\x13\xea\x00\x02\x11\xf8\x07\x00\xfc\xfc\x0f\x00\xfc\xfc\x1f\x00\xfc\xfc\x2f\x00\xfc\xfc\x3f\x00\xfc\xfc\xbd\x00\xfc\xfc\xf8\x00\xfc\xfc\xf4\x01\xfc\xfc\xf0\x03\xfc\xfc\xd0\x07\xfc\xfc\xc0\x0f\xfc\xfc\x80\x0f\xfc\xfc\x00\x2f\xfc\xfc\x00\x3e\xfc\xfc\x00\xbc\xfc\xfc\x00\xfc\xfc\xfc\x00\xf4\xfd\xfc\x00\xf0\xff\xfc\x00\xd0\xff\xfc\x00\xc0\xff\xfc\x00\x80\xff\xfc\x00\x00\xff" +

	// glyph 3271 for rune 'O'
	"\x13\xea\x01\x02\x12\x00\xf8\xbf\x00\x00\xfc\xff\x1f\x00\xfc\x57\xfe\x01\xf8\x02\xe0\x0f\xf0\x03\x00\x7f\xd0\x0b\x00\xf4\x83\x1f\x00\xc0\x0f\x7f\x00\x00\x3f\xfc\x00\x00\xfc\x55\xc5\x1f\x00\xc0\x4f\xf8\x02\x00\xfd\xc0\x0f\x00\xfc\x01\xbf\x00\xf4\x03\xf0\x1f\xf9\x07\x40\xff\xff\x0b\x00\xe0\xff\x07\x00\x00\x10\x00\x00" +

	// glyph 3350 for rune 'P'
	"\x13\xea\x00\x02\x12\xf4\xff\xbf\x01\xd0\xff\xff\x3f\x40\xbf\xaa\xff\x03\xfd\x00\xd0\x3f\xf4\x03\x00\xfd\xd0\x0f\x00\xf0\x57\xf4\x03\x00\xfd\xd0\x0f\x00\xfc\x43\x7f\x55\xfe\x07\xfd\xff\xff\x07\xf4\xff\xff\x02\xd0\x0f\x00\x00\x50\x55\x05" +

	// glyph 3409 for rune 'Q'
	"\x13\xea\x03\x02\x13\x00\xf8\xbf\x00\x00\xf0\xff\x7f\x00\xc0\x7f\xe5\x1f\x00\xbe\x00\xf8\x03\xf0\x03\x00\x7f\x40\x2f\x00\xd0\x0f\xf8\x01\x00\xfc\xc0\x1f\x00\xc0\x0f\xfc\x00\x00\xfc\x51\x15\x7f\x00\x00\x3f\x84\x2f\x00\xd0\x0f\xf0\x03\x00\x7f\x00\xbf\x00\xf4\x03\xc0\x7f\xe4\x3f\x00\xf4\xff\xff\x0f\x00\xf8\xff\xfd\x03\x00\x00\x00\xbf\x00\x00\x00\xc0\x0f\x00\x00\x00\xf4" +

	// glyph 3501 for rune 'R'
	"\x13\xea\x00\x02\x12\xf8\xff\x6f\x00\xf0\xff\xff\x1f\xc0\xaf\xaa\xff\x00\x7f\x00\xe0\x0f\xfc\x01\x00\x7f\xf0\x07\x00\xf8\xc2\x1f\x00\xd0\x0b\x7f\x00\x80\x2f\xfc\x01\x00\x7f\xf0\x07\x00\xff\xc0\xbf\xfa\xff\x00\xff\xff\xbf\x00\xfc\xff\xbf\x00\xf0\x07\xf0\x03\xc0\x1f\x80\x2f\x00\x7f\x00\xfc\x00\xfc\x01\xe0\x0b\xf0\x07\x00\x3f\xc0\x1f\x00\xf8\x02\x7f\x00\xc0\x1f\xfc\x01\x00\xfe\xf0\x07\x00\xf0\x07" +

	// glyph 3600 for rune 'S'
	"\x13\xe9\x01\x02\x11\x00\x00\x05\x00\x00\xfd\xff\x1b\x80\xff\xff\x3f\xe0\x6f\x40\x3e\xf4\x03\x00\x10\xf8\x01\x00\x00\xf1\x07\x00\x00\xe0\x0b\x00\x00\xd0\x7f\x00\x00\x80\xff\x2b\x00\x00\xfe\xff\x0b\x00\xd0\xff\x7f\x00\x00\xd0\xff\x00\x00\x00\xfc\x02\x00\x00\xf0\x57\x24\x00\x00\xbe\xf8\x06\xd0\x3f\xf8\xff\xff\x1f\xd0\xff\xff\x02\x00\x50\x05\x00" +

	// glyph 3686 for rune 'T'
	"\x13\xea\x00\x00\x13\xf0\xff\xff\xff\x7f\xf4\xff\xff\xff\xbf\xa0\xaa\xfe\xaa\x6a\x00\x00\xfc\x00\x00\x55\x55\x55\x55\x05" +

	// glyph 3716 for rune 'U'
	"\x13\xea\x01\x02\x12\xfc\x00\x00\xf8\xf1\x03\x00\xf0\x57\x55\x55\x55\xfc\x00\x00\xfc\x84\x2f\x00\xd0\x0f\xfc\x06\xe0\x1f\xd0\xff\xff\x2f\x00\xf8\xff\x1f\x00\x00\x54\x01\x00" +

	// glyph 3759 for rune 'V'
	"\x13\xea\x00\x01\x12\xfc\x00\x00\xc0\x8f\x1f\x00\x00\xfc\xf4\x02\x00\xd0\x0b\x3f\x00\x00\x7e\xe0\x07\x00\xf0\x03\xbd\x00\x40\x2f\xc0\x0f\x00\xf8\x01\xfc\x00\xc0\x0f\x40\x1f\x00\xbc\x00\xf0\x02\xd0\x07\x00\x3f\x00\x3f\x00\xe0\x07\xf0\x03\x00\xbd\x40\x2f\x00\xc0\x0f\xf8\x00\x00\xf8\xc0\x0f\x00\x40\x2f\xbc\x00\x00\xf0\xe3\x07\x00\x00\x3e\x3f\x00\x00\xd0\xfb\x02\x00\x00\xfc\x1f\x00\x00\xc0\xff\x00\x00\x00\xf4\x0b\x00" +

	// glyph 3863 for rune 'W'
	"\x13\xea\x00\x00\x13\xbc\x00\xfc\x01\xf8\xf8\x00\xfd\x02\xf8\xf8\x00\xfe\x02\xfc\xf4\x00\xfe\x03\xbc\xf0\x01\xef\x03\xbc\xf0\x02\xdf\x03\x7c\xf0\x02\xcf\x07\x3d\xf0\x43\x8f\x0b\x3e\xe0\x43\x8f\x0b\x3e\xd0\x83\x4b\x0f\x2f\xd0\x83\x07\x0f\x2f\xc0\xc7\x07\x0f\x1f\xc0\xc7\x03\x1f\x1f\xc0\xcb\x03\x5e\x0f\x80\xdb\x03\x6e\x0f\x80\xdf\x02\xad\x0f\x40\xef\x01\xbc\x0b\x00\xff\x01\xfc\x07\x00\xff\x00\xfc\x07\x00\xff\x00\xf8\x03\x00\xfe\x00\xf8\x03\x00\xbe\x00\xf4\x03" +

	// glyph 3978 for rune 'X'
	"\x13\xea\x00\x00\x13\xf0\x07\x00\x00\x3f\xd0\x0f\x00\xc0\x2f\x80\x3f\x00\xd0\x0f\x00\x7f\x00\xf0\x03\x00\xfc\x00\xf8\x01\x00\xf4\x03\xfd\x00\x00\xe0\x0b\x3f\x00\x00\xc0\x9f\x1f\x00\x00\x40\xff\x0f\x00\x00\x00\xfe\x03\x00\x00\x00\xfc\x02\x00\x00\x00\xfe\x03\x00\x00\x40\xff\x0f\x00\x00\xc0\x8f\x1f\x00\x00\xe0\x0b\x3f\x00\x00\xf4\x03\xfd\x00\x00\xfc\x01\xfc\x01\x00\xbe\x00\xf0\x03\x40\x3f\x00\xe0\x0b\xc0\x1f\x00\xc0\x1f\xe0\x0b\x00\x40\x3f\xf4\x03\x00\x00\xbe" +

	// glyph 4093 for rune 'Y'
	"\x13\xea\x00\x01\x13\xfc\x00\x00\xc0\x1f\xfe\x00\x00\xf8\x03\x7f\x00\x00\x3f\x40\x3f\x00\xe0\x0b\x80\x2f\x00\xfd\x00\xc0\x0f\xc0\x1f\x00\xd0\x0b\xf4\x02\x00\xf0\x07\x3f\x00\x00\xf0\xe3\x07\x00\x00\xf8\xfe\x00\x00\x00\xfc\x0f\x00\x00\x00\xfd\x02\x00\x00\x00\x3f\x00\x40\x55\x55" +

	// glyph 4162 for rune 'Z'
	"\x13\xea\x00\x01\x12\xf0\xff\xff\xff\x13\xa8\xaa\xaa\xff\x00\x00\x00\xf0\x07\x00\x00\x80\x2f\x00\x00\x00\xfd\x00\x00\x00\xf0\x07\x00\x00\xc0\x2f\x00\x00\x00\xfe\x00\x00\x00\xf0\x03\x00\x00\xc0\x1f\x00\x00\x00\xbe\x00\x00\x00\xf4\x03\x00\x00\xc0\x1f\x00\x00\x00\xbe\x00\x00\x00\xf4\x03\x00\x00\xc0\x0f\x00\x00\x00\x7e\x00\x00\x00\xf4\x02\x00\x00\xc0\xaf\xaa\xaa\x1a\xfd\xff\xff\xff\x06" +

	// glyph 4258 for rune '['
	"\x13\xe7\x03\x04\x0f\xf4\xff\x7f\xd1\xaf\xaa\xd0\x0b\x00\x54\x55\x55\x55\x55\x45\xbf\xaa\x42\xff\xff\x17" +

	// glyph 4284 for rune '\\'
	"\x13\xe8\x02\x02\x11\x54\x00\x00\x00\xf8\x01\x00\x00\xf0\x03\x00\x00\xe0\x07\x00\x00\xc0\x0f\x00\x00\x80\x1f\x00\x00\x00\x3f\x00\x00\x00\x7e\x00\x00\x00\xfc\x00\x00\x00\xf8\x01\x00\x00\xf0\x03\x00\x00\xe0\x07\x00\x00\xc0\x0f\x00\x00\x80\x1f\x00\x00\x00\x3f\x00\x00\x00\x7e\x00\x00\x00\xfc\x00\x00\x00\xf8\x01\x00\x00\xf0\x03\x00\x00\xe0\x07\x00\x00\xc0\x0f\x00\x00\x80\x1f\x00\x00\x00\x3f\x00\x00\x00\x7e\x00\x00\x00\xfc\x00\x00\x00\x54" +

	// glyph 4393 for rune ']'
	"\x13\xe7\x03\x05\x0f\xfc\xff\x1f\xff\xff\x8b\xaa\xfa\x02\x00\xbd\x55\x55\x55\x55\x55\xa1\xaa\xbe\xfc\xff\x2f\xff\xff\x07" +

	// glyph 4423 for rune '^'
	"\x13\xea\xf3\x02\x11\x00\xc0\x1f\x00\x00\xf0\x3f\x00\x00\xf8\xff\x00\x00\xfd\xf8\x02\x00\x3f\xf0\x07\xc0\x1f\xc0\x0f\xe0\x0f\x80\x3f\xf4\x03\x00\xbf\x54\x01\x00\x54" +

	// glyph 4464 for rune '_'
	"\x13\x00\x03\xff\x14\xa0\xaa\xaa\xaa\xaa\x46\xff\xff\xff\xff\xbf\xf0\xff\xff\xff\xff\x07" +

	// glyph 4486 for rune '`'
	"\x13\xe8\xef\x05\x0d\xf4\x03\x80\x3f\x00\xfc\x02\xd0\x0f\x00\xfe\x00\xf0\x0b\x00\x15" +

	// glyph 4507 for rune 'a'
	"\x13\xef\x01\x02\x11\x40\xff\x6f\x00\x80\xff\xff\x07\x40\x55\xf9\x1f\x00\x00\x40\x3f\x00\x00\x00\x7e\x00\x00\x00\xbd\x00\x40\x55\xbe\x00\xfe\xff\xbf\xd0\xff\xff\xbf\xf0\x07\x00\xbd\xf8\x01\x00\xbd\xfc\x00\x00\xbe\xfc\x00\x00\xbf\xf8\x01\x80\xbf\xf4\x07\xe0\xbf\xe0\xff\xff\xbd\x40\xff\x7f\xbc\x00\x50\x01\x00" +

	// glyph 4584 for rune 'b'
	"\x13\xe8\x01\x03\x12\xa8\x00\x00\x00\xfc\x00\x00\x00\x55\xf1\x43\xbe\x00\xf0\xf3\xff\x0b\xf0\xff\xe6\x2f\xf0\x2f\x00\x3e\xf0\x0f\x00\xbc\xf0\x07\x00\xf4\xf0\x07\x00\xf0\xf0\x03\x00\xf0\x15\x3f\x00\x00\x0f\x7f\x00\x40\x0f\xbf\x00\x80\x0f\xff\x00\xc0\x07\xff\x07\xf4\x03\xbf\xff\xff\x00\x2f\xfe\x3f\x00\x00\x50\x01\x00" +

	// glyph 4663 for rune 'c'
	"\x13\xef\x01\x02\x10\x00\x80\xfe\x06\x00\xff\xff\x0f\xf0\xaf\xf9\x03\xbf\x00\xc0\xe0\x0f\x00\x00\xfc\x00\x00\x10\xbd\x00\x00\x50\xf4\x03\x00\x00\xfc\x00\x00\x00\x7e\x00\x00\x40\x3f\x00\x00\x80\x7f\x00\x38\x80\xff\xff\x0f\x40\xff\xff\x01\x00\x54\x00" +

	// glyph 4725 for rune 'd'
	"\x13\xe8\x01\x02\x11\x00\x00\x00\x29\x00\x00\x00\x7e\x55\x01\x90\x2f\xf8\x01\xfd\xff\xf9\x41\xff\xe6\xff\xc1\x2f\x00\xff\xd1\x0f\x00\xfd\xe1\x07\x00\xfc\xf1\x03\x00\xfc\xf1\x03\x00\xf8\x55\xfc\x00\x00\x7f\xf8\x02\x00\x7f\xf0\x03\x80\x7f\xe0\x0f\xe0\x7f\xc0\xff\xff\x7e\x00\xfe\x2f\x7d\x00\x50\x01\x00" +

	// glyph 4800 for rune 'e'
	"\x13\xef\x01\x02\x12\x00\x90\x6f\x00\x00\xf8\xff\x2f\x00\xf8\x6b\xfe\x02\xf8\x03\xc0\x1f\xf0\x03\x00\xfc\xd0\x0b\x00\xe0\xc3\x0f\x00\x80\x1f\xbf\xaa\xaa\x7f\xfc\xff\xff\xff\xf1\x57\x55\x55\xc1\x0f\x00\x00\x00\x3e\x00\x00\x00\xf4\x02\x00\x00\xc0\x1f\x00\x00\x01\xfd\x02\x90\x0b\xd0\xff\xff\x2f\x00\xf8\xff\x2f\x00\x00\x54\x01\x00" +

	// glyph 4882 for rune 'f'
	"\x13\xe9\x00\x02\x11\x00\x00\xf9\xff\x00\x80\xff\xff\x00\xc0\x6f\x55\x00\xd0\x07\x00\x00\xe0\x03\x00\x01\xc0\x0f\x00\xc4\xff\xff\xff\x11\xa9\xfe\xaa\x02\x00\xfc\x00\x40\x55\x55\x05" +

	// glyph 4927 for rune 'g'
	"\x13\xef\x08\x02\x11\x00\xe4\x1b\x54\x00\xfe\xff\xbd\xc0\xbf\xf9\xbf\xd0\x0f\x80\xbf\xf0\x03\x00\xbf\xf0\x03\x00\xbe\xf4\x01\x00\xbd\xf8\x01\x00\xbc\x45\x1f\x00\xd0\x4b\x2f\x00\xd0\x0b\x3f\x00\xf0\x0b\xbe\x00\xf4\x0b\xfc\x47\xfe\x0b\xf0\xff\xef\x0b\x80\xff\xc7\x0b\x00\x00\xc0\x1b\x00\x00\x40\x1f\x00\x00\xc0\x0f\x00\x00\xe0\x0f\xa0\xaa\xfe\x03\xf0\xff\xff\x00\xf0\xff\x2f\x00" +

	// glyph 5021 for rune 'h'
	"\x13\xe8\x00\x03\x11\xa8\x00\x00\x00\x3f\x00\x00\x50\x15\x3f\xf8\x1f\xc0\xdf\xff\x3f\xf0\xff\xfa\x2f\xfc\x07\xd0\x0f\xbf\x00\xf0\xc3\x0f\x00\xf8\xc5\x0f\x00\xf4\x55\x55\x05" +

	// glyph 5064 for rune 'i'
	"\x13\xe9\x00\x05\x0c\x00\xfc\x05\x80\x0a\x00\x10\x55\x15\xff\x7f\x00\xfc\x55\x55\x55\x01" +

	// glyph 5086 for rune 'j'
	"\x13\xe9\x08\x03\x0e\x00\x00\x7d\x00\x00\xbe\x01\x00\xa4\x01\x00\x00\x04\x40\x55\x05\xe0\xff\x1f\x00\x00\x7f\x55\x55\x55\x05\x00\xc0\x0b\x00\xd0\x0b\x00\xf0\x07\xaa\xfe\x43\xff\xff\x40\xff\x1b\x00" +

	// glyph 5135 for rune 'k'
	"\x13\xe9\x00\x03\x12\xf4\x01\x00\x00\xf4\x02\x00\x00\x55\xf4\x02\x00\x15\xf4\x01\xc0\x1f\xf4\x01\xf0\x07\xf4\x01\xfd\x01\xf4\x41\x3f\x00\xf4\xd1\x0f\x00\xf4\xfa\x03\x00\xf4\xff\x07\x00\xf4\xff\x0f\x00\xf4\x8b\x3f\x00\xf4\x02\xbf\x00\xf4\x02\xfc\x01\xf4\x02\xf4\x03\xf4\x02\xe0\x0f\xf4\x02\xc0\x2f\xf4\x02\x00\x7f\xf4\x02\x00\xfd" +

	// glyph 5217 for rune 'l'
	"\x13\xea\x00\x04\x0c\xf8\xff\x46\xaa\x2f\x00\xbc\x55\x55\x55\x55\x05" +

	// glyph 5234 for rune 'm'
	"\x13\xef\x00\x01\x12\x54\xe4\x02\xbe\x80\xfb\xff\xfc\x3f\xf8\xdb\xff\xe6\x8b\x1f\xf0\x07\xfc\xf8\x00\x3f\x80\x8f\x0f\xe0\x03\xf8\xe1\x03\xf8\x00\x7d\x55\x55" +

	// glyph 5273 for rune 'n'
	"\x13\xef\x00\x03\x11\x54\x90\x6f\x00\x2f\xfe\xbf\xc0\xef\xeb\x7f\xf0\x1f\x40\x3f\xfc\x02\xc0\x0f\x3f\x00\xe0\x17\x3f\x00\xd0\x57\x55\x15" +

	// glyph 5307 for rune 'o'
	"\x13\xef\x01\x02\x11\x00\xe4\x6f\x00\x00\xff\xff\x07\xc0\xbf\xf9\x1f\xe0\x0f\x80\x3f\xf0\x03\x00\x7f\xf4\x02\x00\xbd\xf8\x01\x00\xfc\xfc\x01\x00\xfc\xfc\x00\x00\xfc\xfc\x01\x00\xfc\xf8\x01\x00\xfc\xf8\x02\x00\xfd\xf4\x03\x00\xbe\xf0\x07\x00\x3f\xd0\x2f\xd0\x2f\x80\xff\xff\x0f\x00\xfd\xff\x02\x00\x40\x05\x00" +

	// glyph 5384 for rune 'p'
	"\x13\xef\x08\x03\x12\x54\x90\x2b\x00\xbc\xfc\xff\x02\xfc\xbf\xf9\x0b\xfc\x0f\xc0\x0f\xfc\x03\x00\x2f\xfc\x03\x00\x3f\xfc\x01\x00\x3e\xfc\x01\x00\x3d\xfc\x01\x00\x7d\xf1\x07\x00\xf4\xf0\x0b\x00\xf8\xf0\x0f\x00\xfc\xf0\x1f\x00\x7d\xf0\xbf\x40\x3f\xf0\xfb\xff\x0f\xf0\xe3\xff\x02\xf0\x03\x15\x00\xf0\x03\x00\x00\x54\x15" +

	// glyph 5463 for rune 'q'
	"\x13\xef\x08\x02\x11\x00\xe4\x0b\x14\x00\xff\xbf\x7c\xc0\xbf\xfe\x7f\xe0\x0b\xe0\x7f\xf0\x03\x80\x7f\xf4\x01\x40\x7f\xf8\x00\x00\x7f\xfc\x00\x00\x7f\x85\x0f\x00\xf0\x17\xbd\x00\xd0\x1f\xfc\x00\xf0\x1f\xf4\x03\xfc\x1f\xe0\xff\xbf\x1f\x80\xff\x4f\x1f\x00\x54\x41\x1f\x00\x00\x40\x5f\x55\x01" +

	// glyph 5535 for rune 'r'
	"\x13\xef\x00\x04\x11\x54\x90\xaf\xc0\xcf\xff\xbf\xfc\xff\xfa\xcf\xff\x00\x90\xfc\x03\x00\xc0\x0f\x00\x00\x55\x55\x15" +

	// glyph 5564 for rune 's'
	"\x13\xef\x01\x03\x10\x00\xe8\x6b\x00\xfc\xff\x3f\xf0\x6f\xe9\x43\x3f\x00\x10\xf4\x01\x00\x80\x1f\x00\x00\xf4\x07\x00\x00\xff\x5b\x00\x80\xff\xbf\x00\x40\xfe\x3f\x00\x00\xf4\x0b\x00\x00\xfc\x85\x0b\x40\x7f\xf8\xff\xff\x42\xfe\xff\x07\x00\x50\x01\x00" +

	// glyph 5626 for rune 't'
	"\x13\xea\x00\x04\x0f\x54\x00\x00\xfc\x00\x00\x15\x7f\x55\x05\xff\xff\x4f\xfc\x00\x00\x55\x55\xf1\x07\x00\xd0\x6f\x55\xc1\xff\xff\x03\xfd\xff\x03" +

	// glyph 5662 for rune 'u'
	"\x13\xef\x01\x02\x11\x50\x00\x00\x14\xf4\x01\x00\xbd\x55\x55\xd1\x0b\x00\xf4\xc2\x0b\x00\xf8\xc2\x0f\x00\xfd\x82\x7f\x80\xff\x02\xff\xff\xff\x02\xf8\xff\xf2\x02\x40\x15\x00\x00" +

	// glyph 5706 for rune 'v'
	"\x13\xef\x00\x02\x12\x54\x00\x00\x50\xf1\x03\x00\xe0\x87\x0f\x00\xc0\x0f\xbd\x00\x40\x2f\xf0\x03\x00\x3e\x80\x1f\x00\xfc\x00\xbc\x00\xf4\x01\xf0\x03\xf0\x03\x40\x1f\xc0\x0b\x00\xfc\x80\x1f\x00\xe0\x03\x3f\x00\x40\x2f\xbd\x00\x00\xfc\xf8\x00\x00\xe0\xf7\x03\x00\x00\xff\x07\x00\x00\xfc\x0f\x00\x00\xd0\x2f\x00\x00" +

	// glyph 5784 for rune 'w'
	"\x13\xef\x00\x00\x13\x54\x00\x00\x00\x54\xfc\x00\x00\x00\xfc\xf8\x01\x00\x00\xfc\xf4\x02\x00\x00\xbc\xf0\x03\x00\x00\x7d\xf0\x03\xfc\x00\x3e\xe0\x03\xfd\x02\x3f\xd0\x07\xfe\x03\x2f\xc0\x0b\xef\x03\x1f\xc0\x0f\xcf\x4b\x0f\x80\x4f\x8f\x8f\x0f\x40\x9f\x4b\xcf\x0b\x00\xef\x03\xef\x07\x00\xff\x03\xfe\x03\x00\xfe\x02\xfd\x03\x00\xfd\x01\xfc\x02\x00\xfc\x00\xfc\x01" +

	// glyph 5874 for rune 'x'
	"\x13\xef\x00\x01\x12\x50\x01\x00\x50\x01\xbd\x00\x80\x2f\x80\x2f\x00\xfd\x00\xf0\x07\xf0\x03\x00\xfc\xc0\x1f\x00\x40\x3f\xbe\x00\x00\xd0\xff\x02\x00\x00\xf8\x0f\x00\x00\x00\x7f\x00\x00\x00\xfc\x0f\x00\x00\xe0\xfb\x03\x00\x40\x3f\xbe\x00\x00\xfc\x80\x1f\x00\xf0\x07\xf0\x07\xc0\x1f\x00\xfc\x00\xbe\x00\x40\x3f\xf4\x03\x00\xe0\x0b" +

	// glyph 5956 for rune 'y'
	"\x13\xef\x08\x02\x11\x54\x00\x00\x54\xfc\x00\x00\xbc\x55\x55\x85\x1f\x00\xc0\x4b\x3f\x00\xc0\x0b\xff\xaa\xea\x0b\xfc\xff\xff\x0f\x90\xff\xff\x0f\x00\x00\xc0\x1f\x00\x00\x40\x2f\x00\x00\x80\x1f\x00\x00\xc0\x0f\xa0\xaa\xfa\x0b\xf4\xff\xff\x03\xf4\xff\xbf\x00" +

	// glyph 6020 for rune 'z'
	"\x13\xef\x00\x03\x11\x50\x55\x55\x05\xfd\xff\xff\x47\xaa\xaa\xfe\x00\x00\x40\x3f\x00\x00\xf4\x03\x00\x00\x3f\x00\x00\xf0\x07\x00\x00\x7f\x00\x00\xe0\x0b\x00\x00\xbe\x00\x00\xd0\x0f\x00\x00\xfd\x00\x00\xc0\x0f\x00\x00\xfc\x01\x00\x80\x2f\x00\x00\xf0\xff\xff\x7f\x01" +

	// glyph 6086 for rune '{'
	"\x13\xe7\x04\x03\x10\x00\x00\xfd\x0f\x00\xf8\xff\x00\xc0\xbf\x05\x00\xfc\x00\x00\xd0\x0b\x10\x00\xf4\x01\x14\x00\xf8\x01\x04\x00\x3f\x00\x54\xfe\x02\xc0\xff\x0b\x00\xfc\x7f\x00\x40\xe5\x2f\x00\x00\xf0\x03\x00\x00\x7e\x00\x01\x40\x1f\x40\x01\x40\x2f\x40\x00\xd0\x0f\x00\x00\xfc\x56\x00\x80\xff\x0f\x00\xe0\xff\x00\x00\x00\x05" +

	// glyph 6167 for rune '|'
	"\x13\xe7\x03\x08\x0b\x10\xfc\x55\x55\x55\x55\x55\x55\x51\x01" +

	// glyph 6182 for rune '}'
	"\x13\xe7\x04\x03\x11\xfc\x2f\x00\x00\xff\x3f\x00\x40\xe9\x1f\x00\x00\xe0\x07\x00\x00\xf4\x02\x00\x00\xbc\x00\x50\x05\x00\x3f\x00\x00\xc0\x2f\x00\x00\xd0\x7f\x05\x00\xd0\xff\x1b\x00\xf4\x6f\x01\x00\xbf\x00\x00\xc0\x0f\x00\x00\xf0\x02\x40\x15\x00\xbd\x00\x00\x80\x2f\x00\x50\xf9\x07\x00\xfc\xff\x00\x00\xff\x0b\x00\x00\x01\x00\x00\x00" +

	// glyph 6265 for rune '~'
	"\x13\xf1\xf9\x01\x12\x00\x69\x00\x00\x00\xfc\x3f\x00\x14\xe0\xeb\x0f\x80\x0f\x1f\xf4\x02\xbc\xf4\x00\xbd\xd0\x87\x0f\x80\xff\x2f\x00\x00\xe0\xff\x00\x00\x00\x50\x01" +

	"")
