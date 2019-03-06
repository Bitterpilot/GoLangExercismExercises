// Package strand implements functions to manipulate DNA strings
package strand

import (
	"strings"
)

// ToRNA given a DNA strand, its transcribed RNA strand is formed
// by replacing each nucleotide with its complement:
//
// G -> C
// C -> G
// T -> A
// A -> U
func ToRNA(dna string) string {
	rna := strings.NewReplacer("G", "C",
		"C", "G",
		"T", "A",
		"A", "U")
	return rna.Replace(dna)
}
