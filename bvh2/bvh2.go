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
	if rc == contains2.Contains {
		_ = rq
	} else if rc == contains2.Exclude {
		n.Parent = root_node
		h.move_up(n)
	}

	return h
}

func (h *Of[T, V]) move_up(n *node[T, V]) {
	parent := n.Parent

	parent_siblings := []*node[T, V]{}
	for _, pn := range parent.siblings() {
		rc, rq := contains2.RectRect(pn.Bound, n.Bound)
		if rc == contains2.Contains {
			if rq == contains2.Outside {
				n.Children = append(n.Children, pn)
			}
		} else if rc == contains2.Partial {
			parent_siblings = append(parent_siblings, pn)
		}
	}
	parent.Children = parent_siblings

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
	if h.root == nil {
		return func(yield func(V) bool) {}
	}
	return func(yield func(V) bool) {
	iterate:
		for root := h.root; root != nil; {
			rc, _ := contains2.RectVector(root.Bound, point)
			if rc == contains2.Contains {
				if !yield(root.Value) {
					return
				}
				for _, child := range root.Children {
					rc, _ := contains2.RectVector(child.Bound, point)
					if rc == contains2.Contains {
						root = child
						continue iterate
					}
				}
			}

			break
		}
	}
}

func (h *Of[T, V]) put(parent, n *node[T, V]) {
	rc, rq := contains2.RectRect(parent.Bound, n.Bound)
	switch rc {
	case contains2.Contains:
		switch rq {
		case contains2.Inside:
			parent_children := []*node[T, V]{}
			for _, child := range parent.Children {
				crc, crq := contains2.RectRect(child.Bound, n.Bound)
				switch crc {
				case contains2.Contains:
					switch crq {
					case contains2.Inside:
						h.put(child, n)
						return
					case contains2.Outside:
						h.put(n, child)
					}
				case contains2.Partial:
					parent_children = append(parent_children, child)
				}
			}
			parent.Children = append(parent_children, n)
			n.Parent = parent
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
