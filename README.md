
"encoding/json"的天坑
 //json反序列化的天坑，坑死老娘了
 
 type DataType struct {
   name string
 }
 var data *DataType
 err = json.Unmarshal([]byte(dataBytes), data)   // 此时无论如何都会报错
 
 正确的用法是:
 var data DataType  //不要用指针
 err = json.Unmarshal([]byte(dataBytes), &data)    //反序列化成功,一点点的不同结果千差万别



tcp代理(golang实现）
最近需要在生产环境中调试一个http的问题，但是网管说不能安装抓包工具......

想到的一种替代方案是，增加个反向代理，然后把所有的request和response打印一下即可

1. nginx/haproxy来做反向代理，肯定可行，但是我配置的不太熟练

2. 自己写个简单的即可，正好golang写这类程序很快，所以就花了一会儿重新熟悉下go的语法（有阵子没看忘掉了......）然后写了个

主要的思路就是接收客户端的request,然后把里面的Host header改成真正的服务器的Host,然后转发，并把接收到的response回发个客户端即可

下面是全部的代码，用法很简单 ./tcpproxy 127.0.0.1 1080 www.baidu.com:80

复制代码
 1 // TcpProxy project main.go
 2 package main
 3 
 4 import (
 5     "fmt"
 6     "net"
 7     "os"
 8     "regexp"
 9     "strconv"
10 )
11 
12 func main() {
13     if len(os.Args) < 4 {
14         fmt.Println("missing message!")
15         return
16     }
17     ip := os.Args[1]
18     port, err := strconv.Atoi(os.Args[2])
19     if err != nil {
20         fmt.Println("error happened ,exit")
21         return
22     }
23     addr := os.Args[3]
24     host := "Host: " + addr
25 
26     Service(ip, port, addr, host)
27 }
28 
29 func Service(ip string, port int, dstaddr string, dsthost string) {
30     // listen and accept
31     listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(ip), port, ""})
32     if err != nil {
33         fmt.Println("listen error: ", err.Error())
34         return
35     }
36     fmt.Println("init done...")
37 
38     for {
39         client, err := listen.AcceptTCP()
40         if err != nil {
41             fmt.Println("accept error: ", err.Error())
42             continue
43         }
44         go Channal(client, dstaddr, dsthost)
45     }
46 }
47 
48 func Channal(client *net.TCPConn, addr string, rhost string) {
49 
50     tcpAddr, _ := net.ResolveTCPAddr("tcp4", addr)
51     conn, err := net.DialTCP("tcp", nil, tcpAddr)
52     if err != nil {
53         fmt.Println("connection error: ", err.Error())
54         client.Close()
55         return
56     }
57 
58     go ReadRequest(client, conn, rhost)
59     ReadResponse(conn, client)
60 }
61 
62 func ReadRequest(lconn *net.TCPConn, rconn *net.TCPConn, dsthost string) {
63     for {
64         buf := make([]byte, 10240)
65         n, err := lconn.Read(buf)
66         if err != nil {
67             break
68         }
69 
70         mesg := changeHost(string(buf[:n]), dsthost)
71         // print request
72         fmt.Println(mesg)
73         rconn.Write([]byte(mesg))
74     }
75     lconn.Close()
76 }
77 
78 func ReadResponse(lconn *net.TCPConn, rconn *net.TCPConn) {
79     for {
80         buf := make([]byte, 10240)
81         n, err := lconn.Read(buf)
82         if err != nil {
83             break
84         }
85 
86         fmt.Println(string(buf[:n]))
87         rconn.Write(buf[:n])
88     }
89     lconn.Close()
90 }
91 
92 func changeHost(request string, newhost string) string {
93     reg := regexp.MustCompile(`Host[^\r\n]+`)
94     return reg.ReplaceAllString(request, newhost)
95 }
复制代码
