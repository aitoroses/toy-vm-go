package vm

import (
	"errors"
	"fmt"
)

// Register type
type Register int

// Instruction type
type Instruction int

// VM context
type VM struct {
	running   bool
	program   []Instruction
	stack     [256]int
	registers [NumOfRegisters]Register
	sp        int
	ip        int
}

// New instance
func New() *VM {
	vm := new(VM)
	vm.sp = -1
	vm.ip = 0
	vm.running = true
	return vm
}

// GetProgramPointer return ip
func (s *VM) GetProgramPointer() int {
	return s.ip
}

// GetStackPointer return sp
func (s *VM) GetStackPointer() int {
	return s.sp
}

// Fetch next instruction
func (s *VM) Fetch() (Instruction, error) {
	if s.ip >= len(s.program) {
		return 0, errors.New("Trying to read an instruction out of program memory, does your program have an HLT instruction?")
	}
	return s.program[s.ip], nil
}

func (s *VM) toHeap(a int) {
	s.sp++
	s.stack[s.sp] = a
}

func (s *VM) fromHeap() int {
	value := s.stack[s.sp]
	s.sp--
	return value
}

func (s *VM) getProgramInst(offset int) int {
	return int(s.program[s.ip+offset])
}

func (s *VM) next() {
	s.ip++
}

func (s *VM) nextOffset(offset int) {
	s.ip = s.ip + offset
}

func (s *VM) write(register Register, value int) {
	s.registers[register] = Register(value)
}

func (s *VM) read(register Register) int {
	return int(s.registers[register])
}

// Eval instruction
func (s *VM) Eval(inst Instruction) {
	switch inst {
	case PSH:
		s.toHeap(s.getProgramInst(1))
		s.next()
		break

	case POP:
		register := Register(s.getProgramInst(1))
		value := s.fromHeap()
		s.write(register, value)
		s.next()
		break

	case ADD:
		s.toHeap(s.fromHeap() + s.fromHeap())
		break

	case SET:
		address := s.getProgramInst(1)
		s.ip = address
		s.next()
		break

	case PRINT:
		register := Register(s.getProgramInst(1))
		value := s.read(register)
		fmt.Printf("%d", value)
		s.next()
		break

	case MOV:
		register1 := Register(s.getProgramInst(1))
		register2 := Register(s.getProgramInst(2))
		s.write(register2, s.read(register1))
		s.nextOffset(2)
		break

	case HLT:
		s.running = false
		break
	}
}

// Run a program
func (s *VM) Run(program []Instruction) {
	s.program = program

	// start loop
	for s.running == true {
		instr, err := s.Fetch()
		if err != nil {
			fmt.Println(err.Error())
			s.running = false
			return
		}
		s.Eval(instr)
		s.ip++
	}
}
