## Net Debug

## Command 



- ifconfig [eth_x]

`config ip`
- ifconfig [eth_x] [ip_addr] netmask [255.255.255.0] 

`modify mac`
- ifconfig [eth_x]  hw ether [xx:xx:xx:xx:xx:xx]

- ifconfig [eth_x] mtu [Value]

`config gw`
- route add default gw [ip_addr] dev [eth_x]

`config dns`
- cat "nameserver 8.8.8.8" >> /etc/resolve.conf

- route -nv 

> Destination[目标网段或主机] \
Gateway[网关地址 '*'不需要路由] \
Genmask[网络掩码] \
Flags[标记] \
Metric[路由距离，中转次数] \
Ref[路由项引用次数]  \
Use[被软件查找次数]  \
Iface[表象输出接口] 


        route -v         verbose
              -n         num replace host_name to show addr
              
              -net       dst net
              -host      dst host
              
              del        del a route
              add        add a route
- route print (windows)
- route add default gw  [xxx.xxx.xxx.xxx]
- route del default gw  [xxx.xxx.xxx.xxx]
- route add -host  [xxx.xxx.xxx.xxx] gw  [xxx.xxx.xxx.xxx] [eth_x]
- route del -host  [xxx.xxx.xxx.xxx] gw  [xxx.xxx.xxx.xxx] [eth_x]
- route add -net  [xxx.xxx.xxx.xxx/xx] gw  [xxx.xxx.xxx.xxx] [eth_x]
- route del -net  [xxx.xxx.xxx.xxx/xx] gw  [xxx.xxx.xxx.xxx] [eth_x]
- route add -host  [xxx.xxx.xxx.xxx] reject
- route del -host  [xxx.xxx.xxx.xxx] reject
- route add -net  [xxx.xxx.xxx.xxx] netmask  [xxx.xxx.xxx.xxx] reject
- route del -net  [xxx.xxx.xxx.xxx] netmask  [xxx.xxx.xxx.xxx] reject


## 802.3 eth frames

    pre-symbol(leader)       7 octet(8 bits)
    frame-start              1 octet
    dst-mac                  6 octet
    src-mac                  6 octet
    eth-frame-type           2 octet
    load                     46-1500 octet
    redundancy-check         4 octet
    frame-interval           12 octed
    

- 100M frame speed min-frame-length 148809.5 f/s
- Eth effectiveness 

        Efficiency = payload size / frame size
        Eth pack 1500/1538 = 97.53%
        
- Eth bit efficiency 

        net bit rate = efficiency * writebitrate
        
        non 802.1Q  100 BASE-TX max bit rate 97.53 Mbit/s
        
- Width Uint

        1 Gbit/s = 1000 Mbit/s
        1 Mbit/s = 1000 Kbit/s
        1 Kbit/s = 1000 bit/s
        
- Storage Uint

        1 GB = 1024 MB
        1 MB = 1024 KB
        1 KB = 1025 B
        1 B = 8 bit
        
        
## Linux ethtool

- ethtool [eth_x] `check net status`
- ethtool -i [eth_x] `check driver version`
- ethtool -S [eth_x] `check send/recv statistics`
- ethtool -s [eth_x] speed 10 duplex full autoneg off `force mode`
- ethtool -A [eth_x] `flow control`


## iptables (filter,addr_swap)

> package filter , application proxy , status check

- raw , mangle , nat ---> filter

        5 tables ==> filter nat mangle raw security
        5 chains ==> preroutine input forwarding output postroutine
        
        
- filter

        iptables -A  INPUT  -p icmp -j DROP
        

- nat

    1. SNAT 
    2. DNAT
    3. MASQUERADE
    
            SANT 
            
            iptables -t nat -A POSTROUTINE -d  [xxx.xxx.xxx.xxx] -p [protocal]
            --dport [port] -j SNAT --to  [xxx.xxx.xxx.xxx:xxxx]
            
            eg :iptables -t nat -A POSTROUTINE -d 192.168.1.1 -p udp
            --dport 6666 -j SNAT --to 192.168.2.1:6666
            
            DNAT
            
            iptables -t nat -A PREROUTINE -d  [xxx.xxx.xxx.xxx] -p [protocal]
            --dport [port] -j SNAT --to  [xxx.xxx.xxx.xxx:xxxx]
            
            eg :iptables -t nat -A PREROUTINE -d 192.168.1.1 -p udp
            --dport 6666 -j SNAT --to 192.168.2.1:6666
            
            MASQUERADE
            
            iptables -t nat -A POSTROUTINE -o [eth_x] -j MASQUERADE
            
            
- iptables -t nat -L -vn

- ip (iproute2)

        ip -s -s link ls eth0
        ip link set dev [eth_x] up/down
        ip addr add  [xxx.xxx.xxx.xxx] dev [eth_x]
        ip addr show [eth_x]
        ip link set dev [eth_x] mtu [value]
        ip link set dev [eth_x] address [xx:xx:xx:xx:xx:xx] 
        ip addr flush [eth_x]
        ip route add default via  [xxx.xxx.xxx.xxx] dev [eth_x]
        ip route del default via  [xxx.xxx.xxx.xxx] dev [eth_x]
        ip route add [xxx.xxx.xxx.xxx/xx] via  [xxx.xxx.xxx.xxx] dev [eth_x]
        ip route add [host] via  [xxx.xxx.xxx.xxx] dev [eth_x]
        ip route del [host] via  [xxx.xxx.xxx.xxx] dev [eth_x]
        
        
- natstat

        -a show all listened port
        -s show all protocal statistics
        -r show kernal route msg (-rn)
        -i show net iface list
        -tan 
        
        
        
- traceroute/tracepath
- iwconfig/iwpriv/iwlist/wpa_supplicant
- iperf/netperf
- tcpdump
- tc

## common problem

- ping 

        1.check arp cache is had dist mac addr
        2.if not, send arp req package, recv ack package , 
        and store into arp cache
        3.send icmp req package
        4.recv icmp res package ---> net good 
        
        icmp statistics --> natstat -s 
        
        ------
        failed ping
        
        ping 
        1.not recv icmp reply,
        check network topology, is had loop ?
        2.arp table expired
        3.ip conflict
        
- loss package

        ethtool -S [eth_x]
        
        natstat -s [eth_x]
        
        check link-layer, check which layer loss
        check is Reverse mechanism or not ?
        
        
        ifconfig loss ?
        ifcongfig [eth_x]:droped: xxx
        maybe:
        1.tx dropped: xxx dm816x dma buffer not efficient
        2.system not support that msg
        
        net payload overload ?
        ethtool -k [eth_x]
        
        tcp-segmentation-offload:on
        udp-segmentation-offload:on
        
        flow overload , network card can't deal that, maybe cpu do
           
- ip conflict

        arping
        -b keep boardcast
        -f ack first res && quit
        -q not show warning
        -U active ARP ,update neighbor
        -c data package cnt ,sended
        -w over time 
        -I network card
        -s cofnig src ip addr
        -h show help
        -V version
        




