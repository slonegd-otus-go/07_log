package otus

import "time"

func SetRealNow() {
	now = time.Now
}

func SetNow(v string) {
	t, err := time.Parse("2006-01-02", v)
	if err != nil {
		panic(err)
	}
	now = func() time.Time {
		return t
	}
}

