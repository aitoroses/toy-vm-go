package VmStruct

import (
	"errors"
	"fmt"

	Instr "github.com/aitoroses/toy-vm-go/vm_instructions"
)

// VMStruct variable map
type VMStruct struct {
	running   bool
	program   []int
	stack     [256]int
	registers [NumOfRegisters]int
	sp        int
	ip        int
}

// New instance
func New() *VMStruct {
	vm := new(VMStruct)
	vm.sp = -1
	vm.ip = 0
	vm.running = true
	return vm
}

// GetProgramPointer return ip
func (s *VMStruct) GetProgramPointer() int {
	return s.ip
}

// GetStackPointer return sp
func (s *VMStruct) GetStackPointer() int {
	return s.sp
}

// Fetch next instruction
func (s *VMStruct) Fetch() (int, error) {
	if s.ip >= len(s.program) {
		return 0, errors.New("Trying to read an instruction out of program memory, does your program have an HLT instruction?")
	}
	return s.program[s.ip], nil
}

func (s *VMStruct) toHeap(a int) {
	s.sp++
	s.stack[s.sp] = a
}

func (s *VMStruct) fromHeap() int {
	value := s.stack[s.sp]
	s.sp--
	return value
}

func (s *VMStruct) getProgramInst(offset int) int {
	return s.program[s.ip+offset]
}

func (s *VMStruct) next() {
	s.ip++
}

func (s *VMStruct) nextOffset(offset int) {
	s.ip = s.ip + offset
}

func (s *VMStruct) write(register int, value int) {
	s.registers[register] = value
}

func (s *VMStruct) read(register int) int {
	return s.registers[register]
}

// Eval instruction
func (s *VMStruct) Eval(inst int) {
	switch inst {
	case Instr.PSH:
		s.toHeap(s.getProgramInst(1))
		s.next()
		break

	case Instr.POP:
		register := s.getProgramInst(1)
		value := s.fromHeap()
		s.write(register, value)
		s.next()
		break

	case Instr.ADD:
		s.toHeap(s.fromHeap() + s.fromHeap())
		break

	case Instr.SET:
		address := s.getProgramInst(1)
		s.ip = address
		s.next()
		break

	case Instr.PRINT:
		register := s.getProgramInst(1)
		value := s.read(register)
		fmt.Print(value)
		s.next()
		break

	case Instr.MOV:
		register1 := s.getProgramInst(1)
		register2 := s.getProgramInst(2)
		s.write(register2, s.read(register1))
		s.nextOffset(2)
		break

	case Instr.HLT:
		s.running = false
		break
	}
}

// Run a program
func (s *VMStruct) Run(program []int) {
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
