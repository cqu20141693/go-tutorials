package __18

import (
	"bytes"
	"encoding/hex"
	"fmt"
	cmpp "github.com/bigwhite/gocmpp"
	"testing"
	"unicode/utf8"
)

// 根据长宽获取面积
func GetArea(weight int, height int) int {
	return weight * height
}

/*
 Test 开头： 普通单元测试
 testing.T 类型入参
*/
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func TestUnitTest(t *testing.T) {

	// 定义输入
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"Hello, world"}, "dlrow ,olleH"},
		{"test2", args{" "}, " "},
		{"test3", args{"!12345"}, "54321!"},
	}

	// 使用输出
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ret := Reverse(test.args.s); ret != test.want {
				t.Errorf("Reverse() = %v, want %v", ret, test.want)
			}
		})
	}
}

func FuzzReverse(f *testing.F) {
	// 创建输入数据
	testcases := []string{"Hello, world", " ", "!12345"}
	// 初始化Fuzz test context
	for _, tc := range testcases {
		f.Add(tc)
	}
	// Fuzz test
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
func ReverseRunes(s string) string {
	fmt.Printf("input: %q\n", s)
	r := []rune(s)
	fmt.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func FuzzReverseRunes(f *testing.F) {
	// 创建输入数据
	testcases := []string{"Hello, world", "\x91", "!12345"}
	// 初始化Fuzz test context
	for _, tc := range testcases {
		f.Add(tc)
	}
	// Fuzz test
	f.Fuzz(func(t *testing.T, orig string) {
		rev := ReverseRunes(orig)
		doubleRev := ReverseRunes(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func BenchmarkGetArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetArea(40, 50)
	}
}

func FuzzSubmit(f *testing.F) {
	data := []byte{
		0x00, 0x00, 0x00, 0xbd, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x31, 0x33, 0x35, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x39, 0x30, 0x30, 0x30, 0x30,
		0x31, 0x30, 0x32, 0x31, 0x30, 0x00, 0x00, 0x00, 0x00, 0x31, 0x35, 0x31, 0x31, 0x30, 0x35, 0x31,
		0x33, 0x31, 0x35, 0x35, 0x35, 0x31, 0x30, 0x31, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x39, 0x30, 0x30, 0x30, 0x30,
		0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x31, 0x33, 0x35, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x6d, 0x4b, 0x8b, 0xd5, 0x00, 0x67, 0x00, 0x6f, 0x00,
		0x63, 0x00, 0x6d, 0x00, 0x70, 0x00, 0x70, 0x00, 0x20, 0x00, 0x73, 0x00, 0x75, 0x00, 0x62, 0x00,
		0x6d, 0x00, 0x69, 0x00, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	f.Add(data[8:])
	f.Fuzz(func(t *testing.T, data []byte) {
		p := &cmpp.Cmpp2SubmitReqPkt{}
		_ = p.Unpack(data)
	})
}

func FuzzHex(f *testing.F) {
	for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, in []byte) {
		enc := hex.EncodeToString(in)
		out, err := hex.DecodeString(enc)
		if err != nil {
			t.Fatalf("%v: decode: %v", in, err)
		}
		if !bytes.Equal(in, out) {
			t.Fatalf("%v: not equal after round trip: %v", in, out)
		}
	})
}
