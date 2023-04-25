package utility

import (
	"bytes"
	"context"
	"fmt"
	"iFIX/ifix/entities"
	"net/url"
	"strings"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/gofrs/uuid"
)

func DownloadBlobToBuffer(cred entities.FileuploadEntity) (bytes.Buffer, error) {
	downloadedData := bytes.Buffer{}
	azrKey, accountName, endPoint, container := GetAccountInfo(cred)
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", cred.Filename))
	fmt.Println(u)
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey)
	if errC != nil {
		return downloadedData, errC
	}
	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background()
	// Here's how to download the blob
	downloadResponse, err := blockBlobUrl.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		return downloadedData, err
	}
	// NOTE: automatically retries are performed if the connection fails
	bodyStream := downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20})

	// read the body into a buffer

	_, err = downloadedData.ReadFrom(bodyStream)
	return downloadedData, err

}
func UploadBytesToBlob(b []byte, cred entities.FileuploadEntity, fileName string, fileType string) (entities.FileuploadEntity, error) {

	azrKey, accountName, endPoint, container := GetAccountInfo(cred)
	cred.Originalfile = fileName
	cred.Filename = GetBlobName(fileName)
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", cred.Filename))
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey)
	if errC != nil {
		return cred, errC
	}

	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{}))

	ctx := context.Background()
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType: fileType,
		},
	}

	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	cred.Path = blockBlobUrl.String()
	return cred, errU
}

func GetAccountInfo(cred entities.FileuploadEntity) (string, string, string, string) {
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", cred.Credentialaccount)
	return cred.Credentialkey, cred.Credentialaccount, azrPrimaryBlobServiceEndpoint, cred.Credentialpassword
}

func GetBlobName(fileName string) string {
	t := time.Now()
	uuid, _ := uuid.NewV4()
	splitedFileName := strings.Split(fileName, ".")
	ext := splitedFileName[len(splitedFileName)-1]
	return fmt.Sprintf("%s-%v.%s", t.Format("20060102"), uuid, ext)
}
