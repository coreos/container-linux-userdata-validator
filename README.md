# CoreOS Userdata Validator

This code powers the public service at https://coreos.com/validate/.

## Building

The included multi-stage Dockerfile can be used to build working images. Just run the following:

    docker build .

## Updating dependencies

The following glide commands can be used to update the dependencies of this project:

    glide update --strip-vendor
    glide-vc --use-lock-file --no-tests --only-code
