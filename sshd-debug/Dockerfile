FROM golang:1.23.4 as builder

SHELL ["/bin/bash", "-c"]

RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc \
&& echo "export GOPRIVATE=github.com/emqx" >> ~/.bashrc \
&& echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.bashrc \
&& source ~/.bashrc

# Build Delve
RUN GOPROXY=https://goproxy.cn,direct go install github.com/go-delve/delve/cmd/dlv@latest

RUN apt-get update \
&& apt-get install ssh -y \
&& apt-get install vim -y \
&& apt-get install rsync -y \
&& apt-get clean \
&& echo "root:1" | chpasswd \
&& echo "PasswordAuthentication no" >> /etc/ssh/sshd_config  \
&& echo "PermitRootLogin yes" >> /etc/ssh/sshd_config  \
&& echo "service ssh start && tail -f /dev/null" > /root/sshd.sh \
&& chmod +x /root/sshd.sh

RUN mkdir -p ~/.ssh \
&& echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIECqehnJubTofDRo65kPsT0MzyEmmJnocT/o0BJBqPka sx@sxdeMac-mini" > ~/.ssh/authorized_keys

RUN echo "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIECqehnJubTofDRo65kPsT0MzyEmmJnocT/o0BJBqPka sx@sxdeMac-mini" > ~/.ssh/id_ed25519.pub \
&& git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /
EXPOSE 22

# 启动时执行
ENTRYPOINT ["sh","-l"]
CMD ["/root/sshd.sh"]
