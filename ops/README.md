# deployment from scratch (MVP deployment)

* create new IP address on GCP
* create new VM on GCP (e2-standard-8, 1TB root disk, ubuntu 22.04) using IP above
* install docker: https://docs.docker.com/engine/install/ubuntu/
* install node 20: https://github.com/nodesource/distributions#debian-and-ubuntu-based-distributions
* install go using PPA: https://github.com/golang/go/wiki/Ubuntu

```bash
sudo adduser $USER docker
```

We need to point DNS for `testnet.lilypad.tech` at the node and open the following ports:

 * 80
 * 443
 * 8080
 * 8545
 * 8546
 
log out and log in again

```bash
cd /
sudo mkdir app
sudo chown $USER app
cd /app/
git clone https://github.com/lilypad-tech/lilypad
cd lilypad
(cd hardhat && yarn install)
```

Then we create the production keys:

```bash
(cd hardhat && npx hardhat run scripts/generate-addresses.ts)
```

This will print out the various private keys. We need to copy these into the `/app/lilypad/.env` file.

Now we can boot geth and it will fund the various accounts:

```bash
./stack boot
```

Let's check this:

```bash
./stack balances
```

Time to make the following files by copying the respective private key from `/app/lilypad/.env`

Each file should be of the following format:

```
WEB3_PRIVATE_KEY=xxx
```

 * `/app/lilypad/solver.env` (copy `SOLVER_PRIVATE_KEY` from `.env`)
 * `/app/lilypad/mediator.env` (copy `MEDIATOR_PRIVATE_KEY` from `.env`)
 * `/app/lilypad/resource-provider.env` (copy `RESOURCE_PROVIDER_PRIVATE_KEY` from `.env`)
 * `/app/lilypad/job-creator.env` (copy `SOLVER_PRIVATE_KEY` from `.env`)
    * IMPORTANT: this has to be the solver private key because the job creator runs as it

Now - we copy the systemd units and reload systemd:

```bash
sudo cp /app/lilypad/ops/systemd/*.service /etc/systemd/system
sudo systemctl daemon-reload
```

Now we build lilypad:

```bash
cd /app/lilypad
go build .
sudo mv lilypad /usr/bin/lilypad
```

Then we install bacalhau:

```bash
cd /tmp
wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.0.3/bacalhau_v1.0.3_linux_amd64.tar.gz
tar xfv bacalhau_v1.0.3_linux_amd64.tar.gz
sudo mv bacalhau /usr/bin/bacalhau
sudo mkdir -p /app/data/ipfs
sudo chown -R $USER /app/data
```

Finally we start the various systemd services:

```bash
sudo systemctl start bacalhau
sudo systemctl start solver
sudo systemctl start mediator
sudo systemctl start resource-provider
sudo systemctl start job-creator
```