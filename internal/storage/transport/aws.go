package transport

import (
	"bytes"
	"k8scommerce/internal/storage/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MustNewAwsTransport(cfg config.AWSConfig) (Transport, error) {
	t := &awsTransport{
		cfg: cfg,
	}
	err := t.openSession()
	return t, err
}

type awsTransport struct {
	cfg config.AWSConfig

	svc    *s3.S3
	awsCfg *aws.Config
	sess   *session.Session

	putPath     string
	contentType string

	multipart     *s3.CreateMultipartUploadOutput
	competedParts []*s3.CompletedPart
}

func (t *awsTransport) Open(destinationPath, fileName, contentType string) error {
	t.putPath = destinationPath + fileName
	t.contentType = contentType

	if t.sess == nil {
		t.openSession()
	}

	t.svc = s3.New(t.sess, t.awsCfg)

	// do we have access to this bucket?
	// check by getting the bucket's acl
	_, err := t.svc.GetBucketAcl(&s3.GetBucketAclInput{Bucket: &t.cfg.S3Bucket})
	if err != nil {
		return status.Errorf(codes.Internal, "access to bucket %s deined. %s", t.cfg.S3Bucket, err.Error())
	}

	if err := t.createMultipartUpload(); err != nil {
		return err
	}

	return nil
}

func (t *awsTransport) createMultipartUpload() error {
	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(t.cfg.S3Bucket),
		Key:         aws.String(t.putPath),
		ContentType: aws.String(t.contentType),
		ACL:         aws.String("public-read"),
	}

	resp, err := t.svc.CreateMultipartUpload(input)
	if err != nil {
		return status.Errorf(codes.Internal, "could not create CreateMultipartUpload object. %s", err.Error())
	}

	t.multipart = resp
	return nil
}

func (t *awsTransport) StreamPut(buffer []byte, partNumber int) error {
	const maxRetries = 3
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(buffer),
		Bucket:        t.multipart.Bucket,
		Key:           t.multipart.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      t.multipart.UploadId,
		ContentLength: aws.Int64(int64(len(buffer))),
	}

	for tryNum <= maxRetries {
		uploadResult, err := t.svc.UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				err := t.abort()
				if err != nil {
					return err
				}

				if err, ok := err.(awserr.Error); ok {
					return status.Errorf(codes.Internal, "could not connect to S3, aborting. %s", err.Error())
				}
				return status.Errorf(codes.Internal, "max retries reached, aborting. %s", err.Error())
			}
			tryNum++
		} else {
			t.competedParts = append(t.competedParts, &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			})
			return nil
		}
	}
	return nil
}

func (t *awsTransport) abort() error {
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   t.multipart.Bucket,
		Key:      t.multipart.Key,
		UploadId: t.multipart.UploadId,
	}
	_, err := t.svc.AbortMultipartUpload(abortInput)
	if err != nil {
		return status.Errorf(codes.Internal, "multipart abort failed. %s", err.Error())
	}
	return nil
}

func (t *awsTransport) Close() error {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   t.multipart.Bucket,
		Key:      t.multipart.Key,
		UploadId: t.multipart.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: t.competedParts,
		},
	}
	_, err := t.svc.CompleteMultipartUpload(completeInput)
	if err != nil {
		return status.Errorf(codes.Internal, "failed closing multipart upload. %s", err.Error())
	}
	return err
}

func (t *awsTransport) openSession() error {
	creds := credentials.NewStaticCredentials(t.cfg.AccessKeyId, t.cfg.SecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		return status.Error(codes.Internal, "bad aws credentials")
	}
	t.awsCfg = aws.NewConfig().WithRegion(t.cfg.Region).WithCredentials(creds)

	t.sess, err = session.NewSession(t.awsCfg)
	if err != nil {
		return status.Error(codes.Internal, "could not create aws session")
	}
	return nil
}
