To init go module
> go mod init example.com/m

To install github dependency
> go get github.com/go-sql-driver/mysql


1. Add the service to systemd.
 - sudo systemctl enable ```<filename>```.service
1. Activate the service.
 - sudo systemctl start ```<filename>```.service
1. Check if systemd started it.
 - sudo systemctl status ```<filename```.service
1. Stop systemd if so desired.
 - sudo systemctl stop ```<filename>```.service