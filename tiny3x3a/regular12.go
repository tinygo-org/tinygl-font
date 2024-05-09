// File generated using:
// 	go run ./generate -font=tiny3x3a/tiny3x3a.ttf -size=12 -dpi=72 -package=tiny3x3a -weight=Regular

package tiny3x3a

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      1192
// - glyph metadata:  475
// - glyph mask data: 632

var Regular12 = font.Make("" +
	"\x00" + // version: 0
	"\x0c" + // size:   12
	"\x0c" + // height: 12
	"\x09" + // ascent: 9

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xd6\x00" + // "\"" at index 214
	"\xdf\x00" + // "#" at index 223
	"\xeb\x00" + // "$" at index 235
	"\xf9\x00" + // "%" at index 249
	"\x07\x01" + // "&" at index 263
	"\x15\x01" + // "'" at index 277
	"\x1c\x01" + // "(" at index 284
	"\x28\x01" + // ")" at index 296
	"\x34\x01" + // "*" at index 308
	"\x3c\x01" + // "+" at index 316
	"\x4a\x01" + // "," at index 330
	"\x52\x01" + // "-" at index 338
	"\x5a\x01" + // "." at index 346
	"\x61\x01" + // "/" at index 353
	"\x6f\x01" + // "0" at index 367
	"\x7d\x01" + // "1" at index 381
	"\x88\x01" + // "2" at index 392
	"\x96\x01" + // "3" at index 406
	"\xa4\x01" + // "4" at index 420
	"\xeb\x00" + // "5" at index 235
	"\xb2\x01" + // "6" at index 434
	"\xbe\x01" + // "7" at index 446
	"\xca\x01" + // "8" at index 458
	"\xd8\x01" + // "9" at index 472
	"\xe4\x01" + // ":" at index 484
	"\xee\x01" + // ";" at index 494
	"\x1c\x01" + // "<" at index 284
	"\xf9\x01" + // "=" at index 505
	"\x28\x01" + // ">" at index 296
	"\x07\x02" + // "?" at index 519
	"\x13\x02" + // "@" at index 531
	"\x1f\x02" + // "A" at index 543
	"\x2d\x02" + // "B" at index 557
	"\x39\x02" + // "C" at index 569
	"\x47\x02" + // "D" at index 583
	"\x55\x02" + // "E" at index 597
	"\x63\x02" + // "F" at index 611
	"\x71\x02" + // "G" at index 625
	"\x7f\x02" + // "H" at index 639
	"\x8d\x02" + // "I" at index 653
	"\x95\x02" + // "J" at index 661
	"\xa1\x02" + // "K" at index 673
	"\xaf\x02" + // "L" at index 687
	"\xbb\x02" + // "M" at index 699
	"\xc7\x02" + // "N" at index 711
	"\xd3\x02" + // "O" at index 723
	"\xe1\x02" + // "P" at index 737
	"\xed\x02" + // "Q" at index 749
	"\xfb\x02" + // "R" at index 763
	"\xeb\x00" + // "S" at index 235
	"\x09\x03" + // "T" at index 777
	"\x15\x03" + // "U" at index 789
	"\x21\x03" + // "V" at index 801
	"\xdf\x00" + // "W" at index 223
	"\x2d\x03" + // "X" at index 813
	"\x3b\x03" + // "Y" at index 827
	"\x47\x03" + // "Z" at index 839
	"\x55\x03" + // "[" at index 853
	"\x61\x03" + // "\\" at index 865
	"\x6f\x03" + // "]" at index 879
	"\x7b\x03" + // "^" at index 891
	"\x86\x03" + // "_" at index 902
	"\x8e\x03" + // "`" at index 910
	"\x98\x03" + // "a" at index 920
	"\xa2\x03" + // "b" at index 930
	"\xad\x03" + // "c" at index 941
	"\xb5\x03" + // "d" at index 949
	"\xc0\x03" + // "e" at index 960
	"\xca\x03" + // "f" at index 970
	"\xd5\x03" + // "g" at index 981
	"\xa2\x03" + // "h" at index 930
	"\xde\x03" + // "i" at index 990
	"\xee\x01" + // "j" at index 494
	"\xe6\x03" + // "k" at index 998
	"\xf4\x03" + // "l" at index 1012
	"\xff\x03" + // "m" at index 1023
	"\x0a\x04" + // "n" at index 1034
	"\xad\x03" + // "o" at index 941
	"\x15\x04" + // "p" at index 1045
	"\x20\x04" + // "q" at index 1056
	"\x2b\x04" + // "r" at index 1067
	"\x35\x04" + // "s" at index 1077
	"\x07\x01" + // "t" at index 263
	"\x3f\x04" + // "u" at index 1087
	"\x4a\x04" + // "v" at index 1098
	"\x55\x04" + // "w" at index 1109
	"\xad\x03" + // "x" at index 941
	"\x60\x04" + // "y" at index 1120
	"\x6e\x04" + // "z" at index 1134
	"\x78\x04" + // "{" at index 1144
	"\x86\x04" + // "|" at index 1158
	"\x8f\x04" + // "}" at index 1167
	"\x9d\x04" + // "~" at index 1181

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for runes ' ', '5', '<', '>', 'S', 'W', 'h', 'j', 'o', 't', 'x'
	"\x05\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x09\xf7\xfd\x03\x06\xfc\x55\x01" +

	// glyph 214 for rune '"'
	"\x0c\xf7\xfd\x00\x09\xfc\xc0\x5f\x15" +

	// glyph 223 for rune '#'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\xfc\xff\x5f\x15" +

	// glyph 235 for rune '$'
	"\x0c\xf7\x00\x00\x09\x00\xff\x5f\x00\x3f\x50\xfc\x3f\x50" +

	// glyph 249 for rune '%'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\xfc\xff\x5f\x00\xff\x5f" +

	// glyph 263 for rune '&'
	"\x0c\xf7\x00\x00\x09\x00\x3f\x50\xfc\x3f\x50\x00\xff\x5f" +

	// glyph 277 for rune '\''
	"\x09\xf7\xfa\x03\x06\xfc\x05" +

	// glyph 284 for rune '('
	"\x09\xf7\x00\x00\x06\x00\x7f\xf1\x03\x05\xf0\x17" +

	// glyph 296 for rune ')'
	"\x0c\xf7\x00\x03\x09\xfc\x40\x01\xfc\xc5\x0f\x14" +

	// glyph 308 for rune '*'
	"\x0c\xf7\xfd\x03\x09\xfc\x7f\x55" +

	// glyph 316 for rune '+'
	"\x0c\xf7\x00\x00\x09\x00\x3f\x50\xfc\xff\x5f\x00\x3f\x50" +

	// glyph 330 for rune ','
	"\x06\xfd\x03\x00\x03\xfc\x55\x01" +

	// glyph 338 for rune '-'
	"\x0c\xfa\xfd\x00\x09\xfc\xff\x5f" +

	// glyph 346 for rune '.'
	"\x06\xfd\x00\x00\x03\xfc\x05" +

	// glyph 353 for rune '/'
	"\x0c\xf7\x00\x00\x09\x00\xc0\x5f\x00\x3f\x50\xfc\x00\x50" +

	// glyph 367 for rune '0'
	"\x0c\xf7\x00\x00\x09\x00\x3f\x50\xfc\xc0\x5f\x00\x3f\x50" +

	// glyph 381 for rune '1'
	"\x09\xf7\x00\x00\x06\xfc\x40\x01\xfc\x55\x01" +

	// glyph 392 for rune '2'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\x00\x3f\x50\x00\xff\x5f" +

	// glyph 406 for rune '3'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x00\xff\x5f\xfc\xff\x5f" +

	// glyph 420 for rune '4'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\xfc\xff\x5f\x00\xc0\x5f" +

	// glyph 434 for rune '6'
	"\x0c\xf7\x00\x00\x09\xfc\x00\x50\xfc\xff\x5f\x15" +

	// glyph 446 for rune '7'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x00\xc0\x5f\x15" +

	// glyph 458 for rune '8'
	"\x0c\xf7\x00\x00\x09\x00\xff\x5f\xfc\xff\x5f\xfc\x3f\x50" +

	// glyph 472 for rune '9'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x15\x00\xf0\x17" +

	// glyph 484 for rune ':'
	"\x09\xf7\x00\x03\x06\xfc\x05\x50\xfc\x05" +

	// glyph 494 for rune ';'
	"\x09\xfa\x03\x00\x06\x00\x7f\x55\xfc\x40\x01" +

	// glyph 505 for rune '='
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x00\x00\x50\xfc\xff\x5f" +

	// glyph 519 for rune '?'
	"\x0c\xf7\x00\x03\x09\xfc\x7f\x01\xfc\xc5\x0f\x14" +

	// glyph 531 for rune '@'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\x15\x00\xf0\x17" +

	// glyph 543 for rune 'A'
	"\x0c\xf7\x00\x00\x09\x00\x3f\x50\xfc\xff\x5f\xfc\xc0\x5f" +

	// glyph 557 for rune 'B'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\xfc\xff\x5f\x15" +

	// glyph 569 for rune 'C'
	"\x0c\xf7\x00\x00\x09\x00\xff\x5f\xfc\x00\x50\x00\xff\x5f" +

	// glyph 583 for rune 'D'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\xfc\xc0\x5f\xfc\x3f\x50" +

	// glyph 597 for rune 'E'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\xfc\x3f\x50\xfc\xff\x5f" +

	// glyph 611 for rune 'F'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\xfc\x3f\x50\xfc\x00\x50" +

	// glyph 625 for rune 'G'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\xfc\xc0\x5f\xfc\xff\x5f" +

	// glyph 639 for rune 'H'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\xfc\xff\x5f\xfc\xc0\x5f" +

	// glyph 653 for rune 'I'
	"\x09\xf7\x00\x03\x06\xfc\x55\x55" +

	// glyph 661 for rune 'J'
	"\x0c\xf7\x00\x00\x09\x00\xc0\x5f\x15\xff\x0f\x14" +

	// glyph 673 for rune 'K'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\xfc\x3f\x50\xfc\xc0\x5f" +

	// glyph 687 for rune 'L'
	"\x0c\xf7\x00\x00\x09\xfc\x00\x50\x15\xff\xff\x17" +

	// glyph 699 for rune 'M'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x15\x3f\xf0\x17" +

	// glyph 711 for rune 'N'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\xfc\xc0\x5f\x15" +

	// glyph 723 for rune 'O'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\xfc\xc0\x5f\xfc\xff\x5f" +

	// glyph 737 for rune 'P'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x15\x3f\x00\x14" +

	// glyph 749 for rune 'Q'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\xfc\xc0\x5f\xfc\x3f\x50" +

	// glyph 763 for rune 'R'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\xfc\xff\x5f\xfc\xc0\x5f" +

	// glyph 777 for rune 'T'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x00\x3f\x50\x15" +

	// glyph 789 for rune 'U'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\x15\xff\xff\x17" +

	// glyph 801 for rune 'V'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\x15\xc0\x0f\x14" +

	// glyph 813 for rune 'X'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\x00\x3f\x50\xfc\xc0\x5f" +

	// glyph 827 for rune 'Y'
	"\x0c\xf7\x00\x00\x09\xfc\xc0\x5f\x00\x3f\x50\x15" +

	// glyph 839 for rune 'Z'
	"\x0c\xf7\x00\x00\x09\xfc\xff\x5f\x00\x3f\x50\xfc\xff\x5f" +

	// glyph 853 for rune '['
	"\x09\xf7\x00\x00\x06\xfc\x7f\xf1\x03\xc5\xff\x17" +

	// glyph 865 for rune '\\'
	"\x0c\xf7\x00\x00\x09\xfc\x00\x50\x00\x3f\x50\x00\xc0\x5f" +

	// glyph 879 for rune ']'
	"\x0c\xf7\x00\x03\x09\xfc\x7f\x01\xfc\xc5\xff\x17" +

	// glyph 891 for rune '^'
	"\x0c\xf7\xfd\x00\x09\x00\x3f\x50\xfc\xc0\x5f" +

	// glyph 902 for rune '_'
	"\x0c\xfd\x00\x00\x09\xfc\xff\x5f" +

	// glyph 910 for rune '`'
	"\x09\xf7\xfd\x00\x06\xfc\x40\x01\xfc\x05" +

	// glyph 920 for rune 'a'
	"\x09\xfa\x00\x00\x06\x00\x7f\xf1\xff\x05" +

	// glyph 930 for rune 'b'
	"\x09\xf7\x00\x00\x06\xfc\x40\xf1\xff\x55\x01" +

	// glyph 941 for rune 'c'
	"\x09\xfa\x00\x00\x06\xfc\x7f\x55" +

	// glyph 949 for rune 'd'
	"\x09\xf7\x00\x00\x06\x00\x7f\xf1\xff\x55\x01" +

	// glyph 960 for rune 'e'
	"\x09\xfa\x00\x00\x06\xfc\x40\xf1\xff\x05" +

	// glyph 970 for rune 'f'
	"\x09\xf7\x00\x00\x06\x00\x7f\xf1\x03\x55\x01" +

	// glyph 981 for rune 'g'
	"\x09\xfa\x03\x00\x06\xfc\x7f\x55\x15" +

	// glyph 990 for rune 'i'
	"\x06\xfa\x00\x00\x03\xfc\x55\x01" +

	// glyph 998 for rune 'k'
	"\x0c\xf7\x00\x00\x09\xfc\x00\x50\xfc\x3f\x50\xfc\xc0\x5f" +

	// glyph 1012 for rune 'l'
	"\x09\xf7\x00\x00\x06\xfc\x40\x55\x00\x7f\x01" +

	// glyph 1023 for rune 'm'
	"\x0c\xfa\x00\x00\x09\xfc\x3f\x50\xfc\xff\x5f" +

	// glyph 1034 for rune 'n'
	"\x0c\xfa\x00\x00\x09\xfc\x3f\x50\xfc\xc0\x5f" +

	// glyph 1045 for rune 'p'
	"\x09\xfa\x03\x00\x06\xfc\x7f\x55\xfc\x40\x01" +

	// glyph 1056 for rune 'q'
	"\x09\xfa\x03\x00\x06\xfc\x7f\x55\x00\x7f\x01" +

	// glyph 1067 for rune 'r'
	"\x09\xfa\x00\x00\x06\xfc\x7f\xf1\x03\x05" +

	// glyph 1077 for rune 's'
	"\x09\xfa\x00\x00\x06\x00\x7f\xf1\x03\x05" +

	// glyph 1087 for rune 'u'
	"\x0c\xfa\x00\x00\x09\xfc\xc0\x5f\xfc\xff\x5f" +

	// glyph 1098 for rune 'v'
	"\x0c\xfa\x00\x00\x09\xfc\xc0\x5f\x00\x3f\x50" +

	// glyph 1109 for rune 'w'
	"\x0c\xfa\x00\x00\x09\xfc\xff\x5f\x00\xff\x5f" +

	// glyph 1120 for rune 'y'
	"\x0c\xfa\x03\x00\x09\xfc\xc0\x5f\x00\xff\x5f\x00\xc0\x5f" +

	// glyph 1134 for rune 'z'
	"\x09\xfa\x00\x00\x06\xfc\x40\x01\xfc\x05" +

	// glyph 1144 for rune '{'
	"\x0c\xf7\x00\x00\x09\x00\xff\x5f\xfc\x3f\x50\x00\xff\x5f" +

	// glyph 1158 for rune '|'
	"\x09\xf7\x03\x03\x06\xfc\x55\x55\x15" +

	// glyph 1167 for rune '}'
	"\x0c\xf7\x00\x00\x09\xfc\x3f\x50\x00\xff\x5f\xfc\x3f\x50" +

	// glyph 1181 for rune '~'
	"\x0c\xf7\xfd\x00\x09\x00\xff\x5f\xfc\x3f\x50" +

	"")
