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

	// pre-allocate some space
	buckets := make([]bucket, 0, 32) // expandable list for word length 32 (in bytes)

	go fetchWord(fname, wReader)

	for w := range wReader {

		if n, found := bucketSearch(len(w), buckets); found == true {
			if m, wFound := dataSearch(w, buckets[n].data); wFound == true {
				buckets[n].data[m].freq++
			} else {
				buckets[n].data = dataInsert(m, newWord(1, w), buckets[n].data)
			}
		} else {
			buckets = bucketInsert(n, newBucket(len(w), newData()), buckets)
			buckets[n].data = dataInsert(0, newWord(1, w), buckets[n].data)
		}
	}

	for i := 0; i < len(buckets); i++ {
		dataQSort(buckets[i].data)

		buckets[i].head = len(buckets[i].data) - 1
		buckets[i].freq = func(head int) int {
			if head >= 0 {
				return buckets[i].data[buckets[i].head].freq
			}
			return 0
		}(buckets[i].head)
	}

	for i, last := 0, len(buckets)-1; i < 20; i++ {

		bucketQSort(buckets)

		head := buckets[last].head

		fmt.Printf("%02d. %6d %s\n", i+1,
			buckets[last].data[head].freq,
			buckets[last].data[head].value,
		)

		bucketPop(&buckets[last])
	}

}
