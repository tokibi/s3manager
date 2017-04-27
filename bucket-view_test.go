package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	minio "github.com/minio/minio-go"
	"github.com/stretchr/testify/assert"
)

func TestBucketViewHandler(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		s3                   S3Client
		bucketName           string
		expectedStatusCode   int
		expectedBodyContains string
	}{
		"success (empty bucket)": {
			s3: &S3ClientMock{
				Buckets: []minio.BucketInfo{
					{Name: "testBucket"},
				},
			},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "No objects in",
		},
		"success (with file)": {
			s3: &S3ClientMock{
				Buckets: []minio.BucketInfo{
					{Name: "testBucket"},
				},
				Objects: []minio.ObjectInfo{
					{Key: "testFile"},
				},
			},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "testBucket",
		},
		"success (archive)": {
			s3: &S3ClientMock{
				Buckets: []minio.BucketInfo{
					{Name: "testBucket"},
				},
				Objects: []minio.ObjectInfo{
					{Key: "archive.tar.gz"},
				},
			},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "archive",
		},
		"success (image)": {
			s3: &S3ClientMock{
				Buckets: []minio.BucketInfo{
					{Name: "testBucket"},
				},
				Objects: []minio.ObjectInfo{
					{Key: "testImage.png"},
				},
			},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "photo",
		},
		"success (sound)": {
			s3: &S3ClientMock{
				Buckets: []minio.BucketInfo{
					{Name: "testBucket"},
				},
				Objects: []minio.ObjectInfo{
					{Key: "testSound.mp3"},
				},
			},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusOK,
			expectedBodyContains: "music_note",
		},
		"bucket doesn't exist": {
			s3:                   &S3ClientMock{},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusNotFound,
			expectedBodyContains: http.StatusText(http.StatusNotFound),
		},
		"s3 error": {
			s3: &S3ClientMock{
				Err: errors.New("mocked S3 error"),
			},
			bucketName:           "testBucket",
			expectedStatusCode:   http.StatusInternalServerError,
			expectedBodyContains: http.StatusText(http.StatusInternalServerError),
		},
	}

	for tcID, tc := range tests {
		r := mux.NewRouter()
		r.
			Methods(http.MethodGet).
			Path("/buckets/{bucketName}").
			Handler(BucketViewHandler(tc.s3))

		ts := httptest.NewServer(r)
		defer ts.Close()

		url := fmt.Sprintf("%s/buckets/%s", ts.URL, tc.bucketName)
		resp, err := http.Get(url)
		assert.NoError(err, tcID)
		defer func() {
			err = resp.Body.Close()
			assert.NoError(err, tcID)
		}()

		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(err, tcID)

		assert.Equal(tc.expectedStatusCode, resp.StatusCode, tcID)
		assert.Contains(string(body), tc.expectedBodyContains, tcID)
	}
}