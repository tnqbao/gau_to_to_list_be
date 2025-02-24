<h1 style="text-align:center"> Gấu To-do-list API 
 </h1>

***
## Cài đặt mã nguồn
### Clone mã nguồn dự án:
   ``` bash
    git clone https://github.com/tnqbao/gau_to_do_list_be.git
    cd gau_to_do_list_be
   ```
### Chỉnh sửa lại module trong go.mod:
  ``` bash
   go mod edit -module=your-link-github-repo 
  ```
### Cài đặt các thư viện cần thiết:
  ``` bash
    go mod tidy 
  ``` 

### Cấu hình bến môi trường:
* Tạo file .env như mẫu trong thư mục gốc của dự án:
```dotenv
 LIST_DOMAIN=your-domain (mặc định là localhost)
```

### Chạy ở chế độ dev mode:
   ``` bash 
    go run -tags dev ./cmd/
   ```
  
### Chạy ở chế độ build production:
   ``` bash
    go build -tags release ./cmd/
   ```
  <li>Truy cập api tại: <a href="http://localhost:8088" target="_blank">http://localhost:8088</a></li>

***

## Triển khai với Docker

* Build image:
    ``` bash
    docker build -t gau_to_do_list .
    ```
  
* Chạy container:
    ``` bash
    docker run -p 8088:8088 gau_to_do_list
    ```
* **
## Nâng cao - Triển khai tự động (CICD với Docker và Github Actions lên cloud)

### Luồng triển khai:

    1.Cập nhật mã nguồn code lên nhánh master Github
    2.Github Actions tự động build image và push lên Docker Hub
    3.Pull image từ Docker Hub và chạy container trên cloud
### Cấu hình Github Actions:
 * Cấu hình biến môi trường trong repository github:
     * `DOCKERHUB_USERNAME`: Tên đăng nhập Docker Hub
     * `DOCKERHUB_TOKEN`: Mật khẩu Docker Hub
     * `SSH_PRIVATE_KEY`: Private key để kết nối với server cloud
     * `SERVER USER` :  Username deploy trên server
     * `SERVER_IP` : Địa chỉ IP của server
 
 * File yml mặc định : https://github.com/tnqbao/gau-to-do-list_be/blob/master/.github/workflows/production-deployment.yml

### Cấu hình deploy lên server:
 * Tạo file .env như mẫu phía trên
 * Tạo script thực thi deploy_to_do_list.sh 
    ```bash
    #!/bin/bash

    IMAGE_NAME=your-dockerhub-image-name
    CONTAINER_NAME="gau_to_do_list"
    ENV_FILE=".env"
    
    echo "==== Start Update To Do List Service ===="
   
    # Dừng container cũ nếu có
    if docker ps -a --format '{{.Names}}' | grep -q "^$CONTAINER_NAME$"; then
    echo "Stopping current container: $CONTAINER_NAME..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
    else
    echo "No container running with name: $CONTAINER_NAME."
    fi
   
    # Pull image vừa build từ Docker Hub
    echo "Pulling latest image: $IMAGE_NAME..."
    docker pull $IMAGE_NAME
    echo "==== Update  To Do List Service successful ===="
   
   # Chạy container và bind file .env vào container
   docker run -d -p 8088:8088 --env-file $ENV_FILE --name $CONTAINER_NAME -v $(pwd)/$ENV_FILE:/gau_to_do_list/.env $IMAGE_NAME
    ```
 * Cấp quyền thực thi cho script:
    ``` bash
    chmod +x deploy_to_do_list.sh
    ```
   

 <i>Vậy là xong rồi,giờ thì khi bạn khi push code lên Github, Github Actions sẽ tự động build image và deploy lên server cloud!! </i>

***
## Demo
* Xem demo api tại: <a href="https://ljcbrtxa20.apidog.io/" target="_blank">https://ljcbrtxa20.apidog.io/</a>
* File postman collection: https://driver.daudoo.com/Gau%20To%20Do%20List%20Collection.postman_collection.json
***

