package s3

import (
	"context"
	"io"
	"strings"
	"testing"
	"time"
)

func TestCloudProvider(t *testing.T) {
	ctx := context.Background()

	// Use memblob for portable unit testing
	cp, err := NewCloudProvider(ctx, "mem://")
	if err != nil {
		t.Fatalf("failed to create cloud provider: %v", err)
	}
	defer cp.Close()

	content := "s3 file content"
	path := "test-s3.txt"

	// Test Upload
	_, err = cp.Upload(ctx, path, strings.NewReader(content))
	if err != nil {
		t.Fatalf("upload failed: %v", err)
	}

	// Test Open
	rc, err := cp.Open(ctx, path)
	if err != nil {
		t.Fatalf("open failed: %v", err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if string(data) != content {
		t.Errorf("expected %s, got %s", content, string(data))
	}

	// Test GetURL
	url, err := cp.GetURL(ctx, path, time.Hour)
	if err != nil {
		t.Fatalf("get url failed: %v", err)
	}
	// Fallback should be prefix/path
	if !strings.HasSuffix(url, path) {
		t.Errorf("expected url to end with %s, got %s", path, url)
	}

	// Test Delete
	err = cp.Delete(ctx, path)
	if err != nil {
		t.Errorf("delete failed: %v", err)
	}

	// Verify deleted
	_, err = cp.Open(ctx, path)
	if err == nil {
		t.Error("expected error opening deleted file")
	}
}
