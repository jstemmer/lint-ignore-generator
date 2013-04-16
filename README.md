# lint-ignore-generator

lint-ignore-generator converts an [Android lint tool][] report xml file to a
lint configuration file. It creates an ignore entry for each error and file
path that occurs in the report.

## Installation

Make sure you have [Go][] installed.

Install or update using the `go get` command:

	go get -u github.com/jstemmer/lint-ignore-generator

## Usage

	lint-ignore-generator [options]

	Options:
		-f Filter by path
		-i Input: Lint XML report
		-o Output: Lint configuration file

### Example

Add all `actionbarsherlock` occurrences from `lint-report.xml` to `lint.xml`.

	lint-ignore-generator -i lint-report.xml -o lint.xml -f actionbarsherlock

## License

See the LICENSE file.

[Android lint tool]: http://tools.android.com/tips/lint
[Go]: http://golang.org
