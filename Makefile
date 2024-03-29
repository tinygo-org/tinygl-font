all: roboto proggy

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

proggy: \
    proggy/regular16.go \
    proggy/regular20.go \
    proggy/regular24.go \
    proggy/regular28.go \
    proggy/regular32.go \
    proggy/regular36.go \
    proggy/regular40.go \
    proggy/regular44.go \
    proggy/regular48.go

proggy/regular16.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 16 -o proggy/regular16.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular16.go

proggy/regular20.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 20 -o proggy/regular20.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular20.go

proggy/regular24.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 24 -o proggy/regular24.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular24.go

proggy/regular28.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 28 -o proggy/regular28.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular28.go

proggy/regular32.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 32 -o proggy/regular32.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular32.go

proggy/regular36.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 36 -o proggy/regular36.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular36.go

proggy/regular40.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 40 -o proggy/regular40.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular40.go

proggy/regular44.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 44 -o proggy/regular44.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular44.go

proggy/regular48.go: generate/main.go
	go run ./generate -font proggy/ProggyVector-Regular.ttf -size 48 -o proggy/regular48.go -package=proggy $(FONT_FLAGS)
	@go fmt proggy/regular48.go
