package OperatorIndo

import (
	"encoding/json"
	"strings"
)

var data = `{
    "Telkomsel":[
        "0811","0812", "0813", "0821", "0822", "0823", "0851", "0852", "0853"
    ],
    "XL":[
        "0817", "0818", "0819", "0859", "0877", "0878", "0879"
    ],
    "Axis":[
        "0831", "0832", "0833", "0837", "0838"
    ],
    "Indosat":[
        "0814", "0815", "0815", "0816", "0855", "0856", "0857", "0858"
    ],
    "Three":[
        "0894", "0895", "0896", "0897", "0898", "0899"
    ],
    "Smartfren":[
        "0881", "0882", "0883", "0884", "0885", "0886", "0887", "0888", "0888","0889"
	],
	"Ceria":[
		"0828"
	]
}`
var jsonData = []byte(data)

func Check(nohp string) string {
	re := strings.NewReplacer(`+62`, `0`)
	// fmt.Println(re.Replace(nohp))
	nohp = re.Replace(nohp)
	var dt map[string]interface{}
	json.Unmarshal(jsonData, &dt)
	// fmt.Println(nohp[0:4])
	for c, _ := range dt {
		cari := dt[c].([]interface{})
		for _, z := range cari {
			if nohp[0:4] == z {
				// fmt.Println(c)
				nohp = c
			}
		}

	}
	return nohp
}
