package main

import (
	"sync"
	"testing"
	"time"
)

func withMutext(N int) {
	wg := sync.WaitGroup{}
	m := []int{}
	mut := sync.Mutex{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		i := i
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(1) * time.Second)
			mut.Lock()
			m = append(m, i)
			mut.Unlock()
		}(i)
	}
	wg.Wait()

}

func withoutMutex(N int) {
	m := []int{}
	chanInt := make(chan int, N)
	for i := 0; i < N; i++ {
		i := i
		go func(i int) {
			time.Sleep(time.Duration(1) * time.Second)
			chanInt <- i
		}(i)
	}

	for i := 0; i < N; i++ {
		select {
		case v := <-chanInt:
			m = append(m, v)
		}
	}

	close(chanInt)

}

func withoutMutexWG(N int) {
	wg := sync.WaitGroup{}
	m := []int{}
	chanInt := make(chan int, N)
	for i := 0; i < N; i++ {
		wg.Add(1)
		i := i
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(1) * time.Second)
			chanInt <- i
		}(i)
	}

	go func() {
		wg.Wait()
		close(chanInt)
	}()

	for i := range chanInt {
		m = append(m, i)
	}

}

func BenchmarkMethodWithMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withMutext(100000)
	}
}

func BenchmarkMethodWithoutMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutMutex(100000)
	}
}

func BenchmarkMethodWithoutMutexWG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withoutMutexWG(100000)
	}
}
