English | [CN中文](README_ZH.md)

### 一、Swift Demangler in Go

#### 1、What is it?

This is a Swift Fuction (:neutral_face: em, yes, just Swift Function right now) Demangler. You can run the `main.go` to raise a `repl` tool. Then input the string you want to demangle, press the `Enter` :sparkling_heart:

### 二、Detail

#### 1、Installation

- Go 1.5 and up

#### 2、Usage

- Run the `main.go`

```go
go run main.go
```

```powershell
swift-demangle-go> go run main.go
Hello! This is the swift demangler in go version!
>>_TFC4Pack5class4FuncFSSSb
result: Pack.class.Func(Swift.String) -> Swift.Bool
>>
```

- Make sure that your input is legal, or it would possibly fall into a dead circle.

### 三、Contribution

