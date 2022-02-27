## input-processing

#### Testing large files

    $ yes testing with error line | head -c 10GB | go run main.go

Streaming large files seem to work fine even without changing the MaxScanTokenSize in bufio scanner
Each line is not over the default token size. Need to create a line that exceeds.

Default split for scanner is by line. Can use bufio.ScanWords to split up the line to a more manageable chunks. But the line could not contain spaces. A custom split can be created, ideally one that splits the line into sections that could == "error". 

