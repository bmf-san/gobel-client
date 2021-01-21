{{ define "body" }}
<h1>Error</h1>
<p>Error Code:{{ .Code }}</p>
<p>Error Message:{{ .Message }}</p>
{{ end }}