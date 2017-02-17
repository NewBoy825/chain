package protocol

import "chain/protocol/bc"

func topSort(txs []*bc.EntryRef) []*bc.EntryRef {
	if len(txs) == 1 {
		return txs
	}

	nodes := make(map[bc.Hash]*bc.EntryRef)
	for _, tx := range txs {
		nodes[tx.Hash()] = tx
	}

	incomingEdges := make(map[bc.Hash]int)
	children := make(map[bc.Hash][]bc.Hash)
	for node, tx := range nodes {
		hdr := tx.Entry.(*bc.Header)
		spends, _ := hdr.Inputs()
		for _, spRef := range spends {
			sp := spRef.Entry.(*bc.Spend)
			spentOutputID := sp.SpentOutput().Hash()
			if nodes[spentOutputID] != nil {
				if children[spentOutputID] == nil {
					children[spentOutputID] = make([]bc.Hash, 0, 1)
				}
				children[spentOutputID] = append(children[spentOutputID], node)
				incomingEdges[node]++
			}
		}
	}

	var s []bc.Hash
	for node := range nodes {
		if incomingEdges[node] == 0 {
			s = append(s, node)
		}
	}

	// https://en.wikipedia.org/wiki/Topological_sorting#Algorithms
	var l []*bc.EntryRef
	for len(s) > 0 {
		n := s[0]
		s = s[1:]
		l = append(l, nodes[n])

		for _, m := range children[n] {
			incomingEdges[m]--
			if incomingEdges[m] == 0 {
				delete(incomingEdges, m)
				s = append(s, m)
			}
		}
	}

	if len(incomingEdges) > 0 { // should be impossible
		panic("cyclical tx ordering")
	}

	return l
}

func isTopSorted(txs []*bc.EntryRef) bool {
	exists := make(map[bc.Hash]bool)
	seen := make(map[bc.Hash]bool)
	for _, tx := range txs {
		exists[tx.Hash()] = true
	}
	for _, tx := range txs {
		hdr := tx.Entry.(*bc.Header)
		spends, _ := hdr.Inputs()
		for _, spRef := range spends {
			sp := spRef.Entry.(*bc.Spend)
			spentOutputID := sp.SpentOutput().Hash() // xxx ignoring errors
			if exists[spentOutputID] && !seen[spentOutputID] {
				return false
			}
		}
		seen[tx.Hash()] = true
	}
	return true
}
