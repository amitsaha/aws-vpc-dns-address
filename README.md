## Get AWS DNS Server address within a VPC instance

This is basically a golang version of the comment by Dusan Bajic [here](https://stackoverflow.com/questions/39100395/getting-the-dns-ip-used-within-an-aws-vpc).
Having a Golang binary means, I can use this on Linux and Windows.  Running the program will print the DNS server, which you can then 
use for example to set the DNS server in docker to be able to resolve private DNS names.

### Download

Download a binary from the releases tab, extract the zip.

### Example Run

```
 > .\aws-vpc-dns-address-0.1-windows-amd64.exe
172.34.0.2
```

### Usage

This program will likely be run at instance start up either in user data or as part of your configuration management.
Hence, I recommend using a HTTP client to download the zip, extract it and put it somewhere in your system.

On Windows using PowerShell:

```
> Invoke-WebRequest https://github.com/amitsaha/aws-vpc-dns-address/releases/download/v0.1/aws-vpc-dns-address-0.1-windows-amd64.zip -OutFile aws-vpc-dns-address.zip
> Expand-Archive .\aws-vpc-dns-address.zip
> cd .\aws-vpc-dns-address\

> mv .\aws-vpc-dns-address-0.1-windows-amd64 .\aws-vpc-dns-address-0.1-windows-amd64.exe
> .\aws-vpc-dns-address-0.1-windows-amd64.exe
..
```

On Linux, using `wget`:

```
$ wget https://github.com/amitsaha/aws-vpc-dns-address/releases/download/v0.1/aws-vpc-dns-address-0.1-linux-amd64.zip
$ unzip ..
$ mv aws-vpc-dns-address-0.1-linux-amd64 /usr/local/bin
$ sudo chmod +x /usr/local/bin/aws-vpc-dns-address-0.1-linux-amd64
$ aws-vpc-dns-address-0.1-linux-amd64
..
```

### License

See LICENSE
