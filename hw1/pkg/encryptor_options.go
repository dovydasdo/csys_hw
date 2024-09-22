package pkg

type EncryptorOption func(opts *EncryptorOptions)

type EncryptorOptions struct {
	Input        any
	Capabilities []Capability
}

func GetEncryptorOptions(options ...EncryptorOption) EncryptorOptions {
	opts := EncryptorOptions{}
	for _, opt := range options {
		opt(&opts)
	}
	return opts
}

func EncWithInput(input any) EncryptorOption {
	return func(opts *EncryptorOptions) {
		opts.Input = input
	}
}

func EncWithCapability(cpbl Capability) EncryptorOption {
	return func(opts *EncryptorOptions) {
		opts.Capabilities = append(opts.Capabilities, cpbl)
	}
}
