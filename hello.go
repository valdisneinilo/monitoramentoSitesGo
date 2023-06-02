package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	name := showApresentation()

	for {
		showMenu()
		comand := readingComand()
		switch comand {
		case 1:
			monitoring(name)
		case 2:
			fmt.Println("Exibindo Logs...")
			showLogs()
		case 3:
			fmt.Println("Digite o site que deseja fazer o monitoramento: EX: https://www.google.com.br")
			fmt.Println("Lembrese de colocar o http:// ou https:// antes do site ")
			var site string
			fmt.Scan(&site)
			registerSites(site)
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showApresentation( ) string {
	var name string
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&name)
	fmt.Println("Olá,", name, "o que deseja fazer?")
	fmt.Println("")
	return name
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Adicionar site para monitoramento")
	fmt.Println("0- Sair do Programa")
	fmt.Println("")
}

func readingComand() int {
	var radedComand int
	var confirmedComand int

	fmt.Scan(&radedComand)
	fmt.Println("")
	fmt.Println("O comando escolhido foi", radedComand, "?")
	fmt.Println("Digite 1 para SIM ou 2  para NÃO")
	fmt.Println("")
	fmt.Scan(&confirmedComand)
	fmt.Println("")

	if confirmedComand == 1 {
	return radedComand
	} else if confirmedComand == 2 {
		return 0
	}else {
		fmt.Println("Comando não reconhecido")
		return 404
	}
}

func monitoring(name string) {
	var toMonitor int
	const delay = 60 
	sites := readingFile()

	fmt.Println("O programa ira monitorar seus sites a cada 60 segundos")
	fmt.Println("Quantas vezes deseja repetir o monitoramento?")
	fmt.Println("")
	fmt.Scan(&toMonitor)
	fmt.Println("")
	fmt.Println("Monitorando...")

	for i := 0; i < toMonitor; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i + 1 , ":", site)
			testSite(site, name)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string, name string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Carregado com sucesso! status:", resp.StatusCode)
		registraLog(site, true, name)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false, name)
	}
}

func readingFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	toReader := bufio.NewReader(file)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	for {
		line, err := toReader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}else{
			sites = append(sites, line)
		}
		if err == io.EOF {
			break
		}

	}

	file.Close()
	return sites
}

func registraLog(site string, status bool, name string) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString( name + " - " + time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}

func registerSites(site string) {
	file, err := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(  site + "\n")
	file.Close()
}
