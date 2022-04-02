package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/open-quantum-safe/liboqs-go/oqs"
	"github.com/xuri/excelize/v2"
)

func main() {
	// fmt.Println(oqs.EnabledSigs())
	pqcStrings := []string{
		"Dilithium2",
		"Dilithium3",
		"Dilithium5",
		"Falcon-512",
		"Falcon-1024",
		"picnic3_L1",
		"picnic3_L3",
		"picnic3_L5",
		"Rainbow-I-Circumzenithal",
		"Rainbow-III-Circumzenithal",
		"Rainbow-V-Circumzenithal",
		"SPHINCS+-SHA256-128s-robust",
		"SPHINCS+-SHA256-128s-simple",
		"SPHINCS+-SHA256-192s-robust",
		"SPHINCS+-SHA256-192s-simple",
		"SPHINCS+-SHA256-256s-robust",
		"SPHINCS+-SHA256-256s-simple",
	}
	for _, v := range pqcStrings {
		oqsTest(v)
	}

}

func oqsTest(pqcString string) {
	f, err := excelize.OpenFile("pqc.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.NewSheet(pqcString)
	f.SetCellValue(pqcString, "A1", "pk sizes")
	f.SetCellValue(pqcString, "B1", "sk sizes")
	f.SetCellValue(pqcString, "C1", "sig size")
	f.SetCellValue(pqcString, "D1", "sig time")
	f.SetCellValue(pqcString, "E1", "virify time")
	f.SetCellValue(pqcString, "F1", "geneKey time")

	var i int
	for i < 10000 {
		i++
		fmt.Println(pqcString, i)
		signer := oqs.Signature{}
		defer signer.Clean() // clean up even in case of panic
		if err := signer.Init(pqcString, nil); err != nil {
			log.Fatal(err)
		}

		msg := []byte("This is the message to sign")

		start := time.Now()
		pubKey, err := signer.GenerateKeyPair()
		elapsed := time.Since(start)
		f.SetCellValue(pqcString, "F"+strconv.Itoa(i+1), elapsed.Microseconds())
		sk := signer.ExportSecretKey()
		f.SetCellValue(pqcString, "A"+strconv.Itoa(i+1), len(pubKey))
		f.SetCellValue(pqcString, "B"+strconv.Itoa(i+1), len(sk))

		if err != nil {
			log.Fatal(err)
		}

		start = time.Now()
		signature, _ := signer.Sign(msg)
		elapsed = time.Since(start)
		f.SetCellValue(pqcString, "C"+strconv.Itoa(i+1), len(signature))
		f.SetCellValue(pqcString, "D"+strconv.Itoa(i+1), elapsed.Microseconds())

		verifier := oqs.Signature{}
		defer verifier.Clean() // clean up even in case of panic
		if err := verifier.Init(pqcString, nil); err != nil {
			log.Fatal(err)
		}

		start = time.Now()
		_, err = verifier.Verify(msg, signature, pubKey)
		elapsed = time.Since(start)
		f.SetCellValue(pqcString, "E"+strconv.Itoa(i+1), elapsed.Microseconds())

		if err != nil {
			log.Fatal(err)
		}

	}
	f.Save()

}
