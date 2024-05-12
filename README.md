# Persewaan Alat Lingkungan

## About Project
Project ini bertujuan untuk memberikan akses mudah dan terjangkau terhadap alat-alat lingkungan yang diperlukan untuk berbagai kegiatan di komunitas, acara, atau masyarakat. Sebagai kegiatan mendukung praktik-praktik yang dapat memberikan dampak positif di lingkungan dengan memberi pengguna akses ke teknologi hijau dan alat-alat lingkungan.

## Features
Fitur-fitur yang terdapat di project yang dibuat:

### User
* Pengguna dapat mendaftar menggunakan akun.
* Pengguna dapat menyewa alat dengan jangka waktu yang ditentukan.
* Pengguna dapat mencari alat berdasarkan kategori yang dibutuhkan.
* Pengguna dapat melihat berapa lama sewa alat akan berakhir.

### Admin
* Admin dapat menambah mengedit dan menghapus barang.
* Admin menerima sewa pengguna.
* Admin dapat melihat data pengguna yang menyewa dan sudah selesai menyewa.

## Tech Stacks
App Framework: echo\
ORM Library : Gorm\
DB : Postgres di rds\
Deployment : AWS EC2\
Code Structure : Clean Architecture\
Authentication : JWT\
Other Tools / Libraries : Cloudinary

## API Documentation
https://drive.google.com/file/d/16HioeWTaFrix7AeK92B2LFn4g1Av627z/view?usp=sharing

## ERD
https://drive.google.com/file/d/1WjzuNsHjPQ-i9OAavOWqNS3mt28qmAlY/view?usp=sharing

## Setup
Pertama User login dan sign up jika belum memiliki akun. User dapat melihat alat yang disediakan sebelum menyewa. User dapat menyewa alat dengan menentukan jumlah alat, terlebih dahulu. Setelah user yakin akan menyewa alat yang telah dipilih, user melakukan konfirmasi dengan mengirimkan data seperti bukti pembayaran. User tinggal menunggu konfirmasi dari admin apakah diterima atau tidak. Jika iya maka akan ada status yang awalnya memiliki keterangan pending berubah menjadi confirm. Untuk admin, admin dapat membuat category alat, membuat daftar alat baru dan mengkonfirmasi data user yang ingin menyewa. Jika user sudah mengembalikan alat yang telah disewa admin dapat mengubah status sewa tadi menjadi returned.