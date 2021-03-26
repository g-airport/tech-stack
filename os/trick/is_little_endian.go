package trick

import "unsafe"

// Little or Big Endian

func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return b == 0x04
}

//地址：低 ---------------------> 高
//大端(16进制):   01 02 03 04
//小端(16进制):   04 03 02 01
//大端：原来低位的（如04）放在高地址
//小端：原来低位的（如04）放在低地址
//MSB:Most  Significant   Bit ------- 最高有效位
//LSB:Least Significant   Bit ------- 最低有效位

//int byteOrder() {
//    union {
//        short value;
//        char bytes[2];
//    } u;
//
//    u.value = 0x0102;
//
//    if (u.bytes[0] == 1 && u.bytes[1] == 2) {
//        return 1; // big endian
//    } else if (u.bytes[0] == 2 && u.bytes[1] == 1) {
//        return 2; // little endian
//    } else {
//        return -1; // unknown
//    }
//}

//网络字节序：TCP/IP各层协议将字节序定义为Big-Endian，因此TCP/IP协议中使用的字节序通常称之为网络字节序
//常用的X86结构是小端模式，而KEIL C51则为大端模式
//很多的ARM，DSP都为小端模式
//有些ARM处理器还可以随时在程序中(在ARM Cortex 系列使用REV、REV16、REVSH指令 [2]  )进行大小端的切换