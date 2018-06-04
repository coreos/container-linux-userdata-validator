# CoreOS Userdata Validator

This code powers the public service at https://coreos.com/validate/.

## Building

The included multi-stage Dockerfile can be used to build working images. Just run the following:

    docker build .

## Updating dependencies

The following glide commands can be used to update the dependencies of this project:

    glide update --strip-vendor
    glide-vc --use-lock-file --no-tests --only-code

## Deployment

This repository is configured for autobuilding on Quay, so that new git
tags are automatically available as container images.
Autobuilt images are pushed to `quay.io/coreosinc/coreos-userdata-validator`,
which is not available for pulling by the general audience.

Deployable tags are pushed to git in the format `yyyymmdd-rev`.
The corresponding image needs to be manually bumped in the relevant
Helm chart for the CoreOS kubernetes cluster.
