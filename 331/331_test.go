package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_readFile(t *testing.T) {
	var s1, s2 []string

	s1 = append(s1, "15", "16", "284", "39605", "aas", "2854", "2358")
	s2 = append(s2, "13", "17", "999", "3573", "26", "54")
	tests := []struct {
		name     string
		filename string
		want     []string
	}{
		{
			name:     "first readFile test",
			filename: "14",
			want:     s1,
		},
		{
			name:     "second readFile test",
			filename: "15",
			want:     s2,
		},
		{
			name:     "third readFile test",
			filename: "someStrangeFileThatDoesntExistInAnyOfUniverses",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var i, i3 []int

	var s, s2, s3 []string

	i1 := rand.Int()
	i2 := rand.Int()
	i = append(i, i1, i2)
	s = append(s, fmt.Sprint(i1), fmt.Sprint(i2))
	s2 = append(s, "somestring", "lalala")

	tests := []struct {
		name    string
		numbers []string
		want    []int
	}{
		{
			name:    "first test check",
			numbers: s,
			want:    i,
		},
		{
			name:    "second test check",
			numbers: s2,
			want:    i,
		},
		{
			name:    "third test check",
			numbers: s3,
			want:    i3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_find(t *testing.T) {
	var n1, n2 []int
	n1 = append(n1, 1, 152, 99, 2)
	n2 = append(n2, 15, 33, 0, 13)

	var s1, s2 []string
	s1 = append(s1,
		"for 1 : x = 0, y = 0, z = 1",
		"for 1 : x = 0, y = 1, z = 0",
		"for 1 : x = 1, y = 0, z = 0",
		"for 152 : x = 2, y = 2, z = 12",
		"for 152 : x = 2, y = 12, z = 2",
		"for 152 : x = 4, y = 6, z = 10",
		"for 152 : x = 4, y = 10, z = 6",
		"for 152 : x = 6, y = 4, z = 10",
		"for 152 : x = 6, y = 10, z = 4",
		"for 152 : x = 10, y = 4, z = 6",
		"for 152 : x = 10, y = 6, z = 4",
		"for 152 : x = 12, y = 2, z = 2",
		"for 99 : x = 1, y = 7, z = 7",
		"for 99 : x = 3, y = 3, z = 9",
		"for 99 : x = 3, y = 9, z = 3",
		"for 99 : x = 5, y = 5, z = 7",
		"for 99 : x = 5, y = 7, z = 5",
		"for 99 : x = 7, y = 1, z = 7",
		"for 99 : x = 7, y = 5, z = 5",
		"for 99 : x = 7, y = 7, z = 1",
		"for 99 : x = 9, y = 3, z = 3",
		"for 2 : x = 0, y = 1, z = 1",
		"for 2 : x = 1, y = 0, z = 1",
		"for 2 : x = 1, y = 1, z = 0",
	)

	s2 = append(s2,
		"for 33 : x = 1, y = 4, z = 4",
		"for 33 : x = 2, y = 2, z = 5",
		"for 33 : x = 2, y = 5, z = 2",
		"for 33 : x = 4, y = 1, z = 4",
		"for 33 : x = 4, y = 4, z = 1",
		"for 33 : x = 5, y = 2, z = 2",
		"0 - number must be positive",
		"for 13 : x = 0, y = 2, z = 3",
		"for 13 : x = 0, y = 3, z = 2",
		"for 13 : x = 2, y = 0, z = 3",
		"for 13 : x = 2, y = 3, z = 0",
		"for 13 : x = 3, y = 0, z = 2",
		"for 13 : x = 3, y = 2, z = 0",
	)
	tests := []struct {
		name    string
		numbers []int
		want    []string
	}{
		{
			name:    "first find test",
			numbers: n1,
			want:    s1,
		},
		{
			name:    "second find test",
			numbers: n2,
			want:    s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find(tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find() = %v, want %v", got, tt.want)
			}
		})
	}
}
