package swiss_table

import (
	"runtime"
	"strconv"
	"testing"

	cockroachSwiss "github.com/cockroachdb/swiss"
	dolthubSwiss "github.com/dolthub/swiss"
	csmap "github.com/mhmtszr/concurrent-swiss-map"
	mapsutil "github.com/projectdiscovery/utils/maps"
)

func init() {
	runtime.GOMAXPROCS(1)
}

// keyCount defines the number of keys to use in benchmarks.
var keyCounts = []int{100, 1_000, 10_000, 100_000, 500_000, 1_000_000, 2_500_000, 5_000_000, 10_000_000}

// generateKeys returns a slice of keyCount strings: "key0", "key1", ….
func generateKeys(n int) []string {
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		keys[i] = "key" + strconv.Itoa(i)
	}
	return keys
}

func BenchmarkPut(b *testing.B) {
	for _, keyCount := range keyCounts {
		b.Run(strconv.Itoa(keyCount), func(b *testing.B) {
			keys := generateKeys(keyCount)
			mapStd := make(map[string]int)
			mapMapsutil := make(mapsutil.Map[string, int])
			mapOrderedMap := mapsutil.NewOrderedMap[string, int]()
			mapSyncLockMap := mapsutil.NewSyncLockMap[string, int]()
			mapCsmap := csmap.Create(csmap.WithSize[string, int](uint64(keyCount)))
			mapCockroachDB := cockroachSwiss.New[string, int](keyCount)
			mapDolthub := dolthubSwiss.NewMap[string, int](uint32(keyCount))

			b.ResetTimer()
			b.Run("std", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapStd[key] = j
					}
				}
			})

			b.ResetTimer()
			b.Run("mapsutil", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapMapsutil.Set(key, j)
					}
				}
			})

			b.ResetTimer()
			b.Run("mapsutil-OrderedMap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapOrderedMap.Set(key, j)
					}
				}
			})

			b.ResetTimer()
			b.Run("mapsutil-SyncLockMap", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapSyncLockMap.Set(key, j)
					}
				}
			})

			b.ResetTimer()
			b.Run("concurrent-swiss-map", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapCsmap.Store(key, j)
					}
				}
			})

			b.ResetTimer()
			b.Run("cockroachdb", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapCockroachDB.Put(key, j)
					}
				}
			})

			b.ResetTimer()
			b.Run("dolthub", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for j, key := range keys {
						mapDolthub.Put(key, j)
					}
				}
			})
		})
	}
}

func BenchmarkGet(b *testing.B) {
	for _, keyCount := range keyCounts {
		b.Run(strconv.Itoa(keyCount), func(b *testing.B) {
			keys := generateKeys(keyCount)
			mapStd := make(map[string]int)
			mapMapsutil := make(mapsutil.Map[string, int])
			mapOrderedMap := mapsutil.NewOrderedMap[string, int]()
			mapSyncLockMap := mapsutil.NewSyncLockMap[string, int]()
			mapCsmap := csmap.Create(csmap.WithSize[string, int](uint64(keyCount)))
			mapCockroachDB := cockroachSwiss.New[string, int](keyCount)
			mapDolthub := dolthubSwiss.NewMap[string, int](uint32(keyCount))

			for j, key := range keys {
				mapStd[key] = j
				mapMapsutil.Set(key, j)
				mapOrderedMap.Set(key, j)
				mapSyncLockMap.Set(key, j)
				mapCsmap.Store(key, j)
				mapCockroachDB.Put(key, j)
				mapDolthub.Put(key, j)
			}

			b.ResetTimer()
			b.Run("std", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapStd[key]
				}
			})

			b.ResetTimer()
			b.Run("mapsutil", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapMapsutil.Get(key)
				}
			})

			b.ResetTimer()
			b.Run("mapsutil-OrderedMap", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapOrderedMap.Get(key)
				}
			})

			b.ResetTimer()
			b.Run("mapsutil-SyncLockMap", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapSyncLockMap.Get(key)
				}
			})

			b.ResetTimer()
			b.Run("concurrent-swiss-map", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapCsmap.Load(key)
				}
			})

			b.ResetTimer()
			b.Run("cockroachdb", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapCockroachDB.Get(key)
				}
			})

			b.ResetTimer()
			b.Run("dolthub", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_, _ = mapDolthub.Get(key)
				}
			})
		})
	}
}

func BenchmarkDelete(b *testing.B) {
	for _, keyCount := range keyCounts {
		b.Run(strconv.Itoa(keyCount), func(b *testing.B) {
			keys := generateKeys(keyCount)
			mapStd := make(map[string]int)
			mapMapsutil := make(mapsutil.Map[string, int])
			mapOrderedMap := mapsutil.NewOrderedMap[string, int]()
			mapSyncLockMap := mapsutil.NewSyncLockMap[string, int]()
			mapCsmap := csmap.Create(csmap.WithSize[string, int](uint64(keyCount)))
			mapCockroachDB := cockroachSwiss.New[string, int](keyCount)
			mapDolthub := dolthubSwiss.NewMap[string, int](uint32(keyCount))

			for j, key := range keys {
				mapStd[key] = j
				mapMapsutil.Set(key, j)
				mapOrderedMap.Set(key, j)
				mapSyncLockMap.Set(key, j)
				mapCsmap.Store(key, j)
				mapCockroachDB.Put(key, j)
				mapDolthub.Put(key, j)
			}

			b.ResetTimer()
			b.Run("std", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					delete(mapStd, key)
				}
			})

			b.ResetTimer()
			b.Run("mapsutil/no-op", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// no-op: mapsutil.Map does not have a Delete method
				}
			})

			b.ResetTimer()
			b.Run("mapsutil-OrderedMap", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					mapOrderedMap.Delete(key)
				}
			})

			b.ResetTimer()
			b.Run("mapsutil-SyncLockMap", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					mapSyncLockMap.Delete(key)
				}
			})

			b.ResetTimer()
			b.Run("concurrent-swiss-map", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_ = mapCsmap.Delete(key)
				}
			})

			b.ResetTimer()
			b.Run("cockroachdb", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					mapCockroachDB.Delete(key)
				}
			})

			b.ResetTimer()
			b.Run("dolthub", func(b *testing.B) {
				nKeys := len(keys)
				for i := 0; i < b.N; i++ {
					key := keys[i%nKeys]
					_ = mapDolthub.Delete(key)
				}
			})
		})
	}
}
