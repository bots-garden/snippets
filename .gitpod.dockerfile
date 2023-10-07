FROM gitpod/workspace-full
USER gitpod

# ------------------------------------
# Install Go
# ------------------------------------
RUN <<EOF
GO_VERSION="1.21.1"

GOPATH=$HOME/go-packages
GOROOT=$HOME/go
PATH=$GOROOT/bin:$GOPATH/bin:$PATH

curl -fsSL https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz | tar xzs \
    && printf "%s\n" "export GOPATH=/workspace/go" \
                     "export PATH=$GOPATH/bin:$PATH" > $HOME/.bashrc.d/300-go
go version
go install -v golang.org/x/tools/gopls@latest
go install -v github.com/ramya-rao-a/go-outline@latest
go install -v github.com/stamblerre/gocode@v1.0.0

EOF
