package deluge

import (
	"gopkg.in/olivere/elastic.v3"
)

// IngestorOptionFunc is a function that configures an Ingestor. It is used in
// NewIngestor.
type IngestorOptionFunc func(*Ingestor) error

// SetDocument sets the document type for the ingest.
func SetDocument(ctor Constructor) IngestorOptionFunc {
	return func(i *Ingestor) error {
		if ctor != nil {
			i.documentCtor = ctor
		}
		return nil
	}
}

// SetInput sets the input type for the ingest.
func SetInput(input Input) IngestorOptionFunc {
	return func(i *Ingestor) error {
		if input != nil {
			i.input = input
		}
		return nil
	}
}

// SetClient sets the elasticsearch client.
func SetClient(client *elastic.Client) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.client = client
		return nil
	}
}

// SetErrorThreshold sets the error threshold to stop the ingest at.
func SetErrorThreshold(threshold float64) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.threshold = threshold
		return nil
	}
}

// SetActiveConnections sets the number of active connections to elasticsearch
// allowed.
func SetActiveConnections(numActiveConnections int) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.numActiveConnections = numActiveConnections
		return nil
	}
}

// SetNumReplicas sets the number of replicas to enable upon completion of
// the ingest.
func SetNumReplicas(numReplicas int) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.numReplicas = numReplicas
		return nil
	}
}

// SetNumWorkers sets the number of workers in the worker pool.
func SetNumWorkers(numWorkers int) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.numWorkers = numWorkers
		return nil
	}
}

// ClearExistingIndex clears an existing index if specified.
func ClearExistingIndex(clearExisting bool) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.clearExisting = clearExisting
		return nil
	}
}

// SetCompression sets the compression type for the input files. Supports:
// "bzip2", "flate", "gzip", "zlib".
func SetCompression(compression string) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.compression = compression
		return nil
	}
}

// SetIndex sets the index name to create and ingest into.
func SetIndex(index string) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.index = index
		return nil
	}
}

// SetBulkByteSize sets the maximum number of bytes in a bulk index payload.
func SetBulkByteSize(numBytes int64) IngestorOptionFunc {
	return func(i *Ingestor) error {
		i.bulkByteSize = numBytes
		return nil
	}
}
