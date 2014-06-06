package icecast

// Mount depicts a singular icecast mountpoint. A mountpoint can have
// many clients (same as plain icecast) and have many sources (not the
// same as icecast).
type Mount struct {
	Name    string
	sources *Container
}

func NewMount(name string) *Mount {
	return &Mount{
		Name:    name,
		sources: NewContainer(),
	}
}

// AddSource adds a new source to the mountpoint, the mountpoint will
// be responsible for directing the source to the correct output until
// the source is removed or disconnects.
func (m *Mount) AddSource(s *Source) {
	m.sources.Add(s)
}

// SetMetadata sets the metadata of the source bound to the given
// SourceID. Setting an empty string means deleting the current
// metadata.
func (m *Mount) SetMetadata(id SourceID, metadata string) {
	return
}
