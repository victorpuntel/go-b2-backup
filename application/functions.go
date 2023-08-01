package application

import (
	"context"
	"io"
	"os"

	"github.com/kurin/blazer/b2"
)

func InitBucket(ctx context.Context, keyId, appId, bucketName string) (*b2.Bucket, error) {
	b2, err := b2.NewClient(ctx, keyId, appId)
	if err != nil {
		return nil, err
	}

	bucket, err := b2.Bucket(ctx, bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func SaveFile(ctx context.Context, bucket *b2.Bucket, src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	obj := bucket.Object(dst)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, f); err != nil {
		w.Close()
		return err
	}
	return w.Close()
}
