{{define "body"}}
<div class="container" style="max-width: 640px;">
    <h1>Posts</h1> 
    {{ range $i, $v := .Posts }}
    <article>
        <h1><a href="/posts/{{ $v.Title }}">{{ $v.Title }}</a></h1>
        <a href="/posts/categories/{{ $v.Category.Name }}">{{ $v.Category.Name }}</a>
        {{ range $v.Tags }}
        <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
        {{ end }}
        <div>
            <span class="article-date" style="padding-right: 1rem;">Created at:{{ $v.CreatedAt }}</span>
            <span class="article-date" style="padding-right: 1rem;">Updated at:{{ $v.UpdatedAt }}</span>
        </div>
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