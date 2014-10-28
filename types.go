// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package semver

import (
	"errors"
	"net/url"
)

// Repo defines a single repository and target version.
type Repo struct {
	Version

	// The root URL of the repository (excluding subpackages). For example a
	// package imported at:
	//
	//  https://example.com/pkg/subpkg
	//
	// Would have a root repository URL without "/subpkg":
	//
	//  https://github.com/example/pkg
	//
	// As that is where *the repository* lives; not the Go package itself.
	*url.URL

	// The subpath of the repository. It is joined with the repository root URL
	// in order to build the final package path. SubPath == "subpkg" in the
	// above example:
	//
	//  Repo.URL.String + repo.SubPath == "https://example.com/pkg/subpkg"
	//
	SubPath string
}

// Status represents a single status code returned by a Handler's attempt to
// Handle any given request.
type Status int

const (
	// The request was not handled.
	Unhandled Status = iota

	// The request was handled.
	Handled

	// The request was not handled, but was for the package page (e.g. when
	// viewing in a web browser).
	PkgPage
)

var (
	ErrNotPackageURL = errors.New("not a valid package URL")
)

// HTTPError represents a HTTP error generated by a Handler's Relate function.
type HTTPError struct {
	error
	Status int
}