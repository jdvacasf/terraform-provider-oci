package main

// Reads BareMetal entity
type ResourceReader interface {
	Get() error
	SetData()
}

// ResourceSync synchronizes a ResourceData instance and a BareMetal entity.
type ResourceSync interface {
	ResourceReader
	Id() string
	State() string
	Create() error
	Update() error
	Delete() error
}
