package articulate

import (
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/base"
)

func articulateBytes(byteValue []byte) (strResult string, success bool) {
	hasPrintableCharacters := false
	result := make([]byte, 0, len(byteValue))
	for _, character := range byteValue {
		sanitized, replaced := sanitizeByte(character)
		// if any character has been replaced, it was a special character
		if replaced > 0 {
			result = append(result, sanitized...)
			continue
		}
		// if we are here, the character is not a special one, so we need
		// to check if it's ASCII printable
		if character >= 20 && character <= 126 {
			result = append(result, byte(character))

			// ignore space
			if character > 20 {
				hasPrintableCharacters = true
			}
			continue
		}
		// The character is not ASCII
		return "", false
	}

	if hasPrintableCharacters {
		return string(result), true
	}

	return
}

func sanitizeByte(character byte) (replacement []byte, replaced int) {
	if character == '\\' || character == '\r' {
		return
	}
	if character == '"' {
		return []byte{'\''}, 1
	}
	if character == ',' {
		// C++ used `|` to replace commas, but it breaks compressed* fields format
		// (`|` is the delimeter there)
		return []byte{'_'}, 1
	}
	if character == '|' {
		return []byte{';'}, 1
	}
	if character == '\n' {
		return []byte("[n]"), 1
	}
	if character == '\t' {
		return []byte("[t]"), 1
	}
	return []byte{character}, 0
}

// ArticulateBytes32String turns bytes32 encoded string into Go string
func ArticulateBytes32String(hex string) (result string) {
	if len(hex) < 2 {
		return
	}
	input := base.Hex2Bytes(hex[2:])
	if len(input) == 0 {
		return ""
	}
	// Filter out invalid names, four-byte collisions (0x8406d0897da43a33912995c6ffd792f1f2125cd4)
	if input[0] == 0 {
		return ""
	}
	padStart := len(input)
	for i := (len(input) - 1); i >= 0; i-- {
		if input[i] != 0 {
			break
		}
		padStart = i
	}

	byteValue := input[0:padStart]
	result, _ = articulateBytes(byteValue)

	return
}
