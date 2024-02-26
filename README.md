# Timers vs Tickers

Timers - Bunlar tek seferlik görevler için kullanılır. Gelecekteki tek bir olayı temsil eder.
Timers'e ne kadar süre beklemek istediğinizi söylersiniz ve o da o zaman bildirilecek bir kanal sağlar.

Ticker'lar - Ticker'lar, bir eylemi belirli zaman aralıklarında tekrar tekrar gerçekleştirmeniz gerektiğinde son derece yararlıdır.
Bu görevleri uygulamalarımızın arka planında çalıştırmak için goroutine'lerle birlikte ticker'ları kullanabiliriz.

# Basit Bir Zamanlayıcı

Şimdi 5 saniye sonra çalışacak olan gerçekten basit bir timer ile başlayalım.
Zamanın yürütülmesi bir goroutine ile ilişkilidir.
done kullanmanın amacı sadece programın yürütülmesini kontrol etmektir.

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(5 * time.Second)

	done := make(chan bool)

	go func() {
		<-timer.C
		fmt.Println("Zaman doldu!")
		done <- true
	}()

	<-done
}

```

# Turlardan önce Timer'i durdurun

Aşağıdaki kod listesinde, yürütmeyi durdurmak için zamanlayıcıdaki Stop() yöntemine bir çağrı yapabilirsiniz.
Bir zamanlayıcıyı 5 saniye sonra çalıştırılacak şekilde yapılandırdık ancak 17 numaralı satır bunu zamanından önce durdurdu. (timerstop.go)

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 5)

	go func() {
		<-timer.C
		fmt.Println("Zaman doldu!")
	}()

	stopped := timer.Stop()

	if stopped {
		fmt.Println("Zamanlayıcı durdu!")
	}

	time.Sleep(3 * time.Second)
}

```

# Basit Bir Ticker

Her 1 saniyede bir basit bir fmt.Println deyimini tekrar tekrar çalıştırdığımız gerçekten basit bir işlemle başlayalım.

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Go Tickers")
	ticker := time.NewTicker(1 * time.Second)

	for tc := range ticker.C {
		fmt.Println("Golang ", tc)
	}
}

```

# Bir Ticker'ı Durdurma

Aşağıdaki kod listesinde bir ticker her 1 saniyede bir çalıştırılacak şekilde yapılandırılmıştır.
Bir ticker olayının alınıp alınmadığını veya done kanalının onu durdurmak için bir sinyal alıp almadığını izlemek için bir goroutine yazılmıştır.
done kanalına bir sinyal göndermeden önce, 24. satırda, ticker'ın 5 kez çalışmasına izin vermek için time.Sleep(5*time. Second) çağrısı yaptık.

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	nt := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-nt.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	nt.Stop()
	done <- true
	fmt.Println("Ticker durdu!")
}

```

# Bir Arka Plan Ticker'ı

Arka planda çalıştırmak istediğimiz bir görevimiz olsaydı,
nt.C üzerinde yineleme yapan for a döngümüzü bir goroutine'in içine taşıyabilirdik,
bu da uygulamamızın diğer görevleri yürütmesine olanak tanırdı.

Ticker'ı oluşturma ve döngüye sokma kodunu bgTicker() adlı yeni bir fonksiyona taşıyalım ve,
ardından main() fonksiyonumuz içinde go anahtar sözcüğünü kullanarak bunu bir goroutine olarak çağıralım.

```go
package main

import (
	"fmt"
	"time"
)

func bgTicker() {
	nt := time.NewTicker(1 * time.Second)
	for _ = range nt.C {
		fmt.Println("Go")
	}
}

func main() {
	fmt.Println("Go Tickers")

	fmt.Println("Uygulamamın geri kalanı devam edebilir...")

	go bgTicker()

	select {}

}

```

# Bir Arka Plan Ticker'ni Durdurma

```go
package main

import (
	"fmt"
	"time"
)

func bgTicker(stop chan struct{}) {
	nt := time.NewTicker(1 * time.Second)
	count := 0

	for {
		select {
		case <-nt.C:
			count++
			fmt.Println("Go")

			if count >= 5 {
				close(stop)
				return
			}
		case <-stop:
			nt.Stop()
			return
		}
	}
}

func main() {
	fmt.Println("Go Tickers")

	fmt.Println("Uygulamamın geri kalanı devam edebilir...")

	stopTicker := make(chan struct{})
	go bgTicker(stopTicker)

	<-stopTicker
}

```
