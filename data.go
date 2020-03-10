package main

import (
	"bytes"
)

type word struct {
	freq  int
	value []byte
}

func newWord(freq int, w []byte) word {
	return word{freq, w}
}

func newData() []word {
	return make([]word, 0, 32) // expandable words capacity of this bucket
}

// Binary searches 'w' into the 't' haystack
// t is expected to be sorted (ascending)
// returns tuple of:
// - the index of the find or where to insert
// - true if found or false
func dataSearch(w []byte, t []word) (int, bool) {

	if len(t) == 0 {
		return 0, false
	}

	i, j := 0, len(t)-1

	for i < j {
		n := i + (j-i)/2

		if bytes.Compare(w, t[n].value) > 0 {
			i = n + 1
		} else {
			j = n
		}
	}

	if bytes.Compare(w, t[i].value) > 0 {
		return i + 1, false
	}

	return i, bytes.Compare(w, t[i].value) == 0
}

// Inserts 'el' into 't' haystack
// 'pos' is expected to be 0 <= pos <= len(t)
// (no checks!) and usually comes from a deriving routine (see dataSearch)
func dataInsert(pos int, el word, t []word) []word {

	res := append(t, el)
	copy(res[pos+1:], res[pos:])
	res[pos] = el

	return res

}

// Partitions 't' into [lower/eq] and (greater] than pivot
// The pivot selection scheme is but a simple mid position value
// Returns the index of the greater segment
func dataPart(t []word) int {
	p := t[len(t)/2]

	for i, j := -1, len(t); ; {
		for ok := true; ok; {
			i++
			ok = t[i].freq < p.freq
		}
		for ok := true; ok; {
			j--
			ok = t[j].freq > p.freq
		}

		if i >= j {
			return j + 1
		}

		t[i], t[j] = t[j], t[i]
	}
}

// Standard insert sort
func dataInsSort(t []word) {
	if len(t) == 0 {
		return
	}

	for i := 1; i < len(t); i++ {
		for j := i; j > 0 && t[j-1].freq > t[j].freq; j-- {
			t[j], t[j-1] = t[j-1], t[j]
		}
	}
}

//standard quick sort
func dataQSort(t []word) {

	if len(t) < 7 {
		dataInsSort(t)
		return
	}
	// else
	m := dataPart(t)
	dataQSort(t[:m])
	dataQSort(t[m:])
}
