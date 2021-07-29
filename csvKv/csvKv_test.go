package csvKv

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestReadRange(t *testing.T) {
	m := map[string]string{}
	var strList []string
	var err error
	//err := ReadRange("/Users/zz/Desktop/波斯语翻译/波斯语翻译5/3-表格 1.csv", func(row int, v1, v2 string) {
	//	if row == 0 {
	//		return
	//	}
	//	v1, v2 = v2, v1
	//	//if strings.Contains(v1, "Go to Reseller System") {
	//	//	v1 = "Go to Reseller System"
	//	//}
	//	v, ok := m[v1]
	//	if ok {
	//		if v != v2 {
	//			fmt.Println("row:", row, ",异常数据v1:", v1)
	//		}
	//		return
	//	}
	//	strList = append(strList, []string{
	//		fmt.Sprintf("\t\"%s\": {\n", v1),
	//		fmt.Sprintf("\t\tLanguageFA: \"%s\",\n", v2),
	//		"\t},\n",
	//	}...)
	//	m[v1] = v2
	//})
	//if err != nil {
	//	panic(err)
	//}
	//err = ReadRange("/Users/zz/Desktop/波斯语翻译/波斯语翻译5/reseller 推广web-表格 1.csv", func(row int, v1, v2 string) {
	//	if row == 0 {
	//		return
	//	}
	//	v, ok := m[v1]
	//	if ok {
	//		if v != v2 {
	//			fmt.Println("row:", row, ",异常数据v2:", v1, "old:", v, "new:", v2)
	//		}
	//		return
	//	}
	//	strList = append(strList, []string{
	//		fmt.Sprintf("\t\"%s\": {\n", v1),
	//		fmt.Sprintf("\t\tLanguageFA: \"%s\",\n", v2),
	//		"\t},\n",
	//	}...)
	//	m[v1] = v2
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	err = ReadRange("/Users/zz/Desktop/波斯语翻译/翻译7/翻译7/Sheet2-表格 1.csv", func(row int, v1, v2 string) {
		if row == 0 {
			return
		}
		//v1, v2 = v2, v1
		v1 = strings.Replace(v1, "\n", "", 1)
		v, ok := m[v1]
		if ok {
			if v != v2 {
				fmt.Println("row:", row, ",异常数据v3:", v1, "old:", v, "new:", v2)
			}
			return
		}
		strList = append(strList, []string{
			fmt.Sprintf("\t\"%s\": {\n", v1),
			fmt.Sprintf("\t\tLanguageFA: \"%s\",\n", v2),
			"\t},\n",
		}...)
		m[v1] = v2
	})
	if err != nil {
		panic(err)
	}

	strList = append(strList, "}\n")
	strList = append([]string{"package csvKv\n", "var languageMap = map[string]map[string]string{\n"}, strList...)
	file, err := os.Create("temp.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, str := range strList {
		file.WriteString(str)
	}
}
