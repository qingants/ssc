// package main

// import "fmt"

// const (
// 	generator         = 3
// 	primeNumber int64 = 6700417 // prime number discovered by Leonhard Euler
// )

// // GenerateShareKey : generates a key using client private key ,
// // generator and primeNumber
// // this key can be made public
// // shareKey = (g^key)%primeNumber
// func GenerateShareKey(prvKey int64) int64 {
// 	return modularExponentiation(generator, prvKey, primeNumber)
// }

// // GenerateMutualKey : generates a mutual key that can be used by
// // only alice and bob
// // mutualKey = (shareKey^prvKey)%primeNumber
// func GenerateMutualKey(prvKey, shareKey int64) int64 {
// 	return modularExponentiation(shareKey, prvKey, primeNumber)
// }

// // r = (b^e)%mod
// func modularExponentiation(b, e, mod int64) int64 {

// 	// runs in O(log(n)) where n = e
// 	// uses exponentiation by squaring to speed up the process
// 	if mod == 1 {
// 		return 0
// 	}
// 	var r int64 = 1
// 	b = b % mod
// 	for e > 0 {
// 		if e&1 == 1 {
// 			r = (r * b) % mod
// 		}
// 		e = e >> 1
// 		b = (b * b) % mod
// 	}
// 	return r
// }

// func main() {
// 	// alice私钥【保密，不公开】
// 	var alicePrivateKey int64 = 1860
// 	// alice公钥【bob用户保存，窃取也无所谓】
// 	alicePublicKey := GenerateShareKey(alicePrivateKey)

// 	// bob私钥【保密，不公开】
// 	var bobPrivateKey int64 = 992514
// 	// bob公钥【alice用户保存，窃取也无所谓】
// 	bobPublicKey := GenerateShareKey(bobPrivateKey)

// 	// 加密的key
// 	A := GenerateMutualKey(alicePrivateKey, bobPublicKey) //alice的私钥和bob的公钥
// 	B := GenerateMutualKey(bobPrivateKey, alicePublicKey) //bob的私钥和alice的公钥

//		// A和B的值一样，这样就可以通过A值对明文进行对称加密
//		fmt.Println(A)
//		fmt.Println(B)
//	}
package dh64
