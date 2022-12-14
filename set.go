package util

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"math/rand"
	"time"
)

var randomizer *rand.Rand

func init() {
	src := rand.NewSource(time.Now().Unix())
	randomizer = rand.New(src)
}

// Zero-sized value used as a placeholder
type _dummy = struct{}

var dummy = _dummy{}

type Set[T comparable] struct {
	m map[T]_dummy

	// Snap is an eventual duplicate of set's items stored in a slice,
	// useful when some action requires a sequence of elements,
	// not their map representation
	snap []T
}

func (s *Set[T]) MarshalJSON() ([]byte, error) {
	s.checkSnap()

	b := &bytes.Buffer{}

	enc := json.NewEncoder(b)

	err := enc.Encode(s.snap)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (s *Set[T]) UnmarshalJSON(b []byte) error {
	var slice []T

	dec := json.NewDecoder(bytes.NewReader(b))

	err := dec.Decode(&slice)
	if err != nil {
		return err
	}

	s.Clear()
	for _, v := range slice {
		s.Add(v)
	}

	return nil
}

func (s *Set[T]) GobEncode() ([]byte, error) {
	s.checkSnap()

	b := &bytes.Buffer{}

	enc := gob.NewEncoder(b)

	err := enc.Encode(s.snap)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (s *Set[T]) GobDecode(b []byte) error {
	var slice []T

	dec := gob.NewDecoder(bytes.NewReader(b))

	err := dec.Decode(&slice)
	if err != nil {
		return err
	}

	s.Clear()
	for _, v := range slice {
		s.Add(v)
	}

	return nil
}

func (s *Set[T]) Add(val T) bool {
	_, ok := s.m[val]
	s.m[val] = dummy

	if s.snap != nil {
		s.snap = append(s.snap, val)
	}

	return !ok
}

func (s *Set[T]) Has(val T) bool {
	_, ok := s.m[val]
	return ok
}

func (s *Set[T]) Delete(val T) {
	delete(s.m, val)
	s.snap = nil
}

func (s *Set[T]) checkSnap() {
	if s.snap == nil || len(s.snap) != len(s.m) {
		s.hardSnap()
	}
}

func (s *Set[T]) hardSnap() {
	s.snap = MapKeys(s.m)
}

func (s *Set[T]) Clear() {
	s.m = map[T]_dummy{}
	s.snap = nil
}

func (s *Set[T]) Slice() []T {
	slice := make([]T, 0, len(s.m))

	for k := range s.m {
		slice = append(slice, k)
	}

	return nil
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Equal(s1 *Set[T]) bool {
	if s.Len() != s1.Len() {
		return false
	}

	for k := range s.m {
		if !s1.Has(k) {
			return false
		}
	}

	return true
}

func (s *Set[T]) Range(f func(T) bool) {
	for k := range s.m {
		if !f(k) {
			return
		}
	}
}

func (s *Set[T]) Copy() Set[T] {
	s.checkSnap()

	return MakeSet[T](s.snap...)
}

func (s *Set[T]) Pick() T {
	if s.Len() == 0 {
		panic("Pick() on empty Set!")
	}

	s.checkSnap()

	return s.snap[randomizer.Intn(len(s.snap))]
}

func MakeSet[T comparable](init ...T) Set[T] {
	s := Set[T]{m: map[T]_dummy{}}

	for _, val := range init {
		s.Add(val)
	}

	return s
}
