package utils

import (
	"7youo-wms/global"
	"bytes"
	"crypto"
	"encoding/hex"
	"go.uber.org/zap"
	"strconv"
	"unicode"
)

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}

	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			global.G_LOG.Error("string append error: 内存不够了", zap.Any("err", err))
		}
	}()
	b.WriteString(s)
	return b
}
//驼峰字符串转化成下划线写法字符
func UnderScoreName(name string) string {
	buf := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buf.Append("_")
			}
			buf.Append(unicode.ToLower(r))
		} else {
			buf.Append(r)
		}
	}
	return buf.String()
}

func MD5(str []byte) string {
	h := crypto.MD5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
