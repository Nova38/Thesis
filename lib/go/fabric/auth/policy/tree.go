package policy

type Node[T any] struct {
	Path          string
	SubPaths      map[string]*Node[T]
	AllowSubPaths bool
	Policy        *T
}
