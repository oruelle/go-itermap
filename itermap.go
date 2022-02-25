package itermap

// IndexOf : find index of sub in slice
// return index , true
// otherwise return 0, false
func IndexOf(slice []interface{}, sub interface{}) (int, bool) {
	// Find s in a, r
	n := len(slice) - 1
	if n == -1 {
		return 0, false
	}
	for i, v := range slice {
		switch {
		case v == sub:
			return i, true
		case slice[n-i] == sub:
			return n - i, true
		}
	}
	return 0, false
}

type IterMap struct {
	key   []interface{}
	value []interface{}
	index int
	err   error
}

func (s *IterMap) Set(k, v interface{}) {
	i, ok := IndexOf(s.key, k)
	if ok {
		s.value[i] = v
	} else {
		s.key = append(s.key, k)
		s.value = append(s.value, v)
	}
}

func (s *IterMap) Get(k interface{}) (interface{}, bool) {
	i, ok := IndexOf(s.key, k)
	if ok {
		return s.value[i], true
	} else {
		return nil, false
	}
}

func (s *IterMap) StartIter() {
	s.index = -1
}

func (s *IterMap) Next() bool {
	s.index++
	return s.index < len(s.key)
}

func (s *IterMap) Value() (interface{}, interface{}) {
	if s.index < len(s.key) {
		return s.key[s.index], s.value[s.index]
	}

	return "", nil
}

func (s *IterMap) Err() error {
	return s.err
}
