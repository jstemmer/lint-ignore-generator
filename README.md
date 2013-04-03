# lint-ignore-generator

lint-ignore-generator converts the [Android lint tool][] report xml file to a
lint configuration file ignoring each error and file that occurs in the report.

## Installation

Make sure you have [Go][] installed.

Install or update using the `go get` command:

	go get -u github.com/jstemmer/lint-ignore-generator

## Usage

	lint-ignore-generator lint-report.xml lint-config.xml

lint-ignore-generator takes a lint report file in xml format as its input and
generates a lint configuration xml file.

## License

See the LICENSE file.

[Android lint tool]: http://tools.android.com/tips/lint
[Go]: http://golang.org
