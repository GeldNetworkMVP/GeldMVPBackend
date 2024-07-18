package businessFacade

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
	"github.com/GeldNetworkMVP/GeldMVPBackend/services/ipfsServices"
	"github.com/GeldNetworkMVP/GeldMVPBackend/utilities/logs"
)

var (
	fileBaseBucket = os.Getenv("FILEBASE_BUCKET")
)

func UploadFilesToIpfs(fileObj model.TokenPayload) (string, error) {
	var folderPath string

	folderPath = "geldtokens/" + fileObj.PlotID + fileObj.TokenName

	errWhenCreatingFolder := ipfsServices.CreateFolder(fileBaseBucket, folderPath)
	if errWhenCreatingFolder != nil {
		return "", errWhenCreatingFolder
	}

	cid, errWhenUploadingFileToIpfs := InitiateUpload(fileObj.FileType, fileObj.TokenPayload, fileObj.TokenName, folderPath)
	if errWhenUploadingFileToIpfs != nil {
		return "", errWhenUploadingFileToIpfs
	}

	logs.InfoLogger.Println("CID Hash : " + cid)

	return cid, nil
}

func InitiateUpload(fileType string, fileContent string, fileName string, folderName string) (string, error) {
	var fileNameInLocation string
	var dec []byte

	if fileType == "txt" {
		dec = []byte(fileContent)
		fileNameInLocation = fileName + ".txt"
		fileNameInLocation = strings.ToLower(fileNameInLocation)
	} else {
		return "", errors.New("Invalid file type")
	}

	workingDirectory, errWhenGettingTheDirectory := os.Getwd()
	if errWhenGettingTheDirectory != nil {
		logs.ErrorLogger.Println("Error when getting the working directory : ", errWhenGettingTheDirectory.Error())
		return "", errWhenGettingTheDirectory
	}

	filePath := filepath.Join(workingDirectory, fileNameInLocation)
	file, errWhenCreatingFile := os.Create(filePath)
	if errWhenCreatingFile != nil {
		logs.ErrorLogger.Println("Error when creating")
		return "", errWhenCreatingFile
	}
	defer file.Close()

	if _, errWhenWritingToFile := file.Write(dec); errWhenWritingToFile != nil {
		logs.ErrorLogger.Println("Error when writing data into the file : ", errWhenWritingToFile.Error())
		return "", errWhenWritingToFile
	}
	if errWhenSyncing := file.Sync(); errWhenSyncing != nil {
		logs.ErrorLogger.Println("Error when and clearing memory : ", errWhenSyncing.Error())
		return "", errWhenSyncing
	}

	cid, _, errWhenUploadingToIpfs := ipfsServices.UploadFile(filePath, fileNameInLocation, fileBaseBucket, folderName)
	if errWhenUploadingToIpfs != nil {
		logs.ErrorLogger.Println("Error when uploading to IPFS : ", errWhenUploadingToIpfs)
		errWhenRemovingFile := os.Remove(filePath)
		if errWhenRemovingFile != nil {
			logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingFile)
			return "", errWhenRemovingFile
		}
		return "", errWhenUploadingToIpfs
	}

	errWhenClosingTheFile := file.Close()
	if errWhenClosingTheFile != nil {
		logs.ErrorLogger.Println("Error when closing the file : ", errWhenClosingTheFile)
		return "", errWhenClosingTheFile
	}

	errWhenRemovingTheFile := os.Remove(filePath)
	if errWhenRemovingTheFile != nil {
		logs.ErrorLogger.Println("Error when removing the file : ", errWhenRemovingTheFile)
		return "", errWhenRemovingTheFile
	}
	return cid, nil
}
