package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
	"strconv"
	"io/ioutil"
) 

const monitoramentos = 5
const delay = 10



func menu(){
	fmt.Println("")
	fmt.Println("1- Iniciar o monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Encerrar o programa")

	fmt.Println("")

}


func getValue() int {
		var value int
		fmt.Scan(&value)
		return value
}


func iniciandoMonitoramento(){
	fmt.Println("Monitoramento...")
	sites := isReadFile()
	
	for i:=0; i < monitoramentos; i++{
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site )
			isOnline(site)
		}
		time.Sleep(delay * time.Second)
	}
}
		


func isOnline(site string){
	res, err := http.Get(site)

	if err != nil{
		fmt.Println("Erro:", err)
		
	}

	if res.StatusCode == 200{
		fmt.Println("O site:", site, "Está tudo ok, foi carregado com sucesso..." )
		logRegister(site, true)
	}else{
		fmt.Println("O site:", site, "Está apresentando este código", res.StatusCode )
		logRegister(site, false)
	}
	fmt.Println("")
}

func isReadFile() []string {
	var sites []string
	file, err := os.Open("url.txt")
	
	if err != nil{
		fmt.Println("Erro: ao abrir o arquivo função isRead:", err)
	}

	reader := bufio.NewReader(file)
	for{
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)


		if err == io.EOF{
			break
		}
		fmt.Println(line)
	}

	file.Close()
	return sites
}

func logRegister(site string, status bool){
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("Ocorreu um erro ao abrir o arquivo de log:", err)
	}
	fmt.Println(file)

	file.WriteString("Dia e hora desse log: " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func consoleLogs(){
	file, err := ioutil.ReadFile("log.txt")
	if err != nil{
		fmt.Println("Erro no arquivo de console dos logs: ", err)
	}
	fmt.Println(string(file))
}


func main() {
	fmt.Println("Ola Sr.")
	for {
			menu()
			command := getValue()
			switch command {
				case 1: 
					iniciandoMonitoramento()
				case 2:
					fmt.Println("Exibindo o histórico de logs...")
					consoleLogs()
				case 0:
					fmt.Println("Saindo do programa...")
					os.Exit(0)
				default:
					fmt.Println("Não reconheço esse comando...")
					os.Exit(-1)
			}		
	}
}