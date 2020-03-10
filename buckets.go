package main

type bucket struct {
	wordLen int
	data    []word // desc-ordered by value, all of same wordLen length
	head    int    // points to next greatest freq word (last element after sorting)
	freq    int    // copy of [head].freq as helper with sorting
}

func newBucket(length int, data []word) bucket {
	return bucket{length, data, 0, 0}
}

func bucketSearch(wLen int, t []bucket) (int, bool) {

	if len(t) == 0 {
		return 0, false
	}

	i, j := 0, len(t)-1

	for i < j {
		n := i + (j-i)/2

		if wLen > t[n].wordLen {
			i = n + 1
		} else {
			j = n
		}
	}

	if wLen > t[i].wordLen {
		return i + 1, false
	}

	return i, wLen == t[i].wordLen
}

func bucketInsert(pos int, el bucket, t []bucket) []bucket {

	res := append(t, el)
	copy(res[pos+1:], res[pos:])
	res[pos] = el

	return res

}

func bucketPart(t []bucket) int {
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

func bucketInsSort(t []bucket) {
	if len(t) == 0 {
		return
	}

	for i := 1; i < len(t); i++ {
		for j := i; j > 0 && t[j-1].freq > t[j].freq; j-- {
			t[j], t[j-1] = t[j-1], t[j]
		}
	}
}

func bucketQSort(t []bucket) {

	if len(t) < 7 {
		bucketInsSort(t)
		return
	}
	// else
	m := bucketPart(t)
	bucketQSort(t[:m])
	bucketQSort(t[m:])
}

func bucketPop(b *bucket) {
	b.head--
	b.freq = func(head int) int {
		if head >= 0 {
			return b.data[b.head].freq
		}
		return 0
	}(b.head)
}
