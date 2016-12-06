---
layout: "aws"
page_title: "AWS: aws_lightsail_key_pair"
sidebar_current: "docs-aws-resource-lightsail-key-pair"
description: |-
  Provides an Lightsail Key Pair 
---

# aws\_lightsail\_key\_pair

Provides a Lightsail Key Pair, for use with Lightsail Instances. These key pairs
are seperate from EC2 Key Pairs, and must be created or imported for use with
Lightsail.

Note: Lightsail is currently only supported in `us-east-1` region.

## Example Usage

```
# Create a new Lightsail Key Pair
resource "aws_lightsail_key_pair" "lg_key_pair" {
  name = "lg_key_pair"
}
```

## Create new Key Pair, encrypting the private key with a PGP Key

```
resource "aws_lightsail_key_pair" "lg_key_pair" {
  name    = "lg_key_pair"
  pgp_key = "keybase:keybaseusername"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Lightsail Key Pair
* `name_prefix` - (Optional) An optional prefix for the Key Pair name. This will
generate an `aws_lightsail_key_pair` with a name containing the prefix followed
by a random number
* `pgp_key` – (Optional) An optional PGP key to encrypt the resulting private
key material. 
* `public_key` - (Required) The public key material. This public key will be
imported into Lightsail

~> **NOTE:** a PGP key is not required, however it is strongly encouraged. 
Without a PGP key, the private key material will be stored in state unencrypted. 
`pgp_key` is ignored if `public_key` is supplied.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `id` - The name used for this key pair
* `arn` - The ARN of the Lightsail key pair
* `fingerprint` - The MD5 public key fingerprint as specified in section 4 of RFC 4716.
* `created_at` - The timestamp this key pair was created
* `public_key_base64` - the public key, base64 encoded
* `private_key_base64` - the private key, base64 encoded. This is only populated
when creating a new key, and when no `pgp_key` is provided
* `encrypted_private_key_64` – the private key material, base 64 encoded and
encrypted with the given `pgp_key`. This is only populated when creating a new
key and `pgp_key` is supplied

## Import

Lightsail Instances can be imported using their name, e.g.

```
$ terraform import aws_lightsail_key_pair.bar <name>
```
