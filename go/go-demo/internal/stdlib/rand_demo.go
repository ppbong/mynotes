package stdlib

import (
	crand "crypto/rand"
	"encoding/hex"
	mrand "math/rand/v2"

	"github.com/ppbong/go-demo/internal/demo"
)

// Random 演示 math/rand/v2 与 crypto/rand 的区别和使用场景。
func Random() {
	demo.Section("随机数")

	// math/rand/v2 的全局源会自动播种，适合模拟、游戏等非安全场景。
	demo.KV("rand.IntN(100)", mrand.IntN(100))

	// 固定种子的 PCG 源可以生成可复现的序列，便于测试。
	r := mrand.New(mrand.NewPCG(1, 2))
	samples := make([]int, 3)
	for i := range samples {
		samples[i] = r.IntN(100)
	}
	demo.KV("固定种子序列", samples)

	// crypto/rand 用于令牌、密钥等安全敏感场景。
	buf := make([]byte, 8)
	if _, err := crand.Read(buf); err != nil {
		demo.KV("crypto/rand 失败", err)
		return
	}
	demo.KV("安全随机令牌", hex.EncodeToString(buf))
}
