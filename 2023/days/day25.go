package days

import (
	"AdventOfCode/models"
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
    nodes map[string]models.Set[string]
}

func (g *Graph) Add(v string, new []string) {
    neighbors, exists := g.nodes[v]
    if !exists {
        neighbors = models.Set[string]{}
    }

    for _, n := range(new) {
        neighbors.Append(n)

        other_neighbors, exists := g.nodes[n]
        if !exists {
            other_neighbors = models.Set[string]{}
        }
        other_neighbors.Append(v)
        g.nodes[n] = other_neighbors
    }
    g.nodes[v] = neighbors
}

func (g *Graph) Nodes() models.Set[string] {
    keys := []string{}
    for k := range(g.nodes) {
        keys = append(keys, k)
    }
    return models.Set[string]{}.New(keys) 
}

func (g *Graph) Neighbors(v string) models.Set[string] {
    n, _ := g.nodes[v]
    return n
}

func (g Graph) String() string {
	k := ""
    nodes := g.Nodes()
	for _, v := range(nodes.ToArray()) {
        k = k + v + "\n\t"
        k = k + g.Neighbors(v).String()
        k = k + "\n"
	}
	return k
}

func Day_25_parse_input(use_test_file bool) (Graph) {
	var filename string
	if !use_test_file {
		filename = "inputs/Day_25.txt"
	} else {
		filename = "inputs/temp.txt"
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

    graph := Graph{nodes: map[string]models.Set[string]{}}

	for fileScanner.Scan() {
		txt := fileScanner.Text()
        split := strings.Split(txt, ": ")
        v := split[0]
        n := strings.Split(split[1], " ")
        graph.Add(v, n)
	}

	file.Close()
	return graph
}

func map_count(S models.Set[string], g Graph) []int {
    applied := []int{}
    for _, s := range(S.ToArray()) {
        neigbors := g.Neighbors(s)
        diff := neigbors.Difference(S)
        applied = append(applied, len(diff))
    }
    return applied
}

func three_cut_graph(graph Graph) models.Set[string] {
    // Main Idea: remove node from graph S and place into exclusion group G by measuring biggest amount neighbors
    // Keep removing until each remaining node in S only have three neighbors not in G

    S := graph.Nodes()
    total := map_count(S, graph)
    found := true
    for utils.Sum(total) != 3 {
        _, index := utils.MaxArray(total)
        S.Pop(index)
        total = map_count(S, graph)
        if len(total) < 3 {
            found = false
            break
        }
    }
    if !found {
        // Keep trying until we find the answer
        fmt.Println("repeating")
        S = three_cut_graph(graph)
    }
    return S
}

func Day_25_Part_1() {
    // graph := Day_25_parse_input(true)
    graph := Day_25_parse_input(false)

    S := three_cut_graph(graph)

    nodes := graph.Nodes()
    diff := nodes.Difference(S)
    fmt.Println("Product of the two groups sizes is: ", S.Length() * len(diff))
}
