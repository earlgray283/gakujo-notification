package lib

type Map[K comparable, V any] map[K]V

func NewMapFromIter[K comparable, V any](keys []K, values []V) Map[K, V] {
	mp := make(Map[K, V], len(keys))
	for i := range keys {
		mp[keys[i]] = values[i]
	}
	return mp
}

func MapSlice[T, U any](a []T, f func(t T) U) []U {
	b := make([]U, len(a))
	for i, elem := range a {
		b[i] = f(elem)
	}
	return b
}
