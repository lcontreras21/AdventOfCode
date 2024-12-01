package models

import (
	"AdventOfCode/utils"
	"fmt"
)

type Set[T comparable] struct {
    items []T
}

func (s Set[T]) New(items []T) Set[T] {
    new := Set[T]{}
    for _, item := range(items) {
        new.Append(item)
    }
    
    return new
}

func (s *Set[T]) Append(item T) bool {
    if s.Contains(item) {
        return false
    }
    s.items = append(s.items, item)

    return true
}

func (s Set[T]) String() string {
	k := "{"
	for i, v := range s.items {
        if i != 0 {
            k = k + ", "
        }
		k = k + fmt.Sprint(v)
	}
    k = k + "}"
	return k
}

func (s *Set[T]) Clone() Set[T] {
    cloned := Set[T]{items: utils.Clone(s.items)}
    return cloned 
}

func (s *Set[T]) Length() int {
    return len(s.items)
}

func (s *Set[T]) Contains(item T) bool {
    return utils.FindIndex(s.items, item) >= 0
}

func (s *Set[T]) ToArray() []T {
    return s.items
}

func (s *Set[T]) Difference(o Set[T]) []T {
    diff := []T{}
    for _, n := range(s.items) {
        if !o.Contains(n) {
            diff = append(diff, n)
        }
    }
    return diff
}

func (s *Set[T]) Pop(index int) {
    s.items = append(s.items[:index], s.items[index+1:]...)
}
