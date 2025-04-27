package templatex

var IndexHTML = `{{ define "index.html" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .Title }}</title>
</head>
<body>
    {{ template "navbar.html" . }}
    <main>
        {{ template "content.html" . }}
    </main>
    {{ template "footer.html" . }}
</body>
</html>
{{ end }}`

var NavbarHTML = `{{ define "navbar.html" }}
<nav style="background: #333; color: #fff; padding: 10px;">
    <a href="/" style="color: white; text-decoration: none;">Home</a>
</nav>
{{ end }}`

var ContentHTML = `{{ define "content.html" }}
<section style="padding: 20px;">
    <h1>{{ .Title }}</h1>
    <p>Welcome to the content section!</p>
</section>
{{ end }}`

var FooterHTML = `{{ define "footer.html" }}
<footer style="background: #333; color: #fff; text-align: center; padding: 10px; position: fixed; bottom: 0; width: 100%;">
    &copy; 2025 My Website
</footer>
{{ end }}`
