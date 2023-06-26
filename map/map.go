package _map

// Map is a generic map.
type Map[K comparable, V comparable] map[K]V

// Keys returns Map keys.
func (m Map[K, V]) Keys() []K {
	var res []K
	for i := range m {
		res = append(res, i)
	}

	return res
}

// Values returns Map values.
func (m Map[K, V]) Values() []V {
	var res []V
	for _, i := range m {
		res = append(res, i)
	}

	return res
}

func (m Map[K, V]) ContainsKey(key K) bool {
	_, ok := m[key]

	return ok
}

func (m Map[K, V]) ContainsValue(val V) bool {
	for _, i := range m {
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
