package bridge

import (
	"bytes"
	"regexp"

	"github.com/savaki/jq"
)

type Converter interface {
	Convert(*Input) (*Output, error)
}

type JSONConverter struct {
	convertFn func(*Input) (*Output, error)
}

func NewJSONConverter(syntax []byte) JSONConverter {
	re := regexp.MustCompile(`{{[\w\s/.]+}}`)
	return JSONConverter{
		convertFn: func(input *Input) (*Output, error) {
			result := syntax
			for _, match := range re.FindAll(result, -1) {
				query, err := jq.Parse(string(templateToQuery(match)))
				if err != nil {
					return nil, err
				}

				value, err := query.Apply(input.Payload)
				if err != nil {
					return nil, err
				}
				value = bytes.ReplaceAll(value, []byte("\""), []byte(""))

				result = bytes.Replace(result, match, value, -1)
			}
			return &Output{ConvertedPayload: result}, nil
		},
	}
}

func (c JSONConverter) Convert(input *Input) (*Output, error) {
	return c.convertFn(input)
}

func templateToQuery(template []byte) []byte {
	result := template
	result = bytes.ReplaceAll(result, []byte("{{"), []byte(""))
	result = bytes.ReplaceAll(result, []byte("}}"), []byte(""))
	return result
}
