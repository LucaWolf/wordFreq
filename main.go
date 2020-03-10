package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func fetchWord(fname string, ch chan<- []byte) {
	defer close(ch)

	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	wr := bufio.NewScanner(f)
	wr.Split(patchedScanWords)

	for wr.Scan() {
		// the scan may return a slice into the original buffer, allocate a copy for further handling
		// Go detects stack vs heap based on variable scope
		w := bytes.ToLower(wr.Bytes())
		ch <- w
	}

	if err := wr.Err(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {

	var fname string = "./mobydick.txt"

	if len(os.Args) < 2 {
		fmt.Println("Takes an input file name. Defaults to", fname)
	} else {
		fname = os.Args[1]
	}

	wReader := make(chan []byte, 32)

	// create variable length and sorted word buckets as
	//   b[i] ... b[n] ... b[j] where i < n < j; b[i] holds all words of same length
	//   |
	//   wordLength - this bucket holds all words of this length
	//   head/freq are simple helpers for avoinding extra pointer indirection into data[m]
	//   data[0] .. data[n]  are the words kept in sorted order
	// by convention the sorting order (buckets and payload) used is ascending
	buckets := make([]bucket, 0, 32) // expandable list for word length 32 (in bytes)

	go fetchWord(fname, wReader)

	for w := range wReader {

		if n, found := bucketSearch(len(w), buckets); found == true {
			// identify bucket
			if m, wFound := dataSearch(w, buckets[n].data); wFound == true {
				// identify word
				buckets[n].data[m].freq++
			} else {
				// or register new word; sort preserving by value
				buckets[n].data = dataInsert(m, newWord(1, w), buckets[n].data)
			}
		} else {
			// or create new bucket; sort preserving by length
			buckets = bucketInsert(n, newBucket(len(w), newData()), buckets)
			buckets[n].data = dataInsert(0, newWord(1, w), buckets[n].data)
		}
	}

	// sort words in each bucket by freq
	for i := 0; i < len(buckets); i++ {
		dataQSort(buckets[i].data)

		// and select the best freq
		buckets[i].head = len(buckets[i].data) - 1
		buckets[i].freq = func(head int) int {
			if head >= 0 {
				return buckets[i].data[buckets[i].head].freq
			}
			return 0
		}(buckets[i].head)
	}

	for i, last := 0, len(buckets)-1; i < 20; i++ {

		// after 1st run this list is almost sorted ;-)
		bucketQSort(buckets) // order buckets by freq

		head := buckets[last].head

		fmt.Printf("%02d. %6d %s\n", i+1,
			buckets[last].data[head].freq,
			buckets[last].data[head].value,
		)

		bucketPop(&buckets[last]) // extract top candiate and redo above
	}

}
