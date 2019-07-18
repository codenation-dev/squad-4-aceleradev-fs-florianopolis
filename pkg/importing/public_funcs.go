package importing

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/mholt/archiver"
)

// FetchPublicAgentsFile import data from the goverment site
func FetchPublicAgentsFile(uf, month, year string) ([]entity.PublicFunc, error) {
	var downloadFrom string
	switch uf {
	case "SP":
		downloadFrom = fmt.Sprintf("http://www.transparencia.sp.gov.br/PortalTransparencia-Report/historico/remuneracao_%s_%s.rar", month, year)
	default:
		return nil, entity.ErrDownloadingFile
	}

	filename := fmt.Sprintf("%s_%s_%s", uf, year, month)
	splited := strings.Split(downloadFrom, ".")
	extension := splited[len(splited)-1]
	downloadedFile := fmt.Sprintf("%s/%s.%s", entity.CacheFolder, filename, extension)

	if _, err := os.Stat(downloadedFile); err != nil {
		err = downloadHTTPFile(downloadFrom, downloadedFile)
		if err != nil {
			return nil, err
		}
	}

	decompressedFile := fmt.Sprintf("%s/%s.txt", entity.CacheFolder, filename)
	if _, err := os.Stat(decompressedFile); err != nil {
		if uf == "SP" {
			err := os.Remove(entity.CacheFolder + "/Remuneracao.txt")
			if err != nil {
				return nil, err
			}
		}
		err := archiver.Unarchive(downloadedFile, entity.CacheFolder+"/")
		if err != nil {
			return nil, err
		}
	}
	if uf == "SP" {
		_, err := os.Stat(entity.CacheFolder + "/Remuneracao.txt")
		if err == nil {
			err := os.Rename(entity.CacheFolder+"/Remuneracao.txt", decompressedFile)
			if err != nil {
				return nil, err
			}
		}
	}

	return parseData(uf, month, year)

}

func parseData(uf, month, year string) ([]entity.PublicFunc, error) {
	switch uf {
	case "SP":
		filename := fmt.Sprintf("%s_%s_%s", uf, year, month)
		decompressedFile := fmt.Sprintf("%s/%s.txt", entity.CacheFolder, filename)
		return parseSPData(decompressedFile)
	default:
		return []entity.PublicFunc{}, entity.ErrDownloadingFile
	}
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
