package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/victorpuntel/b2-backup/application"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		pError(errors.New("Usage: b2-backup [src] [dst]"), "Invalid arguments")
	}
	src, dst := args[0], args[1]

	err := godotenv.Load()
	pError(err, "Error loading dotenv")

	keyId, appKey, bucketName := os.Getenv("B2_KEY_ID"), os.Getenv("B2_APP_KEY"), os.Getenv("B2_BUCKET_NAME")

	ctx := context.Background()

	bucket, err := application.InitBucket(ctx, keyId, appKey, bucketName)
	pError(err, "InitClient")

	err = application.SaveFile(ctx, bucket, src, dst)
	pError(err, "CopyFile")
}

func pError(err error, desc string) {
	if err != nil {
		log.Fatalf("[%s]: %s", desc, err)
	}
}
