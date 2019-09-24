package main

import (
	_ "bytes"
	_ "encoding/binary"
	"fmt"
	"github.com/bradfitz/slice"
	"io"
	"net"
	"os"
	"strconv"
)

var clientID string
var arrBalls []int

func main() {
	err := initVariables()
	handleError(err)
	conn, err := configTCPSocket()
	handleError(err)
	getBallsFromServer(&conn)
	err = saveBallsToFile()
	handleError(err)
	err = sendFileToServer(&conn)
	handleError(err)
}

func initVariables() error {
	clientID = ""
	if len(os.Args) >= 2 {
		clientID = os.Args[1]
	}
	arrBalls = make([]int, 0)
	var err error
	var fo *os.File
	fo, err = os.Create(clientID + ".txt")
	defer func() {
		err = fo.Close()
	}()
	return err
}

func configTCPSocket() (net.Conn, error) {
	port := "127.0.0.1:3000"
	tcpAdrr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		return nil, err
	}
	fmt.Println(tcpAdrr)
	conn, err := net.DialTCP("tcp4", nil, tcpAdrr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getBallsFromServer(connection *net.Conn) {
	var buff = make([]byte, 4)
	conn := *connection
	for {
		conn.Write([]byte("get"))
		fmt.Println("SENT GET BALL REQUEST")
		readLen, err := conn.Read(buff)
		if err != nil {
			continue
		}
		buff := buff[:readLen]
		data, err := strconv.Atoi(string(buff))
		if data == -1 {
			break
		} else {
			fmt.Println(data)
			arrBalls = append(arrBalls, data)
		}
	}
}

func saveBallsToFile() error {
	var err error
	var fo *os.File
	slice.Sort(arrBalls, func(i, j int) bool {
		return arrBalls[i] < arrBalls[j]
	})
	fo, err = os.OpenFile(clientID+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	var i int
	for i = 0; i < len(arrBalls); i++ {
		data := ""
		if i == len(arrBalls)-1 {
			data = strconv.Itoa(arrBalls[i])
		} else {
			data = strconv.Itoa(arrBalls[i]) + "\n"
		}
		_, err = fo.WriteString(data)
		if err != nil {
			return err
		}
	}
	defer func() {
		err = fo.Close()
	}()
	return nil
}

func sendFileToServer(connection *net.Conn) error {
	var fo *os.File
	var err error
	var conn = *connection
	buff := make([]byte, 1024)

	conn.Write([]byte(clientID))

	fo, err = os.OpenFile(clientID+".txt", os.O_RDONLY, 0644)
	defer func() {
		err = fo.Close()
	}()
	if err != nil {
		return err
	}
	for {
		memset(buff, '0')
		n, err := fo.Read(buff)
		buff = buff[:n]
		if err != nil && err != io.EOF {
			break
		}
		conn.Write(buff)
		if n == 0 {
			break
		}
	}
	for {
		tmp := make([]byte, 4)
		memset(tmp, '0')
		_, err := conn.Read(tmp)
		//buff = buff[:nRead]
		fmt.Printf("Result is %s\n", string(tmp))
		if err != nil || string(tmp) == "done" {
			break
		}
	}
	return err
}

func memset(b []byte, v byte) {
	for i := range b {
		b[i] = v
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
