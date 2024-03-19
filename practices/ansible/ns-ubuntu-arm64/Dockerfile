FROM arm64v8/ubuntu:18.04

ARG SSH_USER=elliot

# Install tools
RUN apt-get update && apt-get install -y openssh-server sudo

# Copy sshd_config
COPY sshd_config /etc/ssh/sshd_config

# Create user and add in sudoers
RUN useradd -m ${SSH_USER} -s /bin/bash && echo "elliot  ALL=(ALL)  NOPASSWD: ALL" >> /etc/sudoers

# Change user
USER ${SSH_USER}

# Start ssh server
RUN sudo service ssh start

# Execute process
CMD [ "tail", "-f", "/dev/null" ]