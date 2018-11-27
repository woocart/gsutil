package main

import (
	"context"
	"gscp/logger"
	"gscp/pipeline"
	"gscp/version"
	"io"
	"os"
	"reflect"

	"cloud.google.com/go/storage"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var log = logger.New("gscp")
var app = kingpin.New("gscp", "Copies data from and to Google Cloud Storage")

var from = app.Arg("from", "where to read from: gs://bucketname/path or - from stdin or /path/ for local file").Required().String()
var to = app.Arg("to", "Where to write to: gs://bucketname/path or - to stdout or /path/ for local file").Required().String()
var metadata = pipeline.Metadata(app.Arg("metadata", "KV pairs to append to uploaded object"))

func main() {
	app.Author("dz0ny")
	app.Version(version.String())
	kingpin.MustParse(app.Parse(os.Args[1:]))

	ctx := context.Background()
	in, err := pipeline.NewPipeline(ctx, *from)
	if err != nil {
		log.Fatal(err)
	}
	out, err := pipeline.NewPipeline(ctx, *to)
	if err != nil {
		log.Fatal(err)
	}

	if in.IsCloud() && out.IsCloud() {
		log.Fatal("Copying beetwen bucket is not yet supported.")
	}
	if reflect.ValueOf(in.Reader).IsNil() {
		log.Fatalf("%s does not exist.", in.Path)
	}

	if (in.IsStdio() && out.IsCloud()) || (in.IsCloud() && out.IsStdio()) {
		if _, err = io.Copy(out.Writer, in.Reader); err != nil {
			log.Fatal(err)
		}

		switch out.Writer.(type) {
		case *storage.Writer:
			wrt := out.Writer.(*storage.Writer)
			if err = wrt.Close(); err != nil {
				log.Fatal(err)
			}
		default:
		}
		if out.IsCloud() {
			if _, err = out.Object.Update(ctx, storage.ObjectAttrsToUpdate{Metadata: metadata}); err != nil {
				log.Fatal(err)
			}
		}
		log.Debugf("Copying done from %s to %s", in, out)
	}

}
