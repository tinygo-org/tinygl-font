all: fonts

fonts: roboto freemono

roboto: \
    roboto/regular16.go \
    roboto/regular20.go \
    roboto/regular24.go \
    roboto/regular28.go \
    roboto/regular32.go \
    roboto/regular36.go \
    roboto/regular40.go \
    roboto/regular44.go \
    roboto/regular48.go

roboto/regular16.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 16 -o roboto/regular16.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular16.go

roboto/regular20.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 20 -o roboto/regular20.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular20.go

roboto/regular24.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 24 -o roboto/regular24.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular24.go

roboto/regular28.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 28 -o roboto/regular28.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular28.go

roboto/regular32.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 32 -o roboto/regular32.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular32.go

roboto/regular36.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 36 -o roboto/regular36.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular36.go

roboto/regular40.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 40 -o roboto/regular40.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular40.go

roboto/regular44.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 44 -o roboto/regular44.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular44.go

roboto/regular48.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 48 -o roboto/regular48.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular48.go

freemono: \
    freemono/regular16.go \
    freemono/regular20.go \
    freemono/regular24.go \
    freemono/regular28.go \
    freemono/regular32.go \
    freemono/regular36.go \
    freemono/regular40.go \
    freemono/regular44.go \
    freemono/regular48.go

freemono/regular16.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 16 -o freemono/regular16.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular16.go

freemono/regular20.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 20 -o freemono/regular20.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular20.go

freemono/regular24.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 24 -o freemono/regular24.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular24.go

freemono/regular28.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 28 -o freemono/regular28.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular28.go

freemono/regular32.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 32 -o freemono/regular32.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular32.go

freemono/regular36.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 36 -o freemono/regular36.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular36.go

freemono/regular40.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 40 -o freemono/regular40.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular40.go

freemono/regular44.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 44 -o freemono/regular44.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular44.go

freemono/regular48.go: generate/main.go
	go run ./generate -font freemono/FreeMono.ttf -size 48 -o freemono/regular48.go -package=freemono $(FONT_FLAGS)
	@go fmt freemono/regular48.go
