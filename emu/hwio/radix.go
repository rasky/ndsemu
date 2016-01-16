package hwio

import "errors"

const (
	cRadixWidth = 8

	cRadixNumNodes   = 1 << cRadixWidth
	cRadixStartShift = 32 - cRadixWidth
	cRadixMask       = (1 << cRadixWidth) - 1
)

var (
	ErrOverlappingRange = errors.New("overlapping range")
)

type radixNode struct {
	children [cRadixNumNodes]interface{}
}

type radixTree struct {
	root radixNode
}

func (t *radixTree) Search(key uint32) interface{} {
	// CAUTION: this code is really hot, and has been hand-optimized to be as
	// fast as possible. Evaluate any changes against the generated assembly.
	var ok bool
	node := &t.root
	for {
		k := uint8(key >> cRadixStartShift)
		c := node.children[k]
		if c == nil {
			return nil
		}
		if node, ok = c.(*radixNode); !ok {
			return c
		}
		key <<= cRadixWidth
	}
}

func (node *radixNode) insert(shift uint, begin, end uint32, v interface{}) error {
	b, e := (begin>>shift)&cRadixMask, (end>>shift)&cRadixMask

	lowmask := ((uint32(1) << shift) - 1)
	putleaf := false
	if shift == 0 || (begin&lowmask == 0 && (end+1)&lowmask == 0) {
		putleaf = true
	}

	for i := b; i <= e; i++ {
		child := node.children[i]

		if putleaf {
			if child != nil && child != v {
				return ErrOverlappingRange
			}
			node.children[i] = v
		} else {
			n2, ok := child.(*radixNode)
			if !ok {
				if child != nil {
					return ErrOverlappingRange
				}
				n2 = &radixNode{}
				node.children[i] = n2
			}
			if err := n2.insert(shift-cRadixWidth, begin, end, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (node *radixNode) remove(shift uint, begin, end uint32) {
	b, e := (begin>>shift)&cRadixMask, (end>>shift)&cRadixMask

	lowmask := ((uint32(1) << shift) - 1)
	leaf := false
	if shift == 0 || (begin&lowmask == 0 && (end+1)&lowmask == 0) {
		leaf = true
	}

	for i := b; i <= e; i++ {
		if leaf {
			node.children[i] = nil
		} else {
			if n2, ok := node.children[i].(*radixNode); ok {
				n2.remove(shift-cRadixWidth, begin, end)
			} else {
				node.children[i] = nil
			}
		}
	}
}

func (t *radixTree) InsertRange(begin, end uint32, v interface{}) error {
	return t.root.insert(cRadixStartShift, begin, end, v)
}

func (t *radixTree) RemoveRange(begin, end uint32) {
	t.root.remove(cRadixStartShift, begin, end)
}
