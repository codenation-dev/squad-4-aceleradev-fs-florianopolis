package importing

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/mholt/archiver"
)

func exist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	return true
}

// ImportPublicFuncFile implements routine to import the file from web,
// decompress it and parses the data
func ImportPublicFuncFile(month, year string) ([]entity.PublicFunc, error) {
	month = strings.Title(month)
	if len(year) == 2 {
		year = "20" + year
	}

	filename := fmt.Sprintf("remuneracao_%s_%s.rar", strings.Title(month), year)
	baseURL := "http: //www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/"
	cache := "squad-4-aceleradev-fs-florianopolis/cmd/data/cache/"

	rarFile := cache + filename
	txtFile := cache + "Remuneracao.txt"

	if !exist(rarFile) {
		err := downloadHTTPFile(baseURL+filename, cache)
		if err != nil {
			log.Fatal("donwloadHTTPFile", err)
		}
		if exist(txtFile) {
			err := os.Remove(txtFile)
			if err != nil {
				log.Fatal("os.Remove", err)
			}
		}
	}

	if !exist(txtFile) {
		err := archiver.Unarchive(rarFile, cache)
		if err != nil {
			log.Fatal("unarchive", err)
		}
	}

	publicFuncs, err := parseSPData(txtFile)
	if err != nil {
		log.Fatal("parseSPData", err)
	}
	return publicFuncs, err
}

func parseSPData(path string) ([]entity.PublicFunc, error) {
	var indexName = 0
	var indexIncome = 3
	var indexDepartament = 2
	var indexFunction = 1

	publicFuncs := []entity.PublicFunc{}

	// Send this job to retrieve only the data we need
	job := func(row []string) bool {
		var completeName, shortName, wageString, departament, function string

		completeName = strings.Replace(row[indexName], "\u0000", "", -1)
		if len(completeName) > 30 {
			shortName = completeName[:30]
		} else {
			shortName = completeName
		}

		wageString = strings.Replace(row[indexIncome], ",", ".", 1)
		wageFloat, err := strconv.ParseFloat(wageString, 64)
		if err != nil {
			panic(err) //TODO: implemntar um erro melhor, mas este job só aceita retornar bool
			// e se eu só setar o bool como "false", vai parecer que terminou o parse sem erros
		}
		departament = strings.Replace(row[indexDepartament], "\u0000", "", -1)
		function = strings.Replace(row[indexFunction], "\u0000", "", -1)

		publicFuncs = append(publicFuncs, entity.PublicFunc{
			CompleteName: completeName,
			ShortName:    shortName,
			Wage:         wageFloat,
			Departament:  departament,
			Function:     function,
		})
		return true
	}

	readCSV(path, job, ';', true)
	return publicFuncs, nil
}
