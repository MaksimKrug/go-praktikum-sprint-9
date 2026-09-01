package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// в задаче не сказано как такое обрабатывать, пусть будет nil
	if size <= 0 {
		return nil
	}

	nums := make([]int, size)
	for i := range nums {
		nums[i] = rand.Int()
	}
	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// непонятно что надо вернуть в таком случае, если по сигнатуре функции ошибку возвращать нельзя
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// непонятно что надо вернуть в таком случае, если по сигнатуре функции ошибку возвращать нельзя
	if len(data) == 0 {
		return 0
	}

	// нет смысла разбивать на чанки если len(data) < CHUNKS
	if len(data) < CHUNKS {
		return maximum(data)
	}

	// ваш код здесь
	var wg sync.WaitGroup
	maximumChunks := make([]int, CHUNKS)

	chunkSize := len(data) / CHUNKS
	for i := range CHUNKS {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data) // последний чанк забирает остаток
		}

		chunk := data[start:end]
		go func() {
			defer wg.Done()
			maximumChunks[i] = maximum(chunk)
		}()

	}
	wg.Wait()

	return maximum(maximumChunks)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	arr := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	maxVal := maximum(arr)
	end := time.Now()
	elapsed := end.Sub(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxVal = maxChunks(arr)
	end = time.Now()
	elapsed = end.Sub(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed)
}
