### 一、Swift Demangler in Go

#### 1、What is it?

​	这是一个Swift 函数的Demangle工具（:neutral_face: em...是的，现在只能Demangle函数）。你可以运行`main.go`文件启动一个`repl`工具，输入你需要demangle的字符，按下回车，就能得到结果啦！:sparkling_heart:

### 二、Detail

#### 1、条件

- Go 1.5 and up

#### 2、用法

- run the `main.go`

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



###  三、贡献

