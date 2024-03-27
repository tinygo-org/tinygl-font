// File generated using:
// 	go run ./generate -font=freemono/FreeMono.ttf -size=16 -dpi=72 -package=freemono

package freemono

import (
	"tinygo.org/x/tinygl-font"
)

// Font statistics:
// - total size:      1981
// - glyph metadata:  475
// - glyph mask data: 1305

var Regular16 = font.Make("" +
	"\x00" + // version: 0
	"\x10" + // size:   16
	"\x12" + // height: 18
	"\x0d" + // ascent: 13

	// Runes 32..126 (95 runes)
	"\x5f\x00" + // number of runes (95)
	"\x20\x00\x00" + // start rune (32)
	"\xc9\x00" + // " " at index 201
	"\xce\x00" + // "!" at index 206
	"\xd9\x00" + // "\"" at index 217
	"\xe6\x00" + // "#" at index 230
	"\xfa\x00" + // "$" at index 250
	"\x11\x01" + // "%" at index 273
	"\x25\x01" + // "&" at index 293
	"\x38\x01" + // "'" at index 312
	"\x41\x01" + // "(" at index 321
	"\x4d\x01" + // ")" at index 333
	"\x5d\x01" + // "*" at index 349
	"\x6d\x01" + // "+" at index 365
	"\x79\x01" + // "," at index 377
	"\x83\x01" + // "-" at index 387
	"\x8a\x01" + // "." at index 394
	"\x91\x01" + // "/" at index 401
	"\xa5\x01" + // "0" at index 421
	"\xb9\x01" + // "1" at index 441
	"\xc8\x01" + // "2" at index 456
	"\xe1\x01" + // "3" at index 481
	"\xf5\x01" + // "4" at index 501
	"\x0a\x02" + // "5" at index 522
	"\x19\x02" + // "6" at index 537
	"\x2e\x02" + // "7" at index 558
	"\x42\x02" + // "8" at index 578
	"\x56\x02" + // "9" at index 598
	"\x6a\x02" + // ":" at index 618
	"\x76\x02" + // ";" at index 630
	"\x87\x02" + // "<" at index 647
	"\x9a\x02" + // "=" at index 666
	"\xa6\x02" + // ">" at index 678
	"\xb9\x02" + // "?" at index 697
	"\xce\x02" + // "@" at index 718
	"\xe9\x02" + // "A" at index 745
	"\x05\x03" + // "B" at index 773
	"\x1d\x03" + // "C" at index 797
	"\x2f\x03" + // "D" at index 815
	"\x49\x03" + // "E" at index 841
	"\x63\x03" + // "F" at index 867
	"\x7b\x03" + // "G" at index 891
	"\x91\x03" + // "H" at index 913
	"\xa3\x03" + // "I" at index 931
	"\xaf\x03" + // "J" at index 943
	"\xc1\x03" + // "K" at index 961
	"\xdb\x03" + // "L" at index 987
	"\xed\x03" + // "M" at index 1005
	"\x07\x04" + // "N" at index 1031
	"\x21\x04" + // "O" at index 1057
	"\x37\x04" + // "P" at index 1079
	"\x49\x04" + // "Q" at index 1097
	"\x63\x04" + // "R" at index 1123
	"\x7b\x04" + // "S" at index 1147
	"\x92\x04" + // "T" at index 1170
	"\xa2\x04" + // "U" at index 1186
	"\xb2\x04" + // "V" at index 1202
	"\xce\x04" + // "W" at index 1230
	"\xea\x04" + // "X" at index 1258
	"\x04\x05" + // "Y" at index 1284
	"\x1a\x05" + // "Z" at index 1306
	"\x31\x05" + // "[" at index 1329
	"\x3c\x05" + // "\\" at index 1340
	"\x53\x05" + // "]" at index 1363
	"\x5d\x05" + // "^" at index 1373
	"\x69\x05" + // "_" at index 1385
	"\x71\x05" + // "`" at index 1393
	"\x78\x05" + // "a" at index 1400
	"\x8d\x05" + // "b" at index 1421
	"\xa9\x05" + // "c" at index 1449
	"\xbc\x05" + // "d" at index 1468
	"\xd6\x05" + // "e" at index 1494
	"\xe9\x05" + // "f" at index 1513
	"\xfa\x05" + // "g" at index 1530
	"\x12\x06" + // "h" at index 1554
	"\x26\x06" + // "i" at index 1574
	"\x37\x06" + // "j" at index 1591
	"\x4a\x06" + // "k" at index 1610
	"\x64\x06" + // "l" at index 1636
	"\x70\x06" + // "m" at index 1648
	"\x80\x06" + // "n" at index 1664
	"\x8f\x06" + // "o" at index 1679
	"\xa2\x06" + // "p" at index 1698
	"\xbe\x06" + // "q" at index 1726
	"\xd8\x06" + // "r" at index 1752
	"\xe9\x06" + // "s" at index 1769
	"\xfc\x06" + // "t" at index 1788
	"\x0b\x07" + // "u" at index 1803
	"\x1a\x07" + // "v" at index 1818
	"\x2f\x07" + // "w" at index 1839
	"\x44\x07" + // "x" at index 1860
	"\x59\x07" + // "y" at index 1881
	"\x75\x07" + // "z" at index 1909
	"\x87\x07" + // "{" at index 1927
	"\x99\x07" + // "|" at index 1945
	"\xa2\x07" + // "}" at index 1954
	"\xb2\x07" + // "~" at index 1970

	// mark the end of the rune tables
	"\x00\x00" +

	// glyph 201 for rune ' '
	"\x0a\x00\x00\x00\x00" +

	// glyph 206 for rune '!'
	"\x0a\xf6\x00\x04\x06\x18\x87\x85\x04\x44\xb0" +

	// glyph 217 for rune '"'
	"\x0a\xf6\xfb\x02\x07\x24\x89\xe3\x34\x0d\xd3\x20\x08" +

	// glyph 230 for rune '#'
	"\x0a\xf6\x01\x02\x08\x50\x08\x14\x01\x55\x90\x69\x64\x16\x88\x80\xbb\x81\x08\x15" +

	// glyph 250 for rune '$'
	"\x0a\xf6\x01\x02\x08\x80\x00\xa8\x86\x00\x21\x00\xa0\x01\x40\x06\x00\x22\x80\xa8\x0a\x20\x10" +

	// glyph 273 for rune '%'
	"\x0a\xf6\x00\x02\x08\xa0\x00\x82\x10\x6a\x00\x90\x92\x06\x00\x0a\x20\x18\xa0\x02" +

	// glyph 293 for rune '&'
	"\x0a\xf8\x00\x02\x08\x70\x02\x05\x00\x02\xc0\x00\x88\x18\x51\x82\x70\xa0\x7a" +

	// glyph 312 for rune '\''
	"\x0a\xf6\xfb\x04\x06\x28\x1b\x87\x00" +

	// glyph 321 for rune '('
	"\x0a\xf7\x02\x05\x07\x30\xc5\x14\xc2\x84\xc1\x20" +

	// glyph 333 for rune ')'
	"\x0a\xf6\x02\x02\x05\x04\x24\x30\x50\x90\xc0\x01\x42\x82\x44\x00" +

	// glyph 349 for rune '*'
	"\x0a\xf6\xfc\x02\x08\x40\x00\x20\x80\x59\x41\x0b\x50\x02\x08\x02" +

	// glyph 365 for rune '+'
	"\x0a\xf8\xff\x01\x08\x00\x02\x45\xba\x0a\x20\x50" +

	// glyph 377 for rune ','
	"\x0a\xfd\x02\x02\x05\x40\xf0\x70\x34\x18" +

	// glyph 387 for rune '-'
	"\x0a\xfb\xfc\x01\x08\xa4\xaa" +

	// glyph 394 for rune '.'
	"\x0a\xfe\x00\x04\x06\x2c\x0f" +

	// glyph 401 for rune '/'
	"\x0a\xf5\x01\x02\x08\x00\x10\x00\x05\x80\x00\x10\x00\x42\x80\x40\x20\x40\x08\x40" +

	// glyph 421 for rune '0'
	"\x0a\xf6\x00\x02\x08\x90\x02\x05\x82\x00\x21\x80\x04\x60\x08\x60\x08\x08\xa8\x01" +

	// glyph 441 for rune '1'
	"\x0a\xf6\x00\x02\x08\x80\x00\x2c\x80\x08\x00\x02\x55\xa1\x6b" +

	// glyph 456 for rune '2'
	"\x0a\xf6\x00\x01\x08\x80\x0a\x30\x20\x00\x40\x00\x50\x00\x20\x00\x08\x00\x02\x80\x00\x20\x40\xa8\xaa" +

	// glyph 481 for rune '3'
	"\x0a\xf6\x00\x02\x08\xa0\x06\x02\x06\x00\x02\x50\x80\x07\x00\x02\x00\x16\x9a\x02" +

	// glyph 501 for rune '4'
	"\x0a\xf6\x00\x02\x08\x00\x06\x50\x02\x88\x80\x20\x21\x20\x58\x0d\x55\x03\x80\x00\x79" +

	// glyph 522 for rune '5'
	"\x0a\xf6\x00\x02\x08\xa4\x0a\x02\x50\x68\x09\x00\x58\xa1\x29" +

	// glyph 537 for rune '6'
	"\x0a\xf6\x00\x02\x08\x00\x2a\x24\x00\x02\x10\x00\x88\x0a\x0a\x88\x00\x06\x02\x42\x26" +

	// glyph 558 for rune '7'
	"\x0a\xf6\x00\x02\x08\xa8\x1a\x01\x08\x40\x01\x20\x01\x14\x00\x42\x00\x01\x10\x00" +

	// glyph 578 for rune '8'
	"\x0a\xf6\x00\x02\x08\xa0\x06\x02\x82\x00\x22\x40\x60\x49\x08\x10\x02\x18\x98\x02" +

	// glyph 598 for rune '9'
	"\x0a\xf6\x00\x02\x08\x90\x06\x05\x86\x00\x86\x40\x83\x9a\x00\x60\x00\x08\xa5\x00" +

	// glyph 618 for rune ':'
	"\x0a\xf9\x00\x03\x06\x60\xf4\x10\x00\xc1\xc2\x03" +

	// glyph 630 for rune ';'
	"\x0a\xf9\x02\x02\x06\x80\x80\x07\x04\x00\x40\xc0\x03\x0b\x0d\x18\x00" +

	// glyph 647 for rune '<'
	"\x0a\xf8\xff\x01\x08\x00\x90\x00\x19\x80\x01\x34\x00\x90\x01\x00\x19\x00\x90" +

	// glyph 666 for rune '='
	"\x0a\xfa\xfd\x01\x09\xa8\xaa\x01\x00\x80\xaa\x1a" +

	// glyph 678 for rune '>'
	"\x0a\xf8\xff\x01\x08\x24\x00\x40\x02\x00\x28\x00\xd0\x00\x28\x40\x02\x64\x00" +

	// glyph 697 for rune '?'
	"\x0a\xf7\x00\x02\x08\xa8\x0a\x02\x08\x00\x02\x60\x40\x06\x10\x00\x00\x00\x01\xc0\x03" +

	// glyph 718 for rune '@'
	"\x0a\xf6\x01\x01\x08\x80\x1a\x60\x20\x20\x50\x10\x64\x14\x56\x14\x52\x14\x53\x10\x68\x20\x00\x60\x00\x80\x2a" +

	// glyph 745 for rune 'A'
	"\x0a\xf7\x00\x00\x09\x80\x1e\x00\x50\x02\x00\x12\x00\x20\x08\x40\x80\x00\xac\x1e\x40\x00\x02\x02\x20\xb4\x40\x0f" +

	// glyph 773 for rune 'B'
	"\x0a\xf7\x00\x01\x09\xb8\x2a\x80\x00\x12\x08\x14\xb0\x3a\x80\x00\x05\x02\x20\x08\x50\xb8\xba\x00" +

	// glyph 797 for rune 'C'
	"\x0a\xf7\x00\x01\x09\x90\x6a\x80\x00\x86\x00\x40\x15\x08\x50\x80\x66\x00" +

	// glyph 815 for rune 'D'
	"\x0a\xf7\x00\x01\x09\xbc\x1a\x50\x40\x41\x01\x08\x05\x10\x14\x40\x51\x00\x41\x01\x08\x05\x14\xbc\x1a\x00" +

	// glyph 841 for rune 'E'
	"\x0a\xf7\x00\x01\x09\xb8\xea\x80\x00\x02\x02\x04\x08\x02\xb0\x0a\x80\x10\x00\x02\x04\x08\x50\xb8\xea\x01" +

	// glyph 867 for rune 'F'
	"\x0a\xf7\x00\x01\x09\xb8\xea\x81\x00\x05\x02\x00\x08\x02\xb0\x0a\x80\x10\x00\x02\x40\xb8\x02\x00" +

	// glyph 891 for rune 'G'
	"\x0a\xf7\x00\x01\x09\x90\xaa\x80\x00\x86\x00\x40\x21\xa0\x8b\x00\x14\x08\x50\x90\xa6\x00" +

	// glyph 913 for rune 'H'
	"\x0a\xf7\x00\x01\x09\x74\xe0\x80\x00\x52\xb0\xaa\x80\x00\x52\x78\xe0\x01" +

	// glyph 931 for rune 'I'
	"\x0a\xf7\x00\x02\x08\xe8\x1a\x20\x50\x55\xe8\x1a" +

	// glyph 943 for rune 'J'
	"\x0a\xf7\x00\x01\x09\x00\xea\x02\x40\x51\x21\x40\x11\x02\x08\xa0\x19\x00" +

	// glyph 961 for rune 'K'
	"\x0a\xf7\x00\x01\x09\xb8\xe0\x81\x80\x00\x82\x00\x88\x00\xb0\x0a\xc0\x50\x00\x02\x02\x08\x14\xb8\xc0\x02" +

	// glyph 987 for rune 'L'
	"\x0a\xf7\x00\x01\x09\xe4\x02\x00\x02\x50\x01\x02\x04\x08\x60\xe8\xaa\x02" +

	// glyph 1005 for rune 'M'
	"\x0a\xf7\x00\x00\x09\xb4\x00\x0f\x0a\x94\x20\x81\x08\x22\x84\x20\x25\x08\x82\x82\x20\x00\x18\x2d\x90\x03" +

	// glyph 1031 for rune 'N'
	"\x0a\xf7\x00\x00\x09\xf4\x80\x0b\x19\x20\x50\x02\x02\x85\x20\x41\x81\x08\x14\x94\x40\x01\x0a\x7c\xc0\x00" +

	// glyph 1057 for rune 'O'
	"\x0a\xf7\x00\x01\x09\x90\x2a\x80\x00\x82\x00\x14\x02\x80\x85\x00\x14\x08\x20\x90\x29\x00" +

	// glyph 1079 for rune 'P'
	"\x0a\xf7\x00\x01\x08\xb8\x2a\x20\x80\x05\x57\x02\x16\x00\x02\x10\xae\x00" +

	// glyph 1097 for rune 'Q'
	"\x0a\xf7\x02\x01\x09\x90\x2a\x80\x00\x82\x00\x14\x02\x80\x85\x00\x14\x09\x20\x90\x29\x00\x1a\x00\x59\x1a" +

	// glyph 1123 for rune 'R'
	"\x0a\xf7\x00\x01\x09\xb8\x2a\x80\x00\x12\x08\x24\xb0\x1e\x80\xa0\x00\x02\x05\x08\x20\xb8\x00\x03" +

	// glyph 1147 for rune 'S'
	"\x0a\xf7\x00\x01\x08\x90\xa9\x20\xc0\x20\x00\x60\x00\x40\x2a\x00\x80\x04\x80\x18\xc0\x94\x29" +

	// glyph 1170 for rune 'T'
	"\x0a\xf7\x00\x01\x09\xac\xab\x21\x08\x04\x20\x40\x55\x90\x2b\x00" +

	// glyph 1186 for rune 'U'
	"\x0a\xf7\x00\x01\x09\x6c\xe0\x51\x00\x52\x15\x08\x20\x90\x29\x00" +

	// glyph 1202 for rune 'V'
	"\x0a\xf7\x00\x00\x09\xb4\x40\x0e\x02\x10\x40\x00\x02\x08\x14\x40\x80\x00\x20\x08\x00\x12\x00\x50\x02\x00\x1c\x00" +

	// glyph 1230 for rune 'W'
	"\x0a\xf7\x00\x00\x09\xb4\x40\x0e\x02\x80\x20\x1c\x04\x82\x52\x20\x21\x01\x21\x25\x50\x82\x02\x18\x28\xc0\x80\x02" +

	// glyph 1258 for rune 'X'
	"\x0a\xf7\x00\x01\x09\x38\xd0\x81\x40\x01\x08\x02\x90\x02\x00\x07\x00\x25\x00\x09\x02\x08\x24\x7c\xe0\x02" +

	// glyph 1284 for rune 'Y'
	"\x0a\xf7\x00\x01\x09\x38\xd0\x81\x40\x01\x08\x02\x50\x02\x00\x06\x00\x04\x50\x90\x2b\x00" +

	// glyph 1306 for rune 'Z'
	"\x0a\xf7\x00\x01\x08\xb0\xaa\x20\x20\x10\x14\x00\x08\x00\x02\x80\x01\x90\x80\x20\x80\xb4\xea" +

	// glyph 1329 for rune '['
	"\x0a\xf6\x02\x04\x07\xa4\x08\x55\x55\xa1\x02" +

	// glyph 1340 for rune '\\'
	"\x0a\xf5\x01\x02\x08\x04\x00\x02\x40\x01\x80\x00\x50\x00\x20\x00\x14\x00\x08\x01\x20\x01\x80" +

	// glyph 1363 for rune ']'
	"\x0a\xf6\x02\x03\x05\x28\x58\x55\x15\x0e" +

	// glyph 1373 for rune '^'
	"\x0a\xf6\xfa\x02\x08\x40\x00\xa4\x00\x82\x20\x50" +

	// glyph 1385 for rune '_'
	"\x0a\x01\x02\x00\x0a\xfc\xff\x1f" +

	// glyph 1393 for rune '`'
	"\x0a\xf6\xf8\x03\x05\x08\x08" +

	// glyph 1400 for rune 'a'
	"\x0a\xf9\x00\x01\x09\x90\x1a\x00\x80\x00\x00\x05\xa8\x1a\x08\x50\x20\x80\x41\xa6\x1d" +

	// glyph 1421 for rune 'b'
	"\x0a\xf6\x00\x00\x09\x24\x00\x00\x05\x00\x41\xa5\x02\x64\x80\x40\x02\x14\x14\x00\x42\x01\x10\x34\x80\xd0\x69\x06" +

	// glyph 1449 for rune 'c'
	"\x0a\xf9\x00\x01\x09\x80\x6a\x80\x00\x43\x01\x04\x02\x00\x81\x00\x05\x59\x0a" +

	// glyph 1468 for rune 'd'
	"\x0a\xf6\x00\x01\x09\x00\x90\x00\x00\x12\xa0\x22\x20\xe0\x20\x00\x83\x00\x08\x02\x30\x14\xd0\x80\x66\x0a" +

	// glyph 1494 for rune 'e'
	"\x0a\xf9\x00\x01\x08\x80\x1a\x20\x90\x08\x80\xac\xaa\x08\x00\x14\x00\x90\xa5" +

	// glyph 1513 for rune 'f'
	"\x0a\xf6\x00\x02\x08\x00\x2a\x20\x00\x05\xe0\x6b\x50\x40\x15\xbe\x06" +

	// glyph 1530 for rune 'g'
	"\x0a\xf9\x03\x01\x09\x90\x8a\x92\x80\x82\x00\x48\x91\x80\x02\xa8\x08\x00\x20\x00\x50\x00\x6a\x00" +

	// glyph 1554 for rune 'h'
	"\x0a\xf6\x00\x01\x09\x18\x00\x80\x00\x10\x98\x06\x70\x50\x80\x00\x52\xe1\x81\x07" +

	// glyph 1574 for rune 'i'
	"\x0a\xf6\x00\x02\x08\xc0\x00\x20\x00\x00\x90\x02\x80\x40\x15\xba\x0a" +

	// glyph 1591 for rune 'j'
	"\x0a\xf6\x03\x02\x08\x00\x09\x00\x01\x00\x90\x2a\x00\x54\x55\x01\x20\xa4\x06" +

	// glyph 1610 for rune 'k'
	"\x0a\xf6\x00\x01\x09\x24\x00\x80\x00\x10\x08\x1a\x20\x08\x80\x08\x00\x2a\x00\x48\x02\x20\x24\xe0\xc0\x07" +

	// glyph 1636 for rune 'l'
	"\x0a\xf6\x00\x02\x08\xa4\x00\x20\x50\x55\xa1\xab" +

	// glyph 1648 for rune 'm'
	"\x0a\xf9\x00\x00\x09\xa4\x86\x02\x87\x52\x20\x04\x58\xd1\x91\x30" +

	// glyph 1664 for rune 'n'
	"\x0a\xf9\x00\x01\x09\x54\x2a\xc0\x41\x01\x02\x48\x85\x07\x1d" +

	// glyph 1679 for rune 'o'
	"\x0a\xf9\x00\x01\x09\x80\x1a\x80\x40\x82\x00\x08\x02\x50\x91\x00\x02\x99\x02" +

	// glyph 1698 for rune 'p'
	"\x0a\xf9\x03\x00\x09\x24\xa9\x00\x1d\x20\x90\x00\x04\x05\x80\x90\x00\x05\x19\x20\x50\xa9\x00\x05\x00\xd1\x0a\x00" +

	// glyph 1726 for rune 'q'
	"\x0a\xf9\x03\x01\x09\x80\x4a\x92\x80\x83\x00\x0c\x02\x20\x08\xc0\x80\x80\x03\xa8\x08\x00\x20\x01\x90\x0b" +

	// glyph 1752 for rune 'r'
	"\x0a\xf9\x00\x01\x09\xa0\xa4\x00\x16\x04\x18\x00\x20\x00\x45\xae\x02" +

	// glyph 1769 for rune 's'
	"\x0a\xf9\x00\x01\x08\x80\x5a\x30\x90\x20\x00\x40\x1a\x00\x80\x14\x80\xa4\x69" +

	// glyph 1788 for rune 't'
	"\x0a\xf7\x00\x01\x08\x20\x00\xe1\x6a\x80\x00\x54\x01\x96\x02" +

	// glyph 1803 for rune 'u'
	"\x0a\xf9\x00\x01\x09\x18\x64\x80\x00\x52\x81\x40\x02\xaa\x1d" +

	// glyph 1818 for rune 'v'
	"\x0a\xf9\x00\x01\x09\x68\xa0\x82\x00\x02\x02\x05\x14\x08\x80\x10\x00\x25\x00\xb0\x00" +

	// glyph 1839 for rune 'w'
	"\x0a\xf9\x00\x01\x09\x28\x80\x22\x00\x45\x70\x08\x84\x22\x60\x88\x80\x92\x01\x0a\x03" +

	// glyph 1860 for rune 'x'
	"\x0a\xf9\x00\x01\x09\x64\xa0\x80\x80\x01\x98\x01\xc0\x02\x80\x19\x80\x80\xc1\x07\x1e" +

	// glyph 1881 for rune 'y'
	"\x0a\xf9\x03\x01\x09\x28\x90\x81\x00\x02\x02\x05\x10\x08\x80\x14\x00\x24\x00\x60\x00\x80\x00\x00\x01\xa0\x0b\x00" +

	// glyph 1909 for rune 'z'
	"\x0a\xf9\x00\x02\x08\xa8\x1a\x02\x02\x20\x00\x06\x50\x00\x05\xc8\xaa\x03" +

	// glyph 1927 for rune '{'
	"\x0a\xf6\x02\x03\x07\x40\x41\x01\x41\x20\x90\x80\x01\x08\x10\x04\x05\x60" +

	// glyph 1945 for rune '|'
	"\x0a\xf7\x02\x04\x05\x58\x55\x15\x01" +

	// glyph 1954 for rune '}'
	"\x0a\xf6\x02\x03\x07\x18\xc0\x00\x42\x01\x02\x28\x04\x20\x14\x0a" +

	// glyph 1970 for rune '~'
	"\x0a\xfa\xfd\x02\x08\x10\x00\x62\x08\x90\x00" +

	"")
