## About Varint


### Little Endian 

- little: low address put low address
- MSB:Most  Significant Bit 
- LSB:Least Significant Bit
- example: int32(123456)


        binary: 1 11100010 01000000

        +-----------+  +-----------+  +-----------+
        |         1 |  |  11100010 |  |  01000000 | 
        +-----------+  +-----------+  +-----------+
        
        little endian:

        +-----------+  +-----------+  +-----------+
        | 1 1000000 |  | 1 1000100 |  | 0 0000111 | 
        +-----------+  +-----------+  +-----------+

        varint:
        
        192 196 7
        



#### Golang

- int32 4 byte --> varint 1~5  byte
- int64 8 byte --> varint 1~10 byte

##### byte

```go
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.

type byte = uint8
```

##### rune

```go
// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.

type rune = int32
```
