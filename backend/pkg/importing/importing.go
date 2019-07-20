// Package importing implementa funções necessárias à importação dos
// dados tanto de clientes quanto de funcionários públicos.
package importing

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"

	"github.com/mholt/archiver"
)

func readCSV(path string, job func([]string) bool, sep rune, hasHeader bool) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open path %v: %v", path, err)
	}
	fmt.Println("CSV opened")
	defer f.Close()

	r := csv.NewReader(bufio.NewReader(f))
	r.Comma = sep
	if hasHeader {
		r.Read()
	}
	counter := 0
	for {
		counter++
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Could not read csv file: %v", err)
		}
		if counter%100000 == 0 {
			fmt.Println("Reading CSV...")
		}
		keepgoing := job(line)
		if keepgoing != true {
			break
		}
	}
}

// ImportClientesCSV imports the data from clientes.csv file
func ImportClientesCSV(path string) ([]model.Cliente, error) {
	var clienteList []model.Cliente

	job := func(row []string) bool {
		clienteList = append(clienteList, model.Cliente{
			Nome:         prepareString(row[0]),
			NomePesquisa: getNameSearch(row[0]),
		})
		return true
	}
	readCSV(path, job, ',', true)

	return clienteList, nil
}

// DownloadHTTPFile downloads the file with the list of SP public agents
// and returns the path to the downloaded file
func DownloadHTTPFile(path, filename string) (string, error) {
	r, err := http.Get(path + filename + ".rar")
	if err != nil {
		return "", err
	}
	destPath := "file.rar"
	dest, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	_, err = io.Copy(dest, r.Body)
	if err != nil {
		return "", err
	}

	return destPath, nil
}

func getNameSearch(name string) string {
	if len(name) > 30 {
		name = name[0:30]
	}
	//name = strings.ReplaceAll(name, " ", "")
	//hash := md5.Sum([]byte(strings.ToUpper(name)))
	//return hex.EncodeToString(hash[:])
	return name
}

func convertValue(value string) float64 {
	value = strings.Replace(value, ",", ".", 1)
	convertedValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Fatal(err)
	}
	return convertedValue
}

func convertToDate(value string) *time.Time {
	const timeFormat = "2006-01-02"
	convertedValue, err := time.Parse(timeFormat, "2019-04-01")
	if err != nil {
		log.Fatal(err)
	}
	return &convertedValue
}

func prepareString(value string) string {
	return strings.Replace(value, "\x00", "", -1)
}

// ImportPublicFunc import from the goverment site
func ImportPublicFunc() ([]model.Funcionario, error) {
	if _, err := os.Stat("file.rar"); os.IsNotExist(err) {
		err := fetchPublicAgentsFile()
		if err != nil {
			log.Fatal(err)
		}
	}

	funcionarioList := []model.Funcionario{}

	job := func(row []string) bool {
		funcionarioList = append(funcionarioList, model.Funcionario{
			MesReferencia:   convertToDate("2019-04-01"),
			Nome:            prepareString(row[0]),
			NomePesquisa:    getNameSearch(row[0]),
			Cargo:           prepareString(row[1]),
			Orgao:           prepareString(row[2]),
			Estado:          "São Paulo",
			SalarioMensal:   convertValue(row[3]),
			SalarioFerias:   convertValue(row[4]),
			PagtoEventual:   convertValue(row[5]),
			LicencaPremio:   convertValue(row[6]),
			AbonoSalario:    convertValue(row[7]),
			RedutorSalarial: convertValue(row[8]),
			TotalLiquido:    convertValue(row[9]),
		})
		return true
	}

	readCSV("Remuneracao.txt", job, ';', true) //TODO: Fazer dinâmico, pode escolher qual mês baixar (ou atual)
	fmt.Println(funcionarioList[:10])
	return funcionarioList, nil
}

func fetchPublicAgentsFile() error { //TODO: acrescentar opção para escolher qual mês baixar
	filename := "Remuneracao_Abril_2019"
	path := "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/"
	destFolder := "."
	compressedFile := "file.rar"

	if _, err := os.Stat("file.rar"); err != nil {
		_, err := DownloadHTTPFile(path, filename)
		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(filename + ".txt"); err != nil {
		err := archiver.Unarchive(compressedFile, destFolder)
		if err != nil {
			return err
		}
	}
	return nil
}
