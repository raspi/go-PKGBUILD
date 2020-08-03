package main

import (
	"fmt"
	"github.com/raspi/go-PKGBUILD"
	"os"
)

var exampleJSON = `
{
  "_meta": {
    "ver": "v1.0.0"
  },
  "maintainer": "John Doe",
  "maintainer_email": "jd@example.org",
  "name": [
    "exampleapp"
  ],
  "version": "v1.0.0",
  "release": 1,
  "release_time": "1970-01-01T02:00:00+02:00",
  "short_description": "my example application",
  "licenses": [
    "Apache 2.0"
  ],
  "url": "https://github.com/examplerepo/exampleapp",
  "changelog_file": "",
  "groups": null,
  "dependencies": {},
  "optional_packages": {},
  "provides": null,
  "options": [
    "!strip",
    "docs",
    "libtool",
    "staticlibs",
    "emptydirs",
    "!zipman",
    "!ccache",
    "!distcc",
    "!buildflags",
    "makeflags",
    "!debug"
  ],
  "install": "$pkgname.install",
  "files": {
  },
  "commands": {
    "prepare": [
      "echo foo \u003e\u003e main.c"
    ],
    "build": [
      "make"
    ],
    "test": [
      "make test"
    ],
    "install": [
      "cd \"$srcdir\"",
      "install -Dm755 \"bin/$pkgname\" -t \"$pkgdir/usr/bin\""
    ]
  }
}
`

func main() {
	n, err := PKGBUILD.FromJson([]byte(exampleJSON))
	if err != nil {
		panic(err)
	}

	n.Files = map[string][]PKGBUILD.Source{
		"x86_64": {
			PKGBUILD.Source{
				URL:   `https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-x86_64.tar.gz`,
				Alias: "",
				Checksums: map[string]string{
					"sha256": `123456789abcdef`,
				},
			},
		},
	}

	errs := n.Validate()
	if len(errs) > 0 {
		for _, err := range errs {
			_, _ = fmt.Fprintf(os.Stderr, `%v`, err)
		}

		os.Exit(1)
	}

	fmt.Println(n)
}
