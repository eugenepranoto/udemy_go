package chainofresponsibility

import (
	"fmt"
	"strings"
)

type StringHandler interface {
	SetNext(h StringHandler)
	Process(s string) string
}

type LowerCaseHandler struct {
	next StringHandler
}

func (l *LowerCaseHandler) Process(s string) string {
	s = strings.ToLower(s)
	if l.next != nil {
		l.next.Process(s)
	}
	return s
}

func (l *LowerCaseHandler) SetNext(h StringHandler) {
	l.next = h
}

type SpaceRemovel struct {
	next StringHandler
}

func (s *SpaceRemovel) SetNext(h StringHandler) {
	s.next = h
}

func (s SpaceRemovel) Process(st string) string {
	st = strings.Replace(st, " ", "", -1)

	if s.next != nil {
		s.next.Process(st)
	}
	return st
}

func init() {
	sr := SpaceRemovel{}
	lc := LowerCaseHandler{}
	lc.SetNext(&sr)

	fmt.Println(lc.Process("THE titanic"))
}
