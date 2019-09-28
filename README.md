# LetItGo

[![Build status](https://img.shields.io/travis/NoUseFreak/letitgo/master?style=flat-square)](https://travis-ci.org/NoUseFreak/letitgo)
[![Release](https://img.shields.io/github/v/release/NoUseFreak/letitgo?style=flat-square)](https://github.com/NoUseFreak/letitgo/releases)
[![Maintained](https://img.shields.io/maintenance/yes/2019?style=flat-square)](https://github.com/NoUseFreak/letitgo)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/nousefreak/letitgo?style=flat-square)](https://hub.docker.com/r/nousefreak/letitgo)
[![License](https://img.shields.io/github/license/NoUseFreak/letitgo?style=flat-square)](https://github.com/NoUseFreak/letitgo/blob/master/LICENSE)
[![Coffee](https://img.shields.io/badge/☕️-Buy%20me%20a%20coffee-blue?style=flat-square&color=blueviolet)](https://www.buymeacoffee.com/driesdepeuter)

LetItGo simplifies automated releases. A simple definition in `.release.yml` in 
the root of your project is all you need.

## Actions

### Homebrew

Currently it is only supported to update [Taps](https://docs.brew.sh/Taps).
It requires `GITHUB_TOKEN` to be set. 

The following example configuration will update `Formula/letitgo.rb`.

```yaml
homebrew:
  - name: letitgo
    description: LetItGo automates releases.
    homepage: https://github.com/NoUseFreak/letitgo
    url: https://github.com/NoUseFreak/letitgo/releases/download/{{ .Version }}/darwin_amd64.zip
    version: "{{ .Version }}"
    tap:
      url: git@github.com:NoUseFreak/homebrew-brew.git
    test: system "#{bin}/{{ .Name }} -h"
```

```bash
$ letitgo homebrew $(git describe --tags --abbrev=0)
```
