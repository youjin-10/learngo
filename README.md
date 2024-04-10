### Go를 학습하며 배운 점

#### signle-file program & multi-file program

- main.go만 있었을때는 go.mod 파일 없이도 실행
- accounts 폴더를 추가하면서 multi-file 프로그램. Go에서 이렇게 파일이 여러 개인 프로젝트는 `package`로 정리되고 그것들은 또 `module`로 모아진다. 모듈의 디펜던시나 언어 버전 등 명세서 역할을 하는 것이 go.mod 파일.
- 내가 accounts 폴더를 추가 후, `go run main.go`를 실행했을때 에러가 났던 이유는 main.go 파일 상단 import를 실행할때 go.mod 파일을 찾게 되는데 이게 없었기 때문.
- `go mod init ~~~` 명령어로 go.mod 파일을 생성하면 이제 accounts 부분이 이 모듈의 일부분이라는 것을 Go가 알게됨.
