package number_of_islands

import (
	"fmt"
)

func FindIslands(input [][]string) int {
	party := &searchParty{}

	party.sea = input
	party.seaWidth = len(input[0])
	party.seaLength = len(input)
	party.found = 0
	party.visited = make(map[string]bool)
	party.seaMap = make([][]int, 0, 0)

	party.seaMap = append(party.seaMap, []int{0, 0})

	for len(party.seaMap) > 0 {
		seaSearch := party.seaMap[0]
		party.seaMap = party.seaMap[1:]

		if _, ok := party.visited[hash(seaSearch)]; !ok {
			party.visited[hash(seaSearch)] = true
		} else {
			continue
		}

		if party.isIsland(seaSearch) {
			party.islandMap = make([][]int, 0, 0)

			party.lookAround(seaSearch)

			for len(party.islandMap) > 0 {
				islandSearch := party.islandMap[0]
				party.islandMap = party.islandMap[1:]

				if _, ok := party.visited[hash(islandSearch)]; !ok {
					party.visited[hash(islandSearch)] = true
				} else {
					continue
				}

				party.lookAround(islandSearch)
			}

			party.found++

		} else {
			party.islandMap = nil

			party.lookAround(seaSearch)
		}
	}

	return party.found
}

type searchParty struct {
	sea                 [][]string
	seaWidth, seaLength int
	islandMap, seaMap   [][]int
	visited             map[string]bool
	found               int
}

func (s *searchParty) lookAround(coord []int) {
	if coord[1]+1 < s.seaWidth {
		right := []int{coord[0], coord[1] + 1}
		s.check(right)
	}

	if coord[0]+1 < s.seaLength {
		down := []int{coord[0] + 1, coord[1]}
		s.check(down)
	}

	if coord[0]-1 >= 0 {
		left := []int{coord[0] - 1, coord[1]}
		s.check(left)
	}

	if coord[1]-1 >= 0 {
		up := []int{coord[0], coord[1] - 1}
		s.check(up)
	}
}

func (s *searchParty) check(coord []int) {
	if _, ok := s.visited[hash(coord)]; !ok {
		if s.isIsland(coord) && s.islandMap != nil {
			s.islandMap = append(s.islandMap, coord)
		} else {
			s.seaMap = append(s.seaMap, coord)
		}
	}
}

func (s *searchParty) isIsland(coord []int) bool {
	return s.sea[coord[0]][coord[1]] == "1"
}

func hash(coord []int) string {
	return fmt.Sprintf("%d-%d", coord[0], coord[1])
}
