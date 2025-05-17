# viewx

**viewx** is a lightweight Go (Golang) library designed to simplify HTML view rendering in web applications. It provides automatic layout management and flexible title configuration out of the box.

When you call `Init`, this library will **automatically create a default folder structure** under views/, including layout files like `footer.html`, `navbar.html`, and `index.html`, as well as a starter page at `page/home.html`. This allows you to focus directly on building your application’s features without worrying about the initial view setup.

---

## Features

- Automatic creation of `views/` `layout/` and `page/` folders
- Support for dynamic `Title` injection
- Layout rendering with partials such as `navbar` and `footer`
- Minimalistic and easy-to-use
- Zero external dependencies

---

## Installation

```bash
go get github.com/mhhidayat/viewx
```

---

## Example Usage

###### Viewx rendering with layout

```go
viewx.Init("MyApp")

router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    viewx.Render(w, "page/home", nil)
    // or you can use: viewx.Render(w, "page.home", nil)
})
```
###### Rendering a single HTML file

```go
router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    viewx.RenderHTML(w, "views/index", nil)
    // or you can use: viewx.RenderHTML(w, "views.index", nil)
})
```

## Folder Structure (After viewx.Init())

```arduino
project/
└── views/
    ├── layout/
    │   ├── footer.html
    │   ├── index.html
    │   └── navbar.html
    └── page/
        └── home.html
```