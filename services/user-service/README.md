# user-service
This project for testing skyshi for position backend engineer

Pada test ini saya menggunankan pendekatan heksagonal arsitektur, berfokus pada pemisahan core logika dari ketergantungan eksternal. Core harus bersih, hanya terdiri dari pustaka standar dan kode pembungkus yang dibangun dalam repositori ini.

# ERD
[this is link desain table](https://drive.google.com/file/d/1SrFASmTv5EnpJK1f0Pto9Ed3IUJkpNNT/view?usp=sharing)


## Quick start

Local development:  

```sh

# 1. pastikan ketersediaan dependency seperti database dll.
$ export MONEYMAGNET_DB_DSN='postgres://user_service:user_service@localhost:5431/user_db?sslmode=disable'
$ make db/migrations/up

# 2. Copy .env.example menjadi .env kemudian masukan value dibawah ini
SERVER_PORT=8080
DEBUG_PORT=4000
JWT_SECRET_KEY=secret
DB_DSN="postgres://user_service:user_service@localhost:5431/user_db?sslmode=disable"
DB_MAX_CONN=100
DB_MIN_CONN=10

# 3. menjalankan aplikasi dengan makefile (lihat file Makefile)
$ make run/api

# command tersebut akan mengeksekusi
$ go run ./app/api-user

# 4. untuk melihat dokumentasi api buka url dibawah ini dibrowser
$ http://localhost:8080/swagger/index.html

```  