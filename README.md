~log
====

Use **squigglelog** to easily create a simple log for your tilde.

1. Install squigglelog with `./install.sh`
2. [Create new posts](https://github.com/thebaer/squigglelog/tree/master/entries#squigglelog-entries) in the `entries/` folder, or use the `squiggle.sh` utility.
3. Generate your ~log:

```bash
./squigglelog -t mysquigglelog -o mylog
```

This will use any template in `templates/` defined with _mysquigglelog_ (see below) to generate your full squigglelog page in ~/public_html/ called mylog.html.

#### templates

Your template should look like this.

```html
{{define "mysquigglelog"}}
<html>
	<head>
		<title>My ~log!</title>
	</head>
	<body>
		<h1>~log</h1>
		<p>Welcome to my ~log.</p>
		{{template "log" .}}
	</body>
</html>
{{end}}
```

There are two important lines in this file

`{{define "mysquigglelog"}}`

This is what you will reference in the `-t` flag passed to the generator. The template file name does not matter.

`{{template "log" .}}`

This must be included verbatim in your template. It will include the `templates/log.html` template, which holds all the posts.
