# CLI Tool

All subcommands:

* [watson encode](#watson-encode)
* [watson decode](#watson-decode)

## watson encode

### Usage

```
watson encode -t=TYPE [-initial-mode=MODE] [FILE]
```

Converts `FILE` of type `TYPE` into Watson and outputs its Watson Representation to the standard output.

If `FILE` is not specified, it uses the standard input.

### Flags

| flag | mandatory | type | default | description |
| ---- | --------- | ---- | ------- | ----------- |
| **-t**    | no        | `json`, `yaml`, `msgpack`, or `cbor` | `yaml` | input file format |
| **-initial-mode** | no | `A` or `S` | `A` | initial mode of the lexer. see [the specification](./spec.md) for more details. |

## watson decode


### Usage

```
watson decode -t=TYPE [-initial-mode=MODE] [-stack-size=SIZE] [FILES...]
```

Converts Watson files `FILES` into another format that is specified by `TYPE` and outputs it to the standard output.

If `FILES` is not specified, it uses the standard input.

If multiple files are specified, they are executed sequencially by the same lexer and VM, that is, the mode of the lexer and the stack of the VM remains unchanged when the VM finished processing one file and continues to another. After processing the last file, a value at the top of the VM's stack is displayed.

### Flags

| flag | mandatory | type | default | description |
| ---- | --------- | ---- | ------- | ----------- |
| **-t**    | no        | `json`, `yaml`, `msgpack`, or `cbor` | `yaml` | input file format |
| **-initial-mode** | no | `A` or `S` | `A` | initial mode of the lexer. see [the specification](./spec.md) for more details. |
| **-stack-size** | no | integer | 1024 | stack size of the VM. see [the specification](./spec.md) for more details. |
