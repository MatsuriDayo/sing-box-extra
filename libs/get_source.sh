#!/bin/bash
set -e

[ ! -z $NO_ENV ] || source libs/get_source_env.sh
pushd ..

####
if [ ! -d "sing-box" ]; then
  git clone --no-checkout https://github.com/MatsuriDayo/sing-box.git
fi
pushd sing-box
git checkout "$COMMIT_SING_BOX"
popd

####
# if [ ! -d "sing-dns" ]; then
#   git clone --no-checkout https://github.com/MatsuriDayo/sing-dns.git
# fi
# pushd sing-dns
# git checkout "$COMMIT_SING_DNS"
# popd

####
if [ ! -d "libneko" ]; then
  git clone --no-checkout https://github.com/MatsuriDayo/libneko.git
fi
pushd libneko
git checkout "$COMMIT_LIBNEKO"
popd

popd
