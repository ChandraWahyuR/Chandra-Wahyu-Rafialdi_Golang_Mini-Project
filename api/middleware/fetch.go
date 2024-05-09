package middleware

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func GetImage(imageFile *multipart.FileHeader, folder string) (string, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	// Create instance for cloudinary
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalf("Error creating Cloudinary instance: %v", err)
	}

	// Upload image
	file, err := imageFile.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Upload to cloudinary
	uploadParams := uploader.UploadParams{Folder: folder}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
