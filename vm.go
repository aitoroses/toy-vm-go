package main

import (
	i "github.com/aitoroses/toy-vm-go/vm_instructions"
	vm "github.com/aitoroses/toy-vm-go/vm_struct"
)

func main() {

	program := []int{
		i.PSH, 5,
		i.PSH, 6,
		i.ADD,
		i.POP, vm.A,
		i.PRINT, vm.A,
		i.HLT,
	}

	s := vm.New()

	s.Run(program)
}
