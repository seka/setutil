package setutil

// A Set containing a single member
type Set interface {
	// Add the specified item to this set if it is not already present (optional operation).
	Add(item interface{}) bool

	// Clear all of the item from this set (optional operation).
	Clear()

	// Returns true if this set contains the specified item.
	Contains(item interface{}) bool

	// Returns true if this set contains all of the item of the specified collection.
	ContainsAll(items ...interface{}) bool

	// The returned Set contains all items of a that are not a member of b.
	Difference(other Set) Set

	// Compares the specified object with this set for equality.
	Equals(other Set) bool

	// The returned set that are members of both input set (this and other).
	Intersection(other Set) Set

	// Removes the specified item from this set if it is present (optional operation).
	Remove(item interface{})

	// Returns the number of items in this set (its cardinality).
	Size() int

	// The returned set contains all items of this and other.
	Union(other Set) Set
}
