[![Let's FIWARE Banner](https://raw.githubusercontent.com/lets-fiware/ngsi-go/gh-pages/img/lets-fiware-logo-non-free.png)](https://www.letsfiware.jp/)
[![NGSI v2](https://img.shields.io/badge/NGSI-v2-5dc0cf.svg)](https://fiware-ges.github.io/orion/api/v2/stable/)
[![NGSI LD](https://img.shields.io/badge/NGSI-LD-d6604d.svg)](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf)

![FIWARE: Tools](https://nexus.lab.fiware.org/repository/raw/public/badges/chapters/deployment-tools.svg)
[![License: MIT](https://img.shields.io/github/license/lets-fiware/ngsi-go.svg)](https://opensource.org/licenses/MIT)
![GitHub all releases](https://img.shields.io/github/downloads/lets-fiware/ngsi-go/total)
[![Support badge](https://img.shields.io/badge/tag-fiware-orange.svg?logo=stackoverflow)](https://stackoverflow.com/questions/tagged/fiware+ngsi-go)
<br/>
![GitHub top language](https://img.shields.io/github/languages/top/lets-fiware/ngsi-go)
![Lines of code](https://img.shields.io/tokei/lines/github/lets-fiware/ngsi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/lets-fiware/ngsi-go)](https://goreportcard.com/report/github.com/lets-fiware/ngsi-go)
![CI](https://github.com/lets-fiware/ngsi-go/workflows/CI/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lets-fiware/ngsi-go/badge.svg?branch=main)](https://coveralls.io/github/lets-fiware/ngsi-go?branch=main)

The NGSI Go is a Unix command-line tool for FIWARE NGSI v2 and NGSI-LD.

| :books: [Documentation](https://ngsi-go.letsfiware.jp/) | :dart: [Roadmap](./roadmap.md) |
|---------------------------------------------------------|--------------------------------|

## Contents

<details>
<summary><strong>Details</strong></summary>

-   [Getting Started with NGSI Go](#getting-started-with-ngsi-go)
-   [Usage](#usage)
-   [Install](#install)
-   [Third party packages](#third-party-packages)
-   [Copyright and License](#copyright-and-license)

</details>

# What is NGSI Go?

The NGSI Go is a UNIX command-line tool FIWARE supporting both [NGSI v2](https://fiware-ges.github.io/orion/api/v2/stable/) and [NGSI-LD](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf), which simplifies syntax.

> "Brave (hero), bearer of the blood of Erdrick, hero of legend! Know that your weapon will not
> serve to vanquish the Dragonload."
>
> — DRAGON WARRIOR (DRAGON QUEST)

## Getting Started with NGSI Go

You can get the version of your context broker instance as shown:

```console
ngsi version -h localhost:1026
```

```json
{
 "orion" : {
   "version" : "2.5.2",
   "uptime" : "0 d, 13 h, 54 m, 48 s",
   "git_hash" : "11e4cbfef30d28347162e5c4ef4de3a5d2797c69",
   "compile_time" : "Thu Dec 17 08:43:46 UTC 2020",
   "compiled_by" : "root",
   "compiled_in" : "5a4a8800b1fa",
   "release_date" : "Thu Dec 17 08:43:46 UTC 2020",
   "doc" : "https://fiware-orion.rtfd.io/en/2.5.2/",
   "libversions": {
      "boost": "1_53",
      "libcurl": "libcurl/7.29.0 NSS/3.53.1 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
      "libmicrohttpd": "0.9.70",
      "openssl": "1.0.2k",
      "rapidjson": "1.1.0",
      "mongodriver": "legacy-1.1.2"
   }
 }
}
```

You can register an alias to access the broker.

```console
ngsi broker add --host letsfiware --brokerHost http://localhost:1026 --ngsiType v2
```

You can get the version by using the alias `letsfiware`.

```console
ngsi version -h letsfiware
```

```json
{
 "orion" : {
   "version" : "2.5.2",
   "uptime" : "0 d, 13 h, 54 m, 48 s",
   "git_hash" : "11e4cbfef30d28347162e5c4ef4de3a5d2797c69",
   "compile_time" : "Thu Dec 17 08:43:46 UTC 2020",
   "compiled_by" : "root",
   "compiled_in" : "5a4a8800b1fa",
   "release_date" : "Thu Dec 17 08:43:46 UTC 2020",
   "doc" : "https://fiware-orion.rtfd.io/en/2.5.2/",
   "libversions": {
      "boost": "1_53",
      "libcurl": "libcurl/7.29.0 NSS/3.53.1 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
      "libmicrohttpd": "0.9.70",
      "openssl": "1.0.2k",
      "rapidjson": "1.1.0",
      "mongodriver": "legacy-1.1.2"
   }
 }
}
```

Once you access the broker, you can omit to specify the broker.

```console
ngsi version
```

If you want to check the current settings, you can run the following command.

```console
ngsi settings list
```

## Usage

```text
NAME:
   ngsi - unix-like command-line tool for FIWARE NGSI and NGSI-LD

USAGE:
   ngsi [global options] command [command options] [arguments...]

VERSION:
   0.5.0 (git_hash:bd14dc0bcd75e80723bb6399bf3db0213df350e6)

COMMANDS:
   help, h  Shows a list of commands or help for one command
   CONVENIENCE:
     admin     admin command for FIWARE Orion
     cp        copy entities
     wc        print number of entities, subscriptions, registrations or types
     man       print urls of document
     ls        list entities
     rm        remove entities
     receiver  notification receiver
     template  create template of subscription or registration
     version   print the version of Context Broker
   MANAGEMENT:
     broker    manage config for broker
     context   manage @context
     settings  manage settings
     token     manage token
   NGSI:
     append   append attributes
     create   create entity(ies), subscription or registration
     delete   delete entity(ies), attribute, subscription or registration
     get      get entity(ies), attribute(s), subscription, registration or type
     list     list types, entities, subscriptions or registrations
     replace  replace entities or attributes
     update   update entities, attribute(s) or subscription
     upsert   upsert entity or entities

GLOBAL OPTIONS:
   --syslog LEVEL  specify logging LEVEL (off, err, info, debug)
   --stderr LEVEL  specify logging LEVEL (off, err, info, debug)
   --config FILE   specify configuration FILE
   --cache FILE    specify cache FILE
   --batch, -B     don't use previous args (batch) (default: false)
   --help          show help (default: false)
   --version, -v   print the version (default: false)

COPYRIGHT:
   (c) 2020-2021 Kazuhito Suda
```

## Install

### Install NGSI Go binary

The NGSI Go binary is installed in `/usr/local/bin`.

#### Installation on UNIX

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.5.0/ngsi-v0.5.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.5.0-linux-amd64.tar.gz -C /usr/local/bin
```

`ngsi-v0.5.0-linux-arm.tar.gz` and `ngsi-v0.5.0-linux-arm64.tar.gz` binaries are also available.

#### Installation on Mac

```console
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.5.0/ngsi-v0.5.0-darwin-amd64.tar.gz
sudo tar zxvf ngsi-v0.5.0-darwin-amd64.tar.gz -C /usr/local/bin
```

### Install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```console
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```

## Third party packages

The NGSI Go makes use of the following package:

| Package                                         | OSS License        |
| ----------------------------------------------- | ------------------ |
| [urfave/cli](https://github.com/urfave/cli)     | MIT License        |

The dependencies of dependencies have been omitted from the list.

## Copyright and License

Copyright (c) 2020-2021 Kazuhito Suda<br>
Licensed under the [MIT license](./LICENSE).
