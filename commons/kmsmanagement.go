package commons

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func Encrypt(key string) []byte {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(GoDotEnvVariable("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(GoDotEnvVariable("AWS_ACCESS_KEY"), GoDotEnvVariable("AWS_SECRET_KEY"), ""),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
	}

	svc := kms.New(sess)
	input := &kms.EncryptInput{
		KeyId:     aws.String(GoDotEnvVariable("AWS_KMS_KEY_ID")),
		Plaintext: []byte(key),
	}

	result, err := svc.Encrypt(input)
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}
	return result.CiphertextBlob
}

func Decrypt(arr []byte) string {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(GoDotEnvVariable("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(GoDotEnvVariable("AWS_ACCESS_KEY"), GoDotEnvVariable("AWS_SECRET_KEY"), ""),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
	}

	svc := kms.New(sess)
	input := &kms.DecryptInput{
		CiphertextBlob: arr,
		KeyId:          aws.String(GoDotEnvVariable("AWS_KMS_KEY_ID")),
	}

	result, err := svc.Decrypt(input)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(result.Plaintext)
}
