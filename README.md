# A toy VM implemented in Go

 A VM implemented in Go with a purely education purpose

# Status
- [x] Bytecode interpreter
- [] Program to bytecode compiler
- [] CLI tool
- [] Implement a higher level language compiler with this VM bytecode as backend
- [] Explore some other VM features

# Currently implemented instruction set

* **PSH [value]**, pushes a value into the stack
* **ADD**, takes 2 values from the stack and adds them together
* **POP [register]**, pops a value from the stack to a register 
* **SET [memory address]**, sets the program pointer to the memory address
* **PRINT [register]**, takes to stdout the value of a register
* **MOV [register A] [register B]**, moves a value from A to B
* **HLT**, ends the program

# Program example

```
PSH 5         // Push 5 to the stack, [5]
PSH 6         // Push 6 to the stack, [5,6]
ADD           // Add two values, [11]
POP A         // Push 6 to the stack, [] A{11}
PRINT A       // Prints 11
HLT           // Finishes
```

# Roadmap

* CLI Support
    * `toy-vm-go --compile ./program.ext`, compile a program to bytecode, which will be just 
    the int[] hexadecimal version of the text program (`program.o`)
    
    * `toy-vm-go ./program.o`, run the bytecode in the virtual machine
    
    * `toy-vm-go -i ./program.ext`, compile and run the program like an interpreter
    
* Compiler
* Bytecode Interpreter



