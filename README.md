# WATSON: Wasted but Amazing Turing-incomplete Stack-based Object Notation

![ci status](https://github.com/genkami/watson/workflows/Test/badge.svg)

![WATSON](./doc/img/watson.gif)

(image from [walfie gif](https://walfiegif.wordpress.com/2020/11/05/this-is-true/))

Watson makes it hard but fun to write configuration files.

## Index
* [Language Specification](./doc/spec.md)
* [CLI Tool](./doc/cli.md)
* [pkg.go.dev](https://pkg.go.dev/github.com/genkami/watson)

## Installation

Download latest binaries from [releases](https://github.com/genkami/watson/releases).

Or you can build binaries from source code.

```
$ git clone git@github.com:genkami/watson.git
$ cd watson/cmd/watson
$ go install
```

## Overview of Language Specification
For complete information, please visit [full specification](./doc/spec.md).

Watson internally has a stack-based virtual machine called Watson VM. Each character of Watson files is considered as an instruction to Watson VM.

### Integer
Integer (Int) is a 64-bit signed integer.

Basic instructions for Int are as follows:

* `B` : pushes a zero to the stack
* `u` : increments a value at the top of the stack
* `b` : shifts a value at the top of the stack to the left by one bit
* `a` : adds two values on the top of the stack

You can create arbitrary integers by using these instructions:

```
$ echo 'BBuaBubaBubbbaBubbbbaBubbbbbaBubbbbbba' | watson decode -t json
123
```

### String
String is a byte array.

There are two instructions that manipulate String values:

* `?` : pushes an empty string
* `!` : appends a lowest byte of the top of the stack to a string at the second top of the stack

Every time an empty string is pushed, the ASCII characters used for stack manipulation are updated. The above six instructions `B`, `u`, `b`, `a`, `?`, and `!`, are changed to `S`, `h`, `a`, `k`, `$`, and `-`, respectively.
Pushing an empty string again resets to the orignal characters.

```
$ echo '?SShaakShaaaakShaaaaakShaaaaaak-SShkShaaaaakShaaaaaak-SShkShakShaaakShaaaaakShaaaaaak-SShkShakShaakShaaakShaaaaakShaaaaaak-' | watson decode -t json
"tako"
```

### Object
Object is a set of key-value pairs.

There are two instructions that manipulate Object values:

* `~` : pushes an empty Object
* `M` : pops three values `v`, `k`, `o` in this order, set `o[k] = v`, and then pushes `o`

Note that once `?` is invoked these are changed to `+` and `g` respectively.

## Examples

[You can see more examples here.](https://github.com/genkami/watson/tree/main/examples)

### Hello World

```
$ echo '
~?ShaaaaaarrShaaaaarrkShaaarrk-
SameeShaaaaaarrShaaaaarrkShaarrkShrrk-
ShaaaaaarrShaaaaakSameeShaaarrkShaarrk-
ShaaaaaarrShaaaaarrkShaaarrkShaarrk-
ShaaaaaarrShaaaaarrkShaaarrkShaarrkSharrkShrrk-$
BubbbbbbBubbbbbaBubbbbaBubbaBubaBua!
BubbbbbbBubbbbbaBubbbaBubbaBubaBua!
BubbbbbbBubbbbbaBubbbbaBuba!
BubbbbbbBubbbbbaBubbbaBubba!
BubbbbbbBubbbbbaBubba!M?
ShaaaaaaShaaaaakShaakShak-
ShaaaaaaShaaaaakShaaakShk-
ShaaaaaaShaaaaakShaaaakShak-
ShaaaaaaShaaaaakShaaaakShakShk-
ShaaaaaaShaaaaakShaaaakShaak-
^!!!!!!!!!!!!!g
' | watson decode -t yaml
first: true
hello: world
```

### Deploying Nginx Using Kubernetes

```
$ echo '
@~?
SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-$
BBubbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-Samee+$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!~?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!MMM?
SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-Samee+$
BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaAAME#?
SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!BBubbbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!~?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!MMM?
SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!~?
SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-ggg$
BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!~?
SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-vSamee+$
BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!~?
SShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!?
SShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbaBubbbbbaBubbbbbba!?
SShkSharrkShaarrkShaaarrkShaaaaarrk-ggg$
BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!~?
SShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!?
SShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbaBubbbbbaBubbbbbba!?
SShkSharrkShaarrkShaaarrkShaaaaarrk-ggg$
BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!~?
SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBuaBubaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!?
SShkShaaaarrkShaaaaarrk-SShaaaarrkShaaaaarrk-SShaaaarrkShaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-g$
BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbaBubbbbbaBubbbbbba!?
SShkShaaaarrkShaaaaarrk-SShaaaarrkShaaaaarrk-SShaaaarrkShaaaaarrk-SShkShaarrkShaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-gg$
BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!~?
SShkSharrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-$
BBuaBubbbbaBubbbbba!BBubbbbaBubbbbba!BBubbbbaBubbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!M?
SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-$
BBuaBubbbbaBubbbbba!BBubbbbaBubbbbba!BBubbbbaBubbbbba!BBuaBubbaBubbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!MMM?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbaBubbbbaBubbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-$
BBuaBubbbaBubbbbbba!BBubaBubbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!M?
SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-vSamee+$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBuaBubaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbbaAAME#?
SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-$
BBubbaBubbbbaBubbbbbba!BBuaBubaBubbbbbba!BBubbbbaBubbbbbba!MsMsMMMM?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBubaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbaBubbbbba!Ms~?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbaBubbbbba!M?
SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-$
BBuaBubaBubbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!M?
SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-Samee+$
BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!~?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!M?
SShkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!MM?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbaBubbbbaBubbbbbaBubbbbbba!MM?
SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-Samee+$
BBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!@~?
SShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-$
BBubbaBubbbbaBubbbbbba!BBuaBubaBubbbbbba!BBubbbbaBubbbbbba!M?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-$
BBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!M?
SShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaarrkShaaaaaarrkg$
BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbbaAAME#sM?
SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-Samee+$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-g$
BBuaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBubaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbba!BBuaBubaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!?
SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaarrkShaaaaarrkShaaaaaarrk-gg$
BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbaBubbbbbaBubbbbbba!BBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SShkSharrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaaarrk-SShaaaarrkShaaaaaarrk-gg?
' | watson decode -t yaml | kubectl apply -f
deployment.apps/nginx created
service/nginx created
```

### Function

```
$ cat examples/function/args.watson
?
SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SSharrkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SSharrkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-%

$ cat examples/function/function.watson
+$
BBuaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!BBubaBubbbbaBubbbbbaBubbbbbba!BBubaBubbaBubbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!?
SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SSharrkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SSharrkShaarrkShaaaaarrkShaaaaaarrk-SSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShaaaaarrk-SSharrkShaarrkShaaaaarrkShaaaaaarrk-SShkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrk-SShkSharrkShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkShaaaaarrkShaaaaaarrk-SShaarrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaaaarrkShaaaaarrkShaaaaaarrk-SShkSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-SSharrkShaarrkShaaarrkShaaaaarrkShaaaaaarrk-g
:$BBubaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubbaBubbbbbaBubbbbbba!%M%

$ watson decode -t yaml examples/function/args.watson examples/function/function.watson
anotherValue: this value is loaded from function.watson
value: this value is loaded from args.watson
```

### Converting JSON to YAML via Watson

```
$ echo '{"foo": "bar", "baz": "quux"}' | watson encode -t json | watson decode -t yaml
baz: quux
foo: bar
```

### Converting a Go Struct into Watson

``` go
package main

import (
	"github.com/genkami/watson"
	"os"
)

type User struct {
	FullName string `watson:"fullName"`
	Nickname string `watson:"nickname,omitempty"`
}

func main() {
	user := User{
		FullName: "Motoaki Tanigo",
		Nickname: "YAGOO",
	}
	enc := watson.NewEncoder(os.Stdout)
	err := enc.Encode(&user)
	if err != nil {
		panic(err)
	}
}
```

output:

```
~?SShakShaakShaaaaakShaaaaaak-SShkShaakShaaaakShaaaaakShaaaaaak-SShaakShaaakShaaaaakShaaaaaak-SShaakShaaakShaaaaakShaaaaaak-SShakShaakShaaakShaaaaaak-SShkShaaaaakShaaaaaak-SShkShaakShaaakShaaaaakShaaaaaak-SShkShaakShaaaaakShaaaaaak-$BBuaBubbaBubbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBubbaBubbbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBuaBubaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBubbbbba!BBubbaBubbbbaBubbbbbba!BBuaBubbbbbaBubbbbbba!BBubaBubbaBubbbaBubbbbbaBubbbbbba!BBuaBubbbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbaBubbbbbba!M?SShakShaakShaaakShaaaaakShaaaaaak-SShkShaaakShaaaaakShaaaaaak-SShkShakShaaaaakShaaaaaak-SShkShakShaaakShaaaaakShaaaaaak-SShakShaakShaaakShaaaaakShaaaaaak-SShkShaaaaakShaaaaaak-SShkShaakShaaakShaaaaakShaaaaaak-SShkShaakShaaaaakShaaaaaak-$BBuaBubbbaBubbbbaBubbbbbba!BBuaBubbbbbba!BBuaBubaBubbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbba!BBuaBubaBubbaBubbbaBubbbbbba!M
```

## Related Projects

* [ratson](https://github.com/Herbstein/ratson) - A Rust implementation of Watson
* [WATSON-as-a-Service](https://watson-as-a-service.vercel.app/) - A web API for Watson ([GitHub](https://github.com/jozsefsallai/watson-as-a-service))
* [guralang](https://github.com/indronna/guralang) - A
