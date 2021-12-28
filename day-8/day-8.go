package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
	"regexp"
	"strings"
)

const UnknownDigit = 0xf

// map segment (wire line) to bit position

var wireMap map[rune]uint8

func init() {
	wireMap = make(map[rune]uint8)

	wireMap['a'] = 1
	wireMap['b'] = 2
	wireMap['c'] = 3
	wireMap['d'] = 4
	wireMap['e'] = 5
	wireMap['f'] = 6
	wireMap['g'] = 7

	/*  unscrambled segment positions

	bit:     7654321
	segment: gfedcba

	           aaaa
	          b    c
	          b    c
	           dddd
	          e    f
	          e    f
	           gggg
	*/

}

func main() {
	// part 1 result
	var digitCount uint32 = 0
	// part 2 result
	var totalOutput uint32 = 0
	// map signal to digit
	var signalMap map[uint8]uint8
	// map digit to signal
	var reverseSignalMap map[uint8]uint8
	// map segment to signal bit
	var segmentMap map[rune]uint8

	fileInput, err := ioutil.ReadFile("day-8-input.txt")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	re, err := regexp.Compile("[||\n]")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	inputStr := re.Split(string(fileInput), -1)

	for i := range inputStr {
		// check for end of input
		if inputStr[i] == "" {
			break
		}
		fmt.Printf("%v %v\n", i, inputStr[i])
		if i%2 == 0 {
			// clear maps
			signalMap = make(map[uint8]uint8)
			reverseSignalMap = make(map[uint8]uint8)
			segmentMap = make(map[rune]uint8)

			// process input singals
			for _, signalPattern := range strings.Fields(inputStr[i]) {
				bitMap := uint8(0)
				for _, signal := range signalPattern {
					bitMap |= 1 << wireMap[signal]
				}
				signalMap[bitMap] = UnknownDigit
				switch len(signalPattern) {
				case 2:
					signalMap[bitMap] = 1
					reverseSignalMap[1] = bitMap
				case 3:
					signalMap[bitMap] = 7
					reverseSignalMap[7] = bitMap
				case 4:
					signalMap[bitMap] = 4
					reverseSignalMap[4] = bitMap
				case 7:
					signalMap[bitMap] = 8
					reverseSignalMap[8] = bitMap
				}
			}
			// dump known digits [1|4|7|8]
			for key, value := range signalMap {
				if value != UnknownDigit {
					fmt.Printf("sig %08b: %v\n", key, value)
				}
			}
			// find segment 'a'
			segmentMap['a'] = reverseSignalMap[7] ^ reverseSignalMap[1]
			fmt.Printf("segmentMap['a'] = %08b\n", segmentMap['a'])

			// find segment 'c' and signal 6
			for key := range signalMap {
				// find candidates for 6
				if bits.OnesCount8(key) == 6 {
					value := (reverseSignalMap[8] ^ key) & reverseSignalMap[1]
					if value != 0 {
						// found 6
						segmentMap['c'] = value
						fmt.Printf("segmentMap['c'] = %08b\n", segmentMap['c'])
						signalMap[key] = 6
						reverseSignalMap[6] = key
						fmt.Printf("sig %08b: 6\n", key)
					}
				}
			}

			// find segment 'f'
			segmentMap['f'] = reverseSignalMap[1] ^ segmentMap['c']
			fmt.Printf("segmentMap['f'] = %08b\n", segmentMap['f'])

			// find segment 'g' and signal 9
			for key := range signalMap {
				// find candidates for 9
				if bits.OnesCount8(key) == 6 {
					// found 9
					value := key ^ reverseSignalMap[4] ^ reverseSignalMap[1] ^ reverseSignalMap[7]

					if bits.OnesCount8(value) == 1 {
						segmentMap['g'] = value
						fmt.Printf("segmentMap['g'] = %08b\n", segmentMap['g'])
						signalMap[key] = 9
						reverseSignalMap[9] = key
						fmt.Printf("sig %08b: 9\n", key)
					}
				}
			}

			// find segment 'e'
			segmentMap['e'] = reverseSignalMap[8] ^ reverseSignalMap[9]
			fmt.Printf("segmentMap['e'] = %08b\n", segmentMap['e'])

			// find segment 'd' and signal 0
			for key := range signalMap {
				// find candidates for 0
				if bits.OnesCount8(key) == 6 {
					if key != reverseSignalMap[9] && key != reverseSignalMap[6] {
						// found 0
						reverseSignalMap[0] = key
						signalMap[key] = 0
						fmt.Printf("sig %08b: 0\n", key)
						segmentMap['d'] = reverseSignalMap[8] ^ reverseSignalMap[0]
						fmt.Printf("segmentMap['d'] = %08b\n", segmentMap['d'])
					}
				}
			}

			// find segement 'b'
			segmentMap['b'] = segmentMap['d'] ^ reverseSignalMap[4] ^ reverseSignalMap[1]
			fmt.Printf("segmentMap['b'] = %08b\n", segmentMap['b'])

			// fill in rest of signal map (not setting reverse since they are not used)

			//signal 2
			signalMap[segmentMap['a']|segmentMap['c']|segmentMap['d']|segmentMap['e']|segmentMap['g']] = 2
			fmt.Printf("sig %08b: 2\n", segmentMap['a']|segmentMap['c']|segmentMap['d']|segmentMap['e']|segmentMap['g'])
			// signal 3
			signalMap[segmentMap['a']|segmentMap['c']|segmentMap['d']|segmentMap['f']|segmentMap['g']] = 3
			fmt.Printf("sig %08b: 3\n", segmentMap['a']|segmentMap['c']|segmentMap['d']|segmentMap['f']|segmentMap['g'])
			// signal 4
			signalMap[segmentMap['a']|segmentMap['b']|segmentMap['d']|segmentMap['f']|segmentMap['g']] = 5
			fmt.Printf("sig %08b: 5\n", segmentMap['a']|segmentMap['b']|segmentMap['d']|segmentMap['f']|segmentMap['g'])

		} else {
			// process output
			var output uint32 = 0

			for _, signalPattern := range strings.Fields(inputStr[i]) {
				bitMap := uint8(0)
				for _, signal := range signalPattern {
					bitMap |= 1 << wireMap[signal]
				}
				if value, ok := signalMap[bitMap]; ok {
					output = (output * 10) + uint32(value)
					// part 1 count outputs [1|4|7|8]
					switch signalMap[bitMap] {
					case 1:
						fallthrough
					case 4:
						fallthrough
					case 7:
						fallthrough
					case 8:
						digitCount++
					}
					fmt.Printf("** output digit \"%7s\" %08b = %v\n", signalPattern, bitMap, value)
				} else {
					log.Fatalf("error interpreting signal output %v\n", signalPattern)
				}
			}
			fmt.Printf("*** output value %v\n", output)
			totalOutput += output
		}
	}
	fmt.Printf("part 1: [1|4|7|8] count %v\n", digitCount)
	fmt.Printf("part 2: total output %v\n", totalOutput)
}
