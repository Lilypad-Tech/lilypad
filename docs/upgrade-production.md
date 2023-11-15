# how to upgrade production

## if only the binary has changed:

### control plane

```
gcloud compute ssh --zone "us-central1-a" "lilypad-v2" --project "mlops-consulting-production"
```

```
cd /app/lilypad
git checkout main
git pull
go build
sudo mv lilypad /usr/bin/lilypad
sudo systemctl restart solver
sudo systemctl restart mediator # not used currently
```

```
journalctl -f -u solver
```

### gpu node(s)

```
ssh node05.lukemarsden.net
```

```
cd /app/lilypad
git checkout main
git pull
go build
sudo mv lilypad /usr/bin/lilypad
sudo systemctl restart resource-provider-gpu
```

```
journalctl -f -u resource-provider-gpu
```

## if the smart contract has changed:

we have to wipe out the whole world and start from scratch (for now)


```
gcloud compute ssh --zone "us-central1-a" "lilypad-v2" --project "mlops-consulting-production"
```

```
cd /app/lilypad
git checkout main
git pull

sudo systemctl stop solver
sudo systemctl stop mediator
./stack boot # THIS WILL WIPE OUT GETH AND ALL THE BLOCKCHAIN DATA
```

lilysaas will now be stopped, go start it again with `(cd /app/lilysaas; docker-compose up -d)`

now we need to edit the code to set the new default controller contract address from `stack boot` above
https://github.com/bacalhau-project/lilypad/blob/main/pkg/options/web3.go#L21
or `vim pkg/options/web3.go`

commit and push and then use the resulting binary for the control plane and resource providers
(if the stars aligned you may find the controller address hasn't changed, in which case you can skip pushing a change to git)

```
go build

sudo mv lilypad /usr/bin/lilypad

sudo systemctl start solver
sudo systemctl start mediator # not used currently
```
```
journalctl -f -u solver
```

### gpu node(s)

```
ssh node05.lukemarsden.net
```

github should have build the binary by now, go download the latest one from https://github.com/bacalhau-project/lilypad/releases

```
cd /app/lilypad
## CHANGE THE HASH BELOW TO THE LATEST FROM RELEASES PAGE!!
wget https://github.com/bacalhau-project/lilypad/releases/download/v2.0.0-f70e4e1/lilypad-linux-amd64
chmod +x lilypad-linux-amd64
sudo mv lilypad-linux-amd64 /usr/bin/lilypad
sudo systemctl restart resource-provider-gpu
```

```
journalctl -f -u resource-provider-gpu
```

