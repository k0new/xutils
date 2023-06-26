package _map

import "sync"

// Map is a concurrency-safe generic map.
type Map[K comparable, V comparable] struct {
	mu sync.RWMutex
	m  map[K]V
}

// New creates new Map object
func New[K, V comparable](size ...int) Map[K, V] {
	s := 0
	if size != nil {
		s = size[0]
	}
	return Map[K, V]{mu: sync.RWMutex{}, m: make(map[K]V, s)}
}

// NewFromMap creates new Map object from existing map.
func NewFromMap[K, V comparable](m map[K]V) Map[K, V] {
	return Map[K, V]{mu: sync.RWMutex{}, m: m}
}

// Set sets or updates record.
func (m Map[K, V]) Set(k K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[k] = v
}

// Delete deletes record by key.
func (m Map[K, V]) Delete(k K) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.m, k)
}

// Get returns value by provided key.
func (m Map[K, V]) Get(k K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.m[k]
}

// AsMap returns Map as golang map.
func (m Map[K, V]) AsMap() map[K]V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.m
}

// Keys returns Map keys.
func (m Map[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var res []K
	for i := range m.m {
		res = append(res, i)
	}

	return res
}

// Values returns Map values.
func (m Map[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var res []V
	for _, i := range m.m {
		res = append(res, i)
	}

	return res
}

func (m Map[K, V]) ContainsKey(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.m[key]

	return ok
}

func (m Map[K, V]) ContainsValue(val V) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, i := range m.m {
		if i == val {
			return true
		}
	}

	return false
}

func Keys[K comparable, V any](m map[K]V) []K {
	var res []K
	for i := range m {
		res = append(res, i)
	}

	return res
}

func Values[K comparable, V any](m map[K]V) []V {
	var res []V
	for _, i := range m {
		res = append(res, i)
	}

	return res
}

func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]

	return ok
}

func ContainsValue[K comparable, V comparable](m map[K]V, val V) bool {
	for _, i := range m {
		if i == val {
			return true
		}
	}

	return false
}
