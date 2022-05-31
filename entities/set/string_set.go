package set

type StringSet struct {
	items []string
}

func (s *StringSet) Contains(value string) bool {
	for _, v := range s.items {
		if v == value {
			return true
		}
	}

	return false
}

func (s *StringSet) Add(values ...string) {
	for _, value := range values {
		if s.Contains(value) {
			continue
		}

		s.items = append(s.items, value)
	}
}

func (s *StringSet) Discard(value string) bool {
	for i := 0; i < len(s.items); i++ {
		if s.items[i] == value {
			s.items = append(s.items[:i], s.items[i+1:]...)

			return true
		}
	}

	return false
}

func (s *StringSet) Elements() []string {
	ret := make([]string, len(s.items))

	copy(ret, s.items)

	return ret
}

func (s *StringSet) Size() int {
	return len(s.items)
}

func (s *StringSet) IsEmpty() bool {
	return len(s.items) == 0
}
