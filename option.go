package reflecthelper

// Option is a collection of argument options used in this package.
type Option struct {
	FloatPrecision int
	FloatFormat    byte
	BitSize        int
	ComplexBitSize int
	BaseSystem     int
	TimeLayouts    []string

	// Toggle Flag
	hasCheckExtractValid  bool
	IgnoreError           bool
	RecoverPanic          bool
	BlockChannelIteration bool
	ConcurrentMode        bool
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

// WithIgnoreError toggles for ignoring error in the struct, slice, array, and map iteration.
// The default behavior for this package is false.
func WithIgnoreError(input bool) FuncOption {
	return func(o *Option) {
		o.IgnoreError = input
	}
}

// WithPanicRecoverer toggles the panic recoverer in all of the packages' functions.
// The default behavior for this package is false.
func WithPanicRecoverer(input bool) FuncOption {
	return func(o *Option) {
		o.RecoverPanic = input
	}
}

// WithBlockChannel toggles the blocking operation of receive from reflect.Value with kind reflect.Chan.
// WithBlockChannel will use the Recv() method instead of TryRecv().
// The default behavior for this package is false.
func WithBlockChannel(input bool) FuncOption {
	return func(o *Option) {
		o.BlockChannelIteration = input
	}
}

// WithConcurrency toggles the concurrency mode in this package, especially in the iteration.
// This toggles to iterate array, slice, map, or struct elements in concurrent mode.
// The default behavior for this package is false.
func WithConcurrency(input bool) FuncOption {
	return func(o *Option) {
		o.ConcurrentMode = input
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
	return o.Default()
}

// Clone clones the current option to the new memory address.
func (o *Option) Clone() *Option {
	newOpt := *o
	return &newOpt
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
	return o.resetCheck()
}

func (o *Option) resetCheck() *Option {
	o.hasCheckExtractValid = false
	return o
}

func (o *Option) toggleOnCheckExtractValid() *Option {
	o.hasCheckExtractValid = true
	return o
}
