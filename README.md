# Contree

Contree is a Go library to handle configuration powered by tree data structure.

## Installation

```bash
go get -u github.com/jakofys/contree
```

## Use

<!-- ### Import from file

To import configuration value:

```go
conf := contree.FromFile("my/conf/file.yml", contree.YAML)
```

Or from reader as:

```go
conf := contree.From(io.Reader, contree.YAML)
``` -->

### Get value

```go
value := conf.Get("path.to.my.value") // string value
```

### Set value

Just set value using path data:

```go
conf.Set("path.to.my.value", "value") // string value
```

<!-- ### Configuration file format

Configuration file supports:

- YAML
- JSON
- DOTENV
- TOML
- XML

Contree is awesome, it understand path in YAML as `path.to.my.value` and can convert it to:

```yaml
path:
    to:
        my:
            value: "myvalue"
``` -->

<!-- ### List value

You can set a list value or get it juts make:

```go
conf.Set("path.to.my.list", []Value{})
list := conf.Get("path.to.my.list")
```

Or you can append value in conf directly using:

```go
err := conf.Append("path.to.my.list", "valueToAppend")
```

> ⚠ Append can return error if the pass is not a list -->

### Interpolation

You can have same behavior as `fmt.Sprintf` except syntax to interpolate making:

```go
conf.Set("path.to.value", "Casanova")
conf.Sprintf("Simply replace data from %path.to.value%") // Output: Simply replace data Casanova
```

<!-- ### Exportation

You can export structure to any format extension file using:

```go
err := conf.Export(io.Writer, contree.YAML) 
``` -->

## Reference

## Author

- [Jacques COFIS](github.com/jakofys) Software Engineer Junior