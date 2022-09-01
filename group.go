package xyerror

// Group is an array of class. It supports to creates a Class inherited from
// many parents.
type Group []Class

// Combine supports creating a group of error classes. This group can be used
// to create the Class with multiparents.
func Combine(cs ...Class) Group {
	return cs
}

// NewClass creates a Class with multiparents.
func (g Group) NewClass(name string) Class {
	var child = NewClass(name)
	child.parent = g
	return child
}
