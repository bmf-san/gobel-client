{{define "body"}}
<div class="container" style="max-width:640px;">
    <h1>Posts - {{ .CategoryName }}</h1>
    {{ range $i, $v := .Posts }}
    <article>
        <h1><a href="/posts{{ $v.Title }}">{{ $v.Title }}</a></h1>
        <a href="/posts/categories/{{ $v.Category.Name }}">{{ $v.Category.Name }}</a>
        <span class="article-date" style="padding-right: 1rem;">{{ $v.CreatedAt }}</span>
        <span class="article-date" style="padding-right: 1rem;">{{ $v.UpdatedAt }}</span>
        {{ range $v.Tags }}
        <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
        {{ end }}
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Itaque repellat provident quo, in sunt
            possimus
            consequuntur similique cupiditate iusto voluptas obcaecati delectus dolorum facilis numquam
            voluptatibus
            cumque libero minus deleniti.</p>
    </article>
    {{ end }}

    {{ template "pagination" .Pagination }}
</div>
{{ end }}