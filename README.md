# go-or

[![Go Reference](https://pkg.go.dev/badge/github.com/ваш_username/go-or.svg)](https://pkg.go.dev/github.com/ваш_username/go-or)
[![Go Report Card](https://goreportcard.com/badge/github.com/ваш_username/go-or)](https://goreportcard.com/report/github.com/ваш_username/go-or)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Функция `Or` объединяет один или более done-каналов в один. Возвращаемый канал закрывается, когда закрывается любой из исходных. Полезно для таймаутов, отмены операций и управления жизненным циклом в конкурентном Go-коде.

## Установка

```bash
go get github.com/ваш_username/go-or
