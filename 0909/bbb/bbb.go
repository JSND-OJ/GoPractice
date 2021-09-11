package main

import (
	"GoBlock/goweb/WEB9/cipher"
	"GoBlock/goweb/WEB9/lzw"
	"fmt"
)

type Send interface { //Send 인터페이스
	ProcData(string) //ProcessData 데이터 가공
}

type Recv interface { //Receive 인터페이스
	RstrData(string) //RestoreData 데이터 복원
}

var sentData string // 전역변수, 아래 send, data
var recvData string

//----기본기능 본래기능 정의 단계
type SendComp struct{} //실제 기본 기능

func (self *SendComp) ProcData(data string) { //실제 기본기능 구성 후 ProcData 호출
	//Send data
	sentData = data
}

//---데코레이터1 - 압축기능구현 단계
type ZipComp struct { //압축컴포넌트 데코레이터로 다른 컴포넌트 보관
	com Send //데코레이터는 com 컴포넌트를 가진다
}

//data 압축하고 Comp호출 형태로 구성
func (self *ZipComp) ProcData(data string) { //com멤버, ZipComp는 data 압축lzw부터 진행한다
	zipData, err := lzw.Write([]byte(data)) // Write는 String을 []byte로 바꾸고 압축결과로 []byte or err가 나온다
	if err != nil {
		panic(err) //err가 있으면 panic으로 종료
	}
	self.com.ProcData(string(zipData)) //err가 없으면 본인(Self)의 SendData에 암호화 압축한 data 호출
}

//여기까지는 압축 데코레이터를 만들었다

//---데코레이터2 -암호화기능구현 단계
type EncryptComp struct { // Encrypt는 com뿐만 아니라 반드시 key값을 가지는 데코레이터를 만든다
	key string
	com Send // 데코레이터는 com 컴포넌트를 가진다
}

func (self *EncryptComp) ProcData(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key) //Key값 멤버로 가진다
	if err != nil {
		panic(err)
	}
	self.com.ProcData(string(encryptData)) //암호화된 데이터 컴포넌트
}

//---여기까지는 본래기능 + 데코레이션1(압축기능) + 데코레이션2(암호기능)을 추가해서 sentData에 모두 응축 단계

//---- 암호풀기 단계
type DecryptComp struct { // DecryptComp는 암호화 푸는 복호화 단계의 컴포넌트, key값을 가진다
	key string
	com Recv
}

func (self *DecryptComp) RstrData(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key) //압축이 풀린 데이터
	if err != nil {
		panic(err)
	}
	self.com.RstrData(string(decryptData))
}

//--- 압축풀기 단계
type UnzipComp struct { // Unzip 압출풀기 데코레이터
	key string
	com Recv
}

func (self *UnzipComp) RstrData(data string) {
	unzipData, err := lzw.Read([]byte(data)) //lzw.Read는 압축해제 기능을 하는 함수
	if err != nil {
		panic(err)
	}
	self.com.RstrData(string(unzipData))
}

//----recvData로 최종 응축 시키는 단계
type ReadComp struct{}

func (self *ReadComp) RstrData(data string) {
	recvData = data
}

//----recvData로 최종응축 시키는 최종 단계

func main() {
	//SendPart
	sender := &EncryptComp{ // sender는 ZipComp와 SendComp를 모두 가진 상태로 하나의 응축된 레퍼런스 내부 패키지같은 역할
		key: "abcde", //key는 "abcde"
		com: &ZipComp{ // 암호화 컴포넌트는 ZipComp을 가지고 ZipComp는 SendComp를 가진다
			com: &SendComp{}, //ZipComp는 SendComp를 가진다
		},
	}
	sender.ProcData("Hello World") //sender보내는 데이터 "Hello World"
	fmt.Println(sentData)          //암호화+압축 데이터는 최종 sentData에 들어있다

	//RecvPart
	receiver := &UnzipComp{ //하나의 응축된 레퍼런스 내부 패키지같은 역할
		com: &DecryptComp{
			key: "abcde",
			com: &ReadComp{},
		},
	}
	receiver.RstrData(sentData)
	fmt.Println(recvData) //최종 Decrypt와 unzip Data를 출력한다.
}
