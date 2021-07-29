//读取csv文件的kv

package csvKv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const LanguageFA = "fa"

func ReadRange(path string, f func(row int, v1, v2 string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	r := csv.NewReader(file)
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) != 2 {
			fmt.Println("recode len not eq 2 and ignore:", record)
			continue
		}
		if record[0] == "" && record[1] == "" {
			continue
		}
		if record[0] == "" || record[1] == "" {
			fmt.Println("异常数据row:", i)
		}
		r1, r2 := record[0], record[1]
		//r1 := strings.Replace(record[0], "\u00a0", "", -1)
		//r2 := strings.Replace(record[1], "\u00a0", "", -1)
		if r1 == "Go to Reseller System\n" {
			r1 = "Go to Reseller System"
		}
		if r2 == "Go to Reseller System\n" {
			r2 = "Go to Reseller System"
		}
		f(i, r1, r2)
		i++
	}
	return nil
}
