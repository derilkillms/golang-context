# golang-context

- Untuk mengetahui berapa jumlah thread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function di package runtime yang bisa kita gunakan untuk mengubah atau mengambil jumlah thread. Secara default, jumlah thread di golang itu sebanyak jumlah cpu di komputer .

- Context merupakan sebuah data yang isinya membawa value, sinyal cancel, sinyal timeout dan sinyal deadline.
