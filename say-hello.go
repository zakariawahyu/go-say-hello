package go_say_hello

/**
Untuk membuat module baru kita bisa menggunakan perintah : go mod init nama module
Golang akan secara otomatis membaut file go.mod yang berisikan informasi nama module dan versi golang yang ktia gunakan
 */

/**
Rilis module
Golang terintegerasi baik dengan git
Untuk merilis module kita hanya menambahkan tags di git
git tag v1.0.0
git push origin --tags
 */

/**
Upgrade module
kita hanya perlu menambahkan tags version saja
 */

/**
Major Upgarde
Major upgrade biasanya terjadi karena ada perubahan pada sisi kode program kita, sehuingga membuatnya tidak compatible (sebaiknya dihindari)
Namun jika terpaksa merubah major, strategy terbaik adalah merubah nama module
 */
func SayHello(name string) string  {
	return "Hello " + name +" world, dan ini upgrade module to v1.1.0"
}
