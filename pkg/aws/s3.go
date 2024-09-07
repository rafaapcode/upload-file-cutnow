package aws_s3

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadSingleFile(bucketName, filepath string, file multipart.File) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath),
		Body:   file,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func UploadMultipleFile(bucketName, filepath string, file multipart.File) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath),
		Body:   file,
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func DeleteMultipleImages(bucketName, filepath string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	srv := s3.New(sess)

	_, err = srv.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath),
	})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Imagens excluídas com sucesso")
	return nil
}
