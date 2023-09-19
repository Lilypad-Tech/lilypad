# deployment from scratch (MVP deployment)

* create new IP address on GCP
* create new VM on GCP (e2-standard-8, 1TB root disk, ubuntu 22.04) using IP above
* install docker: https://docs.docker.com/engine/install/ubuntu/
* install node 20: https://github.com/nodesource/distributions#debian-and-ubuntu-based-distributions
* install go using PPA: https://github.com/golang/go/wiki/Ubuntu

```
sudo adduser $USER docker
```
log out and log in again
```
cd /
sudo mkdir app
sudo chown $USER app
cd /app/
git clone https://github.com/bacalhau-project/lilypad
cd lilypad
```

then run through [https://github.com/bacalhau-project/lilypad/blob/main/CONTRIBUTING.md](https://github.com/bacalhau-project/lilypad/blob/main/CONTRIBUTING.md)
