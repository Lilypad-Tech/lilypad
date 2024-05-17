# Devnet

A publicly accessible Lilypad network.

## Components

- A blockchain node
- A solver service
- A job creator service

### Resource providers

As is, the network cannot execute jobs, it needs resource providers to connect to the network and these in turn will be the ones running the jobs. Follow this instructions to [run a node in the Lilypad network](https://docs.lilypad.tech/lilypad/lilypad-milky-way-reference/run-a-node).

## Deployments

The three components have been isolated into docker containers and these components expose their services to the public via [Cloudflare tunnels](https://www.cloudflare.com/products/tunnel/) from inside the containers. This means there is no need to configure networking at the cloud service provider, just deploy the components using the appropriate build and run arguments to expose the services to the configured subdomains.

The deployment processes for each component have the same three steps:

- Build the docker image
- Push the image to the container registry
- Pull the image from the registry and execute the image in a container

## Blockchain node

There is a GitHub workflow to deploy the blockchain node, but it does not execute on code changes as there is no need to constantly deploy the blockchain node. There is ongoing work to implement a GitHub workflow that will compile and deploy the smart contracts.

### Reset the blockchain

If, for whatever reason, the blockchain node needs to be reset, this can be done in the following manner:

1. `ssh` into the instance/pod running the blockchain node
2. Connect into the `chain` container's shell (`docker exec -it chain /bin/bash`)
3. In the container shell execute the reset script (`./reset`). This will stop the node, clear the blockchain data, restart the node and fund the admin account.
4. Execute `./stack chain-boot` to compile, deploy and fund the smart contracts, token and accounts respectively. Change the configuration in the Hardhat scripts (`/hardhat.config.ts` and `/scripts/print-contract-env.ts`) to make sure the URLs point to the node's exposed subdomains (at the time of this write-up `8545` -> `https://devnet-chain-http.lilypad.tech` and `8546` -> `wss://devnet-chain-ws.lilypad.tech`).

## Solver and Job creator

When code changes reach the `main` branch, a GitHub workflow will trigger the flow that will first deploy and run the `solver` and then deploy and run the `job creator`. (*reminder*) The `solver` has to be up and running (and have registered its public URL to the blockchain) before the `job-creator` starts (the `job creator` will fail if it can't find the `solver`'s websocket server).

## Deploying to a new cloud provider

These steps have been used for [AWS](https://aws.amazon.com/) so maybe some changes have to be put in place for other providers but the overall flow should be similar.

### VM instance

Create a virtual machine to execute the component. The VM should have a running Docker daemon and should be able to pull an image from the chosen container registry.

[`How to install docker`](https://serverfault.com/questions/836198/how-to-install-docker-on-aws-ec2-instance-with-ami-ce-ee-update)

#### tldr;

```sh
sudo yum update -y
sudo yum install docker -y
sudo service docker start
sudo usermod -a -G docker <vm-user>
sudo systemctl enable docker (restart docker when VM instance restarts)
```

To grant access to the [ECR](https://aws.amazon.com/ecr/) registry from the VM execute `aws configure` with credentials that have `pull` permissions.

### ECR repo

Make sure the repo has been created.

### Cloudflare tunnel

Make sure the Cloudflare tunnel has been created and linked to the desired subdomain, and pass the token to the container in the build step.


