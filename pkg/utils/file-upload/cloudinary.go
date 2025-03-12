package fileupload

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(cld *cloudinary.Cloudinary, fileHeader *multipart.FileHeader, folderPath *string) (*uploader.UploadResult, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// customable fileName
	formatFilename := fmt.Sprintf("%v-%s", time.Now(), fileHeader.Filename)
	if folderPath == nil {
		defaultFolderPath := "ordered-golang-web"
		folderPath = &defaultFolderPath
	}

	// Upload file to Cloudinary
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: formatFilename, // Nama file di Cloudinary
		Folder:   *folderPath,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to Cloudinary: %v", err)
	}

	return uploadResult, nil
}
