## Get AWS DNS Server address within a VPC instance

This is basically a golang version of the comment by Dusan Bajic [here](https://stackoverflow.com/questions/39100395/getting-the-dns-ip-used-within-an-aws-vpc).
Having a Golang binary means, I can use this on Linux and Windows.  Running the program will print the DNS server, which you can then 
use for example to set the DNS server in docker to be able to resolve private DNS names.

### Download

Download a binary from the releases tab, extract the zip.

### Run it

```
 > .\aws-vpc-dns-address-0.1-windows-amd64.exe
172.34.0.2
```

### Usage

This program will likely be run at instance start up either in user data or as part of your configuration management.
Hence, I recommend using a HTTP client to download the zip, extract it and put it somewhere in your system.
