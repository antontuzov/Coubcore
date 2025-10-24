package vm

// Gas represents the gas system for smart contract execution
type Gas struct {
	Limit  uint64 // Maximum gas that can be used
	Used   uint64 // Gas that has been used
	Price  uint64 // Price per unit of gas in wei
	Refund uint64 // Gas to be refunded
}

// GasCosts defines the gas costs for different operations
type GasCosts struct {
	Step   uint64 // Base cost for each instruction
	Push   uint64 // Cost for PUSH operations
	Pop    uint64 // Cost for POP operations
	Add    uint64 // Cost for ADD operations
	Sub    uint64 // Cost for SUB operations
	Mul    uint64 // Cost for MUL operations
	Div    uint64 // Cost for DIV operations
	Eq     uint64 // Cost for EQ operations
	Lt     uint64 // Cost for LT operations
	Gt     uint64 // Cost for GT operations
	Jmp    uint64 // Cost for JMP operations
	JmpIf  uint64 // Cost for JMPIF operations
	Store  uint64 // Cost for STORE operations
	Load   uint64 // Cost for LOAD operations
	Log    uint64 // Cost for LOG operations
	Call   uint64 // Cost for CALL operations
	Create uint64 // Cost for contract creation
	SLoad  uint64 // Cost for storage load
	SStore uint64 // Cost for storage store
}

// DefaultGasCosts returns the default gas costs
func DefaultGasCosts() GasCosts {
	return GasCosts{
		Step:   1,
		Push:   3,
		Pop:    2,
		Add:    3,
		Sub:    3,
		Mul:    5,
		Div:    5,
		Eq:     3,
		Lt:     3,
		Gt:     3,
		Jmp:    8,
		JmpIf:  10,
		Store:  20,
		Load:   20,
		Log:    375,
		Call:   40,
		Create: 32000,
		SLoad:  200,
		SStore: 20000,
	}
}

// NewGas creates a new gas tracker
func NewGas(limit, price uint64) *Gas {
	return &Gas{
		Limit:  limit,
		Used:   0,
		Price:  price,
		Refund: 0,
	}
}

// Consume consumes gas for an operation
func (g *Gas) Consume(amount uint64) bool {
	if g.Used+amount > g.Limit {
		return false // Not enough gas
	}

	g.Used += amount
	return true
}

// Refund adds gas to the refund counter
func (g *Gas) RefundGas(amount uint64) {
	g.Refund += amount
}

// Remaining returns the remaining gas
func (g *Gas) Remaining() uint64 {
	if g.Used >= g.Limit {
		return 0
	}
	return g.Limit - g.Used
}

// RefundAmount returns the amount of gas to be refunded
func (g *Gas) RefundAmount() uint64 {
	return g.Refund
}

// TotalUsed returns the total gas used (including refunds)
func (g *Gas) TotalUsed() uint64 {
	if g.Refund > g.Used {
		return 0
	}
	return g.Used - g.Refund
}

// GetGasCost returns the gas cost for an instruction
func (g *GasCosts) GetGasCost(instruction Instruction) uint64 {
	switch instruction.Op {
	case OP_PUSH:
		return g.Push
	case OP_POP:
		return g.Pop
	case OP_ADD:
		return g.Add
	case OP_SUB:
		return g.Sub
	case OP_MUL:
		return g.Mul
	case OP_DIV:
		return g.Div
	case OP_EQ:
		return g.Eq
	case OP_LT:
		return g.Lt
	case OP_GT:
		return g.Gt
	case OP_JMP:
		return g.Jmp
	case OP_JMPIF:
		return g.JmpIf
	case OP_STORE:
		return g.Store
	case OP_LOAD:
		return g.Load
	case OP_LOG:
		return g.Log
	case OP_CALL:
		return g.Call
	case OP_RET:
		return g.Step // Return has minimal cost
	default:
		return g.Step
	}
}
