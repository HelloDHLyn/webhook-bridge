package bridge

type Bridge struct {
	Name      string
	Source    InputSource
	Target    OutputTarget
	Converter Converter
}

func NewBridge(name string, source InputSource, target OutputTarget, converter Converter) Bridge {
	return Bridge{
		Name:      name,
		Source:    source,
		Target:    target,
		Converter: converter,
	}
}
