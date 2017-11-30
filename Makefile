PACKAGES= raylib raygui raymath easings physics rres

GO?= go

all: packages

packages:
	@for pkg in ${PACKAGES}; do \
		echo "Building package github.com/Member1221/raylib-go/$$pkg..."; \
		${GO} build github.com/Member1221/raylib-go/$$pkg || exit 1; \
	done
