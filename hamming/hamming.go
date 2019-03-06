package hamming

import "errors"

// Distance tells if two stings have the same contents
// returning the number of differences
func Distance(a, b string) (int, error) {
	switch {
	case len(a) != len(b):
		return 0, errors.New("lengths do not match")
	default:
		aS := []byte(a)
		bS := []byte(b)
		var count int
		for i, v := range aS {
			if v != bS[i] {
				count++
			}
		}
		return count, nil
	}
}
