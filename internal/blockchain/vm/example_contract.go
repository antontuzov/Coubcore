package vm

// ExampleContract demonstrates a simple smart contract
var ExampleContract = []Instruction{
	// Push two values onto the stack
	{Op: OP_PUSH, Args: 10},
	{Op: OP_PUSH, Args: 5},

	// Add them together
	{Op: OP_ADD},

	// Store the result in memory
	{Op: OP_STORE, Args: "result"},

	// Push another value
	{Op: OP_PUSH, Args: 3},

	// Load the result from memory
	{Op: OP_LOAD, Args: "result"},

	// Multiply the values
	{Op: OP_MUL},

	// Log the final result
	{Op: OP_LOG},

	// Return
	{Op: OP_RET},
}

// SimpleAddContract demonstrates a simple addition contract
var SimpleAddContract = []Instruction{
	// Push two values onto the stack
	{Op: OP_PUSH, Args: 20.0},
	{Op: OP_PUSH, Args: 30.0},

	// Add them together
	{Op: OP_ADD},

	// Return the result
	{Op: OP_RET},
}

// FibonacciContract demonstrates a recursive-like contract (using loops)
var FibonacciContract = []Instruction{
	// Initialize variables
	{Op: OP_PUSH, Args: 0.0}, // a = 0
	{Op: OP_STORE, Args: "a"},
	{Op: OP_PUSH, Args: 1.0}, // b = 1
	{Op: OP_STORE, Args: "b"},
	{Op: OP_PUSH, Args: 10.0}, // n = 10 (calculate 10th fibonacci number)
	{Op: OP_STORE, Args: "n"},
	{Op: OP_PUSH, Args: 0.0}, // i = 0
	{Op: OP_STORE, Args: "i"},

	// Loop start (label 9)
	// Check if i < n
	{Op: OP_LOAD, Args: "i"},
	{Op: OP_LOAD, Args: "n"},
	{Op: OP_LT},

	// Jump to end if condition is false
	{Op: OP_JMPIF, Args: 27.0}, // Jump to instruction 27 if i >= n

	// Calculate next fibonacci number
	{Op: OP_LOAD, Args: "a"},
	{Op: OP_LOAD, Args: "b"},
	{Op: OP_ADD},

	// Update variables
	{Op: OP_STORE, Args: "temp"}, // temp = a + b
	{Op: OP_LOAD, Args: "b"},     // Load b
	{Op: OP_STORE, Args: "a"},    // a = b
	{Op: OP_LOAD, Args: "temp"},  // Load temp
	{Op: OP_STORE, Args: "b"},    // b = temp

	// Increment i
	{Op: OP_LOAD, Args: "i"},
	{Op: OP_PUSH, Args: 1.0},
	{Op: OP_ADD},
	{Op: OP_STORE, Args: "i"},

	// Jump back to loop start
	{Op: OP_JMP, Args: 9.0},

	// Return the result (b)
	{Op: OP_LOAD, Args: "b"},
	{Op: OP_RET},
}
