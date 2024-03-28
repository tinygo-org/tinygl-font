all: fonts

fonts: roboto proggy-clean proggy-small proggy-tiny

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
	go run ./generate -font roboto/Roboto-Regular.ttf -size 16 -o roboto/regular16.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular16.go

roboto/regular20.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 20 -o roboto/regular20.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular20.go

roboto/regular24.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 24 -o roboto/regular24.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular24.go

roboto/regular28.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 28 -o roboto/regular28.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular28.go

roboto/regular32.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 32 -o roboto/regular32.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular32.go

roboto/regular36.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 36 -o roboto/regular36.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular36.go

roboto/regular40.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 40 -o roboto/regular40.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular40.go

roboto/regular44.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 44 -o roboto/regular44.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular44.go

roboto/regular48.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 48 -o roboto/regular48.go -package=roboto -weight=Regular $(FONT_FLAGS)
	@go fmt roboto/regular48.go

proggy-clean: \
    proggy/clean16.go \
    proggy/clean20.go \
    proggy/clean24.go \
    proggy/clean28.go \
    proggy/clean32.go \
    proggy/clean36.go \
    proggy/clean40.go \
    proggy/clean44.go \
    proggy/clean48.go

proggy/clean16.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 16 -o proggy/clean16.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean16.go

proggy/clean20.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 20 -o proggy/clean20.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean20.go

proggy/clean24.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 24 -o proggy/clean24.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean24.go

proggy/clean28.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 28 -o proggy/clean28.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean28.go

proggy/clean32.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 32 -o proggy/clean32.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean32.go

proggy/clean36.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 36 -o proggy/clean36.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean36.go

proggy/clean40.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 40 -o proggy/clean40.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean40.go

proggy/clean44.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 44 -o proggy/clean44.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean44.go

proggy/clean48.go: generate/main.go
	go run ./generate -font proggy/ProggyClean.ttf -size 48 -o proggy/clean48.go -package=proggy -weight=Clean $(FONT_FLAGS)
	@go fmt proggy/clean48.go


proggy-small: \
    proggy/small16.go \
    proggy/small20.go \
    proggy/small24.go \
    proggy/small28.go \
    proggy/small32.go \
    proggy/small36.go \
    proggy/small40.go \
    proggy/small44.go \
    proggy/small48.go

proggy/small16.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 16 -o proggy/small16.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small16.go

proggy/small20.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 20 -o proggy/small20.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small20.go

proggy/small24.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 24 -o proggy/small24.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small24.go

proggy/small28.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 28 -o proggy/small28.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small28.go

proggy/small32.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 32 -o proggy/small32.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small32.go

proggy/small36.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 36 -o proggy/small36.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small36.go

proggy/small40.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 40 -o proggy/small40.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small40.go

proggy/small44.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 44 -o proggy/small44.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small44.go

proggy/small48.go: generate/main.go
	go run ./generate -font proggy/ProggySmall.ttf -size 48 -o proggy/small48.go -package=proggy -weight=Small $(FONT_FLAGS)
	@go fmt proggy/small48.go


proggy-tiny: \
    proggy/tiny16.go \
    proggy/tiny20.go \
    proggy/tiny24.go \
    proggy/tiny28.go \
    proggy/tiny32.go \
    proggy/tiny36.go \
    proggy/tiny40.go \
    proggy/tiny44.go \
    proggy/tiny48.go

proggy/tiny16.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 16 -o proggy/tiny16.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny16.go

proggy/tiny20.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 20 -o proggy/tiny20.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny20.go

proggy/tiny24.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 24 -o proggy/tiny24.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny24.go

proggy/tiny28.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 28 -o proggy/tiny28.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny28.go

proggy/tiny32.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 32 -o proggy/tiny32.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny32.go

proggy/tiny36.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 36 -o proggy/tiny36.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny36.go

proggy/tiny40.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 40 -o proggy/tiny40.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny40.go

proggy/tiny44.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 44 -o proggy/tiny44.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny44.go

proggy/tiny48.go: generate/main.go
	go run ./generate -font proggy/ProggyTiny.ttf -size 48 -o proggy/tiny48.go -package=proggy -weight=Tiny $(FONT_FLAGS)
	@go fmt proggy/tiny48.go
