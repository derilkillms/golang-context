package golang_context

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// Context biasanya dibuat per request(misal setiap ada request masuk ke server melalui http request), context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses.
// Context di golang biasa digunakan untuk mengirim data request atau sinyal ke proses lain atau dari gorutine ke goroutine yang lain. Dengan menggunakan context, ketika kita ingin membatalkan sebuah proses, kita cukup mengirim sinyal ke context nya, maka secara otomatis semua proses akan dibatalkan.

func TestContext(t *testing.T) {
	// Membuat context kosong. Tidak pernah dibatalkan, tidak pernah timeout, dam tidak memiliki value apapun. Biasanya digunakan di main function atau dalam test, atau dalam awal request proses terjadi.
	background := context.Background()
	fmt.Println(background)

	// Membuat context kosong seperti Background(), namun biasanya menggunakan ini ketika belum jelas context apa yang ingin digunakan.
	todo := context.TODO()
	fmt.Println(todo)
}

// Context menganut konsep parent dan child.
// saat kita membuat context, kita bisa membuat child context dari context yang sudah ada.

// Pada saat awal membuat context, context tidak memiliki value. Kita bisa menambahkan sebuah value dengan data Pair (key – value) ke dalam context, saat kita menambahkan value ke context secara otomatis akan tercipta child context baru
// original context nya tidak akan berubah sama sekali. Untuk menambahkan value ke context kita bisa menggunakan function context.WithValue(parent, key, value).

func TestContextWithValue(t *testing.T) {
	parent := context.Background()

	childA := context.WithValue(parent, "Key", "Muhammad")
	childB := context.WithValue(childA, "Key", "Deril")

	fmt.Println(childA)
	fmt.Println(childB)

	// mengambil value nya dan menampilkan ke layar
	fmt.Println(childA.Value("Key"))
	fmt.Println(childB.Value("Key"))
}

//  goroutine yang menggunakan context tetap harus melakukan pengecekan terhadap context nya, jika tidak maka tidak ada gunanya
//  Untuk membuat context dengan cancel signal kita bisa menggunakan function context.WithCancel(parent)

func Satu(ctx context.Context, cancel func(), group *sync.WaitGroup) {
	defer group.Done()
	canceled := ctx.Err()

	fmt.Println("Goroutine 1 start")
	if canceled != nil {
		fmt.Println(canceled)
		return
	}
	fmt.Println("Goroutin 1 done")
	cancel()
}

func Dua(ctx context.Context, cancel func(), group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println("Goroutine 2 start")

	// membuat goroutine 2 lebih lambat dengan menambahkan Sleep selama 1 detik.
	time.Sleep(1 * time.Second)
	canceled := ctx.Err()

	if canceled != nil {
		fmt.Println("Goroutine 2", canceled)
		return
	}
	fmt.Println("Goroutin 2 done")
	cancel()
}

func TestContextWithCancel(t *testing.T) {
	group := sync.WaitGroup{}
	parent := context.Background()
	data, cancel := context.WithCancel(parent)
	ctx := context.WithValue(data, "data", "Muhammad Deril")

	fmt.Println("Start 2 goroutine")
	group.Add(2)
	go Satu(ctx, cancel, &group)
	go Dua(ctx, cancel, &group)
	group.Wait()
}

// Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel ke context secara otomatis dengan pengaturan timeout.
// Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan function context.WithTimeout(parent, duration).

func TestContextWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("overslept")
		}
	}
}

// Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan deadline.
// Pengaturan deadline sidikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang, nah jika deadline ditentukan kapan waktu timeout nya, misal jam 12 siang hari ini.

func TestContextWithDeadline(t *testing.T) {

	// Untuk penggunaan deadline sebenar nya sama saja dengan timeout sebelumnya, yang berbeda hanya method dan parameter nya saja
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("overslept")
		}
	}
}
