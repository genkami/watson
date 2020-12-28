# Watson Specification

Watson (Wasted but Amazing Turing-incomplete Stack-based Object Notation) is a configuration file format like YAML or JSON.

Watson internally has a stack-based virtual machine called Watson VM. Each byte of Watson files are considered as an instruction to the Watson VM.

The correspondence between instructions and its text representation varies depending on the state of Watson's lexer. We call the state *mode* and the correspondence *Watson Representation*.

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
WIP

## Watson Representation
WIP
