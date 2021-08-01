package main

type Node struct {
	Val       int     //values inside node
	Neighbors []*Node //edges
}

// func main() {
// 	s := make()
// }

// BFS Go/Golang
func cloneGraph(node *Node) *Node {

	if node == nil {
		return node
	}

	// BFS, using queue data structure
	q := make([]*Node, 0) //type node with unlimited length

	// Map to store visited nodes
	visited := make(map[int]*Node) //the value and the circle

	q = append(q, node)
	newNode := &Node{Val: node.Val}
	visited[node.Val] = newNode

	for len(q) != 0 {
		temp := q[0] // read the first element from the queue
		q = q[1:]    // pop the element out of the queue

		for _, v := range temp.Neighbors {

			// if the node is not visited
			if _, ok := visited[v.Val]; !ok {
				visited[v.Val] = &Node{Val: v.Val}
				q = append(q, v)
			}
			visited[temp.Val].Neighbors = append(visited[temp.Val].Neighbors, visited[v.Val])
		}

	}
	return newNode
}

/*
Slices
---------

append = adds values to the end of the array

slice[low:high]
-- another way to look at it slice[include:not-included]

slice := make([]int, 9)
-- [] int = type of value of the slice
-- 9 = the length of the slice

queue in go
==
queue := make([]int, 0)
// Push to the queue
queue = append(queue, 1)
// Top (just get next element, don't remove it)
x = queue[0]
// Discard top element
queue = queue[1:]
// Is empty ?
if len(queue) == 0 {
    fmt.Println("Queue is empty !")
}



*/
