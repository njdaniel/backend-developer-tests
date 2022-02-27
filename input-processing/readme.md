## input-processing

#### Testing large files

    $ yes testing with error line | head -c 10GB | go run main.go

Streaming large files seem to work fine even without changing the MaxScanTokenSize in bufio scanner
Each line is not over the default token size. Need to create a line that exceeds.


