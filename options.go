package bwt

type OptionsBwt struct {
	// End-of-Text code
	ext string
	// Sample rate of suffix index (compresses the array)
	sampleRate int
}

// End-of-Text code
func WithEndOfText(ext rune) func(*OptionsBwt) {
	return func(s *OptionsBwt) {
		s.ext = string(ext)
	}
}

// Sample rate of suffix index (compresses the array)
func WithSampleRate(mod int) func(*OptionsBwt) {
	return func(s *OptionsBwt) {
		s.sampleRate = mod
	}
}

func (s *OptionsBwt) SampleRate() int {
	return s.sampleRate
}

func buildOptions(options []func(*OptionsBwt)) *OptionsBwt {
	opts := &OptionsBwt{}
	for _, o := range options {
		o(opts)
	}

	if opts.sampleRate <= 0 {
		opts.sampleRate = 1
	}

	if opts.ext == "" {
		opts.ext = "\003"
	}

	return opts
}
