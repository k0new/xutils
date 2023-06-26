# xutils

`xutils` is a Go package that provides utility functions and data structures for slices and maps. It includes
concurrency-safe implementations of a generic map and a number of methods for manipulating and performing operations on
slices.

## Installation

To use the `xutils` package in your Go project, you can import it using the following import path:

```go
import "github.com/k0new/xutils"
```

Make sure to have Go installed and properly configured before adding the package to your project.

## _map

The `_map` package within `xutils` provides a concurrency-safe generic map implementation. It allows you to store and
manipulate key-value pairs while ensuring safe concurrent access.

### Features

- Concurrency-safe operations: The `_map.Map` type ensures that concurrent access to the map is safe by utilizing a
  read-write mutex (`sync.RWMutex`).
- Creation and initialization: The package provides functions to create a new `Map` object and initialize it with an
  optional initial size or from an existing `map`.
- Basic operations: The package supports essential map operations such as setting or updating records, deleting records
  by key, retrieving values by key, and getting the number of elements in the map.
- Conversion functions: Conversion functions are available to obtain the map as a Go built-in map (`map[K]V`), get the
  keys or values of the map as slices, and check for the presence of a key or value in the map.
- Helper functions: Additionally, the package includes standalone helper functions to perform similar operations on
  regular Go built-in maps.

## _slice

The `_slice` package within `xutils` provides utility functions and methods for manipulating and performing operations
on slices.

### Features

- Concurrency-safe operations: The `Slice` type ensures that concurrent access to the slice is safe by utilizing a
  read-write mutex (`sync.RWMutex`).
- Creation and initialization: The package provides functions to create a new `Slice` object and initialize it with an
  optional initial size or from an existing slice.
- Basic operations: The package includes methods to append elements to the slice, obtain the length of the slice, find
  the maximum and minimum values, check if a value exists in the slice, and calculate the sum of the elements.
- Conversion functions: Conversion functions are available to obtain the slice as a Go built-in slice (`[]T`).
- Helper functions: Additionally, the package includes standalone helper functions to perform similar operations on
  regular Go built-in slices.

## Usage

To use the `xutils` package and its sub-packages, import them into your Go code as needed. Here's an example
illustrating the basic usage of the `_map` and `_slice` packages:

```go
package main

import (
	"fmt"

	"github.com/k0new/xutils/_map"
	"github.com/k0new/xutils/_slice"
)

func main() {
	// Creating and using a concurrency-safe map
	m := _map.New[int, string]()
	m.Set(1, "apple")
	m.Set(2, "banana")
	m.Set(3, "orange")
	fmt.Println("Map length:", m.Len()) // Output: Map length: 3
	keys := m.Keys()
	fmt.Println("Keys:", keys) // Output: Keys: [1 2 3]

	// Creating and using a concurrency-safe slice
	s := _slice.New[int]()
	s.Append(5)
	s.Append(3)
	s.Append(7)
	fmt.Println("Slice length:", s.Len()) // Output: Slice length: 3
	max := s.Max()
	fmt.Println("Maximum value:", max) // Output: Maximum value: 7

	// Using helper functions on regular slices
	regularSlice := []int{2, 4,

		6, 8, 10}
	min := _slice.Min(regularSlice)
	fmt.Println("Minimum value:", min) // Output: Minimum value: 2
	contains := _slice.Contains(regularSlice, 6)
	fmt.Println("Contains 6:", contains) // Output: Contains 6: true
}
```

The above code demonstrates how to create and use a concurrency-safe map and slice from the `_map` and `_slice`
packages, respectively. It also shows how to use helper functions on regular Go built-in maps and slices.

## Contribution and Support

Contributions to this package are welcome. If you encounter any issues or have suggestions for improvements, please open
an issue on the GitHub repository.

## License

This package is open-source and distributed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to
use it in your own projects.