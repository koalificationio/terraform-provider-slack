# terraform-provider-slack

[![Build Status](https://travis-ci.org/koalificationio/terraform-provider-slack.svg?branch=master)](https://travis-ci.org/koalificationio/terraform-provider-slack)

Terraform provider for Slack

## Installing the Provider
Download binary from releases and follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory, run `terraform init` to initialize it.

Check documentation in [./website](./website) folder.

## Developing the Provider

Clone repository to: `$HOME/development/koalificationio/`

```sh
$ mkdir -p $HOME/development/koalificationio/; cd $HOME/development/koalificationio/
$ git clone https://github.com/koalificationio/terraform-provider-slack
...
```

Enter the provider directory and run `make tools`. This will install the needed tools for the provider.

```sh
$ make tools
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-aws
...
```

## Testing the Provider

In order to test the provider, you can run `make test`.

```sh
$ make test
```
