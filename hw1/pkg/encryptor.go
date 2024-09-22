package pkg

type Capability struct {
	Name   string
	CapF   func(input any) (string, error)
	Inputs []any
}

type Encryptor struct {
	Input        any
	Capabilities []Capability
}

func GetEncryptor(opts EncryptorOptions) Encryptor {
	return Encryptor{
		Input:        opts.Input,
		Capabilities: opts.Capabilities,
	}
}

func (e Encryptor) Execute() (map[string][]string, error) {
	result := make(map[string][]string)
	for _, cap := range e.Capabilities {
		capOut := make([]string, 0)
		for _, input := range cap.Inputs {
			out, err := cap.CapF(input)
			if err != nil {
				continue
			}

			capOut = append(capOut, out)
		}
		result[cap.Name] = capOut
	}

	return result, nil
}
