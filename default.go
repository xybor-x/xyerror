package xyerror

// Default predefined errors.
var (
	Error               = NewClass("Error")
	IOError             = NewClass("IOError")
	FloatingPointError  = NewClass("FloatingPointError")
	IndexError          = NewClass("IndexError")
	KeyError            = NewClass("KeyError")
	NotImplementedError = NewClass("NotImplementedError")
	ValueError          = NewClass("ValueError")
	ParameterError      = NewClass("ParameterError")
	TypeError           = NewClass("TypeError")
	AssertionError      = NewClass("AssertionError")
)
