#!/bin/bash

set -eu

echo "deb http://deb.torproject.org/torproject.org xenial main" | sudo tee /etc/apt/sources.list.d/tor.list
echo "deb-src http://deb.torproject.org/torproject.org xenial main" | sudo tee -a /etc/apt/sources.list.d/tor.list
sudo apt-key adv --keyserver keys.gnupg.net --recv-keys A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89
sudo apt update
sudo apt upgrade -yy
sudo apt install -y tor deb.torproject.org-keyring

tmpdir=$(mktemp -d)
trap "rm -rf ${tmpdir}" EXIT

cd ${tmpdir}
wget -O ormesh.tar.gz https://github.com/cmars/ormesh/releases/download/v0.1.5/ormesh_0.1.5_linux_amd64.tar.gz
tar xf ormesh.tar.gz
sudo cp ormesh /usr/bin/ormesh
/usr/bin/ormesh agent privbind
/usr/bin/ormesh agent systemd | sudo tee /etc/systemd/system/ormesh.service

