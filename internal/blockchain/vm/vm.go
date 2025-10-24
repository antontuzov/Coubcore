package vm

import (
	"encoding/json"
	"fmt"
	"log"
)

// Opcode represents a virtual machine operation
type Opcode string

const (
	OP_PUSH  Opcode = "PUSH"  // Push a value onto the stack
	OP_POP   Opcode = "POP"   // Pop a value from the stack
	OP_ADD   Opcode = "ADD"   // Add two values
	OP_SUB   Opcode = "SUB"   // Subtract two values
	OP_MUL   Opcode = "MUL"   // Multiply two values
	OP_DIV   Opcode = "DIV"   // Divide two values
	OP_EQ    Opcode = "EQ"    // Equality comparison
	OP_LT    Opcode = "LT"    // Less than comparison
	OP_GT    Opcode = "GT"    // Greater than comparison
	OP_JMP   Opcode = "JMP"   // Jump to a specific instruction
	OP_JMPIF Opcode = "JMPIF" // Jump if condition is true
	OP_CALL  Opcode = "CALL"  // Call a function
	OP_RET   Opcode = "RET"   // Return from a function
	OP_STORE Opcode = "STORE" // Store a value in memory
	OP_LOAD  Opcode = "LOAD"  // Load a value from memory
	OP_LOG   Opcode = "LOG"   // Log a value
)

// Instruction represents a single VM instruction
type Instruction struct {
	Op   Opcode      `json:"op"`
	Args interface{} `json:"args,omitempty"`
}

// Contract represents a smart contract
type Contract struct {
	Code    []Instruction          `json:"code"`
	Storage map[string]interface{} `json:"storage"`
	Balance int                    `json:"balance"`
	Address string                 `json:"address"`
}

// Stack represents the VM stack
type Stack struct {
	items []interface{}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top item without removing it
func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// VM represents the virtual machine
type VM struct {
	stack    *Stack
	memory   map[string]interface{}
	contract *Contract
	pc       int // Program counter
	running  bool
	gas      *Gas     // Gas tracker
	gasCosts GasCosts // Gas costs for operations
}

// NewVM creates a new virtual machine
func NewVM(contract *Contract) *VM {
	return &VM{
		stack:    &Stack{items: make([]interface{}, 0)},
		memory:   make(map[string]interface{}),
		contract: contract,
		pc:       0,
		running:  false,
		gas:      NewGas(1000000, 1), // Default gas limit and price
		gasCosts: DefaultGasCosts(),
	}
}

// NewVMWithGas creates a new virtual machine with custom gas settings
func NewVMWithGas(contract *Contract, gasLimit, gasPrice uint64) *VM {
	return &VM{
		stack:    &Stack{items: make([]interface{}, 0)},
		memory:   make(map[string]interface{}),
		contract: contract,
		pc:       0,
		running:  false,
		gas:      NewGas(gasLimit, gasPrice),
		gasCosts: DefaultGasCosts(),
	}
}

// Execute runs the contract code
func (vm *VM) Execute() (interface{}, error) {
	vm.running = true
	vm.pc = 0

	for vm.running && vm.pc < len(vm.contract.Code) {
		instruction := vm.contract.Code[vm.pc]

		// Check gas before executing instruction
		gasCost := vm.gasCosts.GetGasCost(instruction)
		if !vm.gas.Consume(gasCost) {
			return nil, fmt.Errorf("out of gas at PC %d", vm.pc)
		}

		err := vm.executeInstruction(instruction)
		if err != nil {
			return nil, fmt.Errorf("error executing instruction at PC %d: %v", vm.pc, err)
		}

		// Only increment PC if the instruction didn't modify it (for jumps)
		if vm.running && vm.pc == len(vm.contract.Code[:vm.pc+1])-1 {
			vm.pc++
		}
	}

	// Return the top of the stack as the result
	if vm.stack.Size() > 0 {
		return vm.stack.Peek(), nil
	}

	return nil, nil
}

// executeInstruction executes a single instruction
func (vm *VM) executeInstruction(instruction Instruction) error {
	switch instruction.Op {
	case OP_PUSH:
		vm.stack.Push(instruction.Args)

	case OP_POP:
		vm.stack.Pop()

	case OP_ADD:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for ADD operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()

		// Handle different types
		switch a := a.(type) {
		case int:
			if b, ok := b.(int); ok {
				vm.stack.Push(a + b)
			} else {
				return fmt.Errorf("type mismatch: cannot add %T and %T", a, b)
			}
		case float64:
			if b, ok := b.(float64); ok {
				vm.stack.Push(a + b)
			} else {
				return fmt.Errorf("type mismatch: cannot add %T and %T", a, b)
			}
		default:
			return fmt.Errorf("unsupported type for ADD operation: %T", a)
		}

	case OP_SUB:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for SUB operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()

		// Handle different types
		switch a := a.(type) {
		case int:
			if b, ok := b.(int); ok {
				vm.stack.Push(a - b)
			} else {
				return fmt.Errorf("type mismatch: cannot subtract %T and %T", a, b)
			}
		case float64:
			if b, ok := b.(float64); ok {
				vm.stack.Push(a - b)
			} else {
				return fmt.Errorf("type mismatch: cannot subtract %T and %T", a, b)
			}
		default:
			return fmt.Errorf("unsupported type for SUB operation: %T", a)
		}

	case OP_MUL:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for MUL operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()

		// Handle different types
		switch a := a.(type) {
		case int:
			if b, ok := b.(int); ok {
				vm.stack.Push(a * b)
			} else {
				return fmt.Errorf("type mismatch: cannot multiply %T and %T", a, b)
			}
		case float64:
			if b, ok := b.(float64); ok {
				vm.stack.Push(a * b)
			} else {
				return fmt.Errorf("type mismatch: cannot multiply %T and %T", a, b)
			}
		default:
			return fmt.Errorf("unsupported type for MUL operation: %T", a)
		}

	case OP_DIV:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for DIV operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()

		// Handle different types
		switch a := a.(type) {
		case int:
			if b, ok := b.(int); ok {
				if b == 0 {
					return fmt.Errorf("division by zero")
				}
				vm.stack.Push(a / b)
			} else {
				return fmt.Errorf("type mismatch: cannot divide %T and %T", a, b)
			}
		case float64:
			if b, ok := b.(float64); ok {
				if b == 0 {
					return fmt.Errorf("division by zero")
				}
				vm.stack.Push(a / b)
			} else {
				return fmt.Errorf("type mismatch: cannot divide %T and %T", a, b)
			}
		default:
			return fmt.Errorf("unsupported type for DIV operation: %T", a)
		}

	case OP_EQ:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for EQ operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()
		vm.stack.Push(a == b)

	case OP_LT:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for LT operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()

		// Handle different types
		switch a := a.(type) {
		case int:
			if b, ok := b.(int); ok {
				vm.stack.Push(a < b)
			} else {
				return fmt.Errorf("type mismatch: cannot compare %T and %T", a, b)
			}
		case float64:
			if b, ok := b.(float64); ok {
				vm.stack.Push(a < b)
			} else {
				return fmt.Errorf("type mismatch: cannot compare %T and %T", a, b)
			}
		default:
			return fmt.Errorf("unsupported type for LT operation: %T", a)
		}

	case OP_GT:
		if vm.stack.Size() < 2 {
			return fmt.Errorf("stack underflow: not enough values for GT operation")
		}

		b := vm.stack.Pop()
		a := vm.stack.Pop()

		// Handle different types
		switch a := a.(type) {
		case int:
			if b, ok := b.(int); ok {
				vm.stack.Push(a > b)
			} else {
				return fmt.Errorf("type mismatch: cannot compare %T and %T", a, b)
			}
		case float64:
			if b, ok := b.(float64); ok {
				vm.stack.Push(a > b)
			} else {
				return fmt.Errorf("type mismatch: cannot compare %T and %T", a, b)
			}
		default:
			return fmt.Errorf("unsupported type for GT operation: %T", a)
		}

	case OP_JMP:
		if args, ok := instruction.Args.(float64); ok {
			vm.pc = int(args) - 1 // -1 because we increment PC at the end of the loop
		} else {
			return fmt.Errorf("invalid argument for JMP: expected number, got %T", instruction.Args)
		}

	case OP_JMPIF:
		if vm.stack.Size() < 1 {
			return fmt.Errorf("stack underflow: not enough values for JMPIF operation")
		}

		condition := vm.stack.Pop()
		if cond, ok := condition.(bool); ok && cond {
			if args, ok := instruction.Args.(float64); ok {
				vm.pc = int(args) - 1 // -1 because we increment PC at the end of the loop
			} else {
				return fmt.Errorf("invalid argument for JMPIF: expected number, got %T", instruction.Args)
			}
		}

	case OP_STORE:
		if vm.stack.Size() < 1 {
			return fmt.Errorf("stack underflow: not enough values for STORE operation")
		}

		value := vm.stack.Pop()
		if key, ok := instruction.Args.(string); ok {
			vm.memory[key] = value
		} else {
			return fmt.Errorf("invalid argument for STORE: expected string, got %T", instruction.Args)
		}

	case OP_LOAD:
		if key, ok := instruction.Args.(string); ok {
			if value, exists := vm.memory[key]; exists {
				vm.stack.Push(value)
			} else {
				vm.stack.Push(nil)
			}
		} else {
			return fmt.Errorf("invalid argument for LOAD: expected string, got %T", instruction.Args)
		}

	case OP_LOG:
		if vm.stack.Size() < 1 {
			return fmt.Errorf("stack underflow: not enough values for LOG operation")
		}

		value := vm.stack.Pop()
		log.Printf("VM Log: %v", value)
		vm.stack.Push(value) // Push the value back

	case OP_RET:
		vm.running = false

	default:
		return fmt.Errorf("unknown opcode: %s", instruction.Op)
	}

	return nil
}

// DeployContract deploys a new contract
func DeployContract(code []Instruction, initialStorage map[string]interface{}) *Contract {
	contract := &Contract{
		Code:    code,
		Storage: initialStorage,
		Balance: 0,
		Address: generateContractAddress(), // In a real implementation, this would be a hash of the code
	}

	return contract
}

// generateContractAddress generates a simple contract address
// In a real implementation, this would be a cryptographic hash of the contract code
func generateContractAddress() string {
	// Simple implementation for demonstration
	return "contract_0x1234567890abcdef"
}

// SerializeContract serializes a contract to JSON
func SerializeContract(contract *Contract) ([]byte, error) {
	return json.Marshal(contract)
}

// DeserializeContract deserializes a contract from JSON
func DeserializeContract(data []byte) (*Contract, error) {
	var contract Contract
	err := json.Unmarshal(data, &contract)
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

// GetGasUsed returns the amount of gas used
func (vm *VM) GetGasUsed() uint64 {
	return vm.gas.Used
}

// GetGasRemaining returns the remaining gas
func (vm *VM) GetGasRemaining() uint64 {
	return vm.gas.Remaining()
}
