package _slice

import (
	"golang.org/x/exp/constraints"

	_map "github.com/k0new/xutils/map"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Slice[T Number] []T

func (s Slice[T]) Max() T {
	var res T

	for _, i := range s {
		if i > res {
			res = i
		}
	}

	return res
}

func (s Slice[T]) Min() T {
	var res T

	for _, i := range s {
		if i < res {
			res = i
		}
	}

	return res
}

func (s Slice[T]) Contains(n T) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}

	return false
}

func (s Slice[T]) Sum() T {
	var res T

	for _, i := range s {
		res += i
	}

	return res
}

func (s Slice[T]) Map() _map.Map[T, T] {
	m := make(_map.Map[T, T], len(s))
	for i, v := range s {
		m[T(i+1)] = v
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

	return res
}
