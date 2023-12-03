package classical_singleton

var instance *singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	counter int
}

func GetSingletonInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.counter++
	return s.counter
}
