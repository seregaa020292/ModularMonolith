package sensitive

import "encoding/json"

type Map struct {
	data   map[string]any
	filers []string
	mask   string
}

func MapUnmarshal(bytes []byte, mask string, filters []string) (map[string]any, error) {
	m := Map{
		mask:   mask,
		filers: filters,
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}
	return m.data, nil
}

func (m *Map) UnmarshalJSON(bytes []byte) error {
	var temp map[string]any
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	for _, key := range m.filers {
		if _, ok := temp[key]; ok {
			temp[key] = m.mask
		}
	}

	m.data = temp

	return nil
}
