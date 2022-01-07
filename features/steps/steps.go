package steps

import (
	"errors"

	"github.com/ess/kennel"
)

var unimplemented = errors.New("this step is not implemented")
var facts = NewFactis()

type Factis struct {
	facts map[string]interface{}
}

func NewFactis() *Factis {
	f := &Factis{}
	f.Reset()
	return f
}

func (f *Factis) Memorize(name string, fact interface{}) error {
	if _, ok := f.facts[name]; ok == true {
		return errors.New("cannot memorize the fact '" + name + "' more than once")
	}

	f.facts[name] = fact
	return nil
}

func (f *Factis) Recall(name string) (interface{}, error) {
	if _, ok := f.facts[name]; ok == false {
		return nil, errors.New("cannot recall unknown fact '" + name + "'")
	}

	return f.facts[name], nil
}

func (f *Factis) Forget(name string) (interface{}, error) {
	fact, err := f.Recall(name)
	if err != nil {
		return nil, err
	}

	delete(f.facts, name)

	return fact, nil
}

func (f *Factis) Reset() {
	f.facts = make(map[string]interface{})
}

func Register() {
}

type factisSetup struct{}

func (steps *factisSetup) StepUp(s kennel.Suite) {
	s.BeforeScenario(func(interface{}) {
		facts.Reset()
	})
}

func init() {
	kennel.Register(
		&factisSetup{},
	)
}
