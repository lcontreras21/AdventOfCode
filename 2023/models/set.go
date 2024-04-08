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
	k := ""
	for _, i := range s.items {
		k = k + fmt.Sprint(i)
	}
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
    return utils.FindIndex(s.items, item) >= 1
}
