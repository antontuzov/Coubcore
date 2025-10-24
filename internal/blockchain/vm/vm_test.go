package vm

import (
	"testing"
)

func TestVMExecution(t *testing.T) {
	// Create a simple contract that adds two numbers
	contractCode := []Instruction{
		{Op: OP_PUSH, Args: 5},
		{Op: OP_PUSH, Args: 3},
		{Op: OP_ADD},
		{Op: OP_RET},
	}

	contract := &Contract{
		Code:    contractCode,
		Storage: make(map[string]interface{}),
		Balance: 0,
		Address: "test_contract",
	}

	vm := NewVM(contract)
	result, err := vm.Execute()

	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	if result != 8 {
		t.Errorf("Expected result 8, got %v", result)
	}
}

func TestVMWithGas(t *testing.T) {
	contractCode := []Instruction{
		{Op: OP_PUSH, Args: 10},
		{Op: OP_PUSH, Args: 5},
		{Op: OP_SUB},
		{Op: OP_RET},
	}

	contract := &Contract{
		Code:    contractCode,
		Storage: make(map[string]interface{}),
		Balance: 0,
		Address: "test_contract",
	}

	// Create VM with limited gas
	vm := NewVMWithGas(contract, 100, 1)
	result, err := vm.Execute()

	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	if result != 5 {
		t.Errorf("Expected result 5, got %v", result)
	}

	// Check gas usage
	usedGas := vm.GetGasUsed()
	if usedGas == 0 {
		t.Error("Expected gas to be used")
	}
}

func TestVMOutOfGas(t *testing.T) {
	contractCode := []Instruction{
		{Op: OP_PUSH, Args: 1},
		{Op: OP_PUSH, Args: 2},
		{Op: OP_ADD},
		{Op: OP_RET},
	}

	contract := &Contract{
		Code:    contractCode,
		Storage: make(map[string]interface{}),
		Balance: 0,
		Address: "test_contract",
	}

	// Create VM with very limited gas
	vm := NewVMWithGas(contract, 1, 1)
	_, err := vm.Execute()

	if err == nil {
		t.Error("Expected out of gas error")
	}
}

func TestVMExampleContract(t *testing.T) {
	contract := &Contract{
		Code:    ExampleContract,
		Storage: make(map[string]interface{}),
		Balance: 0,
		Address: "example_contract",
	}

	vm := NewVM(contract)
	result, err := vm.Execute()

	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	// The example contract calculates (10 + 5) * 3 = 45
	if result != 45 {
		t.Errorf("Expected result 45, got %v", result)
	}
}

func TestSimpleAddContract(t *testing.T) {
	contract := &Contract{
		Code:    SimpleAddContract,
		Storage: make(map[string]interface{}),
		Balance: 0,
		Address: "simple_add_contract",
	}

	vm := NewVM(contract)
	result, err := vm.Execute()

	if err != nil {
		t.Fatalf("Execution failed: %v", err)
	}

	// The simple add contract calculates 20 + 30 = 50
	// Since we're using float64, the result will be float64(50)
	if result != float64(50) {
		t.Errorf("Expected result 50, got %v (type %T)", result, result)
	}
}
