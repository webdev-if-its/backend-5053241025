# backend-nrp

Repo tugas mata kuliah **Pengembangan Backend Dasar**, dibuat dari template [`webdev-if-its/backend-template`](https://github.com/webdev-if-its/backend-template). Ganti judul di atas jadi nama repo kalian sendiri (`backend-nrp`, contoh: `backend-5025201012`).

## Aturan Umum

- Tugas tiap pertemuan disimpan di folder `pertemuan-XX/` pada repo ini.
- Commit message wajib menyebut level yang dicapai: `pertemuan-XX: level N selesai`.
- Deadline push: sebelum pertemuan berikutnya dimulai.
- Semua level dicek otomatis lewat `go test` — baca `pertemuan-XX/SOAL.md` tiap minggu untuk detail levelnya.

## Mengambil Pertemuan Baru Tiap Minggu

Repo ini **tidak otomatis sinkron** dengan template dosen. Begitu ada pertemuan baru, jalankan (ganti `pertemuan-02` sesuai minggu berjalan):

```bash
git fetch https://github.com/webdev-if-its/backend-template.git main
git checkout FETCH_HEAD -- pertemuan-02
```

Perintah ini **aman dijalankan kapan pun** — tidak akan menimpa folder pertemuan lain yang sudah kalian kerjakan, karena hanya mengambil folder yang disebutkan. Setelah itu, commit folder barunya seperti biasa.

Kalau dosen memperbaiki sesuatu di pertemuan yang sudah dirilis (mis. ada bug di test), biasanya cukup ambil ulang file yang diperbaiki saja, bukan seluruh folder — akan diumumkan file mana yang berubah.

---

Bagian di bawah ini **isi bertahap** sesuai level yang sedang kalian kerjakan (lihat `pertemuan-01/SOAL.md`) — heading-nya dicek otomatis, jangan diganti namanya.

## Identitas
- Nama: Kadek Angga Wistara
- NRP: 5053241025
- Kelas: M

## Commit vs Push
Ketika kita melakukan commit, file yang kita commit hanya masih berada di staging stage. artinya "ready untuk di push" tetapi belum di push. jadi file tersebut belum ada di repo, masih berada di local. Sedangkann jika file di push maka 

## Reproducibility
Jika kode yang ada di GO versi terbaru menggunakan syntax atau fitur yang hanya ada di GO versi terbaru tersebut. Maka jika menjalankan kode tersebut dengan versi GO yang lama, kemungkinan akan terjadi syntax error. Tetapi jika kode di versi terbaru tidak ada menggunakan hal baru juga, maka hal tersebut tidak akan menjadi sebuah masalah

## Catatan Merge Conflict
Yang bentrok ada di line 30, karena ada 2 versi berbeda di line yang sama, git jadi bingung mana versi yang benar. Jadi perlu bantuan manusia untuk memutuskan mana versi yang benar

## Kenapa .gitignore Penting
Karena jika didalam 1 tim ada beberapa IDE berbeda, file config dari 1 IDE kemungkinan akan overwrite file config dari IDE berbeda. Yang dimana file config tersebut bisa saja ada hardcoded windows file paths yang tidak ada di computer rekan 1 tim yang menggunakan  IDE berbeda itu

## Refleksi
Saya bingung tentang merge conflict, kenapa  terjadi merge conflict dan bagaimana cara mengatasinya? setelah saya  mencari  tahu  dan melakukannya  sendiri, saya akhirnya tau kenapa dan bagaimana.
