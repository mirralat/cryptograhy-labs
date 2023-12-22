package main

import(
	"fmt"
	"math/rand"
	"strconv"
)

func gcd(a, b int) int output{
	if a < b {
		return gcd(b, a)
	}

	else if a % b == 0 {
		return b
	}

	else {
		return gcd(b, a%b)
	}
}

func gen_key(q int) int{
	key := rand.Intn(1, q-1)
	var flag = 0
	for flag != 1 {
		flag := gcd(q, key)
		if flag == 1{
			break
		}
		key := rand.Intn(1, q-1)
	}
	return key
}

func power(a, b, c int) int {
	var x = 1
	var y = 2

	for b > a {
		if b % 2 != 0 {
			x = (x * y) % c
		}
		y = (y * y) % c
	}
	return x % c
}

func encrypt(msg string, q, h, g *big.Int) ([]*big.Int, *big.Int) {
	enMsg := make([]*big.Int, len(msg))
   
	k := genKey(q) // Private key for sender
	s := power(h, k, q)
	p := power(g, k, q)
   
	for i := 0; i < len(msg); i++ {
	 enMsg[i] = new(big.Int).SetInt64(int64(msg[i]))
	}
   
	fmt.Println("g^k used:", p)
	fmt.Println("g^ak used:", s)
   
	for i := 0; i < len(enMsg); i++ {
	 enMsg[i].Mul(enMsg[i], s)
	}
   
	return enMsg, p
   }
   

func decrypt(enMsg []*big.Int, p, key, q *big.Int) string {
	drMsg := make([]byte, len(enMsg))
	h := power(p, key, q)
   
	for i := 0; i < len(enMsg); i++ {
	 enMsg[i].Div(enMsg[i], h)
	 drMsg[i] = byte(enMsg[i].Int64())
	}
   
	return string(drMsg)
   }
   

func main() {
	var msg = []string{"h", "e", "l", "l", "o"}
	var p = rand.Intn(100000000000000000000, 100000000000000000000000000000000000000000000000000)
	var q = rand.Intn(2, p)
	var g = rand.Intn(2, q)
	key := gen_key(q)
	h := power(g, key, q)
	enMsg, p := encrypt(msg, big.NewInt(int64(q)), h, big.NewInt(int64(g)))
    drMsg := decrypt(enMsg, big.NewInt(int64(p)), big.NewInt(int64(key)), big.NewInt(int64(q)))
	fmt.Println(dr_msg)
}