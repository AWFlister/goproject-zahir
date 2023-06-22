# Simple Golang CRUD

Server sederhana yang mendukung CRUD, dengan sort, pagination, dan filter. 

Dibuat dengan [gorm](gorm.io), [chi router](https://github.com/go-chi/chi), dan sqlite3

Install dengan `git clone` dan `go get ./...`<br>
Jalankan dengan `go run .`<br>
tambahkan args `seed` untuk mengisi database, atau `reset` untuk reset database.


---
## Objects

### Contact

```
{
    "uuid":"string",    // UUID
    "name":"string",    // Nama
    "gender":"string",  // Gender
    "phone":"string",   // Nomor Hape
    "email":"string",   // Email
    "created_at":"string",
    "updated_at":"string"
}
```
---
## Endpoints

### GET /

Sample endpoint. Return sebuah json yang isinya:

```
{"message":"Hello Zahir!"}
```
- Query Params: None
- Data Params: None
- Headers: <br>Content-Type: application/json
- Success: 200

---
### GET /contacts

Return `limit` buah objek Contact, mengikuti filtering dan sort order yang diberikan dalam URL params

Content:

```
{
    "data": [
        <Contact object>,
        <Contact object>,
        ...
    ],
    "page": integer,
    "page_total": integer
}
```
- Query Params:
  - `page`: halaman yang diminta. Integer. Wajib
  - `limit`: banyaknya data yang diminta. Integer. Wajib
  - `sort_by`: cara pengurutan data. String berformat `"<column_name> <ASC|DESC>"`
  - `filter`: pemilihan data berdasarkan keyword pada column tertentu. String berformat `"<column_name>:<keyword>"`. Lebih dari satu filter dapat disambung dengan tanda koma (`,`)
- Data Params: None
- Headers: <br>Content-Type: application/json
- Success: 200
---
### GET /contacts/:id

Return satu buah objek Contact dengan UUID sesuai dengan `id`, jika ada. Jika tidak maka field `uuid` akan kosong

Content:

```
<Contact object>
```
- Query Params: None
- Data Params: None
- Headers: <br>Content-Type: application/json
- Success: 200
---
### POST /contacts

Buat sebuah Contact objek baru dan simpan dalam database. Return Contact objek yang baru dibuat. Field yang tidak tersedia saat request akan dianggap kosong.

Content:

```
<Contact object>
```
- Query Params: None
- Data Params: 
  - `name`: Nama kontak. String
  - `gender`: Gender. String
  - `phone`: Nomor hape. String
  - `email`: Email. String
- Headers: <br>Content-Type: application/json
- Success: 200
---

### PUT /contacts/:id

Update beberapa field dari sebuah objek Contact. Return objek Contact setelah di-update. Field yang tidak di-update tidak perlu dimasukkan.

Content:

```
<Contact object>
```
- Query Params: None
- Data Params: 
  - `name`: Nama kontak. String
  - `gender`: Gender. String
  - `phone`: Nomor hape. String
  - `email`: Email. String
- Headers: <br>Content-Type: application/json
- Success: 200
---

### DELETE /contacts/:id

Menghapus sebuah objek Contact dengan UUID `id` dari database. Penghapusan dilakukan dengan fitur soft delete dari gorm, sehingga tidak sebenarnya terhapus dari database.

Content: None
- Query Params: None
- Data Params: None
- Headers: <br>Content-Type: application/json
- Success: 200