package importing

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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
	baseURL := "http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/"
	cache := entity.CacheFolder

	rarFile := cache + filename
	txtFile := cache + "Remuneracao.txt"

	if !exist(rarFile) {
		err := downloadHTTPFile(baseURL+filename, rarFile)
		if err != nil {
			return nil, err
		}
		if exist(txtFile) {
			err := os.Remove(txtFile)
			if err != nil {
				return nil, err
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

		completeName = row[indexName]
		// completeName = strings.Replace(row[indexName], "\u0000", "", -1)
		// completeName = strings.Replace(row[indexName], "\U0x96", "", -1)
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
		departament = row[indexDepartament]
		// departament = strings.Replace(row[indexDepartament], "\u0000", "", -1)
		// departament = strings.Replace(row[indexDepartament], "\u9600", "", -1)
		function = row[indexFunction]
		// function = strings.Replace(row[indexFunction], "\u0000", "", -1)
		// function = strings.Replace(row[indexFunction], "\u9600", "", -1)

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

	publicFuncs = makeRelevance(publicFuncs)

	return publicFuncs, nil
}

func makeRelevance(pf []entity.PublicFunc) []entity.PublicFunc {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pfWithRelevance := []entity.PublicFunc{}
	for _, p := range pf {
		r := r1.Intn(10)
		p.Relevancia = r + 1
		pfWithRelevance = append(pfWithRelevance, p)

	}
	return pfWithRelevance
}
