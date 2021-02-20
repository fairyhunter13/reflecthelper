package reflecthelper

// Option is a collection of argument options used in this package.
type Option struct {
	FloatPrecision       int
	FloatFormat          byte
	BitSize              int
	ComplexBitSize       int
	BaseSystem           int
	TimeLayouts          []string
	hasCheckExtractValid bool
}

// FuncOption is a function option to set the Option for function arguments.
type FuncOption func(o *Option)

// WithFloatConfiguration assign the float configuration to the option.
func WithFloatConfiguration(floatPrec int, floatFormat byte) FuncOption {
	return func(o *Option) {
		o.FloatPrecision = floatPrec
		o.FloatFormat = floatFormat
	}
}

// WithBitSize sets the bit size of integer and complex.
func WithBitSize(intFloat, complex int) FuncOption {
	return func(o *Option) {
		o.BitSize = intFloat
		o.ComplexBitSize = complex
	}
}

// WithBaseSystem sets the default base system of Option.
func WithBaseSystem(base int) FuncOption {
	return func(o *Option) {
		o.BaseSystem = base
	}
}

// WithTimeLayouts sets the time layouts for the Option.
func WithTimeLayouts(timeLayouts ...string) FuncOption {
	return func(o *Option) {
		o.TimeLayouts = timeLayouts
	}
}

// NewDefaultOption initialize the new default option.
func NewDefaultOption() *Option {
	return new(Option).Default()
}

// NewOption initialize the new empty option.
func NewOption() *Option {
	return new(Option)
}

// Assign assigns the functional options to the Option.
func (o *Option) Assign(fnOpts ...FuncOption) *Option {
	for _, fnOpt := range fnOpts {
		fnOpt(o)
	}
	o.Default()
	return o
}

func (o *Option) isValidFloatFormat() bool {
	switch o.FloatFormat {
	case 'b', 'e', 'E', 'f', 'g', 'G', 'x', 'X':
		return true
	}
	return false
}

const (
	// DefaultFloatPrecision specifies the default precision used in this package.
	// This is the default maximum precision.
	DefaultFloatPrecision = -1
	// DefaultBitSize is the default bit size used for the conversion in this package.
	DefaultBitSize = 64
	// DefaultComplexBitSize is the default bit size for the complex128 type.
	DefaultComplexBitSize = 128
	// DefaultBaseSystem is the default base system used for decimal in this package.
	DefaultBaseSystem = 10
	// DefaultFloatFormat is the default format used for formmating float value in string
	DefaultFloatFormat = 'g'
)

// Default sets the default value of all variables in Option.
func (o *Option) Default() *Option {
	if o.FloatPrecision <= 0 {
		o.FloatPrecision = DefaultFloatPrecision
	}
	if o.BitSize != 32 && o.BitSize != 64 {
		o.BitSize = DefaultBitSize
	}
	if o.ComplexBitSize != 64 && o.ComplexBitSize != 128 {
		o.ComplexBitSize = DefaultComplexBitSize
	}
	if o.BaseSystem <= 0 {
		o.BaseSystem = DefaultBaseSystem
	}
	if !o.isValidFloatFormat() {
		o.FloatFormat = DefaultFloatFormat
	}
	if o.TimeLayouts == nil {
		o.TimeLayouts = make([]string, 0)
	}
	return o
}

// ResetCheck resets all variables regarding checking.
func (o *Option) ResetCheck() {
	o.hasCheckExtractValid = false
}

func (o *Option) toggleOnCheckExtractValid() {
	o.hasCheckExtractValid = true
}
