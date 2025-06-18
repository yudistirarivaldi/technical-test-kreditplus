 Flow Aplikasi 
1. Registrasi Konsumen
User pertama kali harus register agar memiliki akun dan limit pinjaman awal untuk tiap tenor.
Endpoint: POST **/api/auth/register**

Request:
```
{
  "nik": "1234567890123456",
  "full_name": "Budi Santoso",
  "legal_name": "Budi S.",
  "birth_place": "Jakarta",
  "birth_date": "1995-05-01",
  "salary": 5000000,
  "password": "rahasia123"
}
```
```
Setelah register:
Konsumen otomatis dibuatkan limit awal untuk tenor 1, 2, 3, dan 6 bulan.
Contoh:
tenor 1 bulan: 100.000  
tenor 2 bulan: 200.000  
tenor 3 bulan: 500.000  
tenor 6 bulan: 700.000  
```

2. Login Konsumen
User login untuk mendapatkan token JWT yang digunakan untuk semua endpoint selanjutnya.
Endpoint: POST /api/auth/login

Request:
```
{
  "nik": "1234567890123456",
  "password": "rahasia123"
}
```
Response:
```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```

3. Lihat Profil Konsumen
Endpoint: GET /api/consumers

4. Update Data Konsumen (Optional)
Endpoint: PUT /api/consumers/update

Request:
```
{
  "full_name": "Budi Santoso",
  "legal_name": "Budi S. Update",
  "birth_place": "Jakarta",
  "birth_date": "1995-05-01",
  "salary": 7000000
}
```

6. Ajukan Transaksi
User dapat mengajukan transaksi dengan memilih tenor dan nominal cicilan per bulan.
Sistem akan mengecek apakah limit yang tersedia cukup berdasarkan tenor tersebut.
Endpoint: POST /transactions

Request:
```
{
  "tenor": 3,
  "amount": 300000,
  "note": "Cicilan barang elektronik",
  "source_channel": "mobile",
  "down_payment": 100000
}
```

Validasi sistem:
```
Akan mengecek apakah amount <= limit sisa untuk tenor 3.
Jika valid, maka:
Limit dikurangi sebesar amount.
Transaksi dicatat.
```

6. Lihat Riwayat Transaksi
Endpoint: GET /transactions/history
```
[
  {
    "id": 1,
    "tenor": 3,
    "amount": 300000,
    "note": "Cicilan barang elektronik",
    "created_at": "2025-06-18T10:00:00Z"
  }
]
```

