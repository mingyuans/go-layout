# Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

FROM BASE_IMAGE
MAINTAINER Xiangqun Yan <yanxiangqunwork@163.com>

WORKDIR /opt/iam

RUN mkdir -p /var/log/iam

COPY iam-apiserver /opt/iam/bin/

ENTRYPOINT ["/opt/iam/bin/iam-apiserver"]
CMD ["-c", "/etc/iam/production-iam-apiserver.yaml"]
