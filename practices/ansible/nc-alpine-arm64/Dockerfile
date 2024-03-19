FROM arm64v8/alpine:3.17

ARG USER_ANSIBLE=elliot

# Install tools
RUN apk add --no-cache bash sudo openssh ansible

# Create user and add in sudoers
RUN adduser -D ${USER_ANSIBLE} && echo "elliot  ALL=(ALL)  NOPASSWD: ALL" >> /etc/sudoers

# Change user
USER ${USER_ANSIBLE}

CMD [ "tail", "-f", "/dev/null" ]