{{ define "body" }}
<div class="container" style="max-width:640px;">
    <h1>Tags</h1>
    <ul>
        {{ range $i, $v := .Tags }}
        <li><a href="/posts/tags/{{ $v.Name }}"><span>Tag Name</span>:{{ $v.Name }}</a></li>
        {{ end }}
    </ul>
</div>
{{ template "pagination" .Pagination }}
{{ end }}