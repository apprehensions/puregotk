// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

// An opaque data structure representing a set of bookmarks.
type BookmarkFile struct {
}

// Error codes returned by bookmark file parsing.
type BookmarkFileError int

const (

	// URI was ill-formed
	GBookmarkFileErrorInvalidUriValue BookmarkFileError = 0
	// a requested field was not found
	GBookmarkFileErrorInvalidValueValue BookmarkFileError = 1
	// a requested application did
	//     not register a bookmark
	GBookmarkFileErrorAppNotRegisteredValue BookmarkFileError = 2
	// a requested URI was not found
	GBookmarkFileErrorUriNotFoundValue BookmarkFileError = 3
	// document was ill formed
	GBookmarkFileErrorReadValue BookmarkFileError = 4
	// the text being parsed was
	//     in an unknown encoding
	GBookmarkFileErrorUnknownEncodingValue BookmarkFileError = 5
	// an error occurred while writing
	GBookmarkFileErrorWriteValue BookmarkFileError = 6
	// requested file was not found
	GBookmarkFileErrorFileNotFoundValue BookmarkFileError = 7
)
