## CÂU LỆNH ĐỂ VÀO CSDL REDIS (name-container: là tên container REDIS trong docker)

```bash
docker exec -it <name-container> redis-cli
```
## CÂU LỆNH ĐỂ VÀO CSDL MYSQL (name-container: là tên container MySQL trong docker)

```bash
docker exec -it <name-container> mysql -u root -p
```
## Lệnh liệt kê ra tất cả các bảng trong CSDL

```sql
show tables;
``` 
## Lệnh xem định các trường trong một bảng cơ
```sql
DESCRIBE <name-table>\G;
```
