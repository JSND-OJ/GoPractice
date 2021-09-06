package main

import (
	"bufio"
	"fmt"
	"os"
)

func Readfile(filename string) (string, error) {
	file, err := os.Open(filename) // 1.파일열기
	if err != nil {
		return "", err // 2. 에러 나면 에러 반환
	}
	defer file.Close()             // 3. 함수 종료 직전 파일 닫기
	rd := bufio.NewReader(file)    //4. 파일 내용 읽기. bufio.NewReader() 함수로 bufio.Reader 객체를 만든다
	line, _ := rd.ReadString('\n') //bufio.Reader 객체는 구분자까지 문자열을 읽어오는 ReadString() 메서드를 가지고 있음. \n까지 읽음
	return line, nil
}

func Writefile(filename string, line string) error {
	file, err := os.Create(filename) //5. 파일 생성. os.Create는 첫 번째로 파일 핸들을, 두 번째로 에러를 반환
	if err != nil {                  //6. 에러 나면 에러 반환
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) //7. 파일에 문자열 쓰기. 파일 핸들에 문자열과 줄바꿈 문자 '\n'를 쓴다.
	return err
}

const filename string = "data.txt"

func main() {
	line, err := Readfile(filename) //8.파일 읽기 시도
	if err != nil {
		err = Writefile(filename, "This is Writefile") //9.파일 생성
		if err != nil {                                //10. 에러를 처리
			fmt.Println("Failed to create a file", err)
			return
		}
		line, err = Readfile(filename) //11. 다시 읽기 시도
		if err != nil {
			fmt.Println("Failed to read a file")
			return
		}
	}
	fmt.Println("file context:", line)
}
