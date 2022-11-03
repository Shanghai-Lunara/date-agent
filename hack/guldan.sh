#!/usr/bin/env bash

#set -o errexit
#set -o nounset
#set -o pipefail

GENS="$1"
PROJECT="$2"
REPOSITORY="$3"
TAG="$4"

addr=""
user=""
password=""
gitInfo=""
branch=""
image=""


if [ "${GENS}" = "cli" ] || grep -qw "address" <<<"${GENS}"; then
  addr=`env | grep GULDAN | grep ADDRESS | awk '{split($0,a,"="); print a[2]}'`
#  echo ${addr}
fi

if [ "${GENS}" = "cli" ] || grep -qw "username" <<<"${GENS}"; then
  user=`env | grep GULDAN | grep USER | awk '{split($0,a,"="); print a[2]}'`
#  echo ${user}
fi

if [ "${GENS}" = "cli" ] || grep -qw "password" <<<"${GENS}"; then
  password=`env | grep GULDAN | grep PASS | awk '{split($0,a,"="); print a[2]}'`
#  echo ${password}
fi

if [ "${GENS}" = "cli" ] || grep -qw "git" <<<"${GENS}"; then
  gitInfo=`cat .git/config | grep url | awk '{split($0,a,"/"); print a[2]}'`
#  echo ${gitInfo}
fi

if [ -z "$gitInfo" ]; then
  gitInfo=`cat .git/config | grep url | awk '{split($0,a,"/"); print a[5]}'`
fi

if [ "${GENS}" = "cli" ] || grep -qw "branch" <<<"${GENS}"; then
  branch=`git branch | grep \* | awk '{split($0,a," "); print a[2]}'`
#  echo ${branch}
fi

if [ "${GENS}" = "cli" ] || grep -qw "docker" <<<"${GENS}"; then
  image=`docker images --digests | grep ${PROJECT} | grep ${REPOSITORY} | grep ${TAG} | awk '{print $3}' | awk '{split($0,a,":"); print a[2]}'`
#  echo ${image}
fi

if [ "${GENS}" = "cli" ]; then
  guldan-cli \
    -address=${addr} \
    -username=${user} \
    -password=${password} \
    -project=${PROJECT} \
    -repository=${REPOSITORY} \
    -tag=${TAG} \
    -git=${gitInfo} \
    -branch=${branch}  \
    -commitHash=`git rev-parse HEAD` \
    -sha256=${image}
fi
