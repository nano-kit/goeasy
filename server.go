package main

type server interface {
	Name() string
	Run()
}
