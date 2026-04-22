package or

import "sync"

// Or объединяет один или более done-каналов в один.
// Возвращаемый канал закроется, как только закроется любой из входных каналов.
// Если входной список пуст, возвращается уже закрытый канал.
// Функция не закрывает входные каналы и не вызывает паники при повторном закрытии.
//
// Пример:
//
//	done := Or(sig1, sig2, sig3)
//	<-done
func Or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		c := make(chan interface{})
		close(c)
		return c
	}

	done := make(chan interface{})
	var once sync.Once

	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			<-ch
			once.Do(func() { close(done) })
		}(ch)
	}

	return done
}
