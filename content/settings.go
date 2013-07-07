package content

type Settings struct {
	data   map[string]interface{}
	parent *Settings
}

func NewSettings() *Settings {
	return &Settings{data: make(map[string]interface{})}
}

func (s *Settings) Get(key string) interface{} {
	if v, ok := s.data[key]; ok {
		return v
	} else if s.parent != nil {
		return s.parent.Get(key)
	}
	return nil
}

func (s *Settings) Set(key string, val interface{}) {
	if s.data == nil {
		s.data = make(map[string]interface{})
	}
	s.data[key] = val
}
