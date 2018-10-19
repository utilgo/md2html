package md2html

import "github.com/russross/blackfriday"

func Md2Html(md []byte) (res []byte) {
	mathjax := `<script src='https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.0/MathJax.js?config=TeX-MML-AM_CHTML'>MathJax.Hub.Config({tex2jax:{inlineMath:[['$','$'],['\\(','\\)']],processEscapes:true}});</script>`
	res = []byte(
		`<meta name="viewport" content="width=device-width, initial-scale=1.0"><style>` + docStyle + `</style>` + mathjax + string(blackfriday.MarkdownCommon(md)))
	return
}

const docStyle = `
*{margin:0;padding:0}
html{background-color:#111;color:#ccc;font-family:'Courier New',Courier,monospace}
pre{padding:10px;background-color:#222;color:#999;font-size:12px;line-height:18px}
p{padding:10px;color:#f90;font-size:14px}
h1{padding:10px;padding-bottom:0;color:#697faf;letter-spacing:1px;font-weight:500;font-size:16px}

`

const style = `
html {
    background-color: #f4f5f5;
    color: #24292e;
    font-size: 62.5%;
}

body {
    margin: 0;
    letter-spacing: .0625rem;
    font-size: 1.4rem;
    line-height: 1.5rem;
}

h1,h2 {
    margin-top: 10px;
    letter-spacing: .125rem;
    font-weight: 400;
}

hr {
    margin: 10px 0;
    border: none;
    border-top: 1px solid #dde;
}

a {
    color: #06f;
    text-decoration: none;
    cursor: pointer;
}

a:hover {
    text-decoration: underline;
}

ol,ul {
    padding: 0 20px;
    font-family: "微软雅黑";
}

li {
    padding: 2px 0;
}

p {
    line-height: 1.375rem;
}

table {
    width: 100%;
    border-collapse: collapse;
    font-size: .8125rem;
    font-family: "微软雅黑";
}

table td,table th {
    padding: 6px 2%;
    border: 1px solid #dde;
    color: #58595f;
}

blockquote {
    margin: 0;
    padding: 1px 30px;
    border-left: 3px solid #dfe2e5;
    color: #55595e;
}

pre {
    padding: 15px;
    border-radius: 3px;
    background-color: #f6f8fa;
    color: #333;
    font-size: .75rem;
}

blockquote,pre,table {
    margin-bottom: 30px;
}

.markdown img {
    width: 100%;
}

.markdown {
    position: relative;
    overflow: hidden;
    margin: 0 auto;
    padding: 50px 40px;
    width: 900px;
    border-right: 1px solid #eee;
    border-left: 1px solid #eee;
    background-color: #fff;
    text-overflow: ellipsis;
}

@media only screen and (max-width:900px) {
    body {
        padding: 0 5px;
        background-color: #fff;
    }

    .markdown {
        padding: 30px 0;
        width: 100%;
        border: none;
    }
}
`
