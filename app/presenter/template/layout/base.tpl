{{ define "base" }}
<!DOCTYPE html>
<html lang="ja">

<head>
    <title></title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="" />
    <meta name="og:title" content="" />
    <meta name="og:description" content="" />
    <meta property="og:url" content="" />
    <meta property="og:type" content="">
    <meta property="og:image" content="" />
    <meta property="og:site_name" content="" />
    <meta property="og:locale" content="ja_JP" />
    <meta name="twitter:card" content="" />
    <meta name="twitter:site" content="" />
    <link rel="stylesheet" href="https://unpkg.com/sea.css/dist/sea.min.css">
</head>

<body>
    <header>
        <div>
            <nav class="nav sp-hidden">
                <div class="container">
                    <div class="row">
                        <div class="col">
                            <a class="nav-link-logo"><img src="./image/logo.png" alt="logo"></a>
                        </div>
                        <div class="col">
                            <div class="nav-group">
                                <a class="nav-link" href="/">Home</a>
                                <a class="nav-link" href="/posts">Posts</a>
                                <a class="nav-link" href="/categories">Categories</a>
                                <a class="nav-link" href="/tags">Tags</a>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
            <nav class="nav pc-hidden">
                <a class="nav-link-logo"><img src="./image/logo.png" alt="logo" style="width: 15rem;"></a>
                <div id="nav-sp-drawer">
                    <input id="nav-sp-input" type="checkbox" class="sp-hidden">
                    <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
                    <label class="sp-hidden" id="nav-sp-close" for="nav-sp-input"><span></span></label>
                    <div id="nav-sp-content">
                        <h1>Navigation</h1>
                        <a class="nav-link" href="/">Home</a>
                        <a class="nav-link" href="/posts">Posts</a>
                        <a class="nav-link" href="/categories">Categories</a>
                        <a class="nav-link" href="/tags">Tags</a>
                    </div>
                </div>
            </nav>
    </header>

    <div class="headline" style="background-image: url('./image/sea.JPG'); background-size:100%;padding: 10rem 0;">
        <div class="container">
            <div class="row">
                <div class="col">
                    <h1 style="text-align: center; color: #fff;">TOP</h1>
                </div>
            </div>
        </div>
    </div>

    {{ template "body" . }}

    <footer class="sticky-footer" style="margin-top: 3rem;">
        <div class="headline">
            <div class="row">
                <div class="col">
                    <div>
                        <p><a href="/sitemap">Sitemap</a></p>
                        <p><a href="/feed">Feed</a></p>
                    </div>
                </div>
            </div>
        </div>
    </footer>
</body>

</html>
{{ end }}


