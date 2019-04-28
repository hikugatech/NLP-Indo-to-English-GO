package dll

func Match_array(array []string, matcher string) bool {
	for _, e := range array {
		if e == matcher {
			return true
		}
	}
	return false
}

func Search_array(array []string, matcher string) int {
	for y, e := range array {
		if e == matcher {
			return y
		}
	}
	return -1
}

func Search_array_all(array []string, matcher string) []int {
	var in []int
	for y, e := range array {
		if e == matcher {
			in = append(in, y)
		}
	}
	return in
}

func Search_array_float32(array []float32, matcher float32) int {
	for y, e := range array {
		if e == matcher {
			return y
		}
	}
	return -1
}

func Match_array2D(array [][]string, matcher string) bool {
	for _, y := range array {
		for _, e := range y {
			if e == matcher {
				return true
			}
		}
	}
	return false
}

func Start_array_bool(jum int) []bool {
	var array []bool
	for i := 0; i < jum; i++ {
		array = append(array, false)
	}
	return array
}

func Start_array_string(jum int) []string {
	var array []string
	for i := 0; i < jum; i++ {
		array = append(array, "")
	}
	return array
}

func Start_array_int(jum int) []int {
	var array []int
	for i := 0; i < jum; i++ {
		array = append(array, 0)
	}
	return array
}

func Start_array_float32(jum int) []float32 {
	var array []float32
	for i := 0; i < jum; i++ {
		array = append(array, -1)
	}
	return array
}

func Remove_array(a []int, i int) []int {
	a = a[:i] // Truncate slice
	return a
}
