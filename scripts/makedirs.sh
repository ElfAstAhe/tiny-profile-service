#!/usr/bin/env bash

cd ..

mkdir -p api/{proto/tiny-profile-service,rest} \
    cmd/tiny-profile-service \
    configs \
    deployments \
    docs \
    internal/{app,config,domain,facade/{dto,mapper},repository/postgres,transport/{errs,rest,grpc},usecase} \
    migrations/tiny-profile-service \
    pkg \
    scripts

cd scripts
