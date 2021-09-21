package common

import "errors"

// ErrNotFound is returned if storage couldn't find any document.
var ErrNotFound = errors.New("store: document not found")
