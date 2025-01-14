package initializers

import (
	"log"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func ConnectMinio(config *Config) {
	endpoint := config.MinioEndpoint
	accessKeyID := config.MinioAccessKey
	secretAccessKey := config.MinioSecretKey
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Println("✔ Successfully connected to Minio")
	log.Printf("%#v\n", minioClient) // minioClient is now setup
}
