# Simple Share 📁

Simple Share adalah aplikasi sederhana untuk berbagi file antar device dalam satu jaringan lokal (LAN). Cocok untuk transfer file cepat dari HP ke laptop atau antar komputer tanpa perlu internet.

---

## ✨ Fitur

* Upload file dari browser
* Download file dari device lain
* Berjalan di local network (WiFi yang sama)
* Tanpa konfigurasi ribet
* UI sederhana (HTML + Tailwind)

---

## 🛠️ Tech Stack

* Backend: Golang (HTTP Server)
* Frontend: HTML + TailwindCSS

---

## 📦 Struktur Project

```
.
├── main.go        # entrypoint
├── public/        # Folder untuk file yang di-share
├── templates/     # HTML templates
├── .gitignore
└── README.md
```

---

## 🚀 Cara Menjalankan

1. Clone repo:

```bash
git clone https://github.com/username/simple-share.git
cd simple-share
```

2. Jalankan server:

```bash
go run main.go
```

3. Server akan jalan di:

```
http://localhost:8080
```

---

## 📡 Akses dari Device Lain

Agar bisa diakses dari HP atau device lain:

1. Pastikan semua device terhubung ke WiFi yang sama
2. Cari IP address laptop kamu:

```bash
ip a
```

atau

```bash
ifconfig
```

3. Akses dari device lain:

```
http://<IP_ADDRESS>:8080
```

Contoh:

```
http://192.168.1.10:8080
```

---

## Upload File

* Pilih file dari browser
* Klik tombol upload
* File akan masuk ke folder `public/`

---

## Download File

* Semua file di folder `public/` akan tampil di UI
* Klik file untuk download

---

## ⚠️ Catatan

* Hanya untuk jaringan lokal (tidak aman untuk publik)
* Tidak ada authentication
* Jangan expose ke internet tanpa proteksi tambahan

---

## 💡 Ide Pengembangan

* Progress bar upload
* Drag & drop upload
* Multiple file upload
* Preview file (image/video)

---

## ❤️ Kontribusi
Pull request sangat diterima!

---

## 📄 License

MIT License
