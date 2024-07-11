FROM golang:1.22-bookworm
SHELL ["/bin/bash", "-c"]
RUN <<EOF
apt-get update -qq
apt-get install -qq -y curl wget build-essential sudo libdbus-1-dev libxcb-randr0-dev libxcb1-dev libxcb-xtest0-dev dbus dbus-x11
useradd -m -d /home/dev -s /bin/bash -u 1000 dev
adduser dev sudo
echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
EOF
USER dev
WORKDIR /home/dev
LABEL description="This is a docker image that offers tools for \
developing go applications as a non root user"
