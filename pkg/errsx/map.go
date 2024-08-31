package errsx

import (
	"errors"
	"fmt"
	"strings"
)

// Map представляет собой коллекцию ошибок, индексированных по имени.
type Map map[string]error

func NewMap() *Map {
	return new(Map)
}

// Get вернет строку ошибки для данного ключа.
func (m *Map) Get(key string) string {
	if err := (*m)[key]; err != nil {
		return err.Error()
	}

	return ""
}

func (m *Map) Has(key string) bool {
	_, ok := (*m)[key]

	return ok
}

func (m *Map) Exist() bool {
	if m != nil {
		return len(*m) != 0
	}
	return false
}

// Set ассоциирует данную ошибку с данным ключом.
// Карта создается лениво, если она равна nil.
func (m *Map) Set(key string, msg any) {
	if *m == nil {
		*m = make(Map)
	}

	var err error
	switch msg := msg.(type) {
	case error:
		if msg == nil {
			return
		}

		err = msg

	case string:
		err = errors.New(msg)

	default:
		panic("want error or string message")
	}

	(*m)[key] = err
}

func (m *Map) Error() string {
	if m == nil {
		return "<nil>"
	}

	pairs := make([]string, 0, len(*m))
	for key, err := range *m {
		pairs = append(pairs, fmt.Sprintf("%v: %v", key, err))
	}

	return strings.Join(pairs, "; ")
}

func (m *Map) String() string {
	return m.Error()
}

func (m *Map) MarshalJSON() ([]byte, error) {
	errs := make([]string, 0, len(*m))
	for key, err := range *m {
		errs = append(errs, fmt.Sprintf("%q:%q", key, err.Error()))
	}

	return []byte(fmt.Sprintf("{%v}", strings.Join(errs, ", "))), nil
}

func UnwrapToMap(err error) *Map {
	var errs *Map
	if errors.As(err, &errs) {
		return errs
	}
	return nil
}
