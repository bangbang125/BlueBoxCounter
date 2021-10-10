package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	/*fmt.Println("Eikan", r.FormValue("counterOutput_Eikan"))
	fmt.Println("Hagyou", r.FormValue("counterOutput_Hagyou"))
	fmt.Println("Sigoku", r.FormValue("counterOutput_Sigoku"))
	fmt.Println("Hihi", r.FormValue("counterOutput_Hihi"))
	*/

	var a string
	var b string
	var c string
	var d string
	//var Eikan_num,Hagyou_num, Sigoku_num, Hihi_num  int
	//変数宣言

	a = r.FormValue("counterOutput_Eikan")
	b = r.FormValue("counterOutput_Hagyou")
	c = r.FormValue("counterOutput_Sigoku")
	d = r.FormValue("counterOutput_Hihi")

	/*Eikan_num, _=strconv.Atoi(a)
	Hagyou_num, _=strconv.Atoi(b)
	Sigoku_num, _=strconv.Atoi(c)
	Hihi_num, _=strconv.Atoi(d)*/

	//CSVファイル出力
	var dataSet []string = []string{a, b, c, d}
	fmt.Println(dataSet)
	file, err := os.Create("./save.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data_0 := []byte(dataSet[0])
	_, err = file.Write(data_0)

	data_1 := []byte(dataSet[1])
	_, err = file.Write(data_1)

	data_2 := []byte(dataSet[2])
	_, err = file.Write(data_2)

	data_3 := []byte(dataSet[3])
	_, err = file.Write(data_3)

	/*file, err := os.Create("/test.csv")
	if err != nil {
		log.Println(err)
	}

	writer := csv.NewWriter(file) // utf8
	writer.Write(a)
	writer.Flush()*/

	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func main() {
	// register hander
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
