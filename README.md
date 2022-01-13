# Petunjuk Penggunaan dan Dokumentasi Script (Problem 3)

## Dibuat oleh :
Ahmad Mujahid Abdurrahman (ahmadm.abdurrahman@gmail.com)

## Pentujuk Penggunaan
- Pastikan interpreter Go sudah terinstal pada komputer
- Agar dapat menjalankan unit test, lakukan set-up berikut pada Go dengan mengetik perintah pada terminal / cmd :
    ```go
    go env -w GO111MODULE=auto
    ```
- Untuk mengeksekusi program, gunakan perintah :
    ```go
    go run main.go
    ```
- Untuk menjalakan unit test, gunakan perintah :
    ```go
    go test
    ```
    dengan catatan berada di direktori di mana file **main.go** berada.

## Penjelasan terkait program
Pada program ini, terdapat 2 fungsi utama, yaitu :
- gcd()
- countItem()

### gcd()
```go
func gcd(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return gcd(num2, num1%num2)
	}
}
```
Fungsi **gcd()** / *Greatest Common Divisor* digunakan untuk mencari berapa jumlah kotak yang dapat didistribusikan sesuai dengan jumlah apel dan kue yang dimiliki Ainun.

### countItem()
``` go
func countItem(apples int, cakes int, divisor int) (int, int) {
	return apples / divisor, cakes / divisor
}
```
Fungsi **countItem()** ini digunakan untuk menghitung jumlah masing-masing apel dan kue yang terdistribusi di setiap kotak.

## Unit Test
Pada bagian unit test (**main_test.go**), terdapat dua fungsi, yaitu :
- TestGcd()
- TestCountItem()

### TestGcd()
``` go
func TestGcd(t *testing.T) {
	result := gcd(25, 20)
	resultExpt := 5

	if result != 5 {
		t.Errorf("gcd Test failed. Expected result is %d But you got %d", resultExpt, result)
	} else {
		t.Logf("gcd Test passed")
	}
}
```
Fungsi ini digunakan untuk mengecek apakah operasi **gcd()** sudah berjalan sesuai yang diharapkan.

### TestCountItem()
```go
func TestCountItem(t *testing.T) {
	apples := 25
	cakes := 20
	divisor := gcd(apples, cakes)
	eachApples, eachCakes := countItem(apples, cakes, divisor)

	eachApplesExpt, eachCakesExpt := 5, 4

	if eachApples != 5 && eachCakes != 4 {
		t.Errorf("countItem Test failed. Expected result {apples: %d, cakes: %d}. But you got {apples: %d, cakes: =%d}",
			eachApplesExpt, eachCakesExpt, eachApples, eachCakes)
	} else {
		t.Logf("countItem Test passed")
	}
}
```
Fungsi ini digunakan untuk mengecek apakah operasi **countItem()** sudah berjalan sesuai yang diharapkan.

Untuk mengoperasikan unit test, dapat digunakan perintah sebagai berikut :
```go
go test
```
atau
```go
go test -v
```
untuk mendapatkan output yang *vorbose* (lebih mendetail).