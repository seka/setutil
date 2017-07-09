package setutil

// HashSet implements the Set interface, backed by a hash table (actually a Map).
// HashSet permits the null element.
type HashSet struct {
	hint  int
	items map[interface{}]bool
}

// NewHashSet is sole constructor.
func NewHashSet(hint int) Set {
	return &HashSet{
		hint:  hint,
		items: make(map[interface{}]bool, hint),
	}
}

func (s *HashSet) Add(item interface{}) bool {
	if s.Contains(item) {
		return false
	}
	s.items[item] = true
	return true
}

func (s *HashSet) Clear() {
	s.items = make(map[interface{}]bool, s.hint)
}

func (s *HashSet) Contains(item interface{}) bool {
	_, ok := s.items[item]
	return ok
}

func (s *HashSet) ContainsAll(items ...interface{}) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s *HashSet) Difference(other Set) Set {
	difference := NewHashSet(0)
	for item := range s.items {
		if other.Contains(item) {
			continue
		}
		difference.Add(item)
	}
	return difference
}

func (s *HashSet) Equals(other Set) bool {
	for item := range s.items {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

func (s *HashSet) Intersection(other Set) Set {
	intersect := NewHashSet(0)
	for item := range s.items {
		if !other.Contains(item) {
			continue
		}
		intersect.Add(item)
	}
	return intersect
}

func (s *HashSet) Remove(item interface{}) {
	delete(s.items, item)
}

func (s *HashSet) Size() int {
	return len(s.items)
}

func (s *HashSet) Union(other Set) Set {
	union := NewHashSet(0)
	for item := range s.items {
		union.Add(item)
	}
	o := other.(*HashSet)
	for item := range o.items {
		union.Add(item)
	}
	return union
}

var _ Set = (*HashSet)(nil)
