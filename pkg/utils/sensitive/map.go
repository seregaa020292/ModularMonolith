package sensitive

import "encoding/json"

type Map struct {
	Data   map[string]any
	Filers []string
	Mask   string
}

func MapUnmarshal(bytes []byte, mask string, filters []string) (map[string]any, error) {
	m := Map{
		Mask:   mask,
		Filers: filters,
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}
	return m.Data, nil
}

func (m *Map) UnmarshalJSON(bytes []byte) error {
	var temp map[string]any
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	for _, key := range m.Filers {
		if _, ok := temp[key]; ok {
			temp[key] = m.Mask
		}
	}

	m.Data = temp

	return nil
}
