package models

type Blog struct {
	id        int
	title     string
	content   string
	category  string
	tags      []string
	createdAt string
	updatedAt string
}
