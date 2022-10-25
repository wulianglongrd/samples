package samples

import (
	"crypto/md5"
	"encoding/binary"
	"github.com/cespare/xxhash"
	"strings"
)

type ConfigKey struct {
	Kind      uint8
	Name      string
	Namespace string
}

func Md5KeyFind(samples []ConfigKey) {
	m := map[uint64]struct{}{}
	for _, v := range samples {
		m[md5Key(v)] = struct{}{}
	}
	for _, v := range samples {
		_ = m[md5Key(v)]
	}
}

func HashKeyFind(samples []ConfigKey) {
	m := map[uint64]struct{}{}
	for _, v := range samples {
		m[hashKey(v)] = struct{}{}
	}
	for _, v := range samples {
		_ = m[hashKey(v)]
	}
}

func StringKeyFind(samples []ConfigKey) {
	m := map[string]struct{}{}
	for _, v := range samples {
		m[stringKey(v)] = struct{}{}
	}

	for _, v := range samples {
		_ = m[stringKey(v)]
	}
}

func md5Key(key ConfigKey) uint64 {
	hash := md5.New()
	hash.Write([]byte{key.Kind})
	hash.Write([]byte(key.Name))
	hash.Write([]byte(key.Namespace))
	sum := hash.Sum(nil)
	return binary.BigEndian.Uint64(sum)
}

func hashKey(key ConfigKey) uint64 {
	var builder strings.Builder
	k := string(key.Kind)
	l := len(k) + len(key.Name) + len(key.Namespace)

	builder.Grow(l)
	builder.WriteString(k)
	builder.WriteString(key.Name)
	builder.WriteString(key.Namespace)

	return xxhash.Sum64([]byte(builder.String()))
}

func stringKey(key ConfigKey) string {
	var builder strings.Builder
	k := string(key.Kind)
	l := len(k) + len(key.Name) + len(key.Namespace)

	builder.Grow(l)
	builder.WriteString(k)
	builder.WriteString(key.Name)
	builder.WriteString(key.Namespace)
	return builder.String()
}
