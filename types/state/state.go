/*
 *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 */

package state

import (
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
)

type Code = int64

const SubInit Code = 0
const NoSet Code = -987654321

type State struct {
	code    Code
	targets map[Code]struct{}
	sub     map[Code]*State
}

func Of(code Code, dest ...Code) *State {
	s := &State{
		code:    code,
		targets: nil,
		sub:     nil,
	}
	s.Target(dest...)
	return s
}

func (s *State) Target(dest ...Code) *State {
	if s.targets == nil {
		s.targets = make(map[Code]struct{})
	}
	if len(dest) == 0 {
		return s
	}
	for _, k := range dest {
		s.targets[k] = struct{}{}
	}
	return s
}

func (s *State) RegisterSub(sub ...*State) *errors.Error {
	if s.sub == nil {
		s.sub = make(map[Code]*State)
	}
	if len(sub) == 0 {
		err := errors.Verify("no sub states are registered")
		return err
	}
	for _, i := range sub {
		s.sub[i.code] = i
	}
	return nil
}

func (s *State) MustGetSub(sub Code) (*State, error) {
	if len(s.sub) == 0 {
		err := errors.Verify("no sub states are registered")
		return nil, err
	}
	subState, exists := s.sub[sub]
	if !exists {
		err := errors.Verify(fmt.Sprintf("the corresponding substate [%d] is not registered", sub))
		return nil, err
	}
	return subState, nil
}

func (s *State) Verify(dest Code) bool {
	if s.targets == nil {
		return false
	}
	_, ok := s.targets[dest]
	return ok
}

func (s *State) IncludeSub(sub Code) bool {
	_, err := s.MustGetSub(sub)
	if err != nil {
		return false
	}
	return true
}

func (s *State) VerifySub(subSrc Code, subDest Code) bool {
	subState, err := s.MustGetSub(subSrc)
	if err != nil {
		return false
	}

	return subState.Verify(subDest)
}

type Machine struct {
	states map[Code]*State
}

func NewMachine() *Machine {
	return &Machine{states: make(map[Code]*State)}
}

func (m *Machine) Register(s ...*State) *Machine {
	if len(s) == 0 {
		return m
	}
	for _, i := range s {
		m.states[i.code] = i
	}
	return m
}

func (m *Machine) RegisterSub(src Code, sub ...*State) *errors.Error {
	s, err := m.MustGet(src)
	if err != nil {
		return err
	}
	return s.RegisterSub(sub...)
}

func (m *Machine) MustGet(code Code) (*State, *errors.Error) {
	s, ok := m.states[code]
	if !ok {
		err := errors.Verify(fmt.Sprintf("the status is not registered: %d", code))
		return nil, err
	}
	return s, nil
}

func (m *Machine) Verify(src Code, dest Code) bool {
	if len(m.states) == 0 {
		return false
	}
	s, ok := m.states[src]
	if !ok {
		return false
	}
	return s.Verify(dest)
}

func (m *Machine) IncludeSub(code Code, sub Code) bool {
	s, err := m.MustGet(code)
	if err != nil {
		return false
	}
	return s.IncludeSub(sub)
}

func (m *Machine) VerifySub(code Code, subSrc Code, subDest Code) bool {
	s, err := m.MustGet(code)
	if err != nil {
		return false
	}
	return s.VerifySub(subSrc, subDest)
}
