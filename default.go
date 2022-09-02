package xyerror

// Predefined Exceptions.
var (
	BaseException      = NewException("BaseException")
	IOError            = NewException("IOError")
	FloatingPointError = NewException("FloatingPointError")
	IndexError         = NewException("IndexError")
	KeyError           = NewException("KeyError")
	ValueError         = NewException("ValueError")
	ParameterError     = NewException("ParameterError")
	TypeError          = NewException("TypeError")
	AssertionError     = NewException("AssertionError")
)
