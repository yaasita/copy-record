binary = \
	bin/ExcelRecordCopy.exe \

all: $(binary)

bin/ExcelRecordCopy.exe:
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -a -o $@

clean:
	rm $(binary)
