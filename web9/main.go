package main

import (
	"fmt"
	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var receiveData string

type SendComponent struct {
}

func (self *SendComponent) Operator(data string) {
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	byteData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(byteData))
}

type EncryptComponet struct {
	key string
	com Component
}

func (self *EncryptComponet) Operator(data string) {
	byteData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(byteData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(decData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(unzipData))
}

type ReadComponet struct {
}

func (self *ReadComponet) Operator(data string) {
	receiveData = data
}

func main() {
	sender := &EncryptComponet{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponet{},
		}}

	receiver.Operator(sentData)
	fmt.Println(receiveData)
}
