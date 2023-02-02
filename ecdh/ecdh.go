package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("ECC Paramsters --")
	fmt.Printf("Name: %s\n", elliptic.P256().Params().Name)
	fmt.Printf("N: %x\n", elliptic.P256().Params().N)
	fmt.Printf("P: %x\n", elliptic.P256().Params().P)
	fmt.Printf("Gx:%x\n", elliptic.P224().Params().Gx)
	fmt.Printf("Gy:%x\n", elliptic.P256().Params().Gy)
	fmt.Printf("Bitsize: %x\n", elliptic.P256().Params().BitSize)

	privateA, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	privateB, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	fmt.Printf("\nPrivate Key (Alice) %x\n", privateA.D)
	fmt.Printf("\nPrivate Key (Bob) %x\n", privateB.D)

	publicA := privateA.PublicKey
	publicB := privateB.PublicKey
	fmt.Printf("\nPublic key (Alice) (%x, %x)\n", publicA.X, publicA.Y)
	fmt.Printf("\nPublic key (Bob) (%x, %x)\n", publicB.X, publicB.Y)

	a, _ := publicA.Curve.ScalarMult(publicA.X, publicA.Y, privateB.D.Bytes())
	b, _ := publicB.Curve.ScalarMult(publicB.X, publicB.Y, privateA.X.Bytes())

	shared1 := sha256.Sum256(a.Bytes())
	shared2 := sha256.Sum256(b.Bytes())

	fmt.Printf("\nShare key (Alice) %x\n", shared1)
	fmt.Printf("\nShare key (Bob) %x\n", shared2)
}
