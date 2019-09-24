package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync/atomic"
	_ "time"
)

var arr = createBallArray()
var index int32 = -1
var numClient = 0
var size int32

func main() {
	listener, err := configTCPSocket()
	handleError(err)
	listenForConnection(listener)
}

func listenForConnection(listener *net.TCPListener) {
	for {
		conn, err := listener.Accept()
		fmt.Println("CONNECTED")
		if err != nil {
			continue
		}
		go handleClient(&conn)
	}
}

func configTCPSocket() (*net.TCPListener, error) {
	port := "127.0.0.1:3000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		return nil, err
	}
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func handleClient(connection *net.Conn) {
	var conn = *connection
	fmt.Println("HANDLE CLIENT")
	request := make([]byte, 1024)
	var tmp int32
	for {
		if tmp == -1 {
			break
		}
		readLen, err := conn.Read(request)
		fmt.Println("READ REQUEST")
		if err != nil {
			fmt.Println(err)
			break
		}
		request = request[:readLen]
		fmt.Println(string(request))
		if readLen > 0 && string(request) == "get" {
			tmp = -1
			if index < size-1 {
				tmp = arr[atomic.AddInt32(&index, 1)]
				fmt.Printf("Index and Size and value are: %d %d %d\n", index, size, arr[index])
			}
			var res string
			res = strconv.Itoa(int(tmp))
			fmt.Println("Response: " + res)
			conn.Write([]byte(res))
		}
	}
	err := receiveFileFromClient(connection)
	handleError(err)
	fmt.Println("Process done")
}

func receiveFileFromClient(connection *net.Conn) error {
	var conn = *connection
	var err error
	var fo *os.File
	var readLen int
	var prev = -1
	// var check = false
	buff := make([]byte, 100)
	readLen, err = conn.Read(buff)
	if err != nil {
		return err
	}
	buff = buff[:readLen]
	fmt.Println("RECEIVING FILE: " + string(buff))
	fo, err = os.OpenFile(string(buff)+".txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err = fo.Close()
	}()
	for {
		// if check {
		// 	break
		// }
		memset(buff, '0')
		readLen, err = conn.Read(buff)
		fmt.Printf("num read and prev are %d %d\n", readLen, prev)
		if err != nil {
			break
		}
		buff = buff[:readLen]
		if readLen == 0 || readLen < prev {
			fmt.Println("Received file from client")
			conn.Write([]byte("done"))
			fmt.Println("Process done")
			// time.AfterFunc(2*time.Second, func() {
			// 	check = true
			// })
			break
		} else {
			fo.WriteString(string(buff))
		}
		prev = readLen
	}
	return err
}

func memset(b []byte, v byte) {
	for i := range b {
		b[i] = v
	}
}

func createBallArray() []int32 {
	var i int32
	size = rand.Int31n(1001)
	//size = 10
	a := make([]int32, size)
	for i = 0; i < size; i++ {
		a[i] = rand.Int31n(10001)
	}
	return a
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
