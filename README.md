### Go를 학습하며 배운 점

#### single-file program & multi-file program

- main.go만 있었을때는 go.mod 파일 없이도 실행
- accounts 폴더를 추가하면서 multi-file 프로그램. Go에서 이렇게 파일이 여러 개인 프로젝트는 `package`로 정리되고 그것들은 또 `module`로 모아진다. 모듈의 디펜던시나 언어 버전 등 명세서 역할을 하는 것이 go.mod 파일.
- 내가 accounts 폴더를 추가 후, `go run main.go`를 실행했을때 에러가 났던 이유는 main.go 파일 상단 import를 실행할때 go.mod 파일을 찾게 되는데 이게 없었기 때문.
- `go mod init ~~~` 명령어로 go.mod 파일을 생성하면 이제 accounts 부분이 이 모듈의 일부분이라는 것을 Go가 알게됨.

#### asterisk

- \* \+ type: "이거는 포인터구요 그 포인터가 가리키고 있는 값의 타입은 type이요"
- \* \+ variable: "이 포인터가 무엇을 가리키고 있는지 보여줄게"
- 이걸 알고 나니 아래 코드가 이해가지 않았음

```Go
func (a *Accounts) Deposit(amount int) {
    a.balance += amount
}
```

- `a *Accounts`의 a는 포인터이며 struct 필드인 balance에 접근하려면 \*a로 포인터가 가리키는 대상을 먼저 찾아야할 것 같은데 왜 그렇게 하지 않는지 궁금했음
- 포인터로 struct에 접근하는 경우 Go가 알아서 포인터가 가리키는 객체에 접근하는 걸로 해석해서 처리함
- 그래서 만약 \*a.balance 로 작성한다면 Go는 `a.balance`를 포인터로 보고 이것이 가리키는 것이 무엇인지 찾으려고 할 것임
