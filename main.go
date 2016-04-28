package main

import (
	"github.com/aitoroses/toy-vm-go/vm"
)

func main() {
	
	program := []vm.Instruction{
		vm.PSH, 5,
		vm.PSH, 6,
		vm.ADD,
		vm.POP, vm.A,
		vm.PRINT, vm.A,
		vm.HLT,
	}

	s := vm.New()

	s.Run(program)
}
