package main

type NeighboursFinderI interface {
	Find(index int) []int
}

type NeighboursFinder struct {
	width, height int
}

func NewNeighboursFinder(width, height int) *NeighboursFinder {
	return &NeighboursFinder{width: width, height: height}
}

func (n *NeighboursFinder) Find(index int) (nb []int) {
	switch {
	// corner cases
	case n.topLeft(index):
		nb = append(nb, n.c5(index))
		nb = append(nb, n.c7(index))
		nb = append(nb, n.c8(index))
		return
	case n.topRight(index):
		nb = append(nb, n.c4(index))
		nb = append(nb, n.c6(index))
		nb = append(nb, n.c7(index))
		return
	case n.bottomLeft(index):
		nb = append(nb, n.c2(index))
		nb = append(nb, n.c3(index))
		nb = append(nb, n.c5(index))
		return
	case n.bottomRight(index):
		nb = append(nb, n.c1(index))
		nb = append(nb, n.c2(index))
		nb = append(nb, n.c4(index))
		return
		// side cases
	case n.top(index):
		nb = append(nb, n.c4(index))
		nb = append(nb, n.c5(index))
		nb = append(nb, n.c6(index))
		nb = append(nb, n.c7(index))
		nb = append(nb, n.c8(index))
		return
	case n.left(index):
		nb = append(nb, n.c2(index))
		nb = append(nb, n.c3(index))
		nb = append(nb, n.c5(index))
		nb = append(nb, n.c7(index))
		nb = append(nb, n.c8(index))
		return
	case n.right(index):
		nb = append(nb, n.c1(index))
		nb = append(nb, n.c2(index))
		nb = append(nb, n.c4(index))
		nb = append(nb, n.c6(index))
		nb = append(nb, n.c7(index))
		return
	case n.bottom(index):
		nb = append(nb, n.c1(index))
		nb = append(nb, n.c2(index))
		nb = append(nb, n.c3(index))
		nb = append(nb, n.c4(index))
		nb = append(nb, n.c5(index))
		return
		// everything else
	default:
		nb = append(nb, n.c1(index))
		nb = append(nb, n.c2(index))
		nb = append(nb, n.c3(index))
		nb = append(nb, n.c4(index))
		nb = append(nb, n.c5(index))
		nb = append(nb, n.c6(index))
		nb = append(nb, n.c7(index))
		nb = append(nb, n.c8(index))
	}
	return
}

// functions to check for corners and sides
func (n *NeighboursFinder) topLeft(index int) bool    { return index == 0 }
func (n *NeighboursFinder) topRight(index int) bool   { return index == n.width-1 }
func (n *NeighboursFinder) bottomLeft(index int) bool { return index == n.width*(n.width-1) }
func (n *NeighboursFinder) bottomRight(index int) bool {
	return index == (n.width*n.width)-1
}
func (n *NeighboursFinder) top(index int) bool    { return index < n.width }
func (n *NeighboursFinder) left(index int) bool   { return index%n.width == 0 }
func (n *NeighboursFinder) right(index int) bool  { return index%n.width == n.width-1 }
func (n *NeighboursFinder) bottom(index int) bool { return index >= n.width*(n.width-1) }

// functions to get the index of the neighbours
func (n *NeighboursFinder) c1(index int) int { return index - n.width - 1 }
func (n *NeighboursFinder) c2(index int) int { return index - n.width }
func (n *NeighboursFinder) c3(index int) int { return index - n.width + 1 }
func (n *NeighboursFinder) c4(index int) int { return index - 1 }
func (n *NeighboursFinder) c5(index int) int { return index + 1 }
func (n *NeighboursFinder) c6(index int) int { return index + n.width - 1 }
func (n *NeighboursFinder) c7(index int) int { return index + n.width }
func (n *NeighboursFinder) c8(index int) int { return index + n.width + 1 }
