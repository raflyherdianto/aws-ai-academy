# AWS AI Academy — Developer Card & Guess Who Game

Boilerplate monorepo untuk website perkenalan interaktif peserta AWS AI Academy yang dibuat menggunakan **Vue 3 + Tailwind CSS v4 (Frontend)** dan **Golang Gin + SQLite (Backend)**.

## Fitur Utama
1. **Developer Card Generator**: Membuat kartu karakter digital dengan desain *Glassmorphism* modern dan tema warna AWS/Dicoding.
2. **Visual Capture & Upload**: Kartu dirender langsung di browser lalu diekspor menjadi gambar PNG menggunakan `html2canvas` dan disimpan di server.
3. **WhatsApp Link Sharing with Preview**: Menyebarkan kartu perkenalan ke WhatsApp Group dengan preview Open Graph (OG) Image dinamis.
4. **Guess Who Game**: Permainan kuis tebak-tebakan ice-breaking interaktif berbasis profil peserta yang terdaftar.

---

## Struktur Folder Monorepo
* `/backend`: Aplikasi API server menggunakan Golang Gin.
* `/frontend`: Aplikasi Vue.js 3 menggunakan Node 24 + Vite.

---

## Cara Menjalankan di Lokal (Development)

### Persyaratan:
1. **Go (Golang)** terinstal di sistem.
2. **Node.js 24+** & **npm** terinstal di sistem.

### Langkah-langkah:
1. **Instalasi Dependensi**:
   Buka terminal di root monorepo, lalu jalankan:
   ```bash
   npm run install-all
   ```
   *(Atau jalankan `npm install` di root, `npm install` di folder `frontend`, dan `go mod tidy` di folder `backend`)*.

2. **Jalankan Frontend & Backend Bersamaan**:
   Cukup jalankan satu perintah berikut di root monorepo:
   ```bash
   npm run dev
   ```
   * **Frontend** akan berjalan di: `http://localhost:5173`
   * **Backend API** akan berjalan di: `http://localhost:8080`

---

## Alur Penggunaan
1. Masuk ke halaman utama `http://localhost:5173`.
2. Masukkan email Anda yang terdaftar di `mentee.json` untuk verifikasi.
3. Lengkapi form profil (Nama Panggilan, Pekerjaan, Kota/Provinsi, Coder Class, Bio, dan roll stats Coder).
4. Klik **Simpan & Dapatkan Kartu Karakter**. Sistem akan merender kartu Anda, mengambil gambarnya, mengunggah ke backend, dan mengarahkan ke halaman share.
5. Salin link share tersebut dan bagikan ke grup WhatsApp Anda untuk melihat preview Open Graph-nya secara otomatis!
