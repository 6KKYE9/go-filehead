package filehead

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "f.txt")
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestHeadLines(t *testing.T) {
	p := writeTemp(t, "a\nb\nc\nd\n")
	lines, err := HeadLines(p, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 || lines[0] != "a" || lines[1] != "b" {
		t.Errorf("HeadLines=%v", lines)
	}
}

func TestHeadBytes(t *testing.T) {
	p := writeTemp(t, "hello world")
	b, err := HeadBytes(p, 5)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hello" {
		t.Errorf("HeadBytes=%q", string(b))
	}
	// 超出文件长度
	b, _ = HeadBytes(p, 100)
	if string(b) != "hello world" {
		t.Errorf("HeadBytes over=%q", string(b))
	}
}
