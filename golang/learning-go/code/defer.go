package code

import (
	"io"
	"log"
	"os"
)

func FileOpenExample() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// defer는 호출하는 함수를 둘러싼 함수가 종료될 때까지 수행을 연기한다.
	// defer는 LIFO 로 마지막에 defer로 등록된 것이 가장 먼저 실행된다.
	defer f.Close()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
