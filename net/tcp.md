## TCP 

- TCP Head RFC793,1323

        src, dst port  16 byte
        seq number 32 byte  src --> dst octet (msg segment first data byte)
        ack number 32 byte only ack flag == 1 , make sense (src end next data byte)
        head length 4 byte head take 32byte -> 20 byte ~ 60 byte
        
        U、A、P、R、S、F 6 byte
        urgent pointer, ack, psh(as fast as post to application-layer)
        rst (reset connect), syn(start conn), fin(release conn)
        
        windows size 16 byte in order to flow control , unit as octet
        (expectation recv data length)
        
        tcp checksum 16 byte to tcp head && data ,dst end judge
        
        urgent point 16 byte as a offset presents urgent-data behind one byte's syn number
        
        option segment windows-expansion factor
    
- UDP Head RFC793,1323

         src, dst port  16 byte
         length segment 16 byte UDP head + UDP data
         checksum segment 16 byte (is a opthion segment ,but tcp required)
         
- IP 

        ip head 20 byte not include ip option segment
        
        version segment (ipv4 ipv6)
        
        msg head length 4 bit
        
        Type of Service 8 bit (TOS) inlclude 3 priority （COS，Class of Service）
        4 tos seg && 1 unuse seg means min-time-delay ,max-throuthput, max-ha, min-cost
        
        total length segment ip-data
        identification per send && accumulate 1
        ttl over a route , shift down 1
        protocal segment (tcp 6,udp 17)
        head checksum check complete 
        

- 面向连接的单播协议
- 建立连接（ip+port）
- octet flow,transport layer
- 四元组 （src:ip+port;dst:ip+port）
- tcp 3 phase : start,trans,stop(quit)
- ack delay recv (0~N-1) , sent N 
  positive:if one ack loss, before msg can confirm msg segment
- head length : 32 bit as uint , (4 bytes), 4~15 byte
- ACK —— 确认，使得确认号有效 \
  RST —— 重置连接（reset by peer）\
  SYN —— 初如化一个连接的序列号 \
  FIN —— 该报文段的发送方已经结束向对方发送数据 
- when a new connect built or a connect quit, just exchange head, no data

- three hands-shake --> client,server both had good send, good recv


- ISN

            三次握手的一个重要功能是客户端和服务端交换ISN(Initial Sequence Number), 
            以便让对方知道接下来接收数据的时候如何按序列号组装数据
            如果ISN是固定的，攻击者很容易猜出后续的确认号
            ISN = M + F(localhost, localport, remotehost, remoteport)
            M means Count , F means hash
            
- ISN loop

            /** The next routines deal with comparing 32 bit unsigned ints
             * and worry about wraparound (automatic with unsigned arithmetic).*/
            
            static inline int before(__u32 seq1, __u32 seq2){
                return (__s32)(seq1-seq2) < 0;}
            
            #define after(seq2, seq1) before(seq1, seq2)
           

- syn flood attack

            无效连接的监视释放
            延缓TCB分配方法
            Syn Cache
            Syn Cookie
            使用SYN Proxy防火墙
            连接队列 (netstat -s | grep LISTEN)
            

