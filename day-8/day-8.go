package main

import (
	"io/ioutil"
	"log"
	"math/bits"
	"regexp"
	"strings"
)

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
	var digitCount uint32 = 0
	var totalOutput uint32 = 0
	var segmentMap map[uint8]uint8
	var signalMap map[uint8]uint8
	var reverseSignalMap map[uint8]uint8
	var crossMap map[rune]uint8

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
			segmentMap = make(map[uint8]uint8)
			signalMap = make(map[uint8]uint8)
			reverseSignalMap = make(map[uint8]uint8)
			crossMap = make(map[rune]uint8)

			// process input singals
			for _, signalPattern := range strings.Fields(inputStr[i]) {
				bitMap := uint8(0)
				for _, signal := range signalPattern {
					bitMap |= 1 << wireMap[signal]
				}
				signalMap[bitMap] = 0xf
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
		} else {
			// process output
			var output uint32 = 0
			for _, signalPattern := range strings.Fields(inputStr[i]) {
				bitMap := uint8(0)
				for _, signal := range signalPattern {
					bitMap |= 1 << wireMap[signal]
				}
				output = (output * 10) + uint32(segmentMap[bitMap])
				switch segmentMap[bitMap] {
				case 1:
					fallthrough
				case 4:
					fallthrough
				case 7:
					fallthrough
				case 8:
					digitCount++
				}

				log.Printf("** output %08b = %v\n", bitMap, segmentMap[bitMap])
			}
			log.Printf("** output %v\n", output)
			totalOutput += output
			//break
		}
		// dump keys
		for key, value := range signalMap {
			log.Printf("sig %08b: %v\n", key, value)
		}
		// find segment 'a'
		crossMap['a'] = reverseSignalMap[7] ^ reverseSignalMap[1]
		log.Printf("crossMap['a'] = %08b\n", crossMap['a'])

		// find segment 'c'
		for key := range signalMap {
			// find candidates for 6
			if bits.OnesCount8(key) == 6 {
				//log.Printf("test: %b %b %b\n", reverseSignalMap[8], key, reverseSignalMap[1])
				value := (reverseSignalMap[8] ^ key) & reverseSignalMap[1]
				if value != 0 {
					crossMap['c'] = value
					log.Printf("crossMap['c'] = %08b\n", crossMap['c'])
					signalMap[key] = 6
					reverseSignalMap[6] = key
				}
			}
		}

		// find segment 'f'
		crossMap['f'] = reverseSignalMap[1] ^ crossMap['c']
		log.Printf("crossMap['f'] = %08b\n", crossMap['f'])

		// find segment 'g'
		for key := range signalMap {
			// find candidates for 9
			if bits.OnesCount8(key) == 6 {
				value := key ^ reverseSignalMap[4] ^ reverseSignalMap[1] ^ reverseSignalMap[7]

				if bits.OnesCount8(value) == 1 {
					crossMap['g'] = value
					log.Printf("crossMap['g'] = %08b\n", crossMap['g'])
					signalMap[key] = 9
					reverseSignalMap[9] = key
				}
			}
		}

		// find segment 'e'
		crossMap['e'] = reverseSignalMap[8] ^ reverseSignalMap[9]
		log.Printf("crossMap['e'] = %08b\n", crossMap['e'])

		// find segment 'd'
		for key := range signalMap {
			// find candidates for 0
			if bits.OnesCount8(key) == 6 {
				if key != reverseSignalMap[9] && key != reverseSignalMap[6] {
					reverseSignalMap[0] = key
					signalMap[key] = 0xf
					crossMap['d'] = reverseSignalMap[8] ^ reverseSignalMap[0]
					log.Printf("crossMap['d'] = %08b\n", crossMap['d'])
				}
			}
		}

		// find segement 'b'
		crossMap['b'] = crossMap['d'] ^ reverseSignalMap[4] ^ reverseSignalMap[1]
		log.Printf("crossMap['b'] = %08b\n", crossMap['b'])

		// fill in segment map
		segmentMap[crossMap['a']|crossMap['b']|crossMap['c']|crossMap['e']|crossMap['f']|crossMap['g']] = 0
		segmentMap[crossMap['c']|crossMap['f']] = 1
		segmentMap[crossMap['a']|crossMap['c']|crossMap['d']|crossMap['e']|crossMap['g']] = 2
		segmentMap[crossMap['a']|crossMap['c']|crossMap['d']|crossMap['f']|crossMap['g']] = 3
		segmentMap[crossMap['b']|crossMap['c']|crossMap['d']|crossMap['f']] = 4
		segmentMap[crossMap['a']|crossMap['b']|crossMap['d']|crossMap['f']|crossMap['g']] = 5
		segmentMap[crossMap['a']|crossMap['b']|crossMap['d']|crossMap['e']|crossMap['f']|crossMap['g']] = 6
		segmentMap[crossMap['a']|crossMap['c']|crossMap['f']] = 7
		segmentMap[crossMap['a']|crossMap['b']|crossMap['c']|crossMap['d']|crossMap['e']|crossMap['f']|crossMap['g']] = 8
		segmentMap[crossMap['a']|crossMap['b']|crossMap['c']|crossMap['d']|crossMap['f']|crossMap['g']] = 9

		for key, value := range segmentMap {
			log.Printf("seg %v: %08b\n", value, key)
		}
	}
	log.Printf("total output %v [1|4|7|8] count %v\n", totalOutput, digitCount)
}
