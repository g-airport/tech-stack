package trick

import "fmt"

// -------------------------------------------
// ^uint(0), uint is not a fixed length type
// ^uint(0)
// 32 bit --> 0XFFFFFFFF 2^32
// 64 bit --> 0xFFFFFFFFFFFFFFFF 2^64

//for 32 bit sys: ^unit(0)  2^32−1 --> (2^32 - 1)>>63 == 0;32 << 0 --> 32
//for 64 bit sys: ^unit(0): 2^64−1 --> (2^64 - 1)>>63 == 1;32 << 1 --> 64

func Is64() {
	bit := 32 << (^uint(0) >> 63)
	fmt.Printf("is %d bit sys\n", bit)
}
