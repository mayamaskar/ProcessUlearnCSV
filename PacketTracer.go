package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

type headerStruct struct { //Structure of header
	Identifier int16
	Sequence   int16
	Tag        [32]byte
	Timestamp  int32
}

type operationList struct{}

type operationGet struct {
	Key [32]byte
}

type operationPut struct {
	Key    [32]byte
	Expire int32
	Value  [64]byte
}

type operationDel struct {
	Key [32]byte
}

func main() {

	reader := bufio.NewReader(os.Stdin) //read from console
	var packets [][]byte
	var wg sync.WaitGroup
	//sync. WaitGroup provides a goroutine synchronization mechanism in Golang , and is used for waiting for a collection of goroutines to finish.
	var mutex = &sync.Mutex{}

	fmt.Print("Enter No. of packets: ")
	str, _ := reader.ReadString('\n') //read till new line
	fmt.Println("String From Console", str)
	n, err := strconv.ParseInt(strings.TrimSpace(str), 10, 64) //convert
	fmt.Println("nnnnnn", n)

	if err != nil {
		panic(err)
	}

	// get packets from console
	var i int64
	for i = 0; i < n; i++ {
		var packet []byte
		fmt.Println("Enter Packet", i+1)
		str, _ := reader.ReadString('\n')
		strSlice := strings.Split(str, " ")
		//fmt.Println("strSlice", strSlice)
		for _, c := range strSlice {
			i, _ := strconv.ParseInt(strings.TrimSpace(c), 10, 32)
			//fmt.Println("iiiii", byte(i))
			packet = append(packet, byte(i))
		}

		packets = append(packets, packet)
	}
	//fmt.Println("Pakets", packets)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// seperate go routines parsing each packet
	for idx, pkt := range packets {
		//fmt.Println("Index", idx)
		//fmt.Println("Packet", pkt)
		wg.Add(1)
		go parsePacket(idx, pkt, &wg, mutex)
	}

	// wait for packets to get parsed
	wg.Wait()
}

func parsePacket(idx int, packet []byte, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	fmt.Println("Packet", packet)
	buf := bytes.NewBuffer(packet)
	fmt.Println("Buffer", buf) //*C=C<=C=CJEF<<?JANJ+JANJEI=LAHOFMAHBK?IJHK?JEC
	h, err := parseHeader(buf)
	if err != nil {
		printError(err.Error(), idx, mu)
		return
	}

	id := h.Identifier

	p, err := parsePayload(buf, id)
	//fmt.Println("Buffer", p)
	if err != nil {
		printError(err.Error(), idx, mu)
		return
	}

	mu.Lock()
	fmt.Println()
	fmt.Println("Packet Number:", idx+1)
	fmt.Println("Identifier: ", h.Identifier)
	fmt.Println("Sequence: ", h.Sequence)
	fmt.Println("Tag: ", string(h.Tag[:]))
	//fmt.Print("taggggg", h.Tag[:])
	fmt.Println("Timestamp: ", h.Timestamp)
	fmt.Println()
	fmt.Println("Payload:")

	switch p.(type) {
	case operationList:
	case operationGet:
		fmt.Printf("Key: %s\n", p.(operationGet).Key)

	case operationPut:
		fmt.Printf("Key: %s\n", p.(operationPut).Key)
		fmt.Printf("Expire: %v\n", p.(operationPut).Expire)
		fmt.Printf("Value: %s\n", p.(operationPut).Value)

	case operationDel:
		fmt.Printf("Key: %s\n", p.(operationDel).Key)
	}
	mu.Unlock()
}

func printError(err string, idx int, mu *sync.Mutex) {
	mu.Lock()
	fmt.Println()
	fmt.Println("Error while parsing packet ", idx+1)
	fmt.Println(err)
	mu.Unlock()
}

func parsePayload(buf *bytes.Buffer, id int16) (interface{}, error) {
	var list operationList
	var get operationGet
	var put operationPut
	var del operationDel
	var err error

	// parse payload
	switch id {
	case 0:
		err = binary.Read(buf, binary.LittleEndian, &list)
		return list, err
	case 1:
		err = binary.Read(buf, binary.LittleEndian, &get)
		return get, err
	case 2:
		err = binary.Read(buf, binary.LittleEndian, &put)
		return put, err
	default:
		err = binary.Read(buf, binary.LittleEndian, &del)
		return del, err
	}
}

func parseHeader(buf *bytes.Buffer) (headerStruct, error) {
	var hd headerStruct
	err := binary.Read(buf, binary.LittleEndian, &hd)
	if err != nil {
		return headerStruct{}, err
	}
	//fmt.Println("hhhhhh", h)
	return hd, nil

}
