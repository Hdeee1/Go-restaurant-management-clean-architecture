- Aplikasi Apa Ini? Ini adalah Backend (REST API) untuk sistem kasir atau Point of Sales (POS) khusus restoran.

- Tujuan Utama: Mendigitalisasi dan melacak seluruh operasional restoran agar tidak ada pesanan yang terlewat, mempermudah manajemen meja, dan mencatat pemasukan secara akurat.

- Target Pengguna (Aktor): Aplikasi ini ditujukan untuk pegawai restoran, bukan pelanggan.
    Admin/Manajer: Mengelola data pegawai (User) dan daftar makanan/minuman (Food/Menu).
    Pelayan (Waiter): Memilih Table (meja), membuat Order (pesanan), dan mencatat Note (catatan khusus, misal: "pedas", "jangan pakai bawang").
    Kasir: Memeriksa pesanan yang sudah selesai dan menerbitkan Invoice (tagihan/pembayaran).
    
-Alur Logika Dasar: Pelayan menetapkan pesanan ke sebuah Meja $\rightarrow$ Pesanan berisi detail Makanan $\rightarrow$ Dapur memasak berdasarkan Catatan $\rightarrow$ Pelanggan selesai makan $\rightarrow$ Kasir memproses Pembayaran (Invoice).