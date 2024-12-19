package models

//import "github.com/FatWang1/fatwang-go-utils/set"
//
//type StateMachine[T, V comparable] struct {
//	EndStep set.Set[T]
//	Transistor[T, V]
//}
//
//type Transition[T, V comparable] struct {
//	Step  V
//	State T
//	Next  []NextStep[V]
//}
//
//type NextStep[V comparable] struct {
//	Step      V
//	Reachable func(step, lastStep V) bool
//}
//
//type Transistor[T, V comparable] interface {
//	Reachable(step, nextStep V) bool
//	GetStep() V
//	GetState() T
//	GetNextStep() []NextStep[V]
//}
