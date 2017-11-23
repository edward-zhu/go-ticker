package ticker

import (
	"testing"
	"time"
)

func Test_ticker(t *testing.T) {
	const cd = 10
	d := cd * time.Millisecond
	ticker := NewTicker(d)

	prev := time.Now()

	for i := 0; i < 10; i++ {
		<-ticker.C
		sub := time.Now().Sub(prev)
		prev = time.Now()
		if sub < d {
			t.Errorf("time gap too short: %v duration:%v", sub, d)
		}
	}

	go func() {
		time.Sleep(cd / 2 * time.Millisecond)
		ticker.Reset()
	}()

	flag := false
	for i := 0; i < 20; i++ {
		<-ticker.C
		sub := time.Now().Sub(prev)
		prev = time.Now()
		if sub < d {
			t.Errorf("time gap too short: %v duration:%v", sub, d)
		}
		if sub > cd*1.5*time.Millisecond {
			flag = true
			break
		}
	}

	if !flag {
		t.Errorf("there must be one that cost longer than 1.5 * d.")
	}

}
