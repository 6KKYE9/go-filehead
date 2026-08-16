# go-filehead

看文件开头用的。`head` 在 Windows 上不一定有，这个自带：按行预览（带行号）或者按字节预览都行，命令行直接给文件名。

读行用带缓冲的 scanner，单行长文件也不怕；读字节就是老老实实 `Read` 前 N 个，文件不够长就返回全部，不会报错。

## 装

```bash
go build -o filehead ./cmd/filehead
```

## 用

```bash
./filehead -n 5 README.md        # 前 5 行，带行号
./filehead -bytes -n 20 foo.bin  # 前 20 字节
./filehead a.txt b.txt           # 多个文件依次看
```

参数：
- `-n N`：行数或字节数，默认 10
- `-bytes`：按字节而不是按行（默认按行）

## 当库用

```go
import "filehead"

lines, _ := filehead.HeadLines("log.txt", 10)
b, _ := filehead.HeadBytes("x.bin", 64)
text := filehead.Preview(lines)   // 带行号的预览文本
```

## License

MIT
