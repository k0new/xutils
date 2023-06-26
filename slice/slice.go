package _slice

import (
	"sync"

	"golang.org/x/exp/constraints"

	_map "github.com/k0new/xutils/map"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Slice[T Number] struct {
	mu sync.RWMutex
	s  []T
}

func New[T Number](size ...int) Slice[T] {
	s := 0
	if size != nil {
		s = size[0]
	}

	return Slice[T]{mu: sync.RWMutex{}, s: make([]T, s)}
}

func NewFromSlice[T Number](s []T) Slice[T] {
	return Slice[T]{mu: sync.RWMutex{}, s: s}
}

func (s Slice[T]) Append(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.s = append(s.s, val)
}

func (s Slice[T]) AsSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.s
}

func (s Slice[T]) Max() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res T

	for _, i := range s.s {
		if i > res {
			res = i
		}
	}

	return res
}

func (s Slice[T]) Min() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res T

	for _, i := range s.s {
		if i < res {
			res = i
		}
	}

	return res
}

func (s Slice[T]) Contains(n T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, i := range s.s {
		if i == n {
			return true
		}
	}

	return false
}

func (s Slice[T]) Sum() T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var res T

	for _, i := range s.s {
		res += i
	}

	return res
}

func (s Slice[T]) Map() _map.Map[int, T] {
	s.mu.RLock()
	defer s.mu.RUnlock()

	m := _map.New[int, T](len(s.s))
	for i, v := range s.s {
		m.Set(i+1, v)
	}
	return m
}

func Max[T Number](slice []T) T {
	var res T

	for _, i := range slice {
		if i > res {
			res = i
		}
	}

	return res
}

func Min[T Number](slice []T) T {
	var res T

	for _, i := range slice {
		if i < res {
			res = i
		}
	}

	return res
}

func Sum[T Number](slice []T) T {
	var res T

	for _, i := range slice {
		res += i
	}

	return res
}

func Contains[T comparable](slice []T, n T) bool {
	for _, i := range slice {
		if i == n {
			return true
		}
	}

	return false
}

func MakeRange[T Number](min, max T) Slice[T] {
	res := make([]T, int(max-min+1))
	for i := range res {
		res[i] = min + T(i)
	}

	return NewFromSlice(res)
}
