// Copyright 2015 flannel authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netswatch

type zero struct{} // Empty stuct, 0 byte

// Set is a data structure like Python
type Set struct {
	content map[string]zero
}

// NewSet : Create new set
func NewSet() *Set {
	s := &Set{}
	s.content = make(map[string]zero)
	return s
}

// Has : Check item's existence
func (s *Set) Has(v string) bool {
	_, ok := s.content[v]
	return ok
}

// Add : Add item to set
func (s *Set) Add(v string) {
	s.content[v] = zero{}
}

// AddList : Add list to set
func (s *Set) AddList(l *[]string) {
	for _, v := range *l {
		s.Add(v)
	}
}

// Remove : Remove item from set
func (s *Set) Remove(v string) {
	delete(s.content, v)
}

// Size : Retrun size of set
func (s *Set) Size() int {
	return len(s.content)
}

// Clear : Clear set
func (s *Set) Clear() {
	s.content = make(map[string]zero)
}

// Union : Union operation of set
func (s *Set) Union(s2 *Set) *Set {
	ns := NewSet()
	for v := range s.content {
		ns.Add(v)
	}
	for v := range s2.content {
		ns.Add(v)
	}
	return ns
}

// Intersect : Intersect operation of set
func (s *Set) Intersect(s2 *Set) *Set {
	ns := NewSet()
	for v := range s.content {
		if s2.Has(v) {
			ns.Add(v)
		}
	}
	return ns
}

// Difference : Difference operation of set
func (s *Set) Difference(s2 *Set) *Set {
	ns := NewSet()
	for v := range s.content {
		if s2.Has(v) {
			continue
		}
		ns.Add(v)
	}
	return ns
}
