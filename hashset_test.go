package setutil

import (
	"testing"
)

func TestHashSetAdd(t *testing.T) {
	s := NewHashSet(1)
	s.Add("a")
	if s.Size() != 1 {
		t.Fatal("HashSet size should be 1")
	}
	if !s.Contains("a") {
		t.Fatal("HashSet contains should be true")
	}
}

func TestSetNilAdd(t *testing.T) {
	s := NewHashSet(1)
	s.Add(nil)
	if s.Size() != 1 {
		t.Fatal("HashSet size should be 1")
	}
	if !s.Contains(nil) {
		t.Fatal("HashSet contains should be true")
	}
}

func TestHashSetClear(t *testing.T) {
	s := NewHashSet(1)
	s.Add("a")
	if s.Size() != 1 {
		t.Fatal("HashSet size should be 1")
	}
	s.Clear()
	if s.Size() != 0 {
		t.Fatal("HashSet size should be 0")
	}
}

func TestHashSetDifference(t *testing.T) {
	s1 := NewHashSet(2)
	s1.Add("a")
	s1.Add("b")
	s2 := NewHashSet(2)
	s2.Add("a")
	s2.Add("c")
	diff := s1.Difference(s2)
	if ok := diff.Contains("b"); !ok {
		t.Fatal("HashSet diff should be true: expected: &{0 map[b:true]} actual: ", diff)
	}
	if ok := diff.ContainsAll("a", "c"); ok {
		t.Fatal("HashSet diff should be false: expected: &{0 map[b:true]} actual: ", diff)
	}
}

func TestHashSetEquals(t *testing.T) {
	s1 := NewHashSet(1)
	s1.Add("a")
	s2 := NewHashSet(1)
	s2.Add("a")
	if !s1.Equals(s2) {
		t.Fatal("HashSet equals should be true: ", s1, s2)
	}
	s3 := NewHashSet(1)
	s3.Add("b")
	if s1.Equals(s3) {
		t.Fatal("HashSet equals should be false: ", s1, s3)
	}
}

func TestIntersection(t *testing.T) {
	s1 := NewHashSet(1)
	s1.Add("a")
	s1.Add("b")
	s2 := NewHashSet(1)
	s2.Add("a")
	s2.Add("c")
	intersection := s1.Intersection(s2)
	if ok := intersection.ContainsAll("a"); !ok {
		t.Fatal("HashSet intersect should be true: expected: &{0 map[a:true]} actual: ", intersection)
	}
	if intersection.Size() != 1 {
		t.Fatal("HashSet size should be 1:", intersection)
	}
}

func TestHashSetRemove(t *testing.T) {
	s1 := NewHashSet(1)
	s1.Add("a")
	if !s1.Contains("a") {
		t.Fatal("HashSet equals should be true: ", s1)
	}
	s1.Remove("a")
	if s1.Contains("a") {
		t.Fatal("HashSet equals should be false: ", s1)
	}
}

func TestHashSetUnion(t *testing.T) {
	s1 := NewHashSet(2)
	s1.Add("a")
	s1.Add("b")
	s2 := NewHashSet(2)
	s2.Add("a")
	s2.Add("c")
	union := s1.Union(s2)
	if ok := union.ContainsAll("a", "b", "c"); !ok {
		t.Fatal("HashSet diff should be true: expected: &{0 map[a:true, b:true, c:true]} actual: ", union)
	}
	if union.Size() != 3 {
		t.Fatal("HashSet size should be 3:", union)
	}
}
