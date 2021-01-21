{{ define "body" }}
<section>
    <div class="container" style="max-width: 640px;">
        <div class="row">
            <div class="column">
                <a href="/posts/categories/{{ .Post.Category.Name }}">{{ .Post.Category.Name }}</a>
                {{ range .Post.Tags }}
                <a class="tag" href="/posts/tags/{{ .Name }}">{{ .Name }}</a>
                {{ end }}
                <h1>{{ .Post.Title }}</h1>
                <div>
                    <span class="article-date" style="padding-right: 1rem;">Created at:{{ .Post.CreatedAt }}</span>
                    <span class="article-date" style="padding-right: 1rem;">Updated at:{{ .Post.UpdatedAt }}</span>
                </div>
                {{ .Post.HTMLBody | unescape }}
            </div>
        </div>
    </div>
</section>
{{ end }}