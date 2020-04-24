## render_html

A command line utility that renders Go `template/html` templates from the command line.

### Example

With the template, `index.tpl.html`, containing the following content:

```
<html>
<body>
    Version: {{ .build.version }}
    Time: {{ .build.time }}
</body>
</html>
```

Executing `render_html` and passing in the parameters via `-v` flags.

```
render_html -v build.version=1.2 -v build.time=$(date -u) index.tpl.html index.html
```

This will render the template with the parameters passed on the command line and will write
the results into `index.html`, which will contains:

```
<html>
<body>
    Version:  1.2
    Time: Fri Apr 24 21:28:43 UTC 2020
</body>
</html>
```

### Installing

```
go get github.com/kellegous/render_html
```

### Templating Language

The details of the templating language are described in the Go documentation: https://golang.org/pkg/html/template/

Note that when multiple source template files are given, the first template becomes the template that will be render and all other templates are available to be included by their basename. The details of this behavior is also described in the Go documentation as part of the [ParseFiles](https://golang.org/pkg/html/template/#ParseFiles) method.