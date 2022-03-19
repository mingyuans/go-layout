# Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

FROM BASE_IMAGE
MAINTAINER Xiangqun Yan <yanxiangqunwork@163.com>

WORKDIR /opt/api

RUN mkdir -p /var/log/api

COPY apiserver /opt/api/bin/

ENTRYPOINT ["/opt/api/bin/apiserver"]
CMD ["-c", "/etc/api/production-apiserver.yaml"]
