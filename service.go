package bridge

import (
	"net/http"
	"strconv"
)

type Service struct {
	Config *Config
}

func NewService(configPath string) (*Service, error) {
	config, err := newConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}

	return &Service{Config: config}, nil
}

func (s *Service) StartServer() error {
	for _, b := range s.Config.Bridges {
		pattern := s.Config.Server.PathPrefix + "/" + b.Name
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			input, err := b.Source.GetInput(r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			output, err := b.Converter.Convert(input)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err = b.Target.SendOutput(output); err != nil {
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			w.WriteHeader(http.StatusOK)
		})
	}

	return http.ListenAndServe(":"+strconv.Itoa(s.Config.Server.Port), nil)
}
