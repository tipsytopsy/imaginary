package main

import "github.com/tipsytopsy/bimg"

const Version = "0.1.27"

type Versions struct {
	ImaginaryVersion string `json:"imaginary"`
	BimgVersion      string `json:"bimg"`
	VipsVersion      string `json:"libvips"`
}

var CurrentVersions = Versions{Version, bimg.Version, bimg.VipsVersion}
