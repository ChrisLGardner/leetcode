# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.177.0/containers/ruby/.devcontainer/base.Dockerfile

# [Choice] Ruby version: 3, 3.0, 2, 2.7, 2.6
ARG VARIANT="3.0"
FROM mcr.microsoft.com/vscode/devcontainers/ruby:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# This Dockerfile adds a non-root user with sudo access. Use the "remoteUser"
# property in devcontainer.json to use it. On Linux, the container user's GID/UIDs
# will be updated to match your local UID/GID (when using the dockerFile property).
# See https://aka.ms/vscode-remote/containers/non-root-user for details.
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Configure apt and install packages
RUN wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb \
    && dpkg -i packages-microsoft-prod.deb \
    # Update the list of products
    && apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    # Verify git, process tools installed
    && apt-get -y install git openssh-client less iproute2 procps lsb-release \
    #
    # Install ruby-debug-ide and debase
    && gem install ruby-debug-ide \
    && gem install debase \
    #
    # Install PowerShell
    && apt-get install -y powershell \
    #
    # Install Go
    && curl -O https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz \
    && tar xvf go1.16.4.linux-amd64.tar.gz \
    && chown -R root:root ./go \
    && mv go /usr/local \
    && echo "GOPATH=\$HOME/work" >> ~/.profile \
    && echo "PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> ~/.profile \
    && export GOPATH=$HOME/work \
    && export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin \
    && go get -v golang.org/x/tools/gopls \
    && go get -v github.com/ramya-rao-a/go-outline \
    # # Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
    # && groupadd --gid $USER_GID $USERNAME \
    # && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME \
    # # [Optional] Add sudo support for the non-root user
    # && apt-get install -y sudo \
    # && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME\
    # && chmod 0440 /etc/sudoers.d/$USERNAME \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Switch back to dialog for any ad-hoc use of apt-get
ENV DEBIAN_FRONTEND=dialog
