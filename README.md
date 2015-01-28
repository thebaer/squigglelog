~log
====

Use **squigglelog** to easily create a simple log for your tilde. [Create new posts](https://github.com/thebaer/squigglelog/tree/master/entries#squigglelog-entries) in the `entries/` folder, then run this:

```bash
go build squigglelog.go
./squigglelog -template mysquigglelog
```

This will use any template in `templates/` defined with _mysquigglelog_ (see below) to generate your full squigglelog page.

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
