package memstr

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	sizeMap = map[string]int64{
		"k": 1000,
		"K": 1024,
		"m": 1000000,
		"M": 1048576,
		"g": 1000000000,
		"G": 1073741824,
		"t": 1000000000000,
		"T": 1099511627776,
	}

	// ErrBadFormat means the string was incorreclty formatted.
	ErrBadFormat = errors.New("bad format")

	// ErrInvalidSizeString means the string was formatted correctly, but the size was invalid
	ErrInvalidSizeString = errors.New("invalid size string")

	memMatcher = regexp.MustCompile(`([0-9]+\.*[0-9]*)([kKmMgGtT]*)`)
)

// Parse given a string, parse the size of memory it represents ie "1K"= 1024 bytes
func Parse(s string) (int64, error) {
	match := memMatcher.FindStringSubmatch(s)
	if len(match) < 2 {
		return 0, ErrBadFormat
	}
	i, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return 0, err
	}

	// if we did not match a string size i.e. k,m,G then assume the multiple is 1
	var n int64
	if len(match) < 3 {
		n = 1
	} else {
		var ok bool
		// get the multiple from the size string
		n, ok = sizeMap[match[2]]
		if !ok {
			n = 1
		}
	}
	return int64(float64(n) * i), nil
}

// CompareMemory given two size strings, compare them.
// given CompareMemory(x,y)
// 1 : x > y
// 0 : x == y
// -1 : x < y
// -2 : there was an error
func CompareMemory(x, y string) (int, error) {
	var xi, yi int64
	var err error
	xi, err = Parse(x)
	yi, err = Parse(y)

	if err != nil {
		return -2, err
	}
	if xi > yi {
		return 1, nil
	}
	if xi == yi {
		return 0, nil
	}
	if xi < yi {
		return -1, nil
	}
	return -2, nil
}
