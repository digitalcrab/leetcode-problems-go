package main

import (
	"fmt"
)

type Codec struct {
}

// Encode Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var out []byte
	for _, str := range strs {
		// important constraint is that the length of a single string
		// is no more than 200 chars, we can fit it's len into a byte
		out = append(out, byte(len(str)))
		out = append(out, []byte(str)...)
	}
	return string(out)
}

// Decode Decodes a single string to a list of strings.
func (codec *Codec) Decode(in string) []string {
	var out []string

	for len(in) > 0 {
		// first char is the length of the string
		length := int(in[0])

		if len(in) < length+1 { // length + string
			panic("corrupted string")
		}

		// add a word
		out = append(out, in[1:length+1])

		// move to the next, skipping the length and a word
		in = in[length+1:]
	}

	return out
}

func main() {
	var codec Codec
	result := codec.Decode(codec.Encode([]string{"Hello", "World"}))
	fmt.Printf("Encode-Decode result: %v\n", result)
}
