# render-template

A command-line tool that renders text templates with JSON variables.

## Usage

```bash
render-template -t template_file -v variables_file.json
```

## Example

Template file (`example.template`):

```
Here's your list of {{.name}}:

{{range .items -}}
* {{.}}
{{end}}
```

Variables file (`vars.json`):

```json
{
    "name": "animals",
    "items": ["dog", "cat"]
}
```

Running `render-template -t example.template -v vars.json` outputs:

```
Here's your list of animals:

* dog
* cat
```

Uses Go's standard [text/template](https://pkg.go.dev/text/template) package for templating.
