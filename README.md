# Fast-algorithms-of-Reed-Solomon-erasure-codes
This program is the implementation of
Lin, Han and Chung, "Novel Polynomial Basis and Its Application to Reed-Solomon Erasure Codes," FOCS14.
(http://arxiv.org/abs/1404.3458)

## Setup

Use the following Dockerfile:
```dockerfile
# Use the official golang image as a base
FROM golang:1.20

# Install SSH server
RUN apt-get update && apt-get install -y openssh-server screen htop

# Set up SSH key (for simplicity, using a default password, change in production)
RUN mkdir /var/run/sshd
RUN echo 'root:root' | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

RUN sed -i 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' /etc/pam.d/sshd
ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

# Clean up APT cache to reduce image size
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Expose SSH port
EXPOSE 22

# Start SSH service
CMD ["/usr/sbin/sshd", "-D"]
```
NOTE: You can actually skip the SSH relative steps if they are useless for you.

## Run
Run following command one by one:
```bash
gcc -shared -o libreed_solomon.so reed_solomon.c -fPIC
go build ./main.go
./main
```

You should see the following output:
```
Input: Hello, world!
Encoded: Ifmmp-!xpsme"
Decoded: Hello, world!
```