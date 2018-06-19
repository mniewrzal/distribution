package ocischema

import (
	"context"

	"github.com/docker/distribution"
	"github.com/docker/distribution/manifest"
	"github.com/opencontainers/go-digest"
	"github.com/opencontainers/image-spec/specs-go/v1"
)

// builder is a type for constructing manifests.
type Builder struct {
	// bs is a BlobService used to publish the configuration blob.
	bs distribution.BlobService

	// configJSON references
	configJSON []byte

	// layers is a list of layer descriptors that gets built by successive
	// calls to AppendReference.
	layers []distribution.Descriptor

	// Annotations contains arbitrary metadata relating to the targeted content.
	annotations map[string]string

	// For testing purposes
	mediaType string
}

// NewManifestBuilder is used to build new manifests for the current schema
// version. It takes a BlobService so it can publish the configuration blob
// as part of the Build process, and annotations.
func NewManifestBuilder(bs distribution.BlobService, configJSON []byte, annotations map[string]string) distribution.ManifestBuilder {
	mb := &Builder{
		bs:          bs,
		configJSON:  make([]byte, len(configJSON)),
		annotations: annotations,
		mediaType:   v1.MediaTypeImageManifest,
	}
	copy(mb.configJSON, configJSON)

	return mb
}

// For testing purposes, we want to be able to create an OCI image with
// either an MediaType either empty, or with the OCI image value
func (mb *Builder) SetMediaType(mediaType string) {
	if mediaType != "" && mediaType != v1.MediaTypeImageManifest {
		panic("Invalid media type for OCI image manifest")
	}

	mb.mediaType = mediaType
}

// Build produces a final manifest from the given references.
func (mb *Builder) Build(ctx context.Context) (distribution.Manifest, error) {
	m := Manifest{
		Versioned: manifest.Versioned{
			SchemaVersion: 2,
			MediaType:     mb.mediaType,
		},
		Layers:      make([]distribution.Descriptor, len(mb.layers)),
		Annotations: mb.annotations,
	}
	copy(m.Layers, mb.layers)

	configDigest := digest.FromBytes(mb.configJSON)

	var err error
	m.Config, err = mb.bs.Stat(ctx, configDigest)
	switch err {
	case nil:
		// Override MediaType, since Put always replaces the specified media
		// type with application/octet-stream in the descriptor it returns.
		m.Config.MediaType = v1.MediaTypeImageConfig
		return FromStruct(m)
	case distribution.ErrBlobUnknown:
		// nop
	default:
		return nil, err
	}

	// Add config to the blob store
	m.Config, err = mb.bs.Put(ctx, v1.MediaTypeImageConfig, mb.configJSON)
	// Override MediaType, since Put always replaces the specified media
	// type with application/octet-stream in the descriptor it returns.
	m.Config.MediaType = v1.MediaTypeImageConfig
	if err != nil {
		return nil, err
	}

	return FromStruct(m)
}

// AppendReference adds a reference to the current ManifestBuilder.
func (mb *Builder) AppendReference(d distribution.Describable) error {
	mb.layers = append(mb.layers, d.Descriptor())
	return nil
}

// References returns the current references added to this builder.
func (mb *Builder) References() []distribution.Descriptor {
	return mb.layers
}
