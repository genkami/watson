# Watson Specification

Watson (Wasted but Amazing Turing-incomplete Stack-based Object Notation) is a configuration file format like YAML or JSON.

Watson internally has a stack-based virtual machine called Watson VM. Each byte of Watson files are considered as an instruction to the Watson VM.

The correspondence between instructions and its text representation varies depending on the state of Watson's lexer. We call the state *mode* and an ASCII representation of the instructions *Watson Representation*.

## Table of Contents

* [Types](#types)
* [Instructions](#instructions)
* [Watson Representation](#watson-representation)

## Types

Following eight types are available.

* **Int**: a 64-bit signed integer
* **Uint**: a 64-bit unsigned integer
* **Float**: an IEEE-754 64-bit floating-point number
* **String**: a byte array
* **Object**: a set of key-value pairs (key must be String)
* **Array**: an ordered list of values.
* **Bool**: a boolean value
* **Nil**: a null value

## Instructions
As described above, Watson internally has a stack-based virtual machine called Watson VM.
The VM's initial stack is empty. Values represented by Watson files are determined by following steps:

1. Converts contents of a Watson file into a sequence of instructions.
2. Executes all instructions sequentially.
3. Value at the top of the VM's stack is the value that is represented by the Watson file.

## Watson Representation
WIP
