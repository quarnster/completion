package content

type Settings struct {
	data map[string]interface{}
}

func NewSettings() *Settings {
	return &Settings{data: make(map[string]interface{})}
}

func (s *Settings) Get(key string) interface{} {
	return s.data[key]
}

func (s *Settings) Set(key string, val interface{}) {
	if s.data == nil {
		s.data = make(map[string]interface{})
	}
	s.data[key] = val
}

func (s *Settings) Clone() *Settings {
	s2 := NewSettings()
	for k, v := range s.data {
		s2.Set(k, v)
	}
	return s2
}

func (s *Settings) Merge(other *Settings) {
	for k, v := range other.data {
		s.data[k] = v
	}
}
