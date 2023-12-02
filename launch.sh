#!/bin/bash

# check that private key is set
if [ -z "$ETH_PRIV_KEY" ]; then
    echo "Error: The ETH_PRIV_KEY environment variable is not set."
    exit 1
fi

# check that Mumbai RPC is set
if [ -z "$ETH_RPC_URL" ]; then
    echo "Error: The ETH_RPC_URL environment variable is not set."
    exit 1
fi


# check that our own server URL is set
if [ -z "$ISSUER_URL" ]; then
    echo "Error: The ISSUER_URL environment variable is not set."
    exit 1
fi

# clone the issuer repo
echo "Cloning issuer node"
git clone https://github.com/0xPolygonID/issuer-node.git

cd issuer-node

# copy the env samples to real env files
cp .env-issuer.sample .env-issuer
cp .env-api.sample .env-api
cp .env-ui.sample .env-ui

# change the ethereum rpc
sed '/ISSUER_ETHEREUM_URL/d' .env-issuer > .env-issuer.tmp
echo ISSUER_ETHEREUM_URL=$ETH_RPC_URL >> .env-issuer.tmp

# change the auth user na,e
sed '/ISSUER_API_AUTH_USER/d' .env-issuer > .env-issuer.tmp
echo ISSUER_API_AUTH_USER=user-issuer >> .env-issuer.tmp

# change the auth user pass
sed '/ISSUER_API_AUTH_PASSWORD/d' .env-issuer > .env-issuer.tmp
echo ISSUER_API_AUTH_PASSWORD=password-issuer >> .env-issuer.tmp

# i dont think we need this for our demo
# sed '/ISSUER_SERVER_URL/d' .env-issuer > .env-issuer.tmp
# echo ISSUER_SERVER_URL=$ISSUER_URL >> .env-issuer.tmp

mv .env-issuer.tmp .env-issuer

# start docker
sudo systemctl start docker

# start the infra services
make up

# enable vault auth
make add-vault-token

# set our private key
make private_key=$ETH_PRIV_KEY add-private-key

# run the api
make run

# change username in the ui env file
sed '/ISSUER_UI_AUTH_USERNAME/d' .env-ui > .env-ui.tmp
echo ISSUER_UI_AUTH_USERNAME=user-ui >> .env-ui.tmp

# change pass in the ui env file
sed '/ISSUER_UI_AUTH_PASSWORD/d' .env-ui > .env-ui.tmp
echo ISSUER_UI_AUTH_PASSWORD=password-ui >> .env-ui.tmp

mv .env-ui.tmp .env-ui

# set our own server URL in the api env
sed '/ISSUER_API_UI_SERVER_URL/d' .env-api > .env-api.tmp
echo ISSUER_API_UI_SERVER_URL=$ISSUER_URL >> .env-api.tmp
mv .env-api.tmp .env-api

# generate the did for the issuer
make generate-issuer-did 

# run the UI
make run-ui

cd ..

# clone the linkster api
echo "Cloning the linkster API"
git clone https://github.com/getlinkster/linkster-api.git

cd linkster-api

cp .env.example .env

make run