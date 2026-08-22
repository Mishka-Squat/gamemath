package bvh2

import (
	"iter"

	"github.com/Mishka-Squat/gamemath/contains2"
	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/deque"
	"github.com/Mishka-Squat/goex/mathex"
)

type node[T mathex.SignedNumber, V any] struct {
	Parent   *node[T, V]
	Children []*node[T, V]
	Bound    rect2.Of[T]
	Value    V
	// virtual nodes are synthetic bounding boxes introduced to wrap two
	// entries that don't nest inside one another; they carry no Value of
	// their own and Query skips yielding for them.
	virtual bool
}

func (n node[T, V]) siblings() []*node[T, V] {
	if n.Parent == nil {
		return []*node[T, V]{}
	}
	return n.Parent.Children
}

type Of[T mathex.SignedNumber, V any] struct {
	nodes deque.Of[node[T, V]]
	root  *node[T, V]
}

func (h Of[T, V]) Append(bound rect2.Of[T], value V) Of[T, V] {
	h.nodes = h.nodes.Append(node[T, V]{
		Parent: nil,
		Bound:  bound,
		Value:  value,
	})

	n := h.nodes.Last()
	if h.root == nil {
		h.root = n
		return h
	}

	h.put(h.root, n)

	root_node := h.root

	rc, rq := contains2.RectRect(root_node.Bound, bound)
	switch rc {
	case contains2.Equal:
		break
	case contains2.Contains:
		if rq == contains2.Outside {
			// bound fully wraps the current root: n's own bound already covers
			// everything below it, so it can become the new root directly.
			n.Parent = root_node
			h.move_up(n)
		}
	default:
		// bound only partially overlaps, or is disjoint from, the current
		// root, so neither can be a descendant of the other. Wrap both
		// under a new virtual root whose bound is their union, so the root
		// always genuinely contains everything beneath it.
		h.wrap(root_node, n, root_node.Bound.Union(bound))
	}

	return h
}

// wrap introduces a new virtual root node with the given bound, making a
// and b its two children.
func (h *Of[T, V]) wrap(a, b *node[T, V], bound rect2.Of[T]) {
	h.nodes = h.nodes.Append(node[T, V]{
		Bound:   bound,
		virtual: true,
	})
	w := h.nodes.Last()
	w.Children = []*node[T, V]{a, b}
	a.Parent = w
	b.Parent = w
	h.root = w
}

func (h *Of[T, V]) move_up(n *node[T, V]) {
	parent := n.Parent

	parent_siblings := []*node[T, V]{}
	for _, pn := range parent.siblings() {
		rc, rq := contains2.RectRect(pn.Bound, n.Bound)
		if rc == contains2.Contains && rq == contains2.Outside {
			pn.Parent = n
			n.Children = append(n.Children, pn)
		} else {
			// pn isn't wrapped by n (Partial or Exclude): it stays exactly
			// where it was, as a child of the grandparent.
			parent_siblings = append(parent_siblings, pn)
		}
	}

	if parent.Parent != nil {
		parent.Parent.Children = parent_siblings
	}

	parent.Parent = n
	n.Children = append(n.Children, parent)

	if parent == h.root {
		h.root = n
	}
}

func (h Of[T, V]) Query(point vector2.Of[T]) iter.Seq[V] {
	return func(yield func(V) bool) {
		if h.root == nil {
			return
		}
		queryNode(h.root, point, yield)
	}
}

// queryNode visits n and, if its bound contains point, every child whose
// bound also contains point (there may be more than one: siblings can
// overlap). It returns false once yield asks the caller to stop.
func queryNode[T mathex.SignedNumber, V any](n *node[T, V], point vector2.Of[T], yield func(V) bool) bool {
	rc, _ := contains2.RectVector(n.Bound, point)
	if rc != contains2.Contains {
		return true
	}

	if !n.virtual {
		if !yield(n.Value) {
			return false
		}
	}

	for _, child := range n.Children {
		if !queryNode(child, point, yield) {
			return false
		}
	}

	return true
}

func (h *Of[T, V]) split_children(n *node[T, V], children []*node[T, V]) ([]*node[T, V], bool) {
	orphaned_children := []*node[T, V]{}
	for _, child := range children {
		crc, crq := contains2.RectRect(child.Bound, n.Bound)
		switch crc {
		case contains2.Equal:
			h.put(child, n)
			return nil, false
		case contains2.Contains:
			switch crq {
			case contains2.Inside:
				h.put(child, n)
				return nil, false
			case contains2.Outside:
				h.put(n, child)
			}
		default:
			// child isn't wrapped by n (Partial or Exclude): it
			// stays exactly where it was, as a child of parent.
			orphaned_children = append(orphaned_children, child)
		}
	}

	return orphaned_children, true
}

func (h *Of[T, V]) put(parent, n *node[T, V]) {
	rc, rq := contains2.RectRect(parent.Bound, n.Bound)
	switch rc {
	case contains2.Equal:
		if parent_children, ok := h.split_children(n, parent.Children); ok {
			parent.Children = append(parent_children, n)
			n.Parent = parent
		}
	case contains2.Contains:
		switch rq {
		case contains2.Inside: // n is inside parent bounds
			if parent_children, ok := h.split_children(n, parent.Children); ok {
				parent.Children = append(parent_children, n)
				n.Parent = parent
			}
		case contains2.Outside:
			n.Parent = parent
			h.move_up(n)
		}
	case contains2.Exclude:
		// add as sibling here
		// possibly create virtual parent, or convert to global virtual root
		//n.Parent = parent
		//h.move_up(n)
	}
}
