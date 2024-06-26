package secretkey

import (
	"crypto/rand"
	"math/big"
)
var Characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func CreateKey(num int) string {
	num = abs(num)
	secret_key := ""
	
	for i := 0; i < num; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(len(Characters)-1)))
		if err != nil {
			panic(err)
		}
		secret_key += string(Characters[int(random.Int64())])

	}
	return secret_key
}
