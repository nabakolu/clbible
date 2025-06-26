# Maintainer: Lukas Nabakowski lnabakowski at mailo.com
pkgname=clbible
pkgver=0.2
pkgrel=1
pkgdesc="Query bible verses"
arch=('any')
url="https://github.com/nabakolu/clbible"
license=('GPL-3.0')
makedepends=('go' 'git')

source=("clbible.go" "cli.go" "config.go" "parser.go" "LICENSE")
sha256sums=('SKIP' 'SKIP' 'SKIP' 'SKIP' 'SKIP')

build() {
  cd "$srcdir"

  go build *.go
}

package() {
  cd "$srcdir"

  # Install the binary
  install -Dm755 clbible "$pkgdir/usr/bin/clbible"

  # Install the LICENSE file
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
