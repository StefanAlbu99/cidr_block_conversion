FROM alpine:latest

# Install bash
RUN apk add --no-cache bash jq

# Download and install Terraform
RUN wget -O /tmp/terraform.zip https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip && \
    unzip /tmp/terraform.zip -d /usr/local/bin/ && \
    rm /tmp/terraform.zip && \
    chmod +x /usr/local/bin/terraform

WORKDIR /terraform_code

RUN terraform init

ENTRYPOINT ["/bin/bash"]
