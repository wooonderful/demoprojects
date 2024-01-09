FROM golang:1.21.4 as builder

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
&& apt-get clean \
&& echo "root:1" | chpasswd \
&& echo "PasswordAuthentication no" >> /etc/ssh/sshd_config  \
&& echo "PermitRootLogin yes" >> /etc/ssh/sshd_config  \
&& echo "service ssh start && tail -f /dev/null" > /root/sshd.sh \
&& chmod +x /root/sshd.sh

RUN mkdir -p ~/.ssh \
&& echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC9hkAcxUCkUhc3kjP8BSBd/sLSJ90DoqAj+VQKC9v+qHosWnNYISiU32xal98edwh/Cruav12lcHkmjVGUpy/10BB8pJ617QxQv8MS8xeOmCHQ1X2tDO3Ar9/+3J75dbAjwld7674bIhMJa/hJR28sb1ixdi2mO7RS3v2fZSmVILVHwtCFW712YMPagmmK8rp/CIRpov3jEFHIkqwHQEjEMfLWEquUCsAeICiuiPMZNlm6NEvo8OO9DwIwXIedxnHO2WsEoYV1oau2E4Cil2YYncA5XhswDJ9MTTrpgZRhmcOJeHyYUAYodYrdcb+oK8nUbgC7vhnDOf05iJawa14Osv3KaS54xisRZeLqQAaX5ZPMtqQ701K6BI5NFB0UZiIWdQnJM7v3arF4/Hec0E7G5qHEWFCPxGUXHbf8foSbW+rJ2CrsKalVXLBnWDYtJI7heYrQR0inozRPv30HvL/q0l1Bf6OVevtwZPtghDZBSwZCuybgXmCYEcItrpsckWM= s@sdeMacBook-Pro.local" > ~/.ssh/authorized_keys

RUN echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC9hkAcxUCkUhc3kjP8BSBd/sLSJ90DoqAj+VQKC9v+qHosWnNYISiU32xal98edwh/Cruav12lcHkmjVGUpy/10BB8pJ617QxQv8MS8xeOmCHQ1X2tDO3Ar9/+3J75dbAjwld7674bIhMJa/hJR28sb1ixdi2mO7RS3v2fZSmVILVHwtCFW712YMPagmmK8rp/CIRpov3jEFHIkqwHQEjEMfLWEquUCsAeICiuiPMZNlm6NEvo8OO9DwIwXIedxnHO2WsEoYV1oau2E4Cil2YYncA5XhswDJ9MTTrpgZRhmcOJeHyYUAYodYrdcb+oK8nUbgC7vhnDOf05iJawa14Osv3KaS54xisRZeLqQAaX5ZPMtqQ701K6BI5NFB0UZiIWdQnJM7v3arF4/Hec0E7G5qHEWFCPxGUXHbf8foSbW+rJ2CrsKalVXLBnWDYtJI7heYrQR0inozRPv30HvL/q0l1Bf6OVevtwZPtghDZBSwZCuybgXmCYEcItrpsckWM= s@sdeMacBook-Pro.local" > ~/.ssh/id_rsa.pub \
&& git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /
EXPOSE 22

# 启动时执行
ENTRYPOINT ["sh","-l"]
CMD ["/root/sshd.sh"]
