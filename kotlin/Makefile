.PHONY: install compile run
install:
	brew update
	brew install kotlin
compile:
	kotlinc hello.kt -include-runtime -d hello.jar
run:
	java -jar hello.jar