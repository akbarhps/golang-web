# Golang Web

Sumber Tutorial:
[Udemy](https://www.udemy.com/course/pemrograman-go-lang-pemula-sampai-mahir/learn/lecture/26413070#overview) |
[Slide](https://docs.google.com/presentation/d/1h_8nk-Ani4SykMq5lhgubzdtAtZfzPce7FHUuAKlwmE/edit#slide=id.p)


## Pengenalan Web
---


### Kenapa Web?

- Saat ini web digunakan oleh jutaan, bahkan mungkin milyaran orang setiap hari
- Dengan web, kita bisa melakukan belajar online, mendengarkan musik online, nonton video online, belanja online, sampai memesan makanan secara online
- Namun perlu diperhatikan, Web bukanlah Internet


### Internet

- Internet adalah mekanisme komunikasi antar komputer
- Awal internet ada, untuk komunikasi antar komputer, kita membutuhkan jaringan kabel telepon
- Namun sekarang, semenjak berjamurnya jaringan wifi dan sejenisnya, komunikasi antar komputer menjadi lebih cepat dan mudah

![Internet](https://user-images.githubusercontent.com/69947442/140602364-45a8a624-f607-48cd-ad6f-c9a88c3e3ce8.png)


### Web

- Web merupakan kumpulan informasi yang tersedia dalam sebuah komputer yang terkoneksi secara terus menerus melalui internet
- Web bisa berisi informasi dalam bentuk apapun, seperti teks, gambar, audio, video dan lain-lain
- Web berjalan di aplikasi yang bernama Web Server, yaitu aplikasi yang digunakan untuk menyimpan dan menyampaikan isi informasi Web


### Diagram Web

![Diagram Web](https://user-images.githubusercontent.com/69947442/140602382-4bf3031b-616b-457f-982c-26849642ec84.png)


### Web Host

- Pemilik Web, biasanya tidak menjalankan aplikasi Web Server di komputer pribadi nya
- Biasanya mereka akan menyewa komputer di tempat penyedia data center (kumpulan komputer) yang terjamin keandalan dan kecepatan koneksi internetnya
- Pihak penyedia komputer untuk Web Server biasa disebut Web Host


### Domain

- Saat komputer Web terhubung ke internet, biasanya dia memiliki alamat
- Alamat ini bernama ip address, formatnya misal nya 172.217.194.94
- Karena alamat ip address sangat menyulitkan untuk diingat
- Untung saja ada yang namanya nama domain
- Nama domain adalah alamat yang bisa digunakan sebagai alias ke ip address
- Misal seperti google.co.id, blibli.com, dan lain-lain
- Dengan nama domain, sebagai manusia kita akan mudah mengingat dibandingkan ip address
- Namun, saat kita menggunakan nama domain, sebenarnya komputer tetap akan mengakses web menggunakan alamat ip address


### Web Browser

- Jika Web Server adalah aplikasi yang digunakan untuk menyimpan informasi Web
- Web Browser adalah aplikasi yang digunakan untuk mengakses Web melalui internet
- Kita bisa saja mengakses Web secara langsung tanpa bantuan Web Browser, namun Web Server hanya akan memberikan informasi bahasa mesin seperti HTML, JavaScript, CSS, Gambar, Video dan lain-lain
- Dengan menggunakan Web Browser, semua bahasa mesin tersebut bisa ditampilkan secara visual sehingga kita bisa menyerap informasinya dengan lebih mudah


## Client dan Server
---

- Web adalah aplikasi berbasis Client dan Server, sekarang pertanyaannya, apa itu Client dan Server?
- Sederhananya client server merupakan konsep arsitektur aplikasi yang menghubungkan dua pihak, sistem client dan sistem server
- Sistem client dan sistem server yang saling berkomunikasi melalui jaringan komputer, internet, atau juga bisa di komputer yang sama


### Diagram Client dan Server

![Diagram Client dan Server](https://user-images.githubusercontent.com/69947442/140602420-18a2308d-4aa6-47e3-ade0-5e36eb72d3cd.png)


### Tugas Client dan Server

- Aplikasi Client bertugas mengirim request ke Server, dan menerima response dari Server
- Sedangkan, aplikasi Server bertugas menerima request dari Client, memproses data, dan mengembalikan hasil proses data ke Client


### Keuntungan Client dan Server

- Perubahan aplikasi bisa dilakukan dengan mudah di server, tanpa harus membuat perubahan di client, apalagi jika client nya di lokasi yang sulit dijangkau
- Bisa digunakan oleh banyak client pada saat yang bersamaan, walaupun server tidak banyak
- Bisa diakses dari mana saja, asal terhubung satu jaringan dengan server


### Contoh Client dan Server

- Web adalah salah satu contoh arsitektur client server
- Aplikasi yang bertugas sebagai Client adalah Web Browser (Chrome, Firefox, Opera, Edge dan lain-lain)
- Aplikasi yang bertugas sebagai Server adalah Web Server, dimana di dalam web server terdapat kode program Web kita


## Golang Web
---

- Go-Lang saat ini populer dijadikan salah satu pilihan bahasa pemrograman untuk membuat Web, terutama Web API (Backend)
- Selain itu, di Go-Lang juga sudah disediakan package untuk membuat Web, bahkan di sertakan pula package untuk implementasi unit testing untuk Web
- Hal ini menjadikan pembuatan Web menggunakan Go-Lang lebih mudah, karena tidak butuh menggunakan library atau framework


### Diagram Cara Kerja Golang Web

![Diagram Cara Kerja Golang Web](https://user-images.githubusercontent.com/69947442/140602467-4bbdc258-1f5b-43ef-94d4-8ac58b62c079.png)


### Cara Kerja Golang Web

1. Web Browser akan melakukan HTTP Request ke Web Server
2. Golang menerima HTTP Request, lalu mengeksekusi request tersebut sesuai dengan yang diminta.
3. Setelah melakukan eksekusi request, Golang akan mengembalikan data dan di render sesuai dengan kebutuhannya, misal HTML, CSS, JavaScript dan lain-lain
4. Golang akan mengembalikan content hasil render tersebut tersebut sebagai HTTP Response ke Web Browser
5. Web Browser menerima content dari Web Server, lalu me-render content tersebut sesuai dengan tipe content nya 


### Package net/http

- Pada beberapa bahasa pemrograman lain, seperti Java misalnya, untuk membuat web biasanya dibutuhkan tambahan library atau framework
- Sedangkan di Go-Lang sudah disediakan package untuk membuat web bernama package net/http
- Sehingga untuk membuat web menggunakan Go-Lang, kita tidak butuh lagi library tambahan, kita bisa menggunakan package yang sudah tersedia
- Walaupun memang saat kita akan membuat web dalam skala besar, direkomendasikan menggunakan framework karena beberapa hal sudah dipermudah oleh web framework
- Namun pada course ini, kita akan fokus menggunakan package net/http untuk membuat web nya, karena semua framework web Go-Lang akan menggunakan net/http sebagai basis dasar framework nya


## Server
---

- Server adalah struct yang terdapat di package net/http yang digunakan sebagai representasi Web Server di Go-Lang
- Untuk membuat web, kita wajib membuat Server
- Saat membuat data Server, ada beberapa hal yang perlu kita tentukan, seperti host dan juga port tempat dimana Web kita berjalan
- Setelah membuat Server, kita bisa menjalankan Server tersebut menggunakan function `ListenAndServe()`


### Kode: Server

```go
func TestListenAndServe(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


## Handler
---

- Server hanya bertugas sebagai Web Server, sedangkan untuk menerima HTTP Request yang masuk ke Server, kita butuh yang namanya Handler
- Handler di Go-Lang di representasikan dalam interface, dimana dalam kontraknya terdapat sebuah function bernama `ServeHTTP()` yang digunakan sebagai function yang akan di eksekusi ketika menerima HTTP Request


### HandlerFunc

- Salah satu implementasi dari interface Handler adalah HandlerFunc
- Kita bisa menggunakan HandlerFunc untuk membuat function handler HTTP


### Kode: Hello World Web

```go
func TestServerWithHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


## ServeMux
---

- Saat membuat web, kita biasanya ingin membuat banyak sekali endpoint URL
- HandlerFunc sayangnya tidak mendukung itu
- Alternative implementasi dari Handler adalah ServeMux
- ServeMux adalah implementasi Handler yang bisa mendukung multiple endpoint


### Kode: ServeMux

```go
func TestServerWithServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello kamu yuang bukan wed")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


### URL Pattern

- URL Pattern dalam ServeMux sederhana, kita tinggal menambahkan string yang ingin kita gunakan sebagai  endpoint, tanpa perlu memasukkan domain web kita
- Jika URL Pattern dalam ServeMux kita tambahkan di akhirnya dengan garis miring, artinya semua url tersebut akan menerima path dengan awalan tersebut, misal /images/ artinya akan dieksekusi jika endpoint nya /images/, /images/contoh, /images/contoh/lagi
- Namun jika terdapat URL Pattern yang lebih panjang, maka akan diprioritaskan yang lebih panjang, misal jika terdapat URL /images/ dan /images/thumbnails/, maka jika mengakses /images/thumbnails/ akan mengakses /images/thumbnails/, bukan /images


### Kode: ServeMux URL Pattern

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello helo bandung")
})
mux.HandleFunc("/lol", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ini lol")
})
```


## Request
---

- Request adalah struct yang merepresentasikan HTTP Request yang dikirim oleh Web Browser
- Semua informasi request yang dikirim bisa kita dapatkan di Request
- Seperti, URL, http method, http header, http body, dan lain-lain


### Kode: Request

```go
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Method)
    fmt.Println(r.RequestURI)
})
```


## HTTP Test
---

- Go-Lang sudah menyediakan package khusus untuk membuat unit test terhadap fitur Web yang kita buat
- Semuanya ada di dalam package net/http/httptest https://golang.org/pkg/net/http/httptest/ 
- Dengan menggunakan package ini, kita bisa melakukan testing handler web di Go-Lang tanpa harus menjalankan aplikasi web nya
- Kita bisa langsung fokus terhadap handler function yang ingin kita test


### `httptest.NewRequest()`

- `NewRequest(method, url, body)` merupakan function yang digunakan untuk membuat http.Request
- Kita bisa menentukan method, url dan body yang akan kita kirim sebagai simulasi unit test
- Selain itu, kita juga bisa menambahkan informasi tambahan lainnya pada request yang ingin kita kirim, seperti header, cookie, dan lain-lain


### `httptest.NewRecorder()`

- `httptest.NewRecorder()` merupakan function yang digunakan untuk membuat ResponseRecorder
- ResponseRecorder merupakan struct bantuan untuk merekam HTTP response dari hasil testing yang kita lakukan


### Kode: HTTP Test

```go
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func TestHTTP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "https://localhost:8080/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)
}
```


### Kode: Mengecek Hasil Test

```go
result := rec.Result()
body, _ := io.ReadAll(result.Body)
bodyString := string(body)

fmt.Println(bodyString)
```


## Query Parameter
---

- Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
- Query parameter biasanya digunakan untuk mengirim data dari client ke server
- Query parameter ditempatkan pada URL
- Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya


### url.URL

- Dalam parameter Request, terdapat attribute URL yang berisi data `url.URL`
- Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan method `Query()` yang akan mengembalikan map


### Kode: Query Parameter

```go
func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
```


### Kode: Test Query Parameter

```go
func TestQueryParam(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Stipen", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```


### Multiple Query Parameter

- Dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
- Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter lainnya
- Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter berikutnya


### Kode: Multiple Query Parameter

```go
func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	fName := r.URL.Query().Get("first_name")
	lName := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "Hello, %s %s!", fName, lName)
}
```


### Kode: Test Multiple Query Parameter

```go
func TestMultipleQueryParameter(t *testing.T) {
	req := httptest.NewRequest("GET", "/?first_name=Stipen&last_name=Debandits", nil)
	rec := httptest.NewRecorder()

	MultipleQueryParameter(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```


### Multiple Value Query Parameter

- Sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam map[string][]string
- Artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
- Caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda, misal :
- name=Eko&name=Kurniawan


### Kode: Multiple Value Query Parameter

```go
func MultipleValueQueryParameter(w http.ResponseWriter, r *http.Request) {
	favorite := r.URL.Query()["favorite"]
	fmt.Fprintf(w, "Your favorite food is %s\n", strings.Join(favorite, " "))
}
```


### Kode: Test Multiple Value Query Parameter

```go
func TestMultipleValueQueryParameter(t *testing.T) {
	req := httptest.NewRequest("GET", "/?favorite=Pizza&favorite=Burger", nil)
	rec := httptest.NewRecorder()

	MultipleValueQueryParameter(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```


## Header
---

- Selain Query Parameter, dalam HTTP, ada juga yang bernama Header
- Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
- Jadi dalam Header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa menambahkan informasi header
- Saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh browser, seperti informasi browser, jenis tipe content yang dikirim dan diterima oleh browser, dan lain-lain


### Request Header

- Untuk menangkap request header yang dikirim oleh client, kita bisa mengambilnya di `Request.Header`
- Header mirip seperti Query Parameter, isinya adalah `map[string][]string`
- Berbeda dengan Query Parameter yang case sensitive, secara spesifikasi, Header key tidaklah case sensitive


### Kode: Request Header

```go
func RequestHeader(rec http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprintf(rec, "Content-Type is %s", contentType)
}
```


### Kode: Test Request Header

```go
func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	resp := rec.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```


### Response Header

- Sedangkan jika kita ingin menambahkan header pada response, kita bisa menggunakan function `ResponseWriter.Header()`


### Kode: Response Header

```go
func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Response Header has been changed!")
}
```


### Kode: Test Response Header

```go
func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	resp := rec.Result()
	fmt.Println(resp.Header.Get("Content-Type"))

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```


## Form Post
---

- Saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan method GET atau POST
- Jika menggunakan method GET, maka hasilnya semua data di form akan menjadi query parameter
- Sedangkan jika menggunakan POST, maka semua data di form akan dikirim via body HTTP request
- Di Go-Lang, untuk mengambil data Form Post sangatlah mudah


### `Request.PostForm`

- Semua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute Request.PostForm
- Namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method `Request.ParseForm()` terlebih dahulu, method ini digunakan untuk melakukan parsing data body apakah bisa di parsing menjadi form data atau tidak, jika tidak bisa di parsing, maka akan menyebabkan error


### Kode: Form Post

```go
func PostForm(w http.ResponseWriter, r *http.Request) {
	// parse and get form values
	// err := r.ParseForm()
	// if err != nil {
	// 	panic(err)
	// }
	// fName := r.PostForm.Get("firstName")
	// lName := r.PostForm.Get("lastName")

	// automatically parse form data and get its values
	fName := r.PostFormValue("firstName")
	lName := r.PostFormValue("lastName")
	fmt.Fprintf(w, "%s %s", fName, lName)
}
```


### Kode: Test Form Post

```go
func TestPostForm(t *testing.T) {
	reqBody := strings.NewReader("firstName=Joe&lastName=Mama")
	req := httptest.NewRequest("POST", "/", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	PostForm(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


## Response Code
---

- Dalam HTTP, terdapat yang namanya response code
- Response code merupakan representasi kode response
- Dari response code ini kita bisa melihat apakah sebuah request yang kita kirim itu sukses diproses oleh server atau gagal
- Ada banyak sekali response code yang bisa kita gunakan saat membuat web
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Status 


### Mengubah Response Code

- Secara default, jika kita tidak menyebutkan response code, maka response code nya adalah 200 OK
- Jika kita ingin mengubahnya, kita bisa menggunakan function `ResponseWriter.WriteHeader(int)`
- Semua data status code juga sudah disediakan di Go-Lang, jadi kita ingin, kita bisa gunakan variable yang sudah disediakan : https://github.com/golang/go/blob/master/src/net/http/status.go 


### Kode: Response Code

```go
func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Name Can't Be Blank!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}
```


### Kode: Test Response Code

```go
func TestResponseCode(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=Joe", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	if rec.Code != http.StatusOK {
		panic("StatusCode: " + strconv.Itoa(res.StatusCode) + " Status: " + res.Status)
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}

func TestResponseCodeInvalidInput(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	res := rec.Result()
	if rec.Code != http.StatusBadRequest {
		panic("Response Code Should Be 400 Bad Request")
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


## Cookie
---


### Stateless

- HTTP merupakan stateless antara client dan server, artinya server tidak akan menyimpan data apapun untuk mengingat setiap request dari client
- Hal ini bertujuan agar mudah melakukan scalability di sisi server
- Lantas bagaimana caranya agar server bisa mengingat sebuah client? Misal ketika kita sudah login di website, server otomatis harus tahu jika client tersebut sudah login, sehingga request selanjutnya, tidak perlu diminta untuk login lagi
- Untuk melakukan hal ini, kita bisa memanfaatkan Cookie


### Cookie

- Cookie adalah fitur di HTTP dimana server bisa memberi response cookie (key-value) dan client akan menyimpan cookie tersebut di web browser
- Request selanjutnya, client akan selalu membawa cookie tersebut secara otomatis
- Dan server secara otomatis akan selalu menerima data cookie yang dibawa oleh client setiap kalo client mengirimkan request


### Membuat Cookie

- Cookie merupakan data yang dibuat di server dan sengaja agar disimpan di web browser
- Untuk membuat cookie di server, kita bisa menggunakan function `http.SetCookie()`


### Kode: Membuat Cookie

```go
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "X-USERNAME",
		Value: r.URL.Query().Get("name"),
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cookie is set ", cookie)
}
```


### Kode: Mengambil Cookie

```go
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-USERNAME")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintf(w, "Cookie not found!")
	} else {
		fmt.Fprintf(w, "Welcome Back %s", cookie.Value)
	}
}
```


### Kode: Test Cookie Manual

```go
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		SetCookie(w, r)
	})
	mux.HandleFunc("/get-cookie", func(w http.ResponseWriter, r *http.Request) {
		GetCookie(w, r)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


### Kode: Test Set Cookie `httptest`

```go
func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=John", nil)
	rec := httptest.NewRecorder()

	SetCookie(rec, req)

	cookie := rec.Result().Cookies()
	fmt.Println(cookie)
}
```


### Kode: Test Get Cookie `httptest`

```go
func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	cookie := http.Cookie{
		Name:  "X-USERNAME",
		Value: "John",
	}

	req.AddCookie(&cookie)
	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestGetCookieInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	GetCookie(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusBadRequest {
		panic("Status Code Should Be 404 Bad Request!")
	}

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


## FileServer
---

- Go-Lang memiliki sebuah fitur yang bernama FileServer
- Dengan ini, kita bisa membuat Handler di Go-Lang web yang digunakan sebagai static file server
- Dengan menggunakan FileServer, kita tidak perlu manual me-load file lagi
- FileServer adalah Handler, jadi bisa kita tambahkan ke dalam `http.Server` atau `http.ServeMux`


### Kode: FileServer

```go
func TestFileServer(t *testing.T) {
	dir := http.Dir("resources")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/static/", fileServer)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


### 404 Not Found

- Jika kita coba jalankan, saat kita membuka misal /static/index.js, maka akan dapat error 404 Not Found
- Kenapa ini terjadi?
- Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan url nya, jadi jika kita membuat /static/index.js, maka FileServer akan mencari ke file /resources/static/index.js
- Hal ini menyebabkan 404 Not Found karena memang file nya tidak bisa ditemukan
- Oleh karena itu, kita bisa menggunakan function `http.StripPrefix()` untuk menghapus prefix di url


### Kode: FileServer dengan StripPrefix

```go
prefix := http.StripPrefix("/static", fileServer)
mux.Handle("/static/", prefix)
```


### Golang Embed

- Di Go-Lang 1.16 terdapat fitur baru yang bernama Go-Lang embed
- Dalam Go-Lang embed kita bisa embed file ke dalam binary distribution file, hal ini mempermudah sehingga kita tidak perlu meng-copy static file lagi
- Go-Lang Embed juga memiliki fitur yang bernama embed.FS, fitur ini bisa diintegrasikan dengan FileServer


### Kode: FileServer Golang Embed

```go
//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	fileServer := http.FileServer(http.FS(resources))

	mux := http.NewServeMux()
	prefix := http.StripPrefix("/static", fileServer)
	mux.Handle("/static/", prefix)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


### 404 Not Found

- Jika kita coba jalankan, dan coba buka /static/index.js, maka kita akan mendapatkan error 404 Not Found
- Kenapa ini terjadi? Hal ini karena di Go-Lang embed, nama folder ikut menjadi nama resource nya, misal resources/index.js, jadi untuk mengaksesnya kita perlu gunakan URL /static/resources/index.js
- Jika kita ingin langsung mengakses file index.js tanpa menggunakan resources, kita bisa menggunakan function fs.Sub() untuk mendapatkan sub directory


### Kode: FileServer Go Embed

```go
dir, _ := fs.Sub(resources, "resources")
fileServer := http.FileServer(http.FS(dir))
```


## ServeFile
---

- Kadang ada kasus misal kita hanya ingin menggunakan static file sesuai dengan yang kita inginkan
- Hal ini bisa dilakukan menggunakan function http.ServeFile()
- Dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http response


### Kode: ServeFile

```go
func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "./resources/index.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}
```


### Kode: Test ServeFile

```go
func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


### ServeFile Golang Embed

- Parameter function http.ServeFile hanya berisi string file name, sehingga tidak bisa menggunakan Go-Lang Embed
- Namun bukan berarti kita tidak bisa menggunakan Go-Lang embed, karena jika untuk melakukan load file, kita hanya butuh menggunakan package fmt dan ResponseWriter saja


### Kode: ServeFile Golang Embed

```go
//go:embed resources/index.html
var resourceIndex string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(w, resourceIndex)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}
```


### Kode: Test ServeFile Golang Embed

```go
func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


## Template
---


### Web Dinamis

- Sampai saat ini kita hanya membahas tentang membuat response menggunakan String dan juga static file
- Pada kenyataannya, saat kita membuat web, kita pasti akan membuat halaman yang dinamis, bisa berubah-ubah sesuai dengan data yang diakses oleh user
- Di Go-Lang terdapat fitur HTML Template, yaitu fitur template yang bisa kita gunakan untuk membuat HTML yang dinamis


### HTML Template

- Fitur HTML template terdapat di package html/template
- Sebelum menggunakan HTML template, kita perlu terlebih dahulu membuat template nya
- Template bisa berubah file atau string
- Bagian dinamis pada HTML Template, adalah bagian yang menggunakan tanda `{{  }}`


### Membuat Template

- Saat membuat template dengan string, kira perlu memberi tahu nama template nya
- Dan untuk membuat text template, cukup buat text html, dan untuk konten yang dinamis, kita bisa gunakan tanda `{{.}}`, contoh :
- `<html><body>{{.}}</body></html>`


### Kode: HTML Template String

```go
func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	// manual error check
	// t, err := template.New("Simple").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	//}

	t := template.Must(template.New("Simple").Parse(templateText))
	t.ExecuteTemplate(w, "Simple", "Hello World Using Template")
}
```


### Kode: Test HTML Template String

```go
func TestSimpleHTML(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	SimpleHTML(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Template Dari File

- Selain membuat template dari string, kita juga bisa membuat template langsung dari file
- Hal ini mempermudah kita, karena bisa langsung membuat file html
- Saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama template nya, misal jika kita punya file simple.html, maka nama template nya adalah simple.html


### Kode: Template Dari File

```gohtml
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.}}</title>
</head>

<body>
    <h1>{{.}}</h1>
</body>

</html>
```

```go
func TemplateFromFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/simple.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Hello World Using Template")
}
```


### Kode: Test Template Dari File

```go
func TestTemplateFromFile(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateFromFile(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Template Directory

- Kadang biasanya kita jarang sekali menyebutkan file template satu persatu
- Alangkah baiknya untuk template kita simpan di satu directory
- Go-Lang template mendukung proses load template dari directory
- Hal ini memudahkan kita, sehingga tidak perlu menyebutkan nama file nya satu per satu


### Kode: Template Directory

```go
func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Hello World Using Template")
}
```


### Kode: Test Template Directory

```go
func TestTemplateDirectory(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateDirectory(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Template Golang Embed

- Sejak Go-Lang 1.16, karena sudah ada Go-Lang Embed, jadi direkomendasikan menggunakan Go-Lang embed untuk menyimpan data template
- Menggunakan Go-Lang embed menjadi kita tidak perlu ikut meng-copy template file lagi, karena sudah otomatis di embed di dalam distribution file


### Kode: Template Golang Embed

```go
//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "simple.gohtml", "Hello World Using Template")
}
```


### Kode: Test Template Golang Embed

```go
func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateEmbed(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


## Template Data
---

- Saat kita membuat template, kadang kita ingin menambahkan banyak data dinamis
- Hal ini bisa kita lakukan dengan cara menggunakan data struct atau map
- Namun perlu dilakukan perubahan di dalam text template nya, kita perlu memberi tahu Field atau Key mana yang akan kita gunakan untuk mengisi data dinamis di template
- Kita bisa menyebutkan dengan cara seperti ini `{{.NamaField}}`


### Kode: Template Data

```gohtml
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>

<body>
    <h1>Hello {{.Name}}</h1>
</body>

</html>
```

```go
type Page struct {
	Title string
	Name  string
}

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))

	data := map[string]interface{}{
		"Title": "Template Dari Map",
		"Name":  "Joe Mama",
	}

	t.ExecuteTemplate(w, "name.gohtml", data)
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/name.gohtml"))

	data := Page{
		Title: "Template Dari Map",
		Name:  "Joe Mama",
	}

	t.ExecuteTemplate(w, "name.gohtml", data)
}
```


### Kode: Test Template Data

```go
func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateDataMap(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateDataStruct(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


## Template Action
---

- Go-Lang template mendukung perintah action, seperti percabangan, perulangan dan lain-lain


### If Else

- `{{if .Value}} T1 {{end}}`, jika Value tidak kosong, maka T1 akan dieksekusi, jika kosong, tidak ada yang dieksekusi
- `{{if .Value}} T1 {{else}} T2 {{end}}`, jika value tidak kosong, maka T1 akan dieksekusi, jika kosong, T2 yang akan dieksekusi
- `{{if .Value1}} T1 {{else if .Value2}} T2 {{else}} T3 {{end}}`, jika Value1 tidak kosong, maka T1 akan dieksekusi, jika Value2 tidak kosong, maka T2 akan dieksekusi, jika tidak semuanya, maka T3 akan dieksekusi


### Kode: Template If Statement

```gohtml
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Template If</title>
</head>

<body>
    {{if .Name}}
    <h1>Hello {{.Name}}</h1>
    {{else}}
    <h1>Hello, World!</h1>
    {{end}}
</body>

</html>
```

```go
func TemplateIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	data := map[string]interface{}{
		"Name": r.URL.Query().Get("name"),
	}

	t.ExecuteTemplate(w, "if.gohtml", data)
}
```


### Kode: Test Template If Statement

```go
func TestTemplateIf(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Joe%20Mama", nil)
	rec := httptest.NewRecorder()

	TemplateIf(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Operator Pembadingan

Go-Lang template juga mendukung operator perbandingan, ini cocok ketika butuh melakukan perbandingan number di if statement, berikut adalah operator nya :
- eq	artinya arg1 == arg2
- ne	artinya arg1 != arg2
- lt	artinya arg1 < arg2
- le	artinya arg1 <= arg2
- gt	artinya arg1 > arg2
- ge	artinya arg1 >= arg2


### Kode: Template Operator Pembadingan

```gohtml
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Template Comparator</title>
</head>

<body>
{{if ge .Score 80}}
    <h1>Good</h1>
{{else if ge .Score 60}}
    <h1>Nice try</h1>
{{else}}
    <h1>Try again</h1>
{{end}}
</body>

</html>
```


### Kenapa Operatornya di Depan?

- Hal ini dikarenakan, sebenarnya operator perbandingan tersebut adalah sebuah function
- Jadi saat kita menggunakan {{eq First Second}}, sebenarnya dia akan memanggil function eq dengan parameter First dan Second : eq(First, Second)


### Kode: Operator Perbandingan

```go
func TemplateOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	score, _ := strconv.Atoi(r.URL.Query().Get("score"))
	data := map[string]interface{}{
		"Score": score,
	}

	t.ExecuteTemplate(w, "comparator.gohtml", data)
}
```


### Kode: Test Operator Perbandingan

```go
func TestTemplateOperator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?score=80", nil)
	rec := httptest.NewRecorder()

	TemplateOperator(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Range

- Range digunakan untuk melakukan iterasi data template
- Tidak ada perulangan biasa seperti menggunakan for di Go-Lang template
- Yang kita bisa lakukan adalah menggunakan range untuk mengiterasi tiap data array, slice, map atau channel
- {{range $index, $element := .Value}} T1 {{end}}, jika value memiliki data, maka T1 akan dieksekusi sebanyak element value, dan kita bisa menggunakan $index untuk mengakses index dan $element untuk mengakses element
- {{range $index, $element := .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value tidak memiliki element apapun, maka T2 yang akan dieksekusi


### Kode: Template Range

```gohtml
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Range</title>
</head>

<body>
    {{range $index, $element := .Items}}
    <h1>{{$index}}: {{$element}}</h1>
    {{else}}
    <h1>No items found</h1>
    {{end}}
</body>

</html>
```


```go
func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	data := map[string]interface{}{
		"Items": []string{"Joe", "Mama", "Thicc"},
	}

	t.ExecuteTemplate(w, "range.gohtml", data)
}
```


### Kode: Test Template Range

```go
func TestTemplateRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	TemplateRange(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### With

- Kadang kita sering membuat nested struct 
- Jika menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue
- Di template terdapat action with, yang bisa digunakan mengubah scope dot menjadi object yang kita mau
- {{with .Value}} T1 {{end}}, jika value tidak kosong, di T1 semua dot akan merefer ke value
- {{with .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value kosong, maka T2 yang akan dieksekusi


### Kode: Template With

```gohtml
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Template With</title>
</head>

<body>
    <h1>Name: {{.Name}}</h1>
    {{with .Address}}
    <h2>Address Street: {{.Street}}</h2>
    <h2>Address City: {{.City}}</h2>
    {{end}}
</body>

</html>
```


```go
func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))

	data := map[string]interface{}{
		"Name": r.URL.Query().Get("name"),
		"Address": map[string]string{
			"Street": "Jalan Surga Gang Neraka",
			"City":   "Akhirat",
		},
	}

	t.ExecuteTemplate(w, "with.gohtml", data)
}
```


### Kode: Test Template With

```go
func TestTemplateWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Joe%20Mama", nil)
	rec := httptest.NewRecorder()

	TemplateWith(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Komentar

- Template juga mendukung komentar
- Komentar secara otomatis akan hilang ketika template text di parsing
- Untuk membuat komentar sangat sederhana, kita bisa gunakan {{/* Contoh Komentar */}}

### Kode: Komentar

```gohtml
{{/* Contoh Komentar */}}
```


## Template Layout
---

- Saat kita membuat halaman website, kadang ada beberapa bagian yang selalu sama, misal header dan footer
- Best practice nya jika terdapat bagian yang selalu sama, disarankan untuk disimpan pada template yang terpisah, agar bisa digunakan di template lain
- Go-Lang template mendukung import dari template lain


### Import Table

Untuk melakukan import, kita bisa menggunakan perintah berikut :
- {{template “nama”}}, artinya kita akan meng-import template “nama” tanpa memberikan data apapun
- {{template “nama” .Value}}, artinya kita akan meng-import template “nama” dengan memberikann data value


### Kode: Template Header

```gohtml
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Template Layout</title>
</head>
<body>
```


### Kode: Template Footer

```gohtml
<footer>Ini adalah footer</footer>
</body>
</html>
```


### Kode: Template Layout

```gohtml
{{template "header.gohtml"}}
<h1>Hello {{.Name}}</h1>
{{template "footer.gohtml"}}
```

```go
func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/layout.gohtml", "./templates/header.gohtml", "./templates/footer.gohtml",
	))

	data := map[string]interface{}{
		"Name": r.URL.Query().Get("name"),
	}

	t.ExecuteTemplate(w, "layout.gohtml", data)
}
```


### Kode: Test Template Layout

```go
func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?name=Joe%20Mama", nil)
	rec := httptest.NewRecorder()

	TemplateLayout(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Template Name

- Saat kita membuat template dari file, secara otomatis nama file nya akan menjadi nama template
- Namun jika kita ingin mengubah nama template nya, kita juga bisa melakukan mengguakan perintah {{define “nama”}} TEMPLATE {{end}}, artinya kita membuat template dengan nama “nama”


### Kode: Template Name

```gohtml
{{define "layout"}}
{{template "header.gohtml"}}
<h1>Hello {{.Name}}</h1>
{{template "footer.gohtml"}}
{{end}}
```


## Template Function
---

- Selain mengakses field, dalam template, function juga bisa diakses
- Cara mengakses function sama seperti mengakses field, namun jika function tersebut memiliki parameter, kita bisa gunakan tambahkan parameter ketika memanggil function di template nya
- {{.FunctionName}}, memanggil field FunctionName atau function FunctionName()
- {{.FunctionName “eko”, “kurniawan”}}, memanggil function FunctionName(“eko”, “kurniawan”)


### Kode: Struct

```go
type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}
```


### Kode: Template Function

```go
func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Joe"}}`))

	data := MyPage{
		Name: "Joe Mama",
	}

	t.ExecuteTemplate(w, "FUNCTION", data)
}
```


### Kode: Test Template Function

```go
func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateFunction(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Global Function

- Go-Lang template memiliki beberapa global function
- Global function adalah function yang bisa digunakan secara langsung, tanpa menggunakan template data
- Berikut adalah beberapa global function di Go-Lang template 
- https://github.com/golang/go/blob/master/src/text/template/funcs.go 


### Kode: Template Function Global

```go
func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Joe Mama",
	})
}
```


### Kode: Test Template Function Global

```go
func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionGlobal(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Menambahkan Function Global

- Kita juga bisa menambah global function
- Untuk menambah global function, kita bisa menggunakan method Funcs pada template
- Perlu diingat, bahwa menambahkan global function harus dilakukan sebelum melakukan parsing template


### Kode: Menambahkan Function Global

```go
func TemplateCreateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Joe Mama",
	})
}
```


### Kode: Test Menambahkan Function Global

```go
func TestTemplateCreateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateCreateFunctionGlobal(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


### Function Pipeline

- Go-Lang template mendukung function pipelines, artinya hasil dari function bisa dikirim ke function berikutnya
- Untuk menggunakan function pipelines, kita bisa menggunakan tanda | , misal :
- {{ sayHello .Name | upper }}, artinya akan memanggil global function sayHello(Name) hasil dari sayHello(Name) akan dikirim ke function upper(hasil)
- Kita bisa menambahkan function pipelines lebih dari satu


### Kode: Function Pipeline

```go
func TemplateFunctionPipeline(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))
	t.ExecuteTemplate(w, "FUNCTION", map[string]interface{}{
		"Name": "Joe Mama",
	})
}
```


### Kode: Test Function Pipeline

```go
func TestTemplateFunctionPipeline(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionPipeline(rec, req)

	res := rec.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
```


## Template Caching
---

- Kode-kode diatas yang sudah kita praktekan sebenarnya tidak efisien
- Hal ini dikarenakan, setiap Handler dipanggil, kita selalu melakukan parsing ulang template nya
- Idealnya template hanya melakukan parsing satu kali diawal ketika aplikasinya berjalan
- Selanjutnya data template akan di caching (disimpan di memory), sehingga kita tidak perlu melakukan parsing lagi
- Hal ini akan membuat web kita semakin cepat


### Kode: Template Caching

```go
//go:embed templates/*.gohtml
var embedTemplates embed.FS

var myTemplates = template.Must(template.ParseFS(embedTemplates, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template Caching")
}
```


### Kode: Test Template Caching

```go
func TestTemplateCaching(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	TemplateCaching(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
```


## XSS (Cross Site Scripting)
---

- XSS adalah salah satu security issue yang biasa terjadi ketika membuat web
- XSS adalah celah keamanan, dimana orang bisa secara sengaja memasukkan parameter yang mengandung JavaScript agar dirender oleh halaman website kita
- Biasanya tujuan dari XSS adalah mencuri cookie browser pengguna yang sedang mengakses website kita
- XSS bisa menyebabkan account pengguna kita diambil alih jika tidak ditangani dengan baik


### Auto Escape

- Berbeda dengan bahasa pemrograman lain seperti PHP, pada Go-Lang template, masalah XSS sudah diatasi secara otomatis
- Go-Lang template memiliki fitur Auto Escape, dimana dia bisa mendeteksi data yang perlu ditampilkan di template, jika mengandung tag-tag html atau script, secara otomatis akan di escape
- Semua function escape bisa diliat disini :
- https://github.com/golang/go/blob/master/src/html/template/escape.go 
- https://golang.org/pkg/html/template/#hdr-Contexts 


###  Kode: File Template

```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
   <h1>{{.Title}}</h1> 
   {{.Body}}
</body>
</html>
```


### Kode: Auto Escape

```go
func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Golang Auto Escape",
		"Body": "<p>Selamat Belajar Golang Web</p>",
	}
	myTemplates.ExecuteTemplate(w, "post.gohtml", data)
}
```


### Kode: Test Auto Escape

```go
func TestTemplateAutoEscape(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	TemplateAutoEscape(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	t.Log(string(body))
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}
```


### Mematikan Auto Escape

- Jika kita mau, auto escape juga bisa kita matikan
- Namun, kita perlu memberi tahu template secara eksplisit ketika kita menambahkan template data
- Kita bisa menggunakan data 
- template.HTML , jika ini adalah data html
- template.CSS, jika ini adalah data css
- template.JS, jika ini adalah data javascript


### Kode: Mematikan Auto Escape

```go
func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Golang Auto Escape",
		// auto escape
		// "Body": "<p>Selamat Belajar Golang Web</p>",

		// disable auto escape
		"Body": template.HTML("<p>Selamat Belajar Golang Web</p>"),
	}
	myTemplates.ExecuteTemplate(w, "post.gohtml", data)
}
```


### Masalah XSS (Cross Site Scripting)

- Saat kita mematikan fitur auto escape, bisa dipastikan masalah XSS akan mengintai kita
- Jadi pastikan kita benar-benar percaya terhadap sumber data yang kita matikan auto escape nya


### Kode: XSS

```go
func TemplateXSS(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Golang XSS",
		"Body":  template.HTML(r.URL.Query().Get("body")),
	}
	myTemplates.ExecuteTemplate(w, "post.gohtml", data)
}
```


### Kode: Test XSS

```go
func TestTemplateXSS(t *testing.T) {
	req := httptest.NewRequest("GET", "/?body=<script>alert(1)</script>", nil)
	rec := httptest.NewRecorder()

	TemplateXSS(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	t.Log(string(body))
}

func TestTemplateXSSServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}
```


## Redirect
---

- Saat kita membuat website, kadang kita butuh melakukan redirect
- Misal setelah selesai login, kita lakukan redirect ke halaman dashboard
- Redirect sendiri sebenarnya sudah standard di HTTP https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections 
- Kita hanya perlu membuat response code 3xx dan menambah header Location
- Namun untungnya di Go-Lang, ada function yang bisa kita gunakan untuk mempermudah ini


### Kode: Redirect

```go
func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You has been redirected, welcome back!")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", http.StatusTemporaryRedirect)
}
```


### Kode: Test Redirect

```go
func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}
```


## Upload File
---

- Saat membuat web, selain menerima input data berupa form dan query param, kadang kita juga menerima input data berupa file dari client
- Go-Lang Web sudah memiliki fitur untuk management upload file 
- Hal ini memudahkan kita ketika butuh membuat web yang menerima input file upload


### MultiPart

- Saat kita ingin menerima upload file, kita perlu melakukan parsing terlebih dahulu menggunakan Request.ParseMultipartForm(size), atau kita bisa langsung ambil data file nya menggunakan Request.FormFile(name), di dalam nya secara otomatis melakukan parsing terlebih dahulu
- Hasilnya merupakan data-data yang terdapat pada package multipart, seperti multipart.File sebagai representasi file nya, dan multipart.FileHeader sebagai informasi file nya


### Kode: Template Upload Form

```gohtml
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload FIle</title>
</head>
<body>
    <h1>Upload File</h1>
    <form action="/upload" method="post" enctype="multipart/form-data">
        <label>Name: <input type="text" name="name"></label><br>
        <label>File: <input type="file" name="file"></label><br>
        <input type="submit" value="Upload">
    </form>
</body>
</html>
```


### Kode: Template Upload Form Success

```gohtml
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload Success</title>
</head>
<body>
    <h1>{{.Name}}</h1>
    <a href="{{.File}}">File</a>
</body>
</html>
```


### Kode: Upload Form

```go
func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form,
	// 32 << 20 means that the maximum memory file size that can be stored in memory is 32 megabytes
	// r.ParseMultipartForm(32 << 20)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	// set uploaded file destination
	fileDest, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// copy uploaded file to destination
	_, err = io.Copy(fileDest, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}
```


### Kode: Test Upload Form Server

```go
func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```


### Kode: Test Upload Form Automatic

```go
//go:embed resources/img.png
var image []byte

func TestUploadForm(t *testing.T) {
	// Create a new multipart form (must be the same as the one used in the html)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// fill form name field
	writer.WriteField("name", "Joe mama")

	// fill form upload file field
	// "file" -> name of the form field
	// "image.png" -> file name to send to the server (this is the name that will be saved on the server)
	file, _ := writer.CreateFormFile("file", "image.png")

	// write embedded file to form upload file field
	file.Write(image)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	rec := httptest.NewRecorder()

	// set request header to multipart form
	req.Header.Set("Content-Type", writer.FormDataContentType())

	Upload(rec, req)

	responseBody, _ := ioutil.ReadAll(rec.Result().Body)
	t.Log(string(responseBody))
}
```


## Download File
---

- Selain upload file, kadang kita ingin membuat halaman website yang digunakan untuk download sesuatu
- Sebenarnya di Go-Lang sudah disediakan menggunakan FileServer dan ServeFile
- Dan jika kita ingin memaksa file di download (tanpa di render oleh browser, kita bisa menggunakan header Content-Disposition)
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition 


### Kode: Download File

```go
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request")
		return
	}

	// change Content-Disposition to force download
	// default Content-Disposition is inline
	// inline means that the file will be displayed in browser
	w.Header().Add("Content-Disposition", "attachment; filename="+fileName+"\"")
	http.ServeFile(w, r, "./resources/"+fileName)
}
```


### Kode: Test Download File

```go
func TestDownloadFileServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DownloadFile)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		t.Error(err)
	}
}
```


## Middleware
---

- Dalam pembuatan web, ada konsep yang bernama middleware atau filter atau interceptor
- Middleware adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler di eksekusi


### Diagram Middleware

![Diagram Middleware](https://user-images.githubusercontent.com/69947442/140635161-2d14998d-34b6-4f51-b7e8-01706cef2185.png)


### Middleware di Golang Web

- Sayangnya, di Go-Lang web tidak ada middleware
- Namun karena struktur handler yang baik menggunakan interface, kita bisa membuat middleware sendiri menggunakan handler


### Kode: Middleware

```go
type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before handler execution")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After handler execution")
}
```


### Kode: Test Middleware

```go
func TestMiddlewareLog(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Handler Executed")
		fmt.Fprint(w, "Hello, World!")
	})

	middleware := &LogMiddleware{
		Handler: mux,
	}

	server := httptest.NewServer(middleware)
	defer server.Close()

	req := httptest.NewRequest("GET", server.URL+"/", nil)
	rec := httptest.NewRecorder()

	middleware.ServeHTTP(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	defer rec.Result().Body.Close()

	t.Log(string(body))
}
```


### Error Handler

- Kadang middleware juga biasa digunakan untuk melakukan error handler
- Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan mengubah panic tersebut menjadi error response
- Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan


### Kode: Error Handler Middleware

```go
type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered From Error: ", err)

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error: %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(w, r)
}
```


### Kode: Test Error Handle Middleware

```go
func TestMiddlewareErrorHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("Something went wrong")
	})

	middleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: middleware,
	}

	server := httptest.NewServer(errorHandler)
	defer server.Close()

	t.Log(server.URL)

	req := httptest.NewRequest("GET", server.URL+"/panic", nil)
	rec := httptest.NewRecorder()

	errorHandler.ServeHTTP(rec, req)

	body, _ := ioutil.ReadAll(rec.Result().Body)
	defer rec.Result().Body.Close()

	t.Log(string(body))
}
```


## Routing Library
---

- Walaupun Go-Lang sudah menyediakan ServeMux sebagai handler yang bisa menghandle beberapa endpoint atau istilahnya adalah routing
- Tapi kebanyakan programmer Go-Lang biasanya akan menggunakan library untuk melakukan routing
- Hal ini dikarenakan ServeMux tidak memiliki advanced fitur seperti path variable, auto binding parameter dan middleware
- Banyak alternatif lain yang bisa kita gunakan untuk library routing selain ServeMux


### Contoh Routing Library

- https://github.com/julienschmidt/httprouter
- https://github.com/gorilla/mux
- Dan masih banyak lagi : https://github.com/julienschmidt/go-http-routing-benchmark 