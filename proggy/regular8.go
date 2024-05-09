// File generated using:
// 	go run ./generate -font=proggy/ProggyVector-Regular.ttf -size=8 -dpi=72 -package=proggy -weight=Regular

package proggy

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      1235
// - glyph metadata:  475
// - glyph mask data: 559

var Regular8 = font.Make("" +
	"\x00" + // version: 0
	"\x08" + // size:   8
	"\x0c" + // height: 12
	"\x08" + // ascent: 8

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xd6\x00" + // "\"" at index 214
	"\xde\x00" + // "#" at index 222
	"\xec\x00" + // "$" at index 236
	"\xf9\x00" + // "%" at index 249
	"\x07\x01" + // "&" at index 263
	"\x15\x01" + // "'" at index 277
	"\x1c\x01" + // "(" at index 284
	"\x26\x01" + // ")" at index 294
	"\x2f\x01" + // "*" at index 303
	"\x39\x01" + // "+" at index 313
	"\x44\x01" + // "," at index 324
	"\x4a\x01" + // "-" at index 330
	"\x50\x01" + // "." at index 336
	"\x56\x01" + // "/" at index 342
	"\x60\x01" + // "0" at index 352
	"\x6e\x01" + // "1" at index 366
	"\x78\x01" + // "2" at index 376
	"\x84\x01" + // "3" at index 388
	"\x90\x01" + // "4" at index 400
	"\x9e\x01" + // "5" at index 414
	"\xab\x01" + // "6" at index 427
	"\xb9\x01" + // "7" at index 441
	"\xc4\x01" + // "8" at index 452
	"\xd2\x01" + // "9" at index 466
	"\xe0\x01" + // ":" at index 480
	"\xe7\x01" + // ";" at index 487
	"\xef\x01" + // "<" at index 495
	"\xf9\x01" + // "=" at index 505
	"\x03\x02" + // ">" at index 515
	"\x0e\x02" + // "?" at index 526
	"\x19\x02" + // "@" at index 537
	"\x27\x02" + // "A" at index 551
	"\x35\x02" + // "B" at index 565
	"\x42\x02" + // "C" at index 578
	"\x4e\x02" + // "D" at index 590
	"\x5c\x02" + // "E" at index 604
	"\x68\x02" + // "F" at index 616
	"\x72\x02" + // "G" at index 626
	"\x80\x02" + // "H" at index 640
	"\x8a\x02" + // "I" at index 650
	"\x93\x02" + // "J" at index 659
	"\x9d\x02" + // "K" at index 669
	"\xaa\x02" + // "L" at index 682
	"\xb4\x02" + // "M" at index 692
	"\xc1\x02" + // "N" at index 705
	"\xcf\x02" + // "O" at index 719
	"\xdd\x02" + // "P" at index 733
	"\xea\x02" + // "Q" at index 746
	"\xfa\x02" + // "R" at index 762
	"\x06\x03" + // "S" at index 774
	"\x13\x03" + // "T" at index 787
	"\x1c\x03" + // "U" at index 796
	"\x28\x03" + // "V" at index 808
	"\x35\x03" + // "W" at index 821
	"\x43\x03" + // "X" at index 835
	"\x51\x03" + // "Y" at index 849
	"\x5d\x03" + // "Z" at index 861
	"\x6a\x03" + // "[" at index 874
	"\x73\x03" + // "\\" at index 883
	"\x80\x03" + // "]" at index 896
	"\x89\x03" + // "^" at index 905
	"\x91\x03" + // "_" at index 913
	"\x98\x03" + // "`" at index 920
	"\x9f\x03" + // "a" at index 927
	"\xa9\x03" + // "b" at index 937
	"\xb6\x03" + // "c" at index 950
	"\xbf\x03" + // "d" at index 959
	"\xca\x03" + // "e" at index 970
	"\xd5\x03" + // "f" at index 981
	"\xe0\x03" + // "g" at index 992
	"\xec\x03" + // "h" at index 1004
	"\xf6\x03" + // "i" at index 1014
	"\xff\x03" + // "j" at index 1023
	"\x0a\x04" + // "k" at index 1034
	"\x17\x04" + // "l" at index 1047
	"\x1f\x04" + // "m" at index 1055
	"\x28\x04" + // "n" at index 1064
	"\x31\x04" + // "o" at index 1073
	"\x3a\x04" + // "p" at index 1082
	"\x47\x04" + // "q" at index 1095
	"\x53\x04" + // "r" at index 1107
	"\x5b\x04" + // "s" at index 1115
	"\x64\x04" + // "t" at index 1124
	"\x6f\x04" + // "u" at index 1135
	"\x77\x04" + // "v" at index 1143
	"\x81\x04" + // "w" at index 1153
	"\x8c\x04" + // "x" at index 1164
	"\x96\x04" + // "y" at index 1174
	"\xa1\x04" + // "z" at index 1185
	"\xaa\x04" + // "{" at index 1194
	"\xb6\x04" + // "|" at index 1206
	"\xbe\x04" + // "}" at index 1214
	"\xcb\x04" + // "~" at index 1227

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x05\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x05\xfa\x00\x02\x03\xc8\x21\x30" +

	// glyph 214 for rune '"'
	"\x05\xfa\xfd\x01\x04\x04\x88\x44" +

	// glyph 222 for rune '#'
	"\x05\xfa\x00\x00\x05\x40\x00\x24\xb4\x07\x22\xb4\x02\x09" +

	// glyph 236 for rune '$'
	"\x05\xfa\x00\x01\x05\x10\xe0\x81\x02\x2d\x60\xe1\x01" +

	// glyph 249 for rune '%'
	"\x05\xfa\x00\x00\x05\x10\x80\x25\xa4\x00\x08\x60\x06\x61" +

	// glyph 263 for rune '&'
	"\x05\xfa\x00\x00\x05\x60\x40\x09\xb0\x40\x49\x08\x46\x6a" +

	// glyph 277 for rune '\''
	"\x05\xfa\xfd\x02\x03\x84\x04" +

	// glyph 284 for rune '('
	"\x05\xfa\x01\x01\x03\x20\x85\x84\x81\x10" +

	// glyph 294 for rune ')'
	"\x05\xfa\x01\x02\x04\x48\x54\x48\x04" +

	// glyph 303 for rune '*'
	"\x05\xfb\xff\x00\x04\x80\x90\x0b\x2e\x20" +

	// glyph 313 for rune '+'
	"\x05\xfb\xff\x00\x05\x40\x00\x08\xe4\x06\x08" +

	// glyph 324 for rune ','
	"\x05\xff\x01\x02\x03\x8c" +

	// glyph 330 for rune '-'
	"\x05\xfd\xfe\x01\x04\xa8" +

	// glyph 336 for rune '.'
	"\x05\xff\x00\x02\x03\x0c" +

	// glyph 342 for rune '/'
	"\x05\xfa\x00\x01\x04\x40\x80\x20\x21\x04" +

	// glyph 352 for rune '0'
	"\x05\xfa\x00\x00\x05\x90\x00\x32\x54\x43\x5a\x34\x02\x2a" +

	// glyph 366 for rune '1'
	"\x05\xfa\x00\x01\x04\x10\x38\x20\x85\x0b" +

	// glyph 376 for rune '2'
	"\x05\xfa\x00\x00\x04\x90\x00\x1c\x60\x60\xd0\x0a" +

	// glyph 388 for rune '3'
	"\x05\xfa\x00\x00\x04\xa0\x00\x0c\x64\x00\x92\x0a" +

	// glyph 400 for rune '4'
	"\x05\xfa\x00\x00\x05\x40\x01\x28\x20\x42\x21\xa4\x07\x20" +

	// glyph 414 for rune '5'
	"\x05\xfa\x00\x00\x04\x50\x81\x00\x0b\xc0\x00\x92\x0a" +

	// glyph 427 for rune '6'
	"\x05\xfa\x00\x00\x05\x40\x01\x02\xa4\x41\x22\x14\x06\x2a" +

	// glyph 441 for rune '7'
	"\x05\xfa\x00\x01\x04\x54\xc0\x80\x20\x51\x00" +

	// glyph 452 for rune '8'
	"\x05\xfa\x00\x00\x05\x90\x41\x22\x20\x02\x26\x14\x06\x2a" +

	// glyph 466 for rune '9'
	"\x05\xfa\x00\x00\x05\x90\x40\x22\x14\x02\x76\x00\x02\x1a" +

	// glyph 480 for rune ':'
	"\x05\xfb\xff\x02\x03\x14\x30" +

	// glyph 487 for rune ';'
	"\x05\xfb\x00\x02\x03\x14\x30\x02" +

	// glyph 495 for rune '<'
	"\x05\xfc\xff\x00\x05\x90\x41\x06\x40\x06" +

	// glyph 505 for rune '='
	"\x05\xfc\xff\x00\x05\xa4\x06\x00\x54\x01" +

	// glyph 515 for rune '>'
	"\x05\xfb\xff\x00\x05\x04\x00\x19\x40\x47\x06" +

	// glyph 526 for rune '?'
	"\x05\xfa\x00\x01\x04\x64\x84\x90\x20\x00\x30" +

	// glyph 537 for rune '@'
	"\x05\xfa\x00\x00\x05\x90\x82\x95\x58\x8a\x52\x88\x02\x96" +

	// glyph 551 for rune 'A'
	"\x05\xfa\x00\x00\x05\x40\x00\x0d\x20\x02\x32\xa4\x83\x90" +

	// glyph 565 for rune 'B'
	"\x05\xfa\x00\x00\x05\x50\x40\x21\x64\x12\x45\xd1\x0a" +

	// glyph 578 for rune 'C'
	"\x05\xfa\x00\x00\x04\x80\xc2\x40\x41\x20\x80\x0a" +

	// glyph 590 for rune 'D'
	"\x05\xfa\x00\x00\x05\x50\x40\x31\x14\x46\x51\x14\x42\x1a" +

	// glyph 604 for rune 'E'
	"\x05\xfa\x00\x00\x04\x50\x91\x10\x59\x24\xd0\x0a" +

	// glyph 616 for rune 'F'
	"\x05\xfa\x00\x01\x04\x54\x08\x71\x21\x04" +

	// glyph 626 for rune 'G'
	"\x05\xfa\x00\x00\x05\x80\x01\x02\x14\x80\x64\x14\x05\x2a" +

	// glyph 640 for rune 'H'
	"\x05\xfb\x00\x00\x05\x14\x16\x99\x51\x58" +

	// glyph 650 for rune 'I'
	"\x05\xfa\x00\x01\x04\x54\x30\x15\x2e" +

	// glyph 659 for rune 'J'
	"\x05\xfa\x00\x00\x04\x50\x01\x58\x91\x06" +

	// glyph 669 for rune 'K'
	"\x05\xfb\x00\x00\x05\x14\x42\x0a\xe4\x40\x21\x14\x06" +

	// glyph 682 for rune 'L'
	"\x05\xfa\x00\x01\x05\x04\x20\x50\xb1\x06" +

	// glyph 692 for rune 'M'
	"\x05\xfa\x00\x00\x05\x04\x85\xb2\x68\x89\x8c\x08\x18" +

	// glyph 705 for rune 'N'
	"\x05\xfa\x00\x00\x05\x10\x40\x53\x64\x45\x59\x54\x47\x71" +

	// glyph 719 for rune 'O'
	"\x05\xfa\x00\x00\x05\x90\x00\x32\x14\x86\x51\x14\x02\x2a" +

	// glyph 733 for rune 'P'
	"\x05\xfa\x00\x00\x05\x50\x40\x22\x14\x46\x2b\x14\x10" +

	// glyph 746 for rune 'Q'
	"\x05\xfa\x01\x00\x05\x90\x00\x32\x14\x86\x51\x14\x02\x3a\x00\x08" +

	// glyph 762 for rune 'R'
	"\x05\xfa\x00\x00\x04\x50\x50\x4c\x21\x7d\x14\x06" +

	// glyph 774 for rune 'S'
	"\x05\xfa\x00\x00\x05\x90\x41\x02\x41\x0a\x80\x91\x0a" +

	// glyph 787 for rune 'T'
	"\x05\xfa\x00\x00\x05\x54\x05\x0c\x55" +

	// glyph 796 for rune 'U'
	"\x05\xfa\x00\x00\x05\x04\x80\x51\x45\x51\xa0\x02" +

	// glyph 808 for rune 'V'
	"\x05\xfa\x00\x00\x05\x04\x44\x51\x20\x12\x68\x00\x03" +

	// glyph 821 for rune 'W'
	"\x05\xfa\x00\x00\x05\x44\x84\x8c\x98\x85\x66\x34\x03\x32" +

	// glyph 835 for rune 'X'
	"\x05\xfa\x00\x00\x05\x04\x00\x22\xd0\x01\x0c\x20\x82\x51" +

	// glyph 849 for rune 'Y'
	"\x05\xfa\x00\x00\x05\x04\x04\x22\xa0\x01\x0c\x05" +

	// glyph 861 for rune 'Z'
	"\x05\xfa\x00\x00\x04\x50\x01\x0c\x18\x20\x20\xd0\x0a" +

	// glyph 874 for rune '['
	"\x05\xfa\x01\x01\x04\x1c\x08\x55\x68" +

	// glyph 883 for rune '\\'
	"\x05\xfa\x00\x00\x04\x14\x80\x00\x05\x20\x40\x01\x08" +

	// glyph 896 for rune ']'
	"\x05\xfa\x01\x01\x04\x94\x80\x55\x64" +

	// glyph 905 for rune '^'
	"\x05\xfa\xfd\x01\x04\x10\x58\x44" +

	// glyph 913 for rune '_'
	"\x05\x00\x01\x00\x05\xa8\x0a" +

	// glyph 920 for rune '`'
	"\x05\xfa\xfc\x01\x03\x24\x04" +

	// glyph 927 for rune 'a'
	"\x05\xfc\x00\x00\x04\x50\x42\x4e\x21\xe8" +

	// glyph 937 for rune 'b'
	"\x05\xfa\x00\x00\x05\x10\x40\x01\x70\x42\x52\xd1\x0a" +

	// glyph 950 for rune 'c'
	"\x05\xfc\x00\x00\x04\x60\x52\x10\xa8" +

	// glyph 959 for rune 'd'
	"\x05\xfa\x00\x00\x04\x00\x06\x36\x85\x81\x0e" +

	// glyph 970 for rune 'e'
	"\x05\xfc\x00\x00\x05\x60\x42\x66\x14\x00\x2a" +

	// glyph 981 for rune 'f'
	"\x05\xfa\x00\x00\x04\x80\x02\x42\x2e\x20\x05" +

	// glyph 992 for rune 'g'
	"\x05\xfc\x02\x00\x04\x60\x53\x18\xe8\x00\x82\x0a" +

	// glyph 1004 for rune 'h'
	"\x05\xfa\x00\x01\x04\x04\x71\x23\x12\x06" +

	// glyph 1014 for rune 'i'
	"\x05\xfa\x00\x01\x03\x20\x40\x83\x05" +

	// glyph 1023 for rune 'j'
	"\x05\xfa\x02\x01\x04\x50\x00\x60\x50\x15\x0e" +

	// glyph 1034 for rune 'k'
	"\x05\xfa\x00\x01\x05\x04\x20\x80\x08\x0f\x58\x20\x06" +

	// glyph 1047 for rune 'l'
	"\x05\xfa\x00\x01\x03\x14\x0c\x56" +

	// glyph 1055 for rune 'm'
	"\x05\xfc\x00\x00\x05\xd8\x86\x88\x05" +

	// glyph 1064 for rune 'n'
	"\x05\xfc\x00\x01\x04\x98\x88\x84\x01" +

	// glyph 1073 for rune 'o'
	"\x05\xfc\x00\x00\x04\x60\x52\x18\xa8" +

	// glyph 1082 for rune 'p'
	"\x05\xfc\x02\x00\x05\x74\x42\x52\xd1\x0a\x05\x40\x00" +

	// glyph 1095 for rune 'q'
	"\x05\xfc\x02\x00\x04\x60\x53\x48\x31\xe8\x00\x06" +

	// glyph 1107 for rune 'r'
	"\x05\xfc\x00\x01\x04\x6c\x08\x05" +

	// glyph 1115 for rune 's'
	"\x05\xfc\x00\x01\x04\x5c\x1c\xc0\xa8" +

	// glyph 1124 for rune 't'
	"\x05\xfa\x00\x01\x04\x04\x08\x2c\x08\xa1\x01" +

	// glyph 1135 for rune 'u'
	"\x05\xfc\x00\x00\x04\x14\x16\xe8" +

	// glyph 1143 for rune 'v'
	"\x05\xfc\x00\x00\x04\x14\x82\x08\x16\x30" +

	// glyph 1153 for rune 'w'
	"\x05\xfc\x00\x00\x05\x08\x88\x4c\xa4\x02\x33" +

	// glyph 1164 for rune 'x'
	"\x05\xfc\x00\x00\x04\x20\x42\x03\x19\x89" +

	// glyph 1174 for rune 'y'
	"\x05\xfc\x02\x00\x04\x14\x16\xe8\x00\x82\x0a" +

	// glyph 1185 for rune 'z'
	"\x05\xfc\x00\x01\x04\xd4\x20\x14\xac" +

	// glyph 1194 for rune '{'
	"\x05\xf9\x01\x01\x04\x40\x60\x20\xa1\x80\x04\x09" +

	// glyph 1206 for rune '|'
	"\x05\xfa\x01\x02\x03\x58\x15\x01" +

	// glyph 1214 for rune '}'
	"\x05\xf9\x01\x01\x04\x04\x24\x20\x30\xa0\x20\x61\x00" +

	// glyph 1227 for rune '~'
	"\x05\xfc\xfe\x00\x04\x20\x10\x0a" +

	"")
