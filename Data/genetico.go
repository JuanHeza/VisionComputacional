package main

//	go run genetico.go
import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	// "github.com/skratchdot/open-golang/open"
)

type DATA struct {
	Vias  []string
	Valor int
}

type OUT struct {
	Evaluados        []DATA
	Generacion       int
	Id               int
	PageRadioButtons []RadioButton
	DropDownMenu     []DropDown
}

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}
type DropDown struct {
	Value string
	Text  string
}

var G string
var x string
var vias [][]string
var i int = 1
var Mejor int = -99
var PD OUT

func main() {
	log.Println("Start")
	http.HandleFunc("/", Home)
	http.HandleFunc("/Gen", Mapping)
	log.Println("Handling")
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	log.Println("Opening At http://localhost:8080")
	server.ListenAndServe()
	// open.StartWith("http://localhost:8080","Chrome")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
func Home(w http.ResponseWriter, r *http.Request) {
	MyRadioButtons := []RadioButton{
		RadioButton{"Muta", "si", false, false, "Con Mutacion"},
		RadioButton{"Muta", "no", false, false, "Sin Mutacion"},
		RadioButton{"select", "txt", false, false, "Leer desde TXT"},
		RadioButton{"select", "gen", false, false, "Generar aleatoriamente"},
	}
	MyDropDown := []DropDown{
		DropDown{"4", "4"},
		DropDown{"8", "8"},
		DropDown{"16", "16"},
		DropDown{"32", "32"},
		DropDown{"64", "64"},
		DropDown{"128", "128"},
	}
	QX := OUT{
		PageRadioButtons: MyRadioButtons,
		DropDownMenu:     MyDropDown,
	}
	t, err := template.ParseFiles("Data.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, QX)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	log.Println("Parsed")
}

func Mapping(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	NumGet := r.Form.Get("select")
	Muta := r.Form.Get("Muta")
	fmt.Println(Muta)
	switch NumGet {
	case "txt":
		vias = GetData()
		break
	case "gen":
		G = r.Form.Get("Cant")
		D, err := strconv.Atoi(G)
		if err != nil {
			log.Println("Not Number")
			panic(err)
		}
		log.Println("D Value", D)
		if D >= 4 {
			vias = Generator(D)
		}
	}
	myuserlist := Values(vias, i, Muta)
	t, err := template.ParseFiles("Data.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, myuserlist)
	if err != nil {
		log.Print("template executing error: ", err)
	}
	log.Println("Parsed")
	i++
	time.Sleep(2000 * time.Millisecond)
}

func Values(dat [][]string, i int, MT string) OUT {
	var A, B int
	// for S:=0;S<2;S++{
	var Data []DATA
	log.Println("Generacion #", i)
	for _, i := range dat {
		fmt.Println(i)
	}
	Resultados := make(map[int]int)
	for W, Z := range dat {
		for i := 0; i < len(Z); i++ {
			if i%4 == 3 {
				if i == 3 {
					A, B = 0, 4
				}
				if i == 7 {
					A, B = 4, 8
				}
				if i == 11 {
					A, B = 8, 12
				}
				if i == 15 {
					A, B = 12, 16
				}
				Gen := make([]int, 4)
				var GenZ []int

				for i := A; i < B; i++ {
					D, err := strconv.Atoi(Z[i])
					if err != nil {
						log.Println("Not Number")
						panic(err)
					}
					GenZ = append(GenZ, D)
				}
				copy(Gen, GenZ)
				X, Q := evaluar(Gen, Resultados[W+1])
				fmt.Printf("%v >> %2d ", Gen, X)
				Resultados[W+1] = Q
			}
		}
		CX := DATA{Z, Resultados[W+1]}
		fmt.Println(" Valor", Resultados[W+1])
		Data = append(Data, CX)
	}
	D := 0
	for _, f := range Resultados {
		D += f
	}

	fmt.Println(" Valor de la generacion", D, "Mejor Anterior", Mejor)
	if D >= Mejor {
		Mejor = D
		QX := OUT{
			Evaluados:  Data,
			Generacion: D,
			Id:         i,
		}
		fmt.Println(Resultados)
		EX := Pares(Resultados)
		fmt.Println(EX)
		vias = NewGen(EX, dat, Resultados)
		if MT != "no" {
			vias = Patch(vias)
		}
		fmt.Println()
		PD = QX
		return PD
	}
	return PD
}

func evaluar(Gen []int, ValGen int) (int, int) {
	var valor = 0
	for i := 0; i < 4; i++ {
		if Gen[i] == 1 {
			valor++
		} else {
			valor--
		}
	}
	for i := 0; i < 3; i++ {
		if Gen[i] == 1 && Gen[i+1] == 1 {
			valor++
		}
		if Gen[i] == 0 && Gen[i+1] == 0 {
			valor--
		}
	}
	var sum = 0
	for i := 0; i < 4; i++ {
		sum += Gen[i]
	}
	if sum > 2 {
		valor++
	}
	if sum < 2 {
		valor--
	}
	ValGen += valor
	return valor, ValGen
}

func Pares(Datos map[int]int) []int {
	var Orden []int
	var Ord []int
	for _, y := range Datos {
		Ord = append(Ord, y)
	}
	sort.Ints(Ord)
	for i := 0; i < len(Datos); i++ {
		for A, B := range Datos {
			if B == Ord[i] && !Visited(A, Orden) {
				Orden = append(Orden, A)
			}
		}
	}
	return Orden
}

func Visited(cheap int, Orden []int) bool {
	for _, i := range Orden {
		if cheap == i {
			return true
		}
	}
	return false
}

func NewGen(data []int, vias [][]string, Mapa map[int]int) [][]string {
	log.Println("NEW GEN")
	var google [][]string
	//var max = Mapa[data[len(data)-1]]
	for i := len(data) - 1; i >= 0; i-- {
		for j := len(data) - 2; j >= 0; j-- {
			var Gen1, Gen2 []string
			if /*Mapa[data[i]] + Mapa[data[j]] > max && */ data[i] != data[j] {
				A := vias[data[i]-1][:8] //inicio
				B := vias[data[i]-1][8:] //final
				C := vias[data[j]-1][:8] //inicio
				D := vias[data[j]-1][8:] //final
				//A&D C&B
				fmt.Println("\tP\t", data[i], "\t-", A, "M\t", data[j], "\t-", D)
				fmt.Println("\tP\t", data[j], "\t-", C, "M\t", data[i], "\t-", B)
				Gen1 = append(Gen1, C...)
				Gen1 = append(Gen1, B...)
				Gen2 = append(Gen2, A...)
				Gen2 = append(Gen2, D...)
				google = append(google, Gen2)
				if len(google) < len(vias) {
					google = append(google, Gen1)
				}
			}
			if len(google) >= len(vias) {
				return google
			}
		}
	}
	log.Panic("BAD EXIT")
	return google
}

func Patch(vias [][]string) [][]string {
	for A, B := range vias {
		switch B[A] {
		case "1":
			B[A] = "0"
			break
		case "0":
			B[A] = "1"
			break
		}
		vias[A] = B
	}
	return vias
}

func GetData() [][]string {
	var Data [][]string
	Lineas, err := os.Open("Genetico.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer Lineas.Close()
	scanner := bufio.NewScanner(Lineas)
	for scanner.Scan() {
		x = scanner.Text()
		fila := strings.Fields(x)
		Data = append(Data, fila)
	}
	return Data
}

func Generator(G int) [][]string {
	var XO int64
	var M int64 = 16
	var XN int64
	var data [][]string
	for i := 0; i < G; i++ {
		var GN []string
		for j := 0; j < 4; j++ {
			Q := time.Now()
			XO = int64(Q.Nanosecond()) / M
			A := Q.Unix()
			fmt.Print("Gen ", i+1, "\t")
			fmt.Print("Unix ", A, "\t")
			fmt.Print("NS ", XO, "\t")

			XN = (A * XO) % M
			fmt.Print("XN ", XN, "\t")
			X := strconv.FormatInt(XN, 2)
			runes := strings.SplitN(X, "", 16)
			for len(runes) < 4 {
				runes = append([]string{"0"}, runes...)
			}
			time.Sleep(40 * time.Microsecond)
			fmt.Println(runes)
			GN = append(GN, runes...)
		}
		data = append(data, GN)
	}
	log.Println("GENERADOR - DATA", len(data))
	return data
}
