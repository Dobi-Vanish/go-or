# go-or

[![Go Reference](https://pkg.go.dev/badge/github.com/dobi-vanish/go-or.svg)](https://pkg.go.dev/github.com/dobi-vanish/go-or)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Функция `Or` объединяет один или более done-каналов в один. Возвращаемый канал закрывается, когда закрывается любой из исходных. Как только любая из горутин получает сигнал, она закрывает общий выходной канал с помощью sync.Once, гарантируя однократное закрытие
## Установка

```bash
go get https://github.com/Dobi-Vanish/go-or
