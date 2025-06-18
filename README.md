***Flow Instalasi Aplikasi***
1. Clone Repository

   ```
   git clone https://github.com/yudistirarivaldi/technical-test-kreditplus.git
   cd technical-test-kreditplus
   ```

2. Siapkan File .env
   Buat file .env atau gunakan yang sudah disediakan:
   
   DB_HOST=mysql untuk running host via docker

   ```
   SERVER_PORT=8080
   DB_HOST=mysql 
   DB_PORT=3306
   DB_USER=dev
   DB_PASS=dev123
   DB_NAME=xyz_multifinance
   JWT_SECRET=supersecretkey123
   ```

3. Jalankan via Docker Compose
   ```
   docker compose up --build
   ```
   Ini akan Build image aplikasi Jalankan container MySQL + aplikasi Otomatis baca .env dan mengatur konfigurasi

4. Stop Container
   ```
   docker compose down
   ```


***Flow Aplikasi***
1. Registrasi Konsumen
   
   User pertama kali harus register agar memiliki akun dan limit pinjaman awal untuk tiap tenor.

   Endpoint: POST /api/auth/register

   Request:
   ```
   {
       "nik": "12391930123",
       "full_name": "Budi Agung",
       "legal_name": "Budi Agung",
       "birth_place": "Jakarta",
       "birth_date": "1990-01-01",
       "salary": 10000000,
       "password": "rahasia123",
       "ktp_photo": "budi-ktp.jpg",
       "selfie_photo": "budi-selfie.jpg"
   }
   ```
   Response:
   ```
   {
       "responseCode": "00",
       "message": "Registration successful"
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

3. Login Konsumen
   
   User login untuk mendapatkan token JWT yang digunakan untuk semua endpoint selanjutnya.

   Endpoint: POST /api/auth/login
   
   Request: 
   ```
   {
       "nik": "12334",
       "password": "rahasia123"
   }
   ```
   
   Response:
   ```
   {
       "responseCode": "00",
       "message": "Login successful",
       "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb25zdW1lcl9pZCI6NiwiZXhwIjoxNzUwMzMyMTE0fQ.T2kVE3uPL6TK6jQk7TuLSRnyFSg_pnnDBn9JEWDD31U"
   }
   ```

5. Lihat Profil Konsumen
   
   Endpoint: GET /api/consumers

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```
   
   Response:
   ```
   {
       "responseCode": "00",
       "message": "Consumer profile retrieved successfully",
       "data": {
           "ID": 6,
           "NIK": "12334",
           "Password": "",
           "FullName": "Budi Agung",
           "LegalName": "Budi Agung",
           "BirthPlace": "Jakarta",
           "BirthDate": "1990-01-01T00:00:00Z",
           "Salary": 10000000,
           "KTPPhoto": "budi-ktp.jpg",
           "SelfiePhoto": "budi-selfie.jpg",
           "CreatedAt": "2025-06-18T11:21:29Z",
           "UpdatedAt": "2025-06-18T11:21:29Z"
       }
   }
   ```

7. Update Data Konsumen (Optional)
   
   Endpoint: PUT /api/consumers/update

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "full_name": "Budi Gunawan",
       "legal_name": "Budi Gunawarman",
       "birth_place": "Jakarta",
       "birth_date": "1990-01-01",
       "salary": 12000000,
       "ktp_photo": "budi-ktp-updated.jpg",
       "selfie_photo": "budi-selfie-updated.jpg"
   }
   ```

   Response:
   ```
   {
       "responseCode": "00",
       "message": "Consumer updated successfully"
   }
   ```
9. Ajukan Transaksi
   
   User dapat mengajukan transaksi dengan memilih tenor dan nominal cicilan per bulan.
   Sistem akan mengecek apakah limit yang tersedia cukup berdasarkan tenor tersebut.

   Endpoint: POST /transactions

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   Request:
   ```
   {
       "contract_number": "TXN-20250618001",
       "otr": 1500000,
       "admin_fee": 50000,
       "installment": 600000,
       "interest": 100000,
       "asset_name": "Kulkas LG",
       "source_channel": "ecommerce",
       "tenor": 6,
       "down_payment": 200000
   }
   ```
   Response:
   ```
   {
       "responseCode": "00",
       "message": "Transaction successfully inserted"
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

   Headers:
   ```Authorization: Bearer <JWT_TOKEN>```

   ```
   {
       "responseCode": "00",
       "message": "Success",
       "data": [
           {
               "ID": 2,
               "ConsumerID": 5,
               "ContractNumber": "TXN-20250618001",
               "OTR": 1500000,
               "AdminFee": 50000,
               "Installment": 600000,
               "Interest": 100000,
               "AssetName": "Kulkas LG",
               "SourceChannel": "ecommerce",
               "Tenor": 6,
               "DownPayment": 200000,
               "CreatedAt": "2025-06-18T15:37:45.964Z",
               "UpdatedAt": "2025-06-18T15:37:45.964Z"
           }
       ]
   }
   ```

