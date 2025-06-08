package main

import (
	"errors"
	"io"

	"golang.org/x/net/html"
)

func Lint(in io.Reader) error {
	node, err := html.Parse(in)
	if err != nil {
		return err
	}

	title := false
	doctype := false
	allImagesContainAltText := true
	for c := range node.Descendants() {
		if !title && c.Data == "title" {
			title = true
		}
		if allImagesContainAltText && c.Data == "img" {
			containsAltText := false
			for _, attr := range c.Attr {
				if attr.Key == "alt" {
					containsAltText = true
				}
			}
			allImagesContainAltText = containsAltText
		}
		if !doctype && c.Type == html.DoctypeNode {
			doctype = true
		}
	}

	if !title {
		return errors.New("Missing title.")
	}
	if !doctype {
		return errors.New("Missing doctype.")
	}
	if !allImagesContainAltText {
		return errors.New("Not all images contain alt text.")
	}
	return nil
}
