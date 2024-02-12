package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/dgraph-io/ristretto"
	"github.com/koding/cache"
	gocache "github.com/patrickmn/go-cache"
)

// go:test go test -bench=. -cpuprofile=cpu.prof

var (
	rrcache *ristretto.Cache
	kcache  *cache.MemoryTTL
	gcache  *gocache.Cache
	fcache  *fastcache.Cache
)

func initRisretto() {
	var (
		err error
	)
	rrcache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 31, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
}

const (
	parsedStringLong = "ThisIsAVeryLongStringWithoutAnySpacesAndItGoesOnAndOnWithoutAnyBreaksOrSpacesAndItKeepsContinuingEndlesslyWithoutAnyInterruptionsOrGapsInBetweenAndItJustGoesOnAndOnAndOnForeverAndEverWithoutStoppingOrPausingOrAddingAnySpacesOrBreaksInTheStringAndItContinuesToEndlessInfinityWithoutAnyInterruptionOrModificationAndItGoesOnAndOnAndOnWithoutAnySpacesOrGapsInBetweenAndItIsJustOneContinuousUnbrokenStringThatExtendsEndlesslyWithoutAnyEndOrLimitAndItContinuesForeverWithoutAnyBreaksOrSpacesOrInterruptionsAndItJustKeepsGoingOnAndOnAndOnWithoutAnyPauseOrModificationOrSpaceOrBreakInTheStringAndItIsJustOneUninterruptedStringThatGoesOnAndOnWithoutAnyEndOrLimitInTheStringAndItJustContinuesWithoutAnySpacesOrBreaksOrInterruptionsAndItGoesOnForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoingOnAndOnWithoutAnySpacesOrBreaksOrInterruptionsAndItContinuesForeverWithoutAnyModificationOrPauseOrSpacesOrBreaksInTheStringAndItIsJustOneLongUnbrokenStringThatStretchesIntoInfinityAndBeyondWithoutAnyEndOrLimitInTheStringAndItJustKeepsGoing"
)

func valueLong(i int) string {
	return strconv.Itoa(i) + parsedStringLong
}

func loadtestRistrettoSet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			rrcache.Set(i, valueLong(i), 1)
			rrcache.Wait()
		}()
	}
	wg.Wait()
}

func loadtestRistrettoGet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		// get value from rrcache
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, found := rrcache.Get(i)
			if !found {
				// panic("missing value")
			}

			rrcache.Del(i)
		}()
		// fmt.Println("Ristretto : ", value)
	}
	wg.Wait()
}

func initKodingCache() {
	// create a cache with 2 second TTL
	kcache = cache.NewMemoryWithTTL(10 * time.Second)
}

func loadtestKodingCacheSet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			kcache.Set(strconv.Itoa(i), valueLong(i))
		}()
	}

	wg.Wait()
}

func loadtestKodingCacheGet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := kcache.Get(strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
			err = kcache.Delete(strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
		}()
		// fmt.Println("Kodingcache : ", val)
	}
	wg.Wait()
}

func initGoCache() {
	gcache = gocache.New(5*time.Minute, 10*time.Minute)
}

func loadtestGoCacheSet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			gcache.Set(strconv.Itoa(i), valueLong(i), time.Second*5)
		}()
	}

	wg.Wait()
}

func loadtestGoCacheGet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, found := gcache.Get(strconv.Itoa(i))
			if !found {
				panic("panic missing value")
			}
			gcache.Delete(strconv.Itoa(i))
		}()
		// fmt.Println("Kodingcache : ", val)
	}
	wg.Wait()
}

func initFastCache() {
	fcache = fastcache.New(1 << 31)
}

func loadtestFastCacheSet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fcache.Set([]byte(strconv.Itoa(i)), []byte(valueLong(i)))
		}()
	}

	wg.Wait()
}

func loadtestFastCacheGet(N int) {
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := fcache.Get(nil, []byte(strconv.Itoa(i)))
			if len(val) == 0 {
				panic(fmt.Sprintf("val is empty at index %v\n", i))
			}
			fcache.Del([]byte(strconv.Itoa(i)))
		}()
		// fmt.Println("Kodingcache : ", val)
	}
	wg.Wait()
}

func BenchmarkKodingCache(b *testing.B) {
	initKodingCache()
	loadtestKodingCacheSet(b.N)
	loadtestKodingCacheGet(b.N)
}

func BenchmarkGoCache(b *testing.B) {
	initGoCache()
	loadtestGoCacheSet(b.N)
	loadtestGoCacheGet(b.N)
}

func BenchmarkRistRetto(b *testing.B) {
	initRisretto()
	loadtestRistrettoSet(b.N)
	loadtestRistrettoGet(b.N)
}

func BenchmarkFastCache(b *testing.B) {
	initFastCache()
	loadtestFastCacheSet(b.N)
	loadtestFastCacheGet(b.N)
}
