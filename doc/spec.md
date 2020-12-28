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

The rest of this section describes all instructions of the Watson VM.

### Inew
Inew pushes an Int 0 to the stack.

Pseudo code:

```
push(0);
```

### Iinc
Iinc pops an Int `x` and then pushes `x + 1`.

Pseudo code:

```
x: Int = pop();
push(x + 1);
```

### Ishl
Ishl pops an Int `x` and then pushes `x << 1`;

Pseudo code:

```
x: Int = pop();
push(x << 1);
```

### Iadd
Iadd pops two Ints `x` and `y`, and then pushes `x + y`.

Pseudo code:

```
y: Int = pop();
x: Int = pop();
push(x + y);
```

### Ineg
Ineg pops an Int `x` and then pushes `-x`.

Pseudo code:

```
x: Int = pop();
push(-x);
```

### Isht
Isht pops two Ints `x` and `y`, and then pushes `x << y`.

Pseudo code:

```
y: Int = pop();
x: Int = pop();
push(x << y);
```

### Itof
Itof pops an Int `x` and then converts `x` into a Float by interpreting `x` as an binary representation of IEEE-754 64-bit floating-point number, then pushes this Float value.

Pseudo code:

```
x: Int = pop();
push(x interpreted as an IEEE-754 64-bit floating-point number);
```

### Itou
Itou pops an Int `x` and then converts `x` into an Uint by interpreting `x` as a 64-bit unsigned integer, then pushes this Uint value.

Pseudo code:

```
x: Int = pop();
push(x interpreted as an Uint);
```

### Finf
Finf pushes a positive infinite value of Float.

Pseudo code:

```
push(Inf);
```

### Fnan
Fnan pushes a Float value `NaN`.

Pseudo code:

```
push(NaN);
```

### Fneg
Fneg pops a Float value `x` and then pushes `-x`.

Pseudo code:

```
x: Float = pop();
push(-x);
```

### Snew
Snew pushes an empty String.

Pseudo code:

```
push("");
```

### Sadd
Sadd pops a String `s` and an Int `x`, appends the lowest 8 bits of `x` to `s`, then pushes `s`.

Pseudo code:

```
x: Int = pop();
s: String = pop();
s += (lowest 8 bits of x);
push(t);
```

### Onew
Onew pushes an empty Object.

Pseudo code:

```
push({});
```

### Oadd
Oadd pops an Object `o`, a String `k`, and an arbitrary value `v`, then sets `v` to `o[k]`, and then pushes `o`.

Pseudo code:

```
v: Any = pop();
k: String = pop();
o: Object = pop();
o[k] = v;
push(o);
```

### Anew
Anew pushes an empty Array.

Pseudo code:

```
push([]);
```

### Aadd
Aadd pops an Array `a` and an arbitrary value `x`, appends `x` to `a`, and then pushes `a`.

Pseudo code:

```
x: Any = pop();
a: Array = pop();
a += x;
push(a);
```

### Bnew
Bnew pushes a false.

Pseudo code:

```
push(false);
```

### Bneg
Bneg pops a Bool `x` and pushes `!x`.

Pseudo code:

```
x: Bool = pop();
push(!x);
```

### Nnew
Nnew puses a nil.

Pseudo code:

```
push(nil);
```

### Gdup
Gdup duplicates a value at the top of the stack.

Pseudo code:

```
x: Any = pop();
push(x);
push(x);
```

### Gpop
Gpop pops an arbitrary value from the stack and discards it.

Pseudo code:

```
_: Any = pop();
```

### Gswp
Gswp pops two arbitrary values `x` and `y`, and then pushes them in reverse order.

Pseudo code:

```
y: Any = pop();
x: Any = pop();
push(y);
push(x);
```

If there is no sufficient values in the stack or values on the stack is not what the VM expects, the VM stops execution and reports an error.

## Watson Representation

The correspondence between VM's instructions and its ASCII representation varies depending on the lexer's *mode*.

Each lexer has its own mode. The mode of a lexer is either `A` or `S`. The initial mode of a lexer is `A` unless otherwise specified.

Every time the lexer processes Snew, it flips its mode.

The complete conversion table between instructions and their Watson Representations are as follows:


| insn\mode |A             |S             |
|-----------|--------------|--------------|
|Inew       |B             |S             |
|Iinc       |u             |h             |
|Ishl       |b             |a             |
|Iadd       |a             |k             |
|Ineg       |A             |r             |
|Isht       |e             |A             |
|Itof       |i             |z             |
|Itou       |'             |i             |
|Finf       |q             |m             |
|Fnan       |t             |b             |
|Fneg       |p             |u             |
|Snew       |?             |$             |
|Sadd       |!             |-             |
|Onew       |~             |+             |
|Oadd       |M             |g             |
|Anew       |@             |v             |
|Aadd       |s             |?             |
|Bnew       |z             |^             |
|Bneg       |o             |!             |
|Nnew       |.             |y             |
|Gdup       |E             |/             |
|Gpop       |#             |e             |
|Gswp       |%             |:             |

Any character that is not in this table is simply ignored.

### Examples 1

```
B
```

is converted into

```
Inew
```

since the initial state of the lexer is A and the corresponding instruction of a character `B` is `Inew`.

### Example 2

```
b?b
```

is converted into

```
Ishl Snew Fnan
```

since the lexer changes its mode to S after processing a character `?` and then converts the last character `b` using the `S` column of the conversion table.

