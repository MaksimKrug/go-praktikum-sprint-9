package main

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSliceGenerator(t *testing.T) {
	// random size
	size := rand.Intn(SIZE)
	arr := generateRandomElements(size)

	if len(arr) != size {
		t.Errorf("len = %d, want %d", len(arr), size)
	}

	for i, v := range arr {
		if v < 0 || v >= SIZE {
			t.Errorf("arr[%d] = %d, вне диапазона [0, %d)", i, v, SIZE)
		}
	}

	// empty slice
	arr = generateRandomElements(0)
	if arr != nil {
		t.Errorf("zero len slize should be nil")
	}

	// less than zero len slice
	arr = generateRandomElements(-42)
	if len(arr) != 0 {
		t.Errorf("< 0 len slize should be nil")
	}
}

func TestMaximum(t *testing.T) {
	// random size
	size := rand.Intn(SIZE)
	arr := generateRandomElements(size)

	arrMax := slices.Max(arr)
	gotMax := maximum(arr)
	if gotMax != arrMax {
		t.Errorf("gotMax %d != arrMax %d", gotMax, arrMax)
	}

	// one element
	arr = generateRandomElements(1)

	gotMax = maximum(arr)
	if gotMax != arr[0] {
		t.Errorf("gotMax %d of one element array != arr[0] %d", gotMax, arr[0])
	}

	// empty slice
	arr = generateRandomElements(-42)
	gotMax = maximum(arr)
	if gotMax != 0 {
		t.Errorf("gotMax %d != 0 for empty slice", gotMax)
	}
}

func TestChunksMaximum(t *testing.T) {
	// random size
	size := rand.Intn(SIZE)
	arr := generateRandomElements(size)

	arrMax := slices.Max(arr)
	gotMax := maxChunks(arr)
	if gotMax != arrMax {
		t.Errorf("gotMax %d != arrMax %d", gotMax, arrMax)
	}

	// one element
	arr = generateRandomElements(1)

	gotMax = maxChunks(arr)
	if gotMax != arr[0] {
		t.Errorf("gotMax %d of one element array != arr[0] %d", gotMax, arr[0])
	}

	// empty slice
	arr = generateRandomElements(-42)
	gotMax = maxChunks(arr)
	if gotMax != 0 {
		t.Errorf("gotMax %d != 0 for empty slice", gotMax)
	}

	// maximum eqal maxChunks
	size = rand.Intn(SIZE)
	arr = generateRandomElements(size)

	maxNotParallel := maximum(arr)
	maxWithParallel := maxChunks(arr)
	if maxNotParallel != maxWithParallel {
		t.Errorf("maximum %d and maxChunks %d returns are not equal", maxNotParallel, maxWithParallel)
	}
}
