package main

import (
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
		log.Printf("%v %v\n", i, inputStr[i])
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
					log.Printf("sig %08b: %v\n", key, value)
				}
			}
			// find segment 'a'
			segmentMap['a'] = reverseSignalMap[7] ^ reverseSignalMap[1]
			log.Printf("segmentMap['a'] = %08b\n", segmentMap['a'])

			// find segment 'c' and signal 6
			for key := range signalMap {
				// find candidates for 6
				if bits.OnesCount8(key) == 6 {
					value := (reverseSignalMap[8] ^ key) & reverseSignalMap[1]
					if value != 0 {
						// found 6
						segmentMap['c'] = value
						log.Printf("segmentMap['c'] = %08b\n", segmentMap['c'])
						signalMap[key] = 6
						reverseSignalMap[6] = key
						log.Printf("sig %08b: 6\n", signalMap[6])
					}
				}
			}

			// find segment 'f'
			segmentMap['f'] = reverseSignalMap[1] ^ segmentMap['c']
			log.Printf("segmentMap['f'] = %08b\n", segmentMap['f'])

			// find segment 'g' and signal 9
			for key := range signalMap {
				// find candidates for 9
				if bits.OnesCount8(key) == 6 {
					// found 9
					value := key ^ reverseSignalMap[4] ^ reverseSignalMap[1] ^ reverseSignalMap[7]

					if bits.OnesCount8(value) == 1 {
						segmentMap['g'] = value
						log.Printf("segmentMap['g'] = %08b\n", segmentMap['g'])
						signalMap[key] = 9
						reverseSignalMap[9] = key
						log.Printf("sig %08b: 9\n", signalMap[9])
					}
				}
			}

			// find segment 'e'
			segmentMap['e'] = reverseSignalMap[8] ^ reverseSignalMap[9]
			log.Printf("segmentMap['e'] = %08b\n", segmentMap['e'])

			// find segment 'd' and signal 0
			for key := range signalMap {
				// find candidates for 0
				if bits.OnesCount8(key) == 6 {
					if key != reverseSignalMap[9] && key != reverseSignalMap[6] {
						// found 0
						reverseSignalMap[0] = key
						signalMap[key] = 0
						log.Printf("sig %08b: 0\n", signalMap[0])
						segmentMap['d'] = reverseSignalMap[8] ^ reverseSignalMap[0]
						log.Printf("segmentMap['d'] = %08b\n", segmentMap['d'])
					}
				}
			}

			// find segement 'b'
			segmentMap['b'] = segmentMap['d'] ^ reverseSignalMap[4] ^ reverseSignalMap[1]
			log.Printf("segmentMap['b'] = %08b\n", segmentMap['b'])

			// fill in rest of signal map
			signalMap[segmentMap['a']|segmentMap['c']|segmentMap['d']|segmentMap['e']|segmentMap['g']] = 2
			log.Printf("sig %08b: 2\n", signalMap[2])
			signalMap[segmentMap['a']|segmentMap['c']|segmentMap['d']|segmentMap['f']|segmentMap['g']] = 3
			log.Printf("sig %08b: 3\n", signalMap[3])
			signalMap[segmentMap['a']|segmentMap['b']|segmentMap['d']|segmentMap['f']|segmentMap['g']] = 5
			log.Printf("sig %08b: 5\n", signalMap[5])

		} else {
			// process output
			var output uint32 = 0
			for _, signalPattern := range strings.Fields(inputStr[i]) {
				bitMap := uint8(0)
				for _, signal := range signalPattern {
					bitMap |= 1 << wireMap[signal]
				}
				output = (output * 10) + uint32(signalMap[bitMap])
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

				log.Printf("** output digit \"%7s\" %08b = %v\n", signalPattern, bitMap, signalMap[bitMap])
			}
			log.Printf("*** output value %v\n", output)
			totalOutput += output
		}
	}
	log.Printf("part 1: [1|4|7|8] count %v\n", digitCount)
	log.Printf("part 2: total output %v\n", totalOutput)
}
