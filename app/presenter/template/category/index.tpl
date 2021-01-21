{{ define "body" }}
<div class="container" style="max-width:640px;">
    <h1>Categories</h1>
    <ul>
        {{ range $i, $v := .Categories }}
        <li><a href="/posts/categories/{{ $v.Name }}"><span>Category Name</span>:{{ $v.Name }}</a></li>
        {{ end }}
    </ul>
</div>
{{ template "pagination" .Pagination }}
{{ end }}