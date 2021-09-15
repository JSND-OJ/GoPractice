package fromtxt

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([3]float64, error) { //읽어 올 파일의 이름을 인자로 받아서 숫자 배열과 에러 값을 반환
	var numbers [3]float64         //반환할 배열을 선언
	file, err := os.Open(filename) //받아 올 파일명으로 파일을 열고
	//여는 도중에 에러가 있으면 에러 보고 후 프로그램 종료
	if err != nil {
		return numbers, err
	}

	i := 0 // 값을 할당할 인덱스의 위치를 추적하기 위한 변수
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //텍스트 파일의 문자열을 float64로 변환
		// 타입을 변환하는 도중에 에러는 없는지
		if err != nil {
			return numbers, err
		}
		i++ // 다음 인덱스로 이동
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil //여기까지 왔으면 아무 에러가 없는 거니까, 숫자배열과 nil 반환
}
