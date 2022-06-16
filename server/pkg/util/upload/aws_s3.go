package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
)

type AwsS3 struct{}

//@object: *AwsS3
//@function: UploadFile
//@description: Upload file to Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-upload-file-to-bucket
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	session := newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := config.C.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.C.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		return "", "", err
	}

	return config.C.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

//@object: *AwsS3
//@function: DeleteFile
//@description: Delete file from Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-delete-bucket-item
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) DeleteFile(key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := config.C.AwsS3.PathPrefix + "/" + key
	bucket := config.C.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return errors.New("function svc.DeleteObject() Filed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

// newSession Create S3 session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(config.C.AwsS3.Region),
		Endpoint:         aws.String(config.C.AwsS3.Endpoint), //minio在这里设置地址,可以兼容
		S3ForcePathStyle: aws.Bool(config.C.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(config.C.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			config.C.AwsS3.SecretID,
			config.C.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
