package main

var Template = `<html>
<head>
<title>Hello, world!</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@1/css/pico.min.css">
</head>
<body>
<main class="container">
    <h1>Hello, world!</h1>
    <h6>Here's some detail about my environment...</h6>
    <table role="grid">
    {{ range .Output }}
    <tr>
        <td>{{ . }}</td>
    </tr>
    {{ end }}
    </table>
</main>
</body>
</html>
`
