package filehead

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// HeadLines 读文件前 n 行返回。文件不足 n 行就返回全部。
func HeadLines(path string, n int) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []string
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for i := 0; i < n && sc.Scan(); i++ {
		out = append(out, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// HeadBytes 读文件前 n 个字节返回。不足就返回全部。
func HeadBytes(path string, n int) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, n)
	m, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf[:m], nil
}

// Preview 把行预览拼成可读文本，行号从 1 开始。
func Preview(lines []string) string {
	var out string
	for i, l := range lines {
		out += fmt.Sprintf("%4d: %s\n", i+1, l)
	}
	return out
}
