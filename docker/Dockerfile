FROM ubuntu:latest 

# add your keyname
ARG KEY_NAME <add your keyname>

# add the nickname for your validator on the network
ARG MONIKER <add your moniker>

# add the network id
ARG NETWORK_ID tomahawk_9624-1

# arbitrary validator metadata
ARG DETAILS <details as string>

# add amount you're delegating
ARG TOKEN_DELEGATION <amount>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update &&\ 
	apt-get install -y apt-utils git git-extras software-properties-common vim \
	wget ca-certificates curl build-essential libssl-dev make openssh-server snap

RUN snap install go --classic

# builds out tomad
RUN git clone https://github.com/tomahawk-network/tomahawk-Testnet
WORKDIR /tomahawk-testnet/cmd/tomad
RUN go install -tags ledger ./... &&\
    mv $HOME/go/bin/tomad /usr/bin/

# generate keys
RUN tomad keys add $KEY_NAME
    # tomad keys add $KEY_NAME --recover &&\
    # tomad keys add $KEY_NAME --ledger 

# initialize node
RUN tomad init $MONIKER --chain-id tomahawk_9624-1

# get genesis file
RUN wget https://github.com/tomahawk-network/tomahawk-Testnet-v2/raw/main/genesis.json -P $HOME/.tomad/config/
RUN sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0001atoma"/g' $HOME/.tomad/config/app.toml
RUN wget https://github.com/tomahawk-network/tomahawk-Testnet-v2/raw/main/Docker/tomad.service -P $HOME/systemd/system/

# run tomad
WORKDIR /tomahawk-testnet/cmd/tomad
RUN sudo systemctl daemon-reload &&\
    sudo systemctl enable tomad.service &&\
    systemctl start tomad && journalctl -u tomad -f

ENTRYPOINT ["/bin/bash"]
