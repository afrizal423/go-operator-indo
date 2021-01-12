<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>


# Operator Indonesia Number Parse With Go
<a href="https://afrizalmy.com"><img src="https://img.shields.io/badge/license-MIT-lightgrey" alt="me"></a>

Package Go yang berfungsi untuk mengecek operator atau provider yang mengacu pada prefix nomor telepon.

## Install
```
go get github.com/afrizal423/go-operator-indo/OperatorIndo
```

## Example
```
package main

import (
	"fmt"

	"github.com/afrizal423/go-operator-indo/OperatorIndo"
)

func main() {
	nohp := "083811111111"
	fmt.Println(OperatorIndo.Check(nohp))
}

```

## Output
```
Axis
```
Output bertipe data <b>String</b>, jadi bisa dilanjutkan langsung simpan ke database.