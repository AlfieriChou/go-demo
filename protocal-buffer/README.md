### 环境搭建

* 采用docker镜像

    * 至于为什么采用docker镜像，原因你懂的。

    ```bash
    docker pull znly/protoc
    ```

    ```bash
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=. -I. person.proto
    ```