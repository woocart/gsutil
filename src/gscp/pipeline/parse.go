package pipeline

import (
	"context"
	"io"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

const (
	Cloud = iota
	Local
	STDIO
)

// Pipeline holds data about IO operation
type Pipeline struct {
	Bucket *storage.BucketHandle
	Type   int
	Path   string
	Reader io.Reader
	Writer io.Writer
}

// NewPipeline creates new IO operation object
func NewPipeline(ctx context.Context, input string) (*Pipeline, error) {

	p := &Pipeline{}

	// GSCloud type
	if strings.HasPrefix(input, "gs://") {
		p.Type = Cloud
		u, err := url.Parse(input)
		if err != nil {
			return nil, err
		}

		// Creates a GS client.
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}

		p.Bucket = client.Bucket(u.Host)
		p.Path = strings.TrimPrefix(u.Path, "/")
		p.Writer = p.Bucket.Object(p.Path).NewWriter(ctx)
		p.Reader, _ = p.Bucket.Object(p.Path).NewReader(ctx)
		return p, nil
	}

	// Read from stdin
	if input == "stdio" {
		p.Type = STDIO
		p.Writer = os.Stdout
		p.Reader = os.Stdin
		return p, nil
	}

	// Read or write to local file
	p.Type = Local
	p.Path = input
	p.Reader, _ = os.Open(input)
	p.Writer, _ = os.Create(input)
	return p, nil
}

// IsCloud returns true if this is Google Storage operation
func (p *Pipeline) IsCloud() bool {
	return p.Type == Cloud
}

// IsLocal returns true if this is local file operation
func (p *Pipeline) IsLocal() bool {
	return p.Type == Local
}

// IsStdio returns true if this is standard input/output operation
func (p *Pipeline) IsStdio() bool {
	return p.Type == STDIO
}

// String returs human redebale representation of pipeline
func (p *Pipeline) String() string {
	return p.Path
}
