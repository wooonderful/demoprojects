FROM golang:1.18.7 as builder

SHELL ["/bin/bash", "-c"]

RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc \
&& echo "export GOPRIVATE=github.com/emqx" >> ~/.bashrc \
&& echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.bashrc \
&& source ~/.bashrc

# Build Delve
RUN GOPROXY=https://goproxy.cn,direct go install github.com/go-delve/delve/cmd/dlv@latest

RUN apt-get update \
&& apt-get install ssh -y \
&& apt-get clean \
&& echo "root:1" | chpasswd \
&& echo "PermitRootLogin yes" >> /etc/ssh/sshd_config  \
&& echo "service ssh start && tail -f /dev/null" > /root/sshd.sh \
&& chmod +x /root/sshd.sh

WORKDIR /
EXPOSE 22

# 启动时执行
ENTRYPOINT ["sh","-l"]
CMD ["/root/sshd.sh"]