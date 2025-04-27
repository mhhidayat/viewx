# viewx

**viewx** adalah library sederhana untuk membantu proses render view (HTML) di aplikasi web Go (Golang) dengan layout otomatis dan pengaturan title yang fleksibel.

Library ini akan **otomatis membuat folder `views/` dan `layout/`**, lengkap dengan file `index.html`, `navbar.html`, dan `home.html` saat `Init` dipanggil, sehingga kamu langsung bisa fokus ke pengembangan fitur.

---

## Fitur

- Otomatis setup folder `views` dan `layout`
- Support dynamic `Title`
- Rendering layout dengan partials seperti `navbar` dan `footer`
- Minimalis dan mudah dipakai

---

## Instalasi

```bash
go get github.com/mhhidayat/viewx