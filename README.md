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
