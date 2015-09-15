# Universal Publishing service files + service file definition
Repository that contains the UP stack service files and services definition for pre-prod and prod environments.

##Development

Get a linux environment: local/VM

- Juila has VM conf for working networking
- Julia - shared directories with host, symlink to $HOME

In VM install required packages (as root):

```bash
## ubuntu
apt-get install docker.io golang git

## Fedora/CentOS
yum install docker.io golang git

## Arch linux
pacman -S docker go git
```

Set up GO env:

```bash
mkdir -p ~/gopath/bin

# edit ~/.bashrc and add:
export GOPATH=$HOME/gopath
export PATH=$PATH:$GOPATH/bin
```

Run docker daemon, allowing insecure connections from our register (as root):

```bash
## Add your user to the docker group:
sudo gpasswd -a [user] docker

## services stop docker
### In new terminal window
sudo docker -d --insecure-registry up-registry.ft.com
```

Install fleet and etcd:

```bash
go get github.com/coreos/fleet
cd $GOPATH/src/github.com/coreos/fleet && ./build
cp $GOPATH/src/github.com/coreos/fleet/bin/* $GOPATH/bin

go get github.com/coreos/etcd
cd $GOPATH/src/github.com/coreos/etcd && ./build
cp $GOPATH/src/github.com/coreos/etcd/bin/* $GOPATH/bin
```

Run etcd and fleet daemons:

```bash
### In new terminal window
$GOPATH/bin/etcd

### In new terminal window
sudo $GOPATH/bin/fleetd
```

