package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	// use bfs algorithm to traverse the graph and get all nodes.
	nodes := getNodes(node)

	// copy nodes, store the old->new mapping information in a hash map
	mapping := make(map[*Node]*Node)

	for _, n := range nodes {
		mapping[n] = &Node{Val: n.Val}
	}

	// copy neighbors(edges)
	for _, n := range nodes {
		newNode := mapping[n]
		for _, neighbor := range n.Neighbors {
			newNeighbor := mapping[neighbor]
			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
		}
	}
	return mapping[node]
}

func getNodes(node *Node) []*Node {
	var queue []*Node
	set := make(map[*Node]bool)
	queue = append(queue, node)
	set[node] = true
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		for _, neighbor := range head.Neighbors {
			if _, ok := set[neighbor]; !ok {
				set[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	var res []*Node
	for k := range set {
		res = append(res, k)
	}
	return res
}

/*
 1.HashMap is used to store all the visited nodes and the key of the cloned node
 HashMap to store the nodes of the original graph.
 Value, stores the corresponding nodes in the clone diagram to prevent them from falling into a
 dead loop, and to obtain the nodes of the clone graph.
2.Add the first node to the queue clone and add the first node to the HashMap named Visited
3.BFS Fetch a node from the head of the queue then traverse all of the
 nodes edges(neighbors)
If an edge point has been visited, the edge point must be in visited, then the edge point is obtained from visited
Otherwise, create a new node to store in visited

--if the negihbor is visited check to see if its in the cloned map, if it isnt add it to the cloned map
*/
