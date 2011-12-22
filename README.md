# noeq.go - A Go client for noeqd

# Install

		$ goinstall github.com/bmizerany/noeq.go

# Use

		package main

		import (
			"github.com/bmizerany/noeq.go"
			"log"
		)

		func main() {
			fl, err := noeq.New("secretToken", "localhost:4444")
			if err != nil {
				log.Fatal(err)
			}

			for n := 0; n < 100; n++ {
				id, err := genOne(fl, 3)
				if err != nil {
					log.Fatal("FATAL: unable to connect to weiqid")
				}
				log.Printf("%d", id)
			}
		}

		// The client attempts to give the user as much insight as possible
		// It will *not* automaticly attempt a reconnect
		// on error. However, it will attempt a reconnect if you ask for
		// another id after an error has occured. This gives the user more
		// control over how/when to attempt a reconnect.
		func genOne(fl *noeq.Client, trys int) (id uint64, err error) {
			for ; trys > 0; trys-- {
				id, err = fl.GenOne()
				if err != nil {
					log.Println("ERROR:", err)
					continue
				}

				return
			}

			return
		}

# Contributing

This is Github. You know the drill. Please make sure you keep your changes in a
branch other than `master` and in nice, clean, atomic commits. If you modify a
`.go` file, please use `gofmt` with no parameters to format it; then hit the
pull-request button.

# Issues

These are tracked in this repos Github [issues tracker](http://github.com/bmizerany/weiqi.go/issues).

# LICENSE

C)opyright (C) 2011 by Blake Mizerany ([@bmizerany](http://twitter.com/bmizerany))

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE. 
