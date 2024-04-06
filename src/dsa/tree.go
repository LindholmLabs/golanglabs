package dsa

type Node struct {
	value *int
	left  *Node
	right *Node
}

type Tree struct {
	n *Node
}

// Insert places the element in the first available slot.
func (t *Tree) Insert(elem int) {
	newNode := &Node{value: &elem}
	if t.n == nil {
		t.n = newNode
		return
	}

	parent := t.Bfs(nil) // Find the parent node to insert the new node
	if parent != nil {
		if parent.left == nil {
			parent.left = newNode
		} else if parent.right == nil {
			parent.right = newNode
		}
	}
}

// Bfs performs a breadth-first search to find a node with an available slot.
func (t *Tree) Bfs(target *int) *Node {
	var q Queue
	q.Enqueue(t.n) // Enqueue the root node

	for !q.IsEmpty() {

		currentInterface, _ := q.Dequeue()
		current, ok := currentInterface.(*Node)

		if !ok {
			return nil
		}

		// Return this node if it has space for a new child
		if current.left == nil || current.right == nil {
			return current
		}

		// Otherwise, enqueue its children to continue the search
		if current.left != nil {
			q.Enqueue(current.left)
		}
		if current.right != nil {
			q.Enqueue(current.right)
		}
	}

	return nil // Return nil if no suitable parent is found (should not happen)
}
