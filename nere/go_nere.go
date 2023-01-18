// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go_nere

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"
)

func ObtenirNAD(in string) string {
	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length.
	//salt := "n=xDS5.dty]aPE2@YB>c-4}z*g!N?X/9"
	salt := []byte("n=xDS5.dty]aPE2@YB>c-4}z*g!N?X/9")

	dk, err := scrypt.Key([]byte(in), salt, 524288, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk)
}
