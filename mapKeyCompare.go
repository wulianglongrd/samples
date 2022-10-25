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

func IstioMd5KeyFind(samples []ConfigKey) {
	m := make(map[uint64]struct{}, len(samples))

	// this is simulate model.initSidecarScopes in push_context.go
	for _, v := range samples {
		m[istioMd5Key(v)] = struct{}{}
	}

	// this is simulate model.DependsOnConfig() in sidecar.go
	for _, v := range samples {
		_ = m[istioMd5Key(v)]
	}
}

func AnotherMd5KeyFind(samples []ConfigKey) {
	m := make(map[uint64]struct{}, len(samples))
	for _, v := range samples {
		m[anotherMd5Key(v)] = struct{}{}
	}
	for _, v := range samples {
		_ = m[anotherMd5Key(v)]
	}
}

func HashKeyFind(samples []ConfigKey) {
	m := make(map[uint64]struct{}, len(samples))
	for _, v := range samples {
		m[hashKey(v)] = struct{}{}
	}
	for _, v := range samples {
		_ = m[hashKey(v)]
	}
}

func StringKeyFind(samples []ConfigKey) {
	m := make(map[string]struct{}, len(samples))
	for _, v := range samples {
		m[stringKey(v)] = struct{}{}
	}

	for _, v := range samples {
		_ = m[stringKey(v)]
	}
}

func istioMd5Key(key ConfigKey) uint64 {
	hash := md5.New()
	hash.Write([]byte{key.Kind})
	hash.Write([]byte(key.Name))
	hash.Write([]byte(key.Namespace))
	sum := hash.Sum(nil)
	return binary.BigEndian.Uint64(sum)
}

func anotherMd5Key(key ConfigKey) uint64 {
	s := configToString(key)
	hash := md5.New()
	sum := hash.Sum([]byte(s))
	return binary.BigEndian.Uint64(sum)
}

func hashKey(key ConfigKey) uint64 {
	s := configToString(key)
	return xxhash.Sum64([]byte(s))
}

func stringKey(key ConfigKey) string {
	return configToString(key)
}

func configToString(key ConfigKey) string {
	var builder strings.Builder
	k := string(key.Kind)
	l := len(k) + len(key.Name) + len(key.Namespace)

	builder.Grow(l)
	builder.WriteString(k)
	builder.WriteString(key.Name)
	builder.WriteString(key.Namespace)
	return builder.String()
}
