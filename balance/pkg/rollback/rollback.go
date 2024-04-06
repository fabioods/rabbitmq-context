package rollback

import (
	"context"
	"log"
)

type Func struct {
	Name string
	Func func()
}

type IRollback interface {
	Add(name string, function func()) IRollback
	Do(ctx context.Context) []string
}

type Rollback struct {
	functions []Func
}

func (s *Rollback) Add(name string, function func()) IRollback {
	s.functions = append(s.functions, Func{
		Name: name,
		Func: function,
	})
	return s
}

func (s *Rollback) Do(ctx context.Context) []string {
	callFuncName := make([]string, 0)
	for i := len(s.functions) - 1; i >= 0; i-- {
		item := s.functions[i]
		log.Println("Rollback: ", item.Name)
		item.Func()
		callFuncName = append(callFuncName, item.Name)
	}
	return callFuncName
}

func New() *Rollback {
	return &Rollback{
		functions: make([]Func, 0),
	}
}
