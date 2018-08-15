## Get AWS DNS Server address within a VPC instance

## The problem

AWS Route 53 private hosted zones enable you to have private DNS names which only resolve from your VPC. This is great
when working from EC2 instances since everything is setup and ready to go. This however becomes a problem when using
docker containers on a systemd system. On such a system, `systemd-resolved` sits in between your host applications
and name resolution. The entry in `/etc/resolv.conf` is basically, `127.0.0.53` which doesn't mean much when you want
name resolution from a docker container which defaults to `8.8.8.8` for name resolution. Hence, we need a way to set
AWS VPC DNS server as an additional DNS server for the docker daemon.

Hence, this small utility. This is basically a golang version of the comment by Dusan Bajic [here](https://stackoverflow.com/questions/39100395/getting-the-dns-ip-used-within-an-aws-vpc).
Having a Golang binary means, I can use this on Linux and Windows.  Running the program will print the DNS server, 
which you can then use for example to set the DNS server in docker to be able to resolve private DNS names.

### Download

Download a binary from the releases tab, extract the zip.

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
172.34.0.2
```

On Linux, using `wget`:

```
$ wget https://github.com/amitsaha/aws-vpc-dns-address/releases/download/v0.1/aws-vpc-dns-address-0.1-linux-amd64.zip
$ unzip ..
$ mv aws-vpc-dns-address-0.1-linux-amd64 /usr/local/bin
$ sudo chmod +x /usr/local/bin/aws-vpc-dns-address-0.1-linux-amd64
$ aws-vpc-dns-address-0.1-linux-amd64
172.34.0.2
```

Then, you can use it, for example:

```
vpc_dns=$(aws-vpc-dns-address)
echo '{"dns":["'"$vpc_dns"'", "8.8.8.8"]}' > /etc/docker/daemon.json
cat /etc/docker/daemon.json
systemctl restart docker
```

### License

See LICENSE
