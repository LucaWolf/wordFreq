package main

import (
	"bytes"
	"testing"
)

func Test_dataInsSort(t *testing.T) {
	w1 := word{0, []byte{'a', 'b', 'c', 'd'}}
	w2 := word{0, []byte{'a', 'b', 'a', 'd'}}
	w3 := word{0, []byte{'b', 'b', 'a', 'z'}}
	w4 := word{0, []byte{'b', 'b', 'a', 'c'}}
	// expect w2, w1, w4, w3

	type args struct {
		t []word
	}

	type res struct {
		pos int
		w   word
	}

	tests := []struct {
		name string
		args args
		r    res
	}{
		// TODO: Add test cases.
		{"insert_1", args{[]word{w1, w2, w3, w4}}, res{3, w4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataInsSort(tt.args.t)
		})

		if bytes.Compare(tt.args.t[tt.r.pos].value, tt.r.w.value) != 0 {
			t.Errorf("test %s produced wrong order", tt.name)
		}
	}
}

func Test_dataPart(t *testing.T) {
	w1 := word{3, nil}
	w2 := word{12, nil}
	w3 := word{5, nil}
	w4 := word{19, nil}
	w5 := word{1, nil}

	type args struct {
		t []word
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// 3,12,*5*,19 -> (3,5) (->12,19)
		{"pivot_odd", args{[]word{w1, w2, w3, w4}}, 2},
		// 3,12,*5*,19,1 -> (3,1,5) (->19,12)
		{"pivot_even", args{[]word{w1, w2, w3, w4, w5}}, 3},
		// 12,5,*1*,3,19 -> (1) (->1,5,12,3,9)
		{"pivot_first", args{[]word{w2, w3, w5, w1, w4}}, 1},
		// 3,5,*19*,1,12 -> (3,5,12,1) (->19)
		{"pivot_last", args{[]word{w1, w3, w4, w5, w2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dataPart(tt.args.t); got != tt.want {
				t.Errorf("dataPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dataSearch(t *testing.T) {

	blank := make([]word, 0, 8)

	data := []word{
		word{0, []byte{'a', 'b', 'c', 'a'}}, // 0
		word{0, []byte{'a', 'b', 'c', 'c'}},
		word{0, []byte{'a', 'b', 'd', 'a'}},
		word{0, []byte{'a', 'b', 'd', 'b'}},
		word{0, []byte{'a', 'b', 'd', 'x'}},
		word{0, []byte{'a', 'b', 'd', 'y'}}, // 5
	}

	type args struct {
		w []byte
		t []word
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
		{"search_blank", args{[]byte{'a', 'b', 'c', 'd'}, blank}, 0, false},
		{"search_1st_insert", args{[]byte{'a', 'b', 'b', 'z'}, data}, 0, false},
		{"search_1st_found", args{[]byte{'a', 'b', 'c', 'a'}, data}, 0, true},
		{"search_2nd_insert", args{[]byte{'a', 'b', 'c', 'b'}, data}, 1, false},
		{"search_2nd_found", args{[]byte{'a', 'b', 'c', 'c'}, data}, 1, true},
		{"search_mid_insert", args{[]byte{'a', 'b', 'd', 'o'}, data}, 4, false},
		{"search_last_found", args{[]byte{'a', 'b', 'd', 'y'}, data}, 5, true},
		{"search_last_insert", args{[]byte{'a', 'b', 'd', 'z'}, data}, 6, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := dataSearch(tt.args.w, tt.args.t)
			if got != tt.want {
				t.Errorf("dataSearch() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("dataSearch() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
