package storage

import (
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Storage est l'implémentation de StorageService pour Amazon S3
type S3Storage struct {
	svc    *s3.S3
	bucket string
}

// NewS3Storage crée une nouvelle instance de S3Storage
func NewS3Storage(bucket string) *S3Storage {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-north-1"), // Remplace par ta région
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	return &S3Storage{
		svc:    s3.New(sess),
		bucket: bucket,
	}
}

// UploadFile télécharge un fichier sur S3 et retourne son URL
func (s *S3Storage) UploadFile(file multipart.File, fileName string) (string, error) {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
		Body:   file,
		ACL:    aws.String("private"),
	})
	if err != nil {
		return "", err
	}

	return s.GetFileURL(fileName), nil
}

// DeleteFile supprime un fichier de S3
func (s *S3Storage) DeleteFile(fileName string) error {
	_, err := s.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
	})
	return err
}

// GetFileURL retourne l'URL du fichier stocké
func (s *S3Storage) GetFileURL(fileName string) string {
	return "https://" + s.bucket + ".s3.amazonaws.com/" + fileName
}
