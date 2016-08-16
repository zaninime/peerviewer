package main

import "errors"

type configRoot struct {
	HTTP    configHTTP     `json:"http"`
	Streams []configStream `json:"streams"`
}

type configHTTP struct {
	Listen string `json:"listen"`
}

type configStream struct {
	Description string                 `json:"description"`
	Listen      string                 `json:"listen"`
	Kind        configStreamKindString `json:"kind"`
}

type configStreamKind int

type configStreamKindString struct {
	Value configStreamKind
}

const (
	configStreamKindVideoWebM configStreamKind = iota
	configStreamKindAudioOpus
)

var errConfigInvalidStreamKind = errors.New("config: invalid stream kind")

func (s configStreamKindString) MarshalJSON() ([]byte, error) {
	m := map[configStreamKind][]byte{
		configStreamKindVideoWebM: []byte(`"video-webm"`),
		configStreamKindAudioOpus: []byte(`"audio-opus"`),
	}
	v, ok := m[s.Value]
	if !ok {
		return nil, errConfigInvalidStreamKind
	}
	return v, nil
}

func (s *configStreamKindString) UnmarshalJSON(value []byte) error {
	m := map[string]configStreamKind{
		"video-webm": configStreamKindVideoWebM,
		"opus-webm":  configStreamKindAudioOpus,
	}
	v, ok := m[string(value)]
	if !ok {
		return errConfigInvalidStreamKind
	}
	s.Value = v
	return nil
}

var configDefault = configRoot{
	HTTP: configHTTP{
		Listen: ":8080",
	},
	Streams: []configStream{
		configStream{
			Description: "Channel 1",
			Listen:      ":60006",
			Kind:        configStreamKindString{Value: configStreamKindVideoWebM},
		},
		configStream{
			Description: "Radio 1",
			Listen:      ":60007",
			Kind:        configStreamKindString{Value: configStreamKindAudioOpus},
		}},
}
