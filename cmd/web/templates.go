package main

import "github.com/Bakhram74/snippetbox.git/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
