package swiss_table

import (
	"strconv"
	"testing"

	cockroachSwiss "github.com/cockroachdb/swiss"
	dolthubSwiss "github.com/dolthub/swiss"
	csmap "github.com/mhmtszr/concurrent-swiss-map"
)

// keyCount defines the number of keys to use in benchmarks.
const keyCount = 1 << 10

// generateKeys returns a slice of keyCount strings: "key0", "key1", ….
func generateKeys(n int) []string {
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		keys[i] = "key" + strconv.Itoa(i)
	}
	return keys
}

func BenchmarkPut(b *testing.B) {
	keys := generateKeys(keyCount)
	mapCsmap := csmap.Create[string, int]()
	mapCockroachDB := cockroachSwiss.New[string, int](keyCount)
	mapDolthub := dolthubSwiss.NewMap[string, int](keyCount)

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
}

func BenchmarkGet(b *testing.B) {
	keys := generateKeys(keyCount)
	mapCsmap := csmap.Create[string, int]()
	mapCockroachDB := cockroachSwiss.New[string, int](keyCount)
	mapDolthub := dolthubSwiss.NewMap[string, int](keyCount)

	for j, key := range keys {
		mapCsmap.Store(key, j)
		mapCockroachDB.Put(key, j)
		mapDolthub.Put(key, j)
	}

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
}

func BenchmarkDelete(b *testing.B) {
	keys := generateKeys(keyCount)
	mapCsmap := csmap.Create[string, int]()
	mapCockroachDB := cockroachSwiss.New[string, int](keyCount)
	mapDolthub := dolthubSwiss.NewMap[string, int](keyCount)

	for i := 0; i < b.N; i++ {
		for j, key := range keys {
			mapCsmap.Store(key, j)
			mapCockroachDB.Put(key, j)
			mapDolthub.Put(key, j)
		}
	}

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
}
