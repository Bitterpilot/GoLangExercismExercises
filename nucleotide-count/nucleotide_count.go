// Package dna implements functions to manipulate DNA strands.
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if DNA contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, l := range d {
		switch {
		case l == 'A':
			h['A']++
		case l == 'C':
			h['C']++
		case l == 'G':
			h['G']++
		case l == 'T':
			h['T']++
		default:
			err := errors.New("unexpected nucleotide")
			return nil, err
		}
	}
	return h, nil
}
